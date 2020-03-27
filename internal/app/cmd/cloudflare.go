package cmd

import (
	"errors"

	"github.com/Wojteek/dynhost/internal/app"
	"github.com/Wojteek/dynhost/internal/app/provider"
	"github.com/urfave/cli/v2"
)

// CloudflareCommand the definition of the command
func CloudflareCommand() *cli.Command {
	return &cli.Command{
		Name:  "cloudflare",
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
			&cli.StringFlag{
				Name:  "data, d",
				Usage: "Set a path of the JSON file with the data of an application",
				Value: "data.json",
			},
			&cli.DurationFlag{
				Name:     "timer, t",
				Usage:    "Set the interval between automatic checking of an external IP address",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			authToken := ctx.String("auth-token")
			zoneID := ctx.String("zone-id")
			dnsID := ctx.String("dns-id")
			hostname := ctx.String("hostname")
			dataPath := ctx.String("data")

			if len(zoneID) == 0 {
				return errors.New("zoneID is required")
			}

			if len(dnsID) == 0 {
				return errors.New("dnsID is required")
			}

			if len(authToken) == 0 {
				return errors.New("authToken is required")
			}

			if len(hostname) == 0 {
				return errors.New("hostname is required")
			}

			if len(dataPath) == 0 {
				return errors.New("data is required")
			}

			timer := ctx.Duration("timer")
			processCommand := &app.ProcessCommand{
				DataPath: dataPath,
				Timer:    timer,
			}

			var changedIPCallback app.ChangedIPCallback = func(currentIP string) error {
				c := provider.NewCloudflare(
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

			if timer == 0 {
				_ = updateIP()
			} else {
				app.Timer(timer, updateIP)
			}

			return nil
		},
	}
}
