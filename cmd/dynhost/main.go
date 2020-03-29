package main

import (
	"github.com/Wojteek/dynhost/internal/app/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cli.VersionPrinter = cmd.VersionPrinter(date, commit)
	app := &cli.App{
		Name: "DynHost",

		Usage: "It synchronizes the external IP address with the DNS record in the OVH or Cloudflare.",

		Version: version,

		Commands: []*cli.Command{
			cmd.OVHCommand(),
			cmd.CloudflareCommand(),
		},

		Before: func(ctx *cli.Context) error {
			if debug := ctx.Bool("debug"); debug {
				log.SetLevel(log.DebugLevel)
			}

			log.SetFormatter(&log.TextFormatter{
				FullTimestamp: true,
			})

			return nil
		},

		Flags: []cli.Flag{
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
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable the debug mode",
				Value: false,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
