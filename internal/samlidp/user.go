package samlidp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name              string   `json:"name"`
	PlaintextPassword *string  `json:"password,omitempty"` // not stored
	HashedPassword    []byte   `json:"hashed_password,omitempty"`
	Groups            []string `json:"groups,omitempty"`
	Email             string   `json:"email,omitempty"`
	CommonName        string   `json:"common_name,omitempty"`
	Surname           string   `json:"surname,omitempty"`
	GivenName         string   `json:"given_name,omitempty"`
	ScopedAffiliation string   `json:"scoped_affiliation,omitempty"`
}

// HandlerListUsers обрабатывает запрос `GET /api/v1/users/` и отвечает списком имен пользователей в формате JSON.
func (s *Server) HandlerListUsers(c *gin.Context) {
	w := c.Writer

	users, err := s.Store.List("/users/")
	if err != nil {
		s.logger.Printf("ERROR: %S", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, struct {
		Users []string `json:"users"`
	}{Users: users})
}

func (s *Server) HandleGetUser(c *gin.Context) {}

func (s *Server) HandlePutUser(c *gin.Context) {}

func (s *Server) HandleDeleteUser(c *gin.Context) {}
