package files

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"

	"github.com/massiveco/headstart/config"
)

//Process files on the host system
func Process(cfg config.Config) {

	for k, v := range cfg.Files {

		createFile(k, v)
	}
}

func createFile(filename string, options config.File) {

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
		case "https":

			contents, _ := fetchFile(options.Source)

			//TODO: Check file sha256 against Options.Hash
			options.Contents = string(contents)

		default:
			println("Unknown source type: ", sourceType, ". Skipping!")
			return
		}
	}

	if options.EncodedContents != "" {
		println("Decoding", filename)
		data, err := base64.StdEncoding.DecodeString(options.EncodedContents)
		if err != nil {
			println("Unable to decode file: ", filename, err)

			return
		}
		options.Contents = string(data[:])
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

func convertFileOptionsToIDs(cfg config.File) (int, int, error) {

	ownerID, err := convertOwnerToUID(cfg.Owner)
	if err != nil {
		return 0, 0, err
	}
	groupID, err := convertGroupToGID(cfg.Group)
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

func fetchFile(url string) ([]byte, error) {
	c := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := c.Get(url)
	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func compareFileToHash(file []byte, hash []byte) bool {

	h := sha256.New()
	_, err := h.Write(file)
	if err != nil {
		println("Unable to hash")
		return false
	}
	filesha := h.Sum(nil)

	fmt.Printf("%x", filesha)
	fmt.Printf("%x", hash)
	println()

	return bytes.Equal(hash, filesha)
}
