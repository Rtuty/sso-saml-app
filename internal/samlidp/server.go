package samlidp

import (
	"crypto"
	"crypto/x509"
	"github.com/gin-gonic/gin"
	"github.com/tenrok/saml"
	"github.com/tenrok/saml/logger"
	"net/http"
	"net/url"
	"sync"
)

type Options struct {
	URL         url.URL
	Key         crypto.PrivateKey
	Logger      logger.Interface
	Certificate *x509.Certificate
	Store       Store
}

type Server struct {
	http.Handler
	idpConfigMu      sync.RWMutex
	logger           logger.Interface
	serviceProviders map[string]*saml.EntityDescriptor
	IDP              saml.IdentityProvider
	Store            Store
}

// New возвращает новый сервер
func New(opts Options) (*Server, error) {
	metadataURL := opts.URL
	metadataURL.Path = metadataURL.Path + "/metadata"

	ssoURL := opts.URL
	ssoURL.Path = ssoURL.Path + "/sso"

	logr := opts.Logger
	if logr == nil {
		logr = logger.DefaultLogger
	}

	s := &Server{
		serviceProviders: map[string]*saml.EntityDescriptor{},
		IDP: saml.IdentityProvider{
			Key:         opts.Key,
			Logger:      logr,
			Certificate: opts.Certificate,
			MetadataURL: metadataURL,
			SSOURL:      ssoURL,
		},
		logger: logr,
		Store:  opts.Store,
	}

	s.IDP.SessionProvider = s
	s.IDP.ServiceProviderProvider = s

	if err := s.initService(); err != nil {
		return nil, err
	}

	return s, nil
}

// InitializeHTTP TODO
func (s *Server) InitializeHTTP(router *gin.Engine) {
	router.GET("/metadata", func(c *gin.Context) {
		s.idpConfigMu.RLock()
		defer s.idpConfigMu.RUnlock()
		s.IDP.ServeMetadata(c.Writer, c.Request)
	})

	router.Any("/sso", func(c *gin.Context) {
		s.idpConfigMu.RLock()
		defer s.idpConfigMu.RUnlock()
		s.IDP.ServeMetadata(c.Writer, c.Request)
	})

	router.Any("login", s.Login)

	v1 := router.Group("/api/v1")
	{
		usr := v1.Group("/users")
		{
			usr.GET("/", s.ListUsers)
			usr.GET("/:id", s.GetUser)
			usr.PUT("/:id", s.PutUser)
			usr.DELETE("/:id", s.DeleteUser)
		}

		srv := v1.Group("/service")
		{
			srv.GET("/", s.ListServices)
			srv.GET("/:id", s.GetService)
			srv.PUT("/:id", s.PutService)
			srv.POST("/:id", s.PutService)
			srv.DELETE("/:id", s.DeleteService)
		}

		ssn := v1.Group("/sessions")
		{
			ssn.GET("/", s.ListSessions)
		}
	}
}
