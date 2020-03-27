package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

const versionTemplate = `DynHost version:
  Version:      %s
  Built:        %s
  Git commit:   %s
`

// VersionPrinter returns the version of the application
func VersionPrinter(date string, commit string) func(ctx *cli.Context) {
	return func(ctx *cli.Context) {
		fmt.Printf(
			versionTemplate,
			ctx.App.Version,
			date,
			commit,
		)
	}
}
