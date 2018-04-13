package files

import (
	"github.com/massiveco/headstart/config"
	"github.com/massiveco/headstart/linux"
)

//Create files on the host system
func Create(config config.Config) {

	for k, v := range config.Files {

		linux.CreateFile(k, v)
	}
}
