package users

import (
	"github.com/massiveco/headstart/config"
	"github.com/massiveco/headstart/linux"
)

//Create users on the host system
func Create(config config.Config) {

	for k, v := range config.Users {

		linux.CreateUser(k, v)
	}
}
