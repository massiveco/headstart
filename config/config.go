package config

import (
	"bytes"
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
func Parse(configStr []byte) Config {

	preamble := string(configStr[0:11])
	if preamble != "#!headstart" {
		log.Fatal("Config file does not appear to be a headstart config. Giving up")
	}
	data := buildTemplateVars()
	tmpl, err := template.New("config").Parse(string(configStr[:]))
	if err != nil {
		log.Fatal("Unable to apply template.  Giving up.")
	}
	var config Config
	rendered := new(bytes.Buffer)
	err = tmpl.Execute(rendered, data)

	err = yaml.Unmarshal(rendered.Bytes(), &config)
	if err != nil {
		log.Fatal("Unable to parse config")
	}

	return config
}

func buildTemplateVars() templateVars {

	hostname, err := os.Hostname()
	if err != nil {
		log.Println("Unable to retrieve hostname")
		return templateVars{}
	}
	return templateVars{
		Hostname: hostname,
	}
}
