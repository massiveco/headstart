package scripts

import (
	"github.com/massiveco/headstart/config"
	"github.com/massiveco/headstart/linux"
)

//Run files on the host system
func Run(config config.Config) {

	for _, script := range config.Scripts {

		linux.Run(script)
	}
}
