package config

import (
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

//Config Headstart config
type Config struct {
	Users  map[string]User  `yaml:"users"`
	Files  map[string]File  `yaml:"files"`
	Groups map[string]Group `yaml:"groups"`
}

// Parse a string into a config struct
func Parse(configStr []byte) Config {

	preamble := string(configStr[0:11])
	if preamble != "#!headstart" {
		log.Fatal("Config file does not appear to be a headstart config. Giving up")
	}

	var config Config
	err := yaml.Unmarshal(configStr, &config)
	if err != nil {
		log.Fatal("Unable to parse config")
	}

	return config
}
