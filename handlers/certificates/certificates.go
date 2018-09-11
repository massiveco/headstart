package certificates

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"

	"github.com/massiveco/headstart/config"
	"github.com/massiveco/serverlessl/client"
)

//Process requested PKI requests
func Process(cfg config.Config) {

	for _, certDetails := range cfg.Certificates {

		serverlesslSvc := client.New(client.Config{
			Name: certDetails.Name,
			Lambda: client.LambdaConfig{
				Region: certDetails.Region,
			},
		})
		caCert, err := serverlesslSvc.FetchCa()
		if err != nil {
			fmt.Println("Unable to fetch Ca ", certDetails.Name)
			return
		}

		csr, key, crt, err := serverlesslSvc.RequestCertificate(client.CertificateDetails{
			CommonName: certDetails.Details.CommonName,
			Group:      certDetails.Details.Group,
			Hosts:      certDetails.Details.Hosts,
		})
		if err != nil {
			fmt.Println("Unable to fetch certificate ", certDetails.Name)
			return
		}

		cert := pem.Block{
			Type:  "CERTIFICATE",
			Bytes: crt,
		}

		if certDetails.Paths.CertificateAuthority != "" {
			ioutil.WriteFile(certDetails.Paths.CertificateAuthority, caCert, 0600)
		}

		if certDetails.Paths.SigningRequest != "" {
			ioutil.WriteFile(certDetails.Paths.SigningRequest, csr, 0600)
		}
		if certDetails.Paths.Key != "" {
			ioutil.WriteFile(certDetails.Paths.Key, key, 0600)
		}
		if certDetails.Paths.Certificate != "" {
			ioutil.WriteFile(certDetails.Paths.Certificate, pem.EncodeToMemory(&cert), 0600)
		}
	}
}
