package ip

import (
	"log"
)

// Providers - the structure of the providers for fetching the external IP address
type Providers struct {
	url  string
	name string
}

// ExternalIP - main structure of the ExternalIP
type ExternalIP struct {
	providers []Providers
}

// NewExternalIP creates the structure of the ExternalIP
func NewExternalIP() *ExternalIP {
	return &ExternalIP{
		providers: []Providers{
			{
				name: "IPify",
				url:  "https://api.ipify.org"},
			{
				name: "AmazonAWS",
				url:  "https://checkip.amazonaws.com",
			},
			{
				name: "MyExternalIP",
				url:  "https://myexternalip.com/raw",
			},
		},
	}
}

// IP gets the external IP address
func (e *ExternalIP) IP() []byte {
	for _, provider := range e.providers {
		log.Printf("Requesting to: %s", provider.url)

		currentIP, err := request(provider.url)

		if err != nil {
			log.Printf("Error %s@provider: %s", provider.name, err)

			continue
		}

		return currentIP
	}

	return nil
}
