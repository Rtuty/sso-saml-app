package samlidp

import (
	"crypto"
	"crypto/x509"
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
