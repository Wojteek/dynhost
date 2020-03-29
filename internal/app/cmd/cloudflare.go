package cmd

import (
	"errors"
	"github.com/Wojteek/dynhost/internal/app"
	"github.com/Wojteek/dynhost/internal/app/service"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// CloudflareCommand the definition of the command
func CloudflareCommand() *cli.Command {
	return &cli.Command{
		Name: "cloudflare",

		Usage: "Cloudflare provider",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "auth-token",
				Usage:    "The authentication token of the Cloudflare API",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "zone-id",
				Usage:    "The zone identifier (Cloudflare DNS)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "dns-id",
				Usage:    "The dns identifier (Cloudflare DNS)",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "hostname",
				Usage:    "The hostname (Cloudflare DNS)",
				Required: true,
			},
		},

		Before: func(ctx *cli.Context) error {
			if zoneID := ctx.String("zone-id"); len(zoneID) == 0 {
				return errors.New("zoneID is required")
			}

			if dnsID := ctx.String("dns-id"); len(dnsID) == 0 {
				return errors.New("dnsID is required")
			}

			if authToken := ctx.String("auth-token"); len(authToken) == 0 {
				return errors.New("authToken is required")
			}

			if hostname := ctx.String("hostname"); len(hostname) == 0 {
				return errors.New("hostname is required")
			}

			if dataPath := ctx.String("data"); len(dataPath) == 0 {
				return errors.New("data is required")
			}

			return nil
		},

		Action: func(ctx *cli.Context) error {
			authToken := ctx.String("auth-token")
			zoneID := ctx.String("zone-id")
			dnsID := ctx.String("dns-id")
			hostname := ctx.String("hostname")
			dataPath := ctx.String("data")
			timer := ctx.Duration("timer")

			processCommand := &app.ProcessCommand{
				DataPath: dataPath,
				Timer:    timer,
			}

			var changedIPCallback app.ChangedIPCallback = func(currentIP string) error {
				c := service.NewCloudflare(
					authToken,
					hostname,
					currentIP,
					zoneID,
					dnsID,
				)

				if _, err := c.UpdateRecordRequest(); err != nil {
					return err
				}

				return nil
			}

			var updateIP = app.UpdateIP(processCommand, changedIPCallback)

			log.WithFields(log.Fields{
				"service": "cloudflare",
				"timer":   timer,
			}).Info("The DynHost is running")

			if timer == 0 {
				_ = updateIP()
			} else {
				app.Timer(timer, updateIP)
			}

			return nil
		},
	}
}
