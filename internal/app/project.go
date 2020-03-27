package app

import (
	"time"
)

// ChangedIPCallback is the structure of callback function of changing IP
type ChangedIPCallback func(currentIP string) error

// ProcessCommand is the structure of main options of command
type ProcessCommand struct {
	DataPath string
	Timer    time.Duration
}
