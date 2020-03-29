package app

import (
	log "github.com/sirupsen/logrus"
	"time"
)

// IPChangedCallback is the structure of callback function of changing IP
type IPChangedCallback func(currentIP string) error

// ServiceCommand is the structure of main options of command
type ServiceCommand struct {
	DataPath string
	Timer    time.Duration
}

// Execute is the method for running the command
func (s *ServiceCommand) Execute(service string, callback IPChangedCallback) {
	var updateIP = UpdateIP(s, callback)

	log.WithFields(log.Fields{
		"service": service,
		"timer":   s.Timer,
	}).Info("The DynHost is running")

	if s.Timer == 0 {
		if err := updateIP(); err != nil {
			log.Fatal(err)
		}
	} else {
		Timer(s.Timer, updateIP)
	}
}
