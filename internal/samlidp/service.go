package samlidp

import (
	"github.com/gin-gonic/gin"
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

func (s *Server) HandleListServices(c *gin.Context) {
	services, err := s.Store.List("/services/")
	if err != nil {
		s.logger.Printf("ERROR: %S", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"services": services})
}