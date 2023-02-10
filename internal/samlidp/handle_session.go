package samlidp

import (
	"github.com/gin-gonic/gin"
	"github.com/tenrok/saml"
	"net/http"
)

func (s *Server) HandleLogin(c *gin.Context) {
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
