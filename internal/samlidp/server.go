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