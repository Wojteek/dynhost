package main

import (
	"github.com/Wojteek/dynhost/internal/app/cmd"
	"github.com/urfave/cli/v2"
	"log"
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
		Version: version,
		Commands: []*cli.Command{
			cmd.OVHCommand(),
			cmd.CloudflareCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
