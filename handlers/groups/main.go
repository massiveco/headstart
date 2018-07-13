package groups

import (
	"github.com/massiveco/headstart/config"
)

//Process groups on the host system
func Process(config config.Config) {

	for k, v := range config.Groups {

		createGroups(k, v)
	}
}

func createGroups(groupName string, option config.Group) {
	println("Creating group ", groupName, option.Sudo)
}
