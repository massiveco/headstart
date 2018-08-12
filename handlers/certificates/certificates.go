package certificates

import (
	"fmt"

	"github.com/massiveco/headstart/config"
)

//Process requested PKI requests
func Process(cfg config.Config) {

	for k, v := range cfg.Certificates {

		fmt.Println(k, v)
	}
}
