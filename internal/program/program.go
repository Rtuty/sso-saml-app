package program

import (
	"context"
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"
	"github.com/spf13/viper"
	"log"
	"modules/internal/logger"
	"modules/internal/samlidp"
	"modules/internal/tools"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

type Program struct {
	exit  chan struct{}
	cfg   *viper.Viper
	srv   *http.Server
	store *samlidp.SqliteStore
}

func New(cfg *viper.Viper, logger *logger.LoggerEx, workDir string) *Program {
	p := new(Program)
	p.cfg = cfg

	host := p.cfg.GetString("server.host")
	port := p.cfg.GetInt("server.port")

	baseURLstr := fmt.Sprintf(`http://%s:%d`, host, port)
	baseURL, err := url.Parse(baseURLstr)
	if err != nil {
		log.Fatalf("cannot parse base URL: %v", err)
		return nil
	}

	certFilename := p.cfg.GetString("pathes.certFile")
	cert, err := p.getCert(certFilename)
	if err != nil {
		log.Fatalf("Get sert ERROR:%s", err)
		return nil
	}

	keyFilename := p.cfg.GetString("pathes.keyFile")
	key, err := p.getKey(keyFilename)
	if err != nil {
		log.Fatalf("Get key ERROR:%s", err)
		return nil
	}

	store, err := samlidp.NewSqliteStore(filepath.Join(workDir, "database.sqlite"))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	p.store = store

	idpServer, err := samlidp.New(samlidp.Options{
		URL:         *baseURL,
		Key:         key,
		Certificate: cert,
		Store:       p.store,
	})
	if err != nil {
		log.Fatalf("create idp: %s", err)
		return nil
	}
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin.DefaultWriter = logger

	router := gin.Default()
	var r = router

	r.Use(gin.Recovery())
	r.Use(static.Serve("/", static.LocalFile(filepath.Join(workDir, "web"), true)))
	r.Use(secure.Secure(secure.Options{
		AllowedHosts:            nil,
		SSLRedirect:             false,
		SSLTemporaryRedirect:    false,
		SSLHost:                 "",
		SSLProxyHeaders:         nil,
		STSSeconds:              0,
		STSIncludeSubdomains:    false,
		FrameDeny:               true, // не показывать сайт в фрейме
		CustomFrameOptionsValue: "",
		ContentTypeNosniff:      true,                 //
		BrowserXssFilter:        true,                 //
		ContentSecurityPolicy:   "default-src 'self'", //
		IsDevelopment:           false,
		BadHostHandler:          nil,
	}))

	idpServer.InitializeHTTP(router)

	addr := fmt.Sprintf("%s:%d", p.cfg.GetString("server.host"), p.cfg.GetInt("server.port"))

	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		WriteTimeout: 5 * time.Minute,
	}

	p.srv = srv

	return p
}

// getCert получает сертификат
func (p *Program) getCert(fileName string) (*x509.Certificate, error) {
	var certData []byte

	if tools.FileExists(fileName) {
		var err error
		certData, err = os.ReadFile(fileName)
		if err != nil {
			log.Fatalf("Reading idp cert. ERROR: %s", err)
			return nil, err
		}
	}

	b, _ := pem.Decode(certData)
	c, err := x509.ParseCertificate(b.Bytes)
	if err != nil {
		log.Fatalf("Parsing idp cert. ERROR: %s", err)
		return nil, err
	}
	return c, nil
}

// getKey получает ключ
func (p *Program) getKey(fileName string) (crypto.PrivateKey, error) {
	var keyData []byte

	if tools.FileExists(fileName) {
		var err error
		keyData, err = os.ReadFile(fileName)
		if err != nil {
			log.Fatalf("Reading idp key. ERROR: %s", err)
			return nil, err
		}
	}

	b, _ := pem.Decode(keyData)
	k, err := x509.ParsePKCS8PrivateKey(b.Bytes)
	if err != nil {
		log.Fatalf("Parsing id key. ERROR: %s", err)
		return nil, err
	}

	return k, nil
}

// Start вызывается при запуске службы
func (p *Program) Start(s service.Service) error {
	p.exit = make(chan struct{})
	go p.run()
	return nil
}

// Stop вызывается при запуске службы
func (p *Program) Stop(s service.Service) error {
	close(p.exit)
	return nil
}

func (p *Program) run() {
	go func() {
		if err := p.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ERROR: %s\n", err)
		}
	}()

	log.Printf("Server is running at %s", p.srv.Addr)

	<-p.exit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	p.srv.Shutdown(ctx)

	p.store.Close()
}
