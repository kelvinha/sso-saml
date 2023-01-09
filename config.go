package main

import "fmt"

var (
	samlCertificatePath = "./myservice.cert"
	samlPrivateKeyPath  = "./myservice.key"
	samlIDPMetadata     = "https://samltest.id/saml/idp"

	webserverPort    = 7000
	webserverRootURL = fmt.Sprintf("http://localhost:%d", webserverPort)
)
