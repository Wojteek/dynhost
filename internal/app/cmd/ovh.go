package cmd

import (
	"errors"
	"github.com/Wojteek/dynhost/internal/app"
	"github.com/Wojteek/dynhost/internal/app/provider"
	"github.com/urfave/cli/v2"
)

// OVHCommand the definition of the command
func OVHCommand() *cli.Command {
	return &cli.Command{
		Name:  "ovh",
		Usage: "OVH provider",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "auth-username",
				Usage:    "The authentication username of the DynHost option",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "auth-password",
				Usage:    "The authentication password of the DynHost option",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "hostname",
				Usage:    "The hostname of the DynHost option",
				Required: true,
			},
			&cli.StringFlag{
				Name:  "data, d",
				Usage: "Set the path of the JSON file with the data of an application",
				Value: "data.json",
			},
			&cli.DurationFlag{
				Name:     "timer, t",
				Usage:    "Set the interval between automatic checking of an external IP address",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			authUsername := ctx.String("auth-username")
			authPassword := ctx.String("auth-password")
			hostname := ctx.String("hostname")
			dataPath := ctx.String("data")

			if len(authUsername) == 0 {
				return errors.New("authUsername is required")
			}

			if len(authPassword) == 0 {
				return errors.New("authPassword is required")
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
				r := provider.OVH{
					IP:       currentIP,
					Hostname: hostname,
					Credentials: provider.CredentialsOVH{
						Username: authUsername,
						Password: authPassword,
					},
				}

				if _, err := r.UpdateRecordRequest(); err != nil {
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
