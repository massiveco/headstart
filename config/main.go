package config

import (
	"log"

	"github.com/massiveco/headstart/aws"
	"gopkg.in/yaml.v2"
)

// User users to be created on the host
type UserOptions struct {
	Name           string   `yaml:"name,omitempty"`
	AuthorizedKeys []string `yaml:"authorized_keys"`
}

// DiskOptions for creating and mounting disks
type DiskOptions struct {
	Fs      string `yaml:"fs"`
	Options string `yaml:"options,omitempty"`
	Mount   string `yaml:"mount"`
}

// FileOptions for creating files on the host
type FileOptions struct {
	Source string `yaml:"source"`
	Hash   string `yaml:"hash"`
	Mode   string `yaml:"mode"`
	Owner  string `yaml:"owner,omitempty"`
	Group  string `yaml:"group,omitempty"`
}

//Config Headstart config
type Config struct {
	Users   map[string]UserOptions `yaml:"users"`
	Disks   map[string]DiskOptions `yaml:"disks"`
	Files   map[string]FileOptions `yaml:"files"`
	Scripts []string               `yaml:"scripts"`
}

// Parse a string into a config struct
func parse(configStr []byte) Config {

	var config Config
	err := yaml.Unmarshal(configStr, &config)
	if err != nil {
		log.Panic(err)
	}

	return config
}

//Load UserData from the AWS Metadata API
func Load(sourceType string) Config {
	var config Config
	switch sourceType {
	case "aws":
		config = parse(aws.LoadUserData())
	default:
		log.Fatal("Unknown sourceType:", sourceType)
	}

	return config
}
