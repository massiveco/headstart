package main

import (
	"github.com/massiveco/headstart/config"
	"github.com/massiveco/headstart/disks"
	"github.com/massiveco/headstart/files"
	"github.com/massiveco/headstart/scripts"
	"github.com/massiveco/headstart/users"
)

var cloudType = "aws"

func main() {

	headstartConfig := config.Load(cloudType)

	users.Create(headstartConfig)
	files.Create(headstartConfig)
	disks.Mount(headstartConfig)
	scripts.Run(headstartConfig)
}
