package scripts

import (
	"github.com/massiveco/headstart/config"
)

//Run files on the host system
func Run(config config.Config) {

	for _, script := range config.Scripts {

		run(script)
	}
}

func run(cmd string) {
	println("Running '", cmd, "'")
}
