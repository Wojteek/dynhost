package app

import (
	"log"
	"time"
)

// Timer creates the NewTicker and runs callback function in an interval
func Timer(timer time.Duration, fn func() error) {
	log.Printf("Set auto-refreshing: %s", timer)

	ticker := time.NewTicker(timer)
	defer ticker.Stop()

	done := make(chan bool)

loop:
	for {
		_ = fn()

		select {
		case <-done:
			log.Println("Finished DynHost!")
			break loop
		case <-ticker.C:
			continue
		}
	}
}
