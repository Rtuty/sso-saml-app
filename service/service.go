package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/tenrok/saml/samlsp"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n\n", samlsp.AttributeFromContext(r.Context(), "cn"))
	fmt.Fprintf(w, "uid: %s\n", samlsp.AttributeFromContext(r.Context(), "uid"))
	fmt.Fprintf(w, "groups: %s\n", samlsp.AttributeFromContext(r.Context(), "eduPersonAffiliation"))
	fmt.Fprintf(w, "email: %s\n", samlsp.AttributeFromContext(r.Context(), "mail"))
	fmt.Fprintf(w, "cn (common_name): %s\n", samlsp.AttributeFromContext(r.Context(), "cn"))
	fmt.Fprintf(w, "sn (surname): %s\n", samlsp.AttributeFromContext(r.Context(), "sn"))
	fmt.Fprintf(w, "givenName: %s\n", samlsp.AttributeFromContext(r.Context(), "givenName"))
}

func main() {
	keyPair, err := tls.LoadX509KeyPair("service.cert", "service.key")
	if err != nil {
		panic(errors.New(fmt.Sprintf("Service keypair = nil. ERROR: %s", err)))
	}

	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		panic(errors.New(fmt.Sprintf("Service keypair parsing error: %s", err)))
	}
}
