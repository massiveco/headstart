package files

import (
	"io/ioutil"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/massiveco/headstart/config"
)

//Create files on the host system
func Create(config config.Config) {

	for k, v := range config.Files {

		createFile(k, v)
	}
}

func createFile(filename string, options config.FileOptions) {

	ownerID, groupID, err := convertFileOptionsToIDs(options)
	if err != nil {
		println("Unable to create '%s', %s", filename, err)
	}

	if options.Source != "" {
		prefix := strings.Split(options.Source, "://")
		if len(prefix) == 0 {
			println("Could not determine SourceType from ", options.Source)
		}
		sourceType := prefix[0]

		switch sourceType {
		default:
			println("Unknown source type: ", sourceType, ". Skipping!")

		}
		println("Source files not supported yet", sourceType)

		return
	}

	if options.Contents != "" {

		println("Writing to", filename)
		err := ioutil.WriteFile(filename, []byte(options.Contents), os.FileMode(options.Mode))
		if err != nil {
			println("Unable to write file: ", err)
			return
		}

		err = os.Chown(filename, ownerID, groupID)
		if err != nil {
			println("Unable to set file permissions: ", err)
			return
		}
	}
}

func convertFileOptionsToIDs(options config.FileOptions) (int, int, error) {

	ownerID, err := convertOwnerToUID(options.Owner)
	if err != nil {
		return 0, 0, err
	}
	groupID, err := convertGroupToGID(options.Group)
	if err != nil {
		return 0, 0, err
	}

	return ownerID, groupID, nil
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
