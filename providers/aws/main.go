package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

var metadataURL = "http://169.254.169.254/latest/user-data"

func main() {}

func loadMetadata() ([]byte, error) {

	c := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := c.Get(metadataURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	userData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

// FetchConfig from this provider
func FetchConfig(_ string) ([]byte, error) {
	return loadMetadata()
}
