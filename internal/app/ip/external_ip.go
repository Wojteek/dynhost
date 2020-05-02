package ip

import (
	log "github.com/sirupsen/logrus"
	"net"
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
func (e *ExternalIP) IP() string {
	for _, provider := range e.providers {
		log.WithFields(log.Fields{
			"provider": provider.name,
		}).Debug("Checking an external IP address")

		responseProvider, err := request(provider.url)

		if err != nil {
			log.WithFields(log.Fields{
				"provider": provider.name,
			}).Error(err)

			continue
		}

		currentIP := net.ParseIP(string(responseProvider))

		if currentIP == nil {
			log.WithFields(log.Fields{
				"provider": provider.name,
				"response": responseProvider,
			}).Error("The IP address has invalid format")

			continue
		}

		return currentIP.String()
	}

	return ""
}
