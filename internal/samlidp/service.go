package samlidp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tenrok/saml"
	"net/http"
	"os"
)

type Service struct {
	Name     string
	Metadata saml.EntityDescriptor // Метадата выдается в XML формате
}

// initService Иннициализирует все сервисы и запускает identity provider для их обработки
func (s *Server) initService() error {
	serviceNames, err := s.Store.List("/services/")
	if err != nil {
		s.logger.Printf("Service initialization error: %s", err)
		return err
	}

	for _, serviceName := range serviceNames {
		service := Service{}
		if err := s.Store.Get(fmt.Sprintf("/services/%s", serviceName), &service); err != nil {
			s.logger.Printf("Service name list = nil. ERROR: %s", err)
			return err
		}

		s.idpConfigMu.Lock()
		s.serviceProviders[service.Metadata.EntityID] = &service.Metadata
		s.idpConfigMu.Unlock()
	}
	return nil
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

// ListServices выдает json список всех сервисов
func (s *Server) ListServices(c *gin.Context) {
	services, err := s.Store.List("/services/")
	if err != nil {
		s.logger.Printf("ERROR: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"services": services})
}

// GetService выдает XML, в которой записана метадата, по указоному id севриса в url
func (s *Server) GetService(c *gin.Context) {
	id := c.Param("id")
	service := Service{}
	err := s.Store.Get(fmt.Sprintf("/services/%s", id), &service)
	if err != nil {
		s.logger.Printf("ERROR: %s", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.XML(http.StatusOK, service.Metadata)
}

//
func (s *Server) PutService(c *gin.Context) {}

//
func (s *Server) DeleteService(c *gin.Context) {}
