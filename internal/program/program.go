package program

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/contrib/secure"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"modules/internal/logger"
	"modules/internal/samlidp"
	"net/http"
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
}
