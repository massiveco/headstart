package config

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Group to be created on the host
type Group struct {
	Sudo bool `yaml:"Sudo,omitempty"`
}

// User users to be created on the host
type User struct {
	Groups         []string `yaml:"groups,omitempty"`
	AuthorizedKeys []string `yaml:"authorized_keys"`
}

// File for creating files on the host
type File struct {
	Source          string      `yaml:"source,omitempty"`
	Contents        string      `yaml:"contents,omitempty"`
	EncodedContents string      `yaml:"encoded_contents,omitempty"`
	Hash            string      `yaml:"hash,omitempty"`
	Mode            os.FileMode `yaml:"mode,omitempty"`
	Owner           string      `yaml:"owner,omitempty"`
	Group           string      `yaml:"group,omitempty"`
}

// Certificate for creating/requesting PKI certificates
type Certificate struct {
	Type    string             `yaml:"type,omitempty"`
	Region  string             `yaml:"region,omitempty"`
	Name    string             `yaml:"name,omitempty"`
	Profile string             `yaml:"profile,omitempty"`
	Paths   CertificatePaths   `yaml:"paths,omitempty"`
	Details CertificateDetails `yaml:"details,omitempty"`
}

// CertificatePaths defines the paths we should write the certificate components to
type CertificatePaths struct {
	Certificate          string `yaml:"cert,omitempty"`
	SigningRequest       string `yaml:"csr,omitempty"`
	Key                  string `yaml:"key,omitempty"`
	CertificateAuthority string `yaml:"ca,omitempty"`
}

// CertificateDetails defines the Details we should write the certificate components to
type CertificateDetails struct {
	Group      string            `yaml:"group,omitempty"`
	CommonName string            `yaml:"commonName,omitempty"`
	Hosts      []string          `yaml:"hosts,omitempty"`
	Type       string            `yaml:"type,omitempty"`
	Region     string            `yaml:"region,omitempty"`
	Name       string            `yaml:"name,omitempty"`
	Paths      map[string]string `yaml:"paths,omitempty"`
}

//Config Headstart config
type Config struct {
	Users        map[string]User  `yaml:"users"`
	Files        map[string]File  `yaml:"files"`
	Groups       map[string]Group `yaml:"groups"`
	Certificates []Certificate    `yaml:"certificates"`
}
type templateVars struct {
	Hostname string
}

// Parse a string into a config struct
func Parse(configBytes []byte) (*Config, error) {

	if len(configBytes) == 0 {
		return nil, errors.New("No config provided")
	}

	preamble := string(configBytes[0:11])
	if preamble != "#!headstart" {

		return nil, errors.New("Config file does not appear to be a headstart config" + preamble)
	}
	configTemplate, err := template.New("config").Parse(string(configBytes[:]))
	if err != nil {
		return nil, errors.New("Unable to apply template")
	}
	var config Config
	rendered := new(bytes.Buffer)
	err = configTemplate.Execute(rendered, buildTemplateData())
	log.Println(rendered)
	err = yaml.Unmarshal(rendered.Bytes(), &config)
	if err != nil {
		return nil, errors.New("Unable to parse config")
	}

	return &config, nil
}

func buildTemplateData() templateVars {

	hostname, err := os.Hostname()
	if err != nil {
		log.Println("Unable to retrieve hostname")
		return templateVars{}
	}
	return templateVars{
		Hostname: hostname,
	}
}
