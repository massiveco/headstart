package disks

import (
	"github.com/massiveco/headstart/config"
	"github.com/massiveco/headstart/linux"
)

//Create files on the host system
func Mount(config config.Config) {

	for k, v := range config.Disks {

		linux.MountDisk(k, v)
	}
}
