package apiserver

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/logger"
	"github.com/zintpavowj/Zero_Trust_Score-based_Network_Level_AC_PIP/internal/app/store"
)

// server ...
type server struct {
	frontend *http.Server
	router   *mux.Router
	logger   *logrus.Logger
	store    store.Store
}

// NewServer ...
func newServer(store store.Store, config *ConfigT) (*server, error) {

	s := &server{
		frontend: &http.Server{
			ReadTimeout:  time.Hour * 1,
			WriteTimeout: time.Hour * 1,
		},
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	// Configure the logger
	if err := s.configureLogger(config); err != nil {
		return nil, err
	}

	s.configureRouter()

	if err := s.configureFrontend(config); err != nil {
		return nil, err
	}

	return s, nil
}

// ServeHTTP ...
func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

// configureFrontend() ...
func (s *server) configureFrontend(config *ConfigT) error {
	var (
		err     error
		ownCert tls.Certificate
	)

	caCertPool := x509.NewCertPool()

	// Load the certificate pair
	ownCert, err = loadX509KeyPair(
		s.logger,
		config.Pip.SSLCert,
		config.Pip.SSLCertKey,
		"PIP",
		"")
	if err != nil {
		return err
	}

	// Read CA certs used to verify certs to be accepted
	for _, acceptedClientCert := range config.Pip.CACertsToVerifyClientRequests {
		err = loadCACertificate(
			s.logger,
			acceptedClientCert,
			"client",
			caCertPool)
		if err != nil {
			return err
		}
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{ownCert},
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		MinVersion:   tls.VersionTLS13,
		MaxVersion:   tls.VersionTLS13,
	}

	s.frontend.Addr = config.Pip.ListenAddr
	s.frontend.TLSConfig = tlsConfig
	s.frontend.Handler = s.router

	return nil
}

// The function configures the API server internal logger
func (s *server) configureLogger(config *ConfigT) error {

	// Configuration of the logger destination: stdout or a file
	f, err := logger.SetLoggerDestination(s.logger, config.SysLogger.LogDestination)
	if err != nil {
		return err
	}
	logFile := f

	// Configuration of the logger formatter: text or json
	if err := logger.SetLoggerFormatter(s.logger, config.SysLogger.LogFormatter); err != nil {
		return err
	}

	// Configuration of the logger level: Trace, Debug, Info, Warning, Error, Fatal or Panic
	if err := logger.SetLoggerLevel(s.logger, config.SysLogger.LogLevel); err != nil {
		return err
	}

	// Configure the function, that will be called in the case of "Ctrl+C".
	// The function closes the log file and exits.
	logger.SetupCloseHandler(s.logger, logFile)

	s.logger.Debugf("logger destination is set to %s", config.SysLogger.LogDestination)
	s.logger.Debugf("logger formatter is set to %s", config.SysLogger.LogFormatter)
	s.logger.Debugf("logger level is set to %s", config.SysLogger.LogLevel)
	return nil
}
