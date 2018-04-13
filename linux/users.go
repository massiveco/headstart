package linux

import "github.com/massiveco/headstart/config"

//CreateUser a user on the host system
func CreateUser(username string, options config.UserOptions) {
	println("Creating the user", username)
}
