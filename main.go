package main

import (
	"log"
	"os"
	"plugin"

	"github.com/massiveco/headstart/config"
	"github.com/massiveco/headstart/handlers/files"
	"github.com/massiveco/headstart/handlers/users"
)

var providerEnv = os.Getenv("HS_PROVIDER")

func main() {

	if providerEnv == "" {
		providerEnv = "local"
	}

	providerPlugin, err := plugin.Open("providers/" + providerEnv + ".so")
	if err != nil {
		log.Fatal(err)
	}

	providerSym, err := providerPlugin.Lookup("FetchConfig")
	if err != nil {
		log.Fatal(err)
	}

	provider := providerSym.(func() ([]byte, error))
	configStr, err := provider()
	if err != nil {
		log.Fatal(err)
	}
	cfg := config.Parse(configStr)

	users.Process(cfg)
	files.Process(cfg)
}
