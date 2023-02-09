package samlidp

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type User struct {
	Name              string   `json:"name"`
	PlaintextPassword *string  `json:"password,omitempty"`
	HashedPassword    []byte   `json:"hashed_password,omitempty"`
	Groups            []string `json:"groups,omitempty"`
	Email             string   `json:"email,omitempty"`
	CommonName        string   `json:"common_name,omitempty"`
	Surname           string   `json:"surname,omitempty"`
	GivenName         string   `json:"given_name,omitempty"`
	ScopedAffiliation string   `json:"scoped_affiliation,omitempty"`
}

// ListUsers обрабатывает запрос `GET /api/v1/users/` и отвечает списком имен пользователей в формате JSON.
func (s *Server) ListUsers(c *gin.Context) {
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

// GetUser обрабатывает запрос `GET /api/v1/users/:id` и отвечает пользовательским объектом в формате JSON. Поле хэшированного пароля исключено.
func (s *Server) GetUser(c *gin.Context) {
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

func (s *Server) PutUser(c *gin.Context) {
	r := c.Request
	w := c.Writer
	id := c.Param("id")

	user := User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		s.logger.Printf("User body = nil. ERROR: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	user.Name = id

	if user.PlaintextPassword != nil {
		var err error
		user.HashedPassword, err = bcrypt.GenerateFromPassword([]byte(*user.PlaintextPassword), bcrypt.DefaultCost)
		if err != nil {
			s.logger.Printf("Get hashed user password. ERROR: %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	} else {
		existingUser := User{}
		err := s.Store.Get(fmt.Sprintf("/users/%s", id), &existingUser)
		switch {
		case err == nil:
			user.HashedPassword = existingUser.HashedPassword
		case err == ErrNotFound:
		//todo
		default:
			s.logger.Printf("ERROR: %s", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	user.PlaintextPassword = nil

	err := s.Store.Put(fmt.Sprintf("/users/%s", id), &user)
	if err != nil {
		s.logger.Printf("ERROR: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser удаляет пользователя
func (s *Server) DeleteUser(c *gin.Context) {
	w := c.Writer
	id := c.Param("id")

	err := s.Store.Delete(fmt.Sprintf("/users/%s", id))
	if err != nil {
		s.logger.Printf("ERROR: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
