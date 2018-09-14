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

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to run command: "+err.Error())
		os.Exit(1)
	}
}

func getAPI() *datadog.API {
	apiKey := os.Getenv("DD_API_KEY")
	appKey := os.Getenv("DD_APP_KEY")
	if apiKey == "" || appKey == "" {
		fmt.Println("DD_API_KEY and DD_APP_KEY required")
		os.Exit(1)
	}
	return datadog.New(apiKey, appKey)
}
