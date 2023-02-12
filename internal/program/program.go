package program

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"modules/internal/logger"
	"modules/internal/samlidp"
	"net/http"
	"net/url"
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

	//todo add certificate

	//todo add key

	store, err := samlidp.NewSqliteStore(filepath.Join(workDir, "database.sqlite"))
	if err != nil {
		log.Fatal(err)
		return nil
	}
	p.store = store

	idpServer, err := samlidp.New(samlidp.Options{
		URL:         *baseURL,
		Key:         nil, //todo
		Certificate: nil, //todo
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