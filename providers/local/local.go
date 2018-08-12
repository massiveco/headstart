package main

import (
	"io/ioutil"
)

func main() {}

// FetchConfig from this provider
func FetchConfig(filename string) ([]byte, error) {
	userData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return userData, nil
}
