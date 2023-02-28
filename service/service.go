package main

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/crewjam/saml/samlsp"
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
		panic(errors.New(fmt.Sprintf("Service keypair parsing ERROR: %s", err)))
	}

	idpMetadataURL, err := url.Parse("http://localhost:8000/metadata")
	if err != nil {
		panic(errors.New(fmt.Sprintf("Service metadata URL parsing have hil result ERROR: %s", err)))
	}

	idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient, *idpMetadataURL)
	if err != nil {
		panic(errors.New(fmt.Sprintf("Service fetch metadata have hil result ERROR: %s", err)))
	}

	rootURL, err := url.Parse("http://localhost:8001")
	if err != nil {
		panic(errors.New(fmt.Sprintf("RootURL service ERROR: %s", err)))
	}

	samlSP, _ := samlsp.New(samlsp.Options{
		URL:         *rootURL,
		Key:         keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate: keyPair.Leaf,
		IDPMetadata: idpMetadata,
	})

	spMetadataBuf, _ := xml.MarshalIndent(samlSP.ServiceProvider.Metadata(), "", " ")
	spURL := *idpMetadataURL
	spURL.Path = "/api/v1/services/sp"
	http.Post(spURL.String(), "text/xml", bytes.NewReader(spMetadataBuf))

	app := http.HandlerFunc(hello)
	http.Handle("/hello/", samlSP.RequireAccount(app))
	http.Handle("/saml/", samlSP)
	http.ListenAndServe(":8001", nil)
}
