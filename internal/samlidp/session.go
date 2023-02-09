package samlidp

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/tenrok/saml"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var sessionMaxAge = time.Hour

func (s *Server) GetSession(w http.ResponseWriter, r *http.Request, req *saml.IdpAuthnRequest) *saml.Session {
	if r.Method == "POST" && r.PostForm.Get("user") != "" {
		user := User{}
		if err := s.Store.Get(fmt.Sprintf("/users/%s", r.PostForm.Get("user")), &user); err != nil {
			return nil
		}

		if err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(r.PostForm.Get("password"))); err != nil {
			return nil
		}

		session := &saml.Session{
			ID:                    base64.StdEncoding.EncodeToString(randomBytes(32)),
			NameID:                user.Email,
			CreateTime:            saml.TimeNow(),
			ExpireTime:            saml.TimeNow().Add(sessionMaxAge),
			Index:                 hex.EncodeToString(randomBytes(32)),
			UserName:              user.Name,
			Groups:                user.Groups[:],
			UserEmail:             user.Email,
			UserCommonName:        user.CommonName,
			UserSurname:           user.Surname,
			UserGivenName:         user.GivenName,
			UserScopedAffiliation: user.ScopedAffiliation,
		}
	}
	if err := s.Store.Put(fmt.Sprintf("/sessions/%s", session.ID), &session); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return nil
	}

	return nil
}
