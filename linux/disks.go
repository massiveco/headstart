package linux

import "github.com/massiveco/headstart/config"

//MountDisk a user on the host system
func MountDisk(deviceName string, options config.DiskOptions) {
	println("Creating the disk ", deviceName, options.Fs)
}
