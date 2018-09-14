package main

import (
	"fmt"
	"os"

	"github.com/porty/ddcli/datadog"
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
		{
			Name:  "metrics",
			Usage: "metrics commands",
			Subcommands: []cli.Command{
				{
					Name:   "active",
					Usage:  "list active metrics from the specified duration",
					Action: activeMetricsFromDuration,
					Flags: []cli.Flag{
						cli.DurationFlag{
							Name:  "duration, d",
							Usage: "Duration, e.g. 1hr, 2h45m",
						},
					},
				},
				{
					Name:   "top500",
					Usage:  "list top 500 custom metrics for the month",
					Action: top500CustomMetrics,
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "format, f",
							Value: "csv",
							Usage: "Format, either csv or md (markdown)",
						},
					},
				},
			},
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
