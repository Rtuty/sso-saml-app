package program

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"modules/internal/logger"
	"modules/internal/samlidp"
	"modules/internal/tools"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type Program struct {
	exit  chan struct{}
	cfg   *viper.Viper
	srv   *http.Server
	store *samlidp.SqliteStore
}

func New(cfg *viper.Viper, logger *logger.LoggerEx, workDir string) *Program {
	p := new(Program)
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

	//todo add key

	store, err := samlidp.NewSqliteStore(filepath.Join(workDir, "database.sqlite"))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	p.store = store

	idpServer, err := samlidp.New(samlidp.Options{
		URL:         *baseURL,
		Key:         nil,  //todo
		Certificate: cert, //todo
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

	//todo return
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
