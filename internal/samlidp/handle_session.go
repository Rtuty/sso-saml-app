package samlidp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tenrok/saml"
	"net/http"
)

// Login обрабатывает `POST /login/` и `GET /login/` запросы.
func (s *Server) Login(c *gin.Context) {
	r := c.Request
	w := c.Writer

	if err := r.ParseForm(); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	session := s.GetSession(w, r, &saml.IdpAuthnRequest{IDP: &s.IDP})
	if session == nil {
		return
	}

	c.JSON(http.StatusOK, session)
}

// ListSessions обрабатывает `GET /sessions/` запрос и выводит список сессий в формате JSON
func (s *Server) ListSessions(c *gin.Context) {
	w := c.Writer
	sessions, err := s.Store.List("/sessions")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, struct {
		Sessions []string `json:"sessions"'`
	}{Sessions: sessions})
}

func (s *Server) GetSessionByID(c *gin.Context) {
	w := c.Writer
	id := c.Param("id")

	session := saml.Session{}

	err := s.Store.Get(fmt.Sprintf("/sessions/%s", id), &session)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, session)
}
