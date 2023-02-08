package samlidp

import (
	"github.com/tenrok/saml"
	"net/http"
	"os"
)

type Service struct {
	Name     string
	Metadata saml.EntitiesDescriptor
}

func (s *Server) GetServiceProvider(r *http.Request, serviceProviderID string) (*saml.EntityDescriptor, error) {
	s.idpConfigMu.RLock()
	defer s.idpConfigMu.Unlock()

	rv, ok := s.serviceProviders[serviceProviderID]
	if !ok {
		return nil, os.ErrNotExist
	}

	return rv, nil
}