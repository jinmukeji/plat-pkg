package mysql

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"os"
)

func NewTLSConfig(rootCert string) (*tls.Config, error) {
	rootCertPool := x509.NewCertPool()
	pem, err := os.ReadFile(rootCert)
	if err != nil {
		return nil, err
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		return nil, errors.New("failed to append PEM")
	}

	return &tls.Config{
		RootCAs: rootCertPool,
	}, nil
}
