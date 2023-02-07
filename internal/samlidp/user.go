package samlidp

import (
	"fmt"
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

// HandleListUsers обрабатывает запрос `GET /api/v1/users/` и отвечает списком имен пользователей в формате JSON.
func (s *Server) HandleListUsers(c *gin.Context) {
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

// HandleGetUser обрабатывает запрос `GET /api/v1/users/:id` и отвечает пользовательским объектом в формате JSON. Поле хэшированного пароля исключено.
func (s *Server) HandleGetUser(c *gin.Context) {
	w := c.Writer
	id := c.Param("id")

	user := User{}

	err := s.Store.Get(fmt.Sprintf("users/%s", id), &user)
	if err != nil {
		s.logger.Printf("ERROR: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	user.HashedPassword = nil
	c.JSON(http.StatusOK, user)
}

func (s *Server) HandlePutUser(c *gin.Context) {}

func (s *Server) HandleDeleteUser(c *gin.Context) {}
