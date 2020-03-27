package app

import (
	"github.com/Wojteek/dynhost/internal/app/ip"
	"log"
	"strings"
	"time"
)

// UpdateIP fetches and updates the IP when the IP was changed
func UpdateIP(p *ProcessCommand, ChangedIPFn ChangedIPCallback) func() error {
	return func() error {
		log.Println("Checking an external IP address...")

		data, _ := GetData(p.DataPath)
		externalIP := strings.TrimSpace(string(ip.NewExternalIP().IP()))

		if externalIP == data.LastIP {
			return nil
		}

		if err := ChangedIPFn(externalIP); err != nil {
			log.Fatal(err)
		}

		var d interface{} = &Data{
			LastIP:    externalIP,
			PrevIP:    data.LastIP,
			ChangedAt: time.Now(),
		}

		if _, err := d.(*Data).SaveData(p.DataPath); err != nil {
			return err
		}

		log.Printf("The IP address has been changed: %s - before: %s", externalIP, data.LastIP)

		return nil
	}
}
