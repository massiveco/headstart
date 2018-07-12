package main

import (
	"io/ioutil"
)

// FetchConfig from this provider
func FetchConfig() ([]byte, error) {

	userData, err := ioutil.ReadFile("./sample_config.yml")
	if err != nil {
		return nil, err
	}

	return userData, nil
}
