package apiserver

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// LoadX509KeyPair() unifies the loading of X509 key pairs for different components
func loadX509KeyPair(sysLogger *logrus.Logger, certfile, keyfile, componentName, certAttr string) (tls.Certificate, error) {
	keyPair, err := tls.LoadX509KeyPair(certfile, keyfile)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("init: loadX509KeyPair(): loading %s X509KeyPair for %s from %s and %s - FAIL: %v",
			certAttr, componentName, certfile, keyfile, err)
	}
	sysLogger.Debugf("loading %s X509KeyPair for %s from %s and %s - OK", certAttr, componentName, certfile, keyfile)
	return keyPair, nil
}

// function unifies the loading of CA certificates for different components
func loadCACertificate(sysLogger *logrus.Logger, certfile string, componentName string, certPool *x509.CertPool) error {

	// Read the certificate file content
	caRoot, err := os.ReadFile(certfile)
	if err != nil {
		return fmt.Errorf("init: loadCACertificate(): loading %s CA certificate from '%s' - FAIL: %w", componentName, certfile, err)
	}
	sysLogger.Debugf("loading %s CA certificate from '%s' - OK", componentName, certfile)

	// Return error if provided certificate is nil
	if certPool == nil {
		return errors.New("init: loadCACertificate(): provided certPool is nil")
	}

	// Append a certificate to the pool
	certPool.AppendCertsFromPEM(caRoot)
	return nil
}
