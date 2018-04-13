package aws

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var configURL = "https://gist.githubusercontent.com/deedubs/64ca227f953a93ef874caf60f5e170e4/raw/4c1dabca69e603747c4397e3b6d29ea684df9ef5/user-data.yml"

//LoadUserData from the AWS METADATA API
func LoadUserData() []byte {

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Get(configURL)
	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()

	userData, err := ioutil.ReadAll(resp.Body)

	return userData
}
