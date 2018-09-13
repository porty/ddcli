package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"context"

	"github.com/porty/ddcli/datadog"
	"github.com/urfave/cli"
	"golang.org/x/sync/errgroup"
)

type downloader interface {
	Download(api *datadog.API, outputDir string) error
}

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

	// dashboardDir := path.Join(outputDir, "dashboards")
	//screenboardDir := path.Join(outputDir, "screenboards")
	monitorsDir := path.Join(outputDir, "monitors")
	createDirectories(path.Join(outputDir, "dashboards"), path.Join(outputDir, "screenboards"), monitorsDir)
	dd := datadog.New(apiKey, appKey)

	dashes, err := dd.GetDashboards()
	if err != nil {
		panic(err)
	}

	screenboards, err := dd.GetScreenboards()
	if err != nil {
		panic(err)
	}

	work := make(chan (downloader))

	g, _ := errgroup.WithContext(context.Background())
	for i := 0; i < 5; i++ {
		g.Go(func() error {
			for d := range work {
				if err := d.Download(dd, outputDir); err != nil {
					return err
				}
			}
			return nil
		})
	}

	if len(dashes) == 0 {
		log.Println("No dashboards")
	} else {
		for _, info := range dashes {
			work <- info
		}
	}

	if len(screenboards) == 0 {
		log.Print("No screenboards")
	} else {
		for _, info := range screenboards {
			// I don't think this will work - if there are errors, the workers will bail and this will wait attempting to put work on a channel no-one is listening to
			work <- info
		}
	}
	close(work)

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

	return g.Wait()
}

func createDirectories(dirs ...string) {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0777); err != nil {
			fmt.Printf("Failed to create output directory '%s': %s\n", dir, err.Error())
			os.Exit(1)
		}
	}
}
