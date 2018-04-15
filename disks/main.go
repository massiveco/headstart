package disks

import (
	"github.com/massiveco/headstart/config"
)

//Mount disks on the host system
func Mount(config config.Config) {

	for k, v := range config.Disks {

		mountDisk(k, v)
	}
}

func mountDisk(deviceName string, options config.DiskOptions) {
	println("Creating the disk ", deviceName, options.Fs)
}
