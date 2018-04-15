package files

import (
	"io/ioutil"
	"os"
	"os/user"
	"strconv"

	"github.com/cloudflare/cfssl/log"
	"github.com/massiveco/headstart/config"
)

//Create files on the host system
func Create(config config.Config) {

	for k, v := range config.Files {

		createFile(k, v)
	}
}

func createFile(filename string, options config.FileOptions) {

	ownerID, err := convertOwnerToUID(options.Owner)
	if err != nil {
		log.Warning("Unable to determine file owner. Skipping creation of", filename)
		return
	}
	groupID, err := convertGroupToGID(options.Group)
	if err != nil {

		log.Warning("Unable to determine file group. Skipping creation of", filename, err)
		return
	}
	if options.Source != "" {

		log.Warning("Source files not supported yet")

		return
	}

	if options.Contents != "" {

		err := ioutil.WriteFile(filename, []byte(options.Contents), os.FileMode(options.Mode))
		if err != nil {
			log.Fatal(err)
		}

		err = os.Chown(filename, ownerID, groupID)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func convertOwnerToUID(owner string) (int, error) {
	var fileowner *user.User
	var err error

	if owner != "" {

		fileowner, err = user.Lookup(owner)
		if err != nil {
			return 0, err
		}
	} else {
		fileowner, err = user.Current()
		if err != nil {
			return 0, err
		}
	}
	userID, _ := strconv.Atoi(fileowner.Uid)
	return userID, nil
}

func convertGroupToGID(group string) (int, error) {
	var filegroup *user.Group
	var err error

	if group != "" {

		filegroup, err = user.LookupGroup(group)
		if err != nil {
			return 0, err
		}
	} else {
		currentuser, err := user.Current()
		filegroup, err = user.LookupGroupId(currentuser.Gid)
		if err != nil {
			return 0, err
		}
	}
	groupID, _ := strconv.Atoi(filegroup.Gid)
	return groupID, nil
}
