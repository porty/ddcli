package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/porty/ddcli/datadog"
	"github.com/urfave/cli"
)

func export(c *cli.Context) error {

	if c.NArg() != 1 {
		fmt.Println("Output directory required")
		os.Exit(1)
	}

	outputDir := c.Args()[0]

	apiKey := os.Getenv("DD_API_KEY")
	appKey := os.Getenv("DD_APP_KEY")
	if apiKey == "" || appKey == "" {
		fmt.Println("DD_API_KEY and DD_APP_KEY required")
		os.Exit(1)
	}

	dashboardDir := path.Join(outputDir, "dashboards")
	screenboardDir := path.Join(outputDir, "screenboards")
	monitorsDir := path.Join(outputDir, "monitors")
	createDirectories(dashboardDir, screenboardDir, monitorsDir)
	dd := datadog.New(apiKey, appKey)

	dashes, err := dd.GetDashboards()
	if err != nil {
		panic(err)
	}

	if len(dashes) == 0 {
		log.Println("No dashboards")
	} else {
		for i, info := range dashes {
			log.Printf("Getting dashboard %d of %d...", i+1, len(dashes))
			dash, err := dd.GetDashboard(info.ID)
			if err != nil {
				log.Printf("Failed to get dashboard #%s: %s", info.ID, err.Error())
				os.Exit(1)
			}
			dest := path.Join(dashboardDir, info.ID+".json")
			b, err := json.MarshalIndent(dash, "", "  ")
			if err != nil {
				log.Print("Failed to JSON marshal dashboard: " + err.Error())
				os.Exit(1)
			}
			if b[len(b)-1] != '\n' {
				b = append(b, '\n')
			}
			if err = ioutil.WriteFile(dest, b, 0664); err != nil {
				log.Printf("Failed to write to file '%s': %s", dest, err.Error())
				os.Exit(1)
			}
		}
	}

	screenboards, err := dd.GetScreenboards()
	if err != nil {
		panic(err)
	}

	if len(screenboards) == 0 {
		log.Print("No screenboards")
	} else {
		for i, info := range screenboards {
			log.Printf("Getting screenboard %d of %d...", i+1, len(screenboards))
			screenboard, err := dd.GetScreenboard(info.ID)
			if err != nil {
				log.Printf("Failed to get screenboard #%d: %s", info.ID, err.Error())
				os.Exit(1)
			}
			dest := path.Join(screenboardDir, fmt.Sprintf("%d.json", info.ID))
			b, err := json.MarshalIndent(screenboard, "", "  ")
			if err != nil {
				log.Print("Failed to JSON marshal screenboard: " + err.Error())
				os.Exit(1)
			}
			if b[len(b)-1] != '\n' {
				b = append(b, '\n')
			}
			if err = ioutil.WriteFile(dest, b, 0664); err != nil {
				log.Printf("Failed to write to file '%s': %s", dest, err.Error())
				os.Exit(1)
			}
		}
	}

	monitors, err := dd.GetMonitors()
	if err != nil {
		panic(err)
	}
	if len(monitors) == 0 {
		log.Print("No monitors")
	} else {
		for _, monitor := range monitors {
			dest := path.Join(monitorsDir, fmt.Sprintf("%d.json", monitor.ID))
			b, err := json.MarshalIndent(monitor, "", "  ")
			if err != nil {
				log.Print("Failed to JSON marshal monitor: " + err.Error())
				os.Exit(1)
			}
			if b[len(b)-1] != '\n' {
				b = append(b, '\n')
			}
			if err = ioutil.WriteFile(dest, b, 0664); err != nil {
				log.Printf("Failed to write to file '%s': %s", dest, err.Error())
				os.Exit(1)
			}
		}
		log.Printf("Exported %d monitors", len(monitors))
	}
	return nil
}

func createDirectories(dirs ...string) {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0777); err != nil {
			fmt.Printf("Failed to create output directory '%s': %s\n", dir, err.Error())
			os.Exit(1)
		}
	}
}
