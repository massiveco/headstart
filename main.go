package main

import (
	"log"
	"os"
	"path"
	"plugin"

	"github.com/massiveco/headstart/config"
	"github.com/massiveco/headstart/handlers/certificates"
	"github.com/massiveco/headstart/handlers/files"
	"github.com/massiveco/headstart/handlers/groups"
	"github.com/massiveco/headstart/handlers/users"
)

var localConfigPath = os.Getenv("HS_LOCAL_PATH")
var providerEnv = os.Getenv("HS_PROVIDER")
var providerPath = os.Getenv("HS_PROVIDER_PATH")

func init() {
	if providerEnv == "" {
		providerEnv = "local"
	}

	if providerPath == "" {
		providerPath = "/usr/lib/headstart/providers/"
	}

	if localConfigPath == "" {
		localConfigPath = "/etc/headstart/config.yml"
	}
}

func main() {
	providerPlugin, err := plugin.Open(path.Join(providerPath, providerEnv+".so"))
	if err != nil {
		log.Fatal(err)
	}

	providerSym, err := providerPlugin.Lookup("FetchConfig")
	if err != nil {
		log.Fatal(err)
	}

	provider := providerSym.(func(string) ([]byte, error))
	configStr, err := provider(localConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	cfg := config.Parse(configStr)

	certificates.Process(cfg)
	groups.Process(cfg)
	users.Process(cfg)
	files.Process(cfg)
}
