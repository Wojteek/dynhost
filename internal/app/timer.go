package app

import (
	log "github.com/sirupsen/logrus"
	"time"
)

// Timer creates the NewTicker and runs callback function in an interval
func Timer(timer time.Duration, fn func() error) {
	ticker := time.NewTicker(timer)
	defer ticker.Stop()

	done := make(chan bool)

loop:
	for {
		_ = fn()

		select {
		case <-done:
			log.Info("Exit the DynHost")
			break loop
		case <-ticker.C:
			continue
		}
	}
}
