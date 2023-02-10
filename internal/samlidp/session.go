package samlidp

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/tenrok/saml"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"text/template"
	"time"
)

var sessionMaxAge = time.Hour

/*
GetSession возвращает указатель на сессию.
Пользователь указывает логин и пароль. Далее проверка на валидность и соответсвие в хранилище (пока локально).
Если данные действительны, он GetSession устанавливает файл cookie и возвращает созданный объект сессии.
TODO: Если сессионный файл cookie уже существует и представляет действительный сеанс, то сеанс возвращается
*/
func (s *Server) GetSession(w http.ResponseWriter, r *http.Request, req *saml.IdpAuthnRequest) *saml.Session {
	if r.Method == "POST" && r.PostForm.Get("user") != "" {
		user := User{}
		if err := s.Store.Get(fmt.Sprintf("/users/%s", r.PostForm.Get("user")), &user); err != nil {
			s.sendLoginForm(w, r, req, "Invalid username or password")
			return nil
		}

		//Чекаем кэш пароля юзверя
		if err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(r.PostForm.Get("password"))); err != nil {
			s.sendLoginForm(w, r, req, "Invalid username or password")
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
		if err := s.Store.Put(fmt.Sprintf("/sessions/%s", session.ID), &session); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return nil
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    session.ID,
			MaxAge:   int(sessionMaxAge.Seconds()),
			HttpOnly: true,
			Secure:   r.URL.Scheme == "https",
			Path:     "/",
		})
		return session
	}

	return nil
}

func (s *Server) sendLoginForm(w http.ResponseWriter, r *http.Request, req *saml.IdpAuthnRequest, toast string) {
	tmpl := template.Must(template.New("saml-post-form").Parse(`` +
		`<html>` +
		`<p>{{.Toast}}</p>` +
		`<form method="post" action="{{.URL}}">` +
		`<input type="text" name="user" placeholder="user" value="" />` +
		`<input type="password" name="password" placeholder="password" value="" />` +
		`<input type="hidden" name="SAMLRequest" value="{{.SAMLRequest}}" />` +
		`<input type="hidden" name="RelayState" value="{{.RelayState}}" />` +
		`<input type="submit" value="Log In" />` +
		`</form>` +
		`</html>`))
	data := struct {
		Toast       string
		URL         string
		SAMLRequest string
		RelayState  string
	}{
		Toast:       toast,
		URL:         req.IDP.SSOURL.String(),
		SAMLRequest: base64.StdEncoding.EncodeToString(req.RequestBuffer),
		RelayState:  req.RelayState,
	}

	if err := tmpl.Execute(w, data); err != nil {
		panic(err)
	}
}