package certificates

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"github.com/massiveco/headstart/config"
	"github.com/massiveco/serverlessl/client"
)

// Process requested PKI requests
func Process(cfg config.Config) {

	for _, certificate := range cfg.Certificates {
		serverlesslSvc := client.New(client.Config{
			Name: certificate.Name,
			Lambda: client.LambdaConfig{
				Region: certificate.Region,
			},
			Profile: certificate.Profile,	
		})
		caCert, err := serverlesslSvc.FetchCa()
		if err != nil {
			fmt.Println("Unable to fetch Ca ", certificate.Name)
			fmt.Println(err)
			return
		}

		csr, key, crt, err := serverlesslSvc.RequestCertificate(client.CertificateDetails{
			CommonName: certificate.Details.CommonName,
			Group:      certificate.Details.Group,
			Hosts:      certificate.Details.Hosts,
		})
		if err != nil {
			fmt.Println("Unable to fetch certificate ", certificate.Name, err)
			return
		}

		cert := pem.Block{
			Type:  "CERTIFICATE",
			Bytes: crt,
		}

		if certificate.Paths.CertificateAuthority != "" {
			ioutil.WriteFile(certificate.Paths.CertificateAuthority, caCert, 0600)
		}
		if certificate.Paths.SigningRequest != "" {
			ioutil.WriteFile(certificate.Paths.SigningRequest, csr, 0600)
		}
		if certificate.Paths.Key != "" {
			ioutil.WriteFile(certificate.Paths.Key, key, 0600)
		}
		if certificate.Paths.Certificate != "" {
			ioutil.WriteFile(certificate.Paths.Certificate, pem.EncodeToMemory(&cert), 0600)
		}
	}
}
