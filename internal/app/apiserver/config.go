package apiserver

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type SysLoggerT struct {
	LogLevel       string `yaml:"level"`
	LogDestination string `yaml:"destination"`
	LogFormatter   string `yaml:"formatter"`
}

// The struct PipT is for parsing the section 'pip' of the config file.
type PipT struct {
	ListenAddr                    string   `yaml:"listen_addr"`
	SSLCert                       string   `yaml:"ssl_cert"`
	SSLCertKey                    string   `yaml:"ssl_cert_key"`
	CACertsToVerifyClientRequests []string `yaml:"ca_certs_to_verify_client_certs"`
}

// The struct PipT is for parsing the section 'pip' of the config file.
type DataBaseT struct {
	URL string `yaml:"url"`
}

// ConfigT struct is for parsing the basic structure of the config file
type ConfigT struct {
	SysLogger                        SysLoggerT `yaml:"system_logger"`
	Pip                              PipT       `yaml:"pip"`
	DataBase                         DataBaseT  `yaml:"database"`
	UseDBCache                       bool
	CACertPoolToVerifyClientRequests *x509.CertPool
	PipCert                          tls.Certificate
}

// Config contains all input from the config file and is is globally accessible
// var Config ConfigT

// The function returns a config with default values
func NewConfig() *ConfigT {
	return &ConfigT{
		SysLogger: SysLoggerT{
			LogLevel:       "info",
			LogDestination: "stdout",
			LogFormatter:   "text",
		},
		Pip: PipT{
			ListenAddr:                    "",
			SSLCert:                       "",
			SSLCertKey:                    "",
			CACertsToVerifyClientRequests: []string{},
		},
		DataBase: DataBaseT{
			URL: "postgres://host:5432/database?sslmode=disable&user=username&password=Pa$$w0rd",
		},
		UseDBCache:                       false,
		CACertPoolToVerifyClientRequests: x509.NewCertPool(),
		PipCert:                          tls.Certificate{},
	}
}

// The function parses a configuration yaml file and overwrites corresponding fields in the given config variable
// Note! The rest of the original config struct fields keep their values, if they are not presented in the config file
func UpdateConfigFromFile(configPath string, config *ConfigT) error {

	// If the config file path was not provided
	if configPath == "" {
		return errors.New("apiserver: UpdateConfigFromFile(): no configuration file is provided")
	}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("apiserver: UpdateConfigFromFile(): unable to open the YAML configuration file '%s': %s", configPath, err.Error())
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Decode configuration from the YAML config file
	err = d.Decode(&config)
	if err != nil {
		return fmt.Errorf("apiserver: UpdateConfigFromFile(): unable to decode the YAML configuration file '%s': %s", configPath, err.Error())
	}

	return nil
}

// The function checks, whether all mandatory fiekds of the configuration have values
func CheckConfig(config *ConfigT) error {
	var fields string = ""

	if config.Pip.ListenAddr == "" {
		fields += "listen_addr,"
	}

	if config.Pip.SSLCert == "" {
		fields += "ssl_cert,"
	}

	if config.Pip.SSLCertKey == "" {
		fields += "ssl_cert_key,"
	}

	if len(config.Pip.CACertsToVerifyClientRequests) == 0 {
		fields += "ca_certs_to_verify_client_certs,"
	}

	if fields != "" {
		return fmt.Errorf("apiserver: CheckConfig(): in the section 'pip' the following required fields are missed: '%s'", strings.TrimSuffix(fields, ","))
	}

	return nil
}
