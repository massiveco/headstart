package users

import (
	"github.com/massiveco/headstart/config"
)

//Create users on the host system
func Create(config config.Config) {

	for k, v := range config.Users {

		createUser(k, v)
	}
}

func createUser(userName string, option config.UserOptions) {
	println("Creating user ", userName)
}
