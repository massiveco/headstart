package linux

import "github.com/massiveco/headstart/config"

//CreateFile a user on the host system
func CreateFile(filename string, options config.FileOptions) {
	println("Creating the file", filename, options.Source)
}
