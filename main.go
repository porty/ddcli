package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:   "export",
			Usage:  "export Datadog config",
			Action: export,
		},
	}

	app.Run(os.Args)
}
