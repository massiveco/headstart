package users

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	"github.com/massiveco/headstart/config"
)

const addUserCommandPath = "/usr/sbin/adduser"
const userModCommandPath = "/usr/sbin/usermod"
const sshKeyDirFormat = "/home/%s/.ssh"

//Process users on the host system
func Process(config config.Config) {

	for username, options := range config.Users {

		args := []string{"-D", username}
		addUserCmd := exec.Command(addUserCommandPath, args...)
		err := addUserCmd.Run()
		if err != nil {
			log.Fatal(err)

		}

		usr, err := user.Lookup(username)
		if err != nil {
			log.Fatal(err)

		}
		grp, err := user.LookupGroup(usr.Gid)
		if err != nil {
			log.Fatal(err)

		}
		usrID, err := strconv.Atoi(usr.Uid)
		if err != nil {
			log.Fatal(err)

		}
		grpID, err := strconv.Atoi(grp.Gid)
		if err != nil {
			log.Fatal(err)

		}

		if len(options.AuthorizedKeys) != 0 {
			keyDir := fmt.Sprintf(sshKeyDirFormat, username)
			err := os.Mkdir(keyDir, os.FileMode(0700))
			if err != nil {
				log.Fatal(err)
			}
			err = os.Chown(keyDir, usrID, grpID)
			if err != nil {
				log.Fatal(err)
			}

			authorizedKeyPath := strings.Join([]string{keyDir, "authorized_keys"}, "/")
			authorizedKeys := strings.Join(options.AuthorizedKeys, "\n")

			err = ioutil.WriteFile(authorizedKeyPath, []byte(authorizedKeys), os.FileMode(0600))
			if err != nil {
				log.Fatal(err)
			}
			err = os.Chown(authorizedKeyPath, usrID, grpID)
			if err != nil {
				log.Fatal(err)
			}
		}

		if len(options.Groups) != 0 {
			groups := strings.Join(options.Groups, ",")

			cmd := exec.Command(userModCommandPath, []string{"-a", "-G", groups, username}...)

			err := cmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func createUser(userName string, option config.User) {
	println("Creating user ", userName, option.AuthorizedKeys)
}
