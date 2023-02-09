package samlidp

import (
	"encoding/xml"
	"errors"
	"github.com/tenrok/saml"
	"io"
)

func randomBytes(n int) []byte {
	rv := make([]byte, n)

	if _, err := saml.RandReader.Read(rv); err != nil {
		panic(err)
	}
	return rv
}

func getSPMetadata(r io.Reader) (spMetadata *saml.EntityDescriptor, err error) {
	var data []byte
	if data, err = io.ReadAll(r); err != nil {
		return nil, err
	}

	if err := xml.Unmarshal(data, &spMetadata); err != nil {
		if err.Error() == "expected element type <EntityDescriptor> but have <EntitiesDescriptor>" {
			entities := &saml.EntitiesDescriptor{}
			if err := xml.Unmarshal(data, &entities); err != nil {
				return nil, err
			}

			for _, e := range entities.EntityDescriptors {
				if len(e.SPSSODescriptors) > 0 {
					return &e, nil
				}
			}

			return nil, errors.New("metadata contained no service provider metadata")
		}

		return nil, err
	}

	return spMetadata, nil
}