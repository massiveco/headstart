package users

import (
	"github.com/massiveco/headstart/config"
)

//Process users on the host system
func Process(config config.Config) {

	for k, v := range config.Users {

		createUser(k, v)
	}
}

func createUser(userName string, option config.User) {
	println("Creating user ", userName, option.Name)
}
