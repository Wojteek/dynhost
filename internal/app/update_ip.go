package app

import (
	"github.com/Wojteek/dynhost/internal/app/ip"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

// UpdateIP fetches and updates the IP when the IP was changed
func UpdateIP(p *ServiceCommand, IPChangedFn IPChangedCallback) func() error {
	return func() error {
		data, _ := GetData(p.DataPath)
		externalIP := strings.TrimSpace(string(ip.NewExternalIP().IP()))

		if len(externalIP) == 0 {
			return nil
		}

		if externalIP == data.CurrentIP {
			return nil
		}

		if err := IPChangedFn(externalIP); err != nil {
			return err
		}

		var d interface{} = &Data{
			CurrentIP: externalIP,
			PrevIP:    data.CurrentIP,
			ChangedAt: time.Now(),
		}

		if _, err := d.(*Data).SaveData(p.DataPath); err != nil {
			return err
		}

		log.WithFields(log.Fields{
			"current_ip":  externalIP,
			"previous_ip": data.CurrentIP,
		}).Info("The IP address has been changed")

		return nil
	}
}
