package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/porty/ddcli/markdown"
	"github.com/urfave/cli"
)

func activeMetricsFromDuration(c *cli.Context) error {
	api := getAPI()
	dur := c.Duration("duration")
	t := time.Now().Add(-1 * dur)

	metrics, err := api.GetMetrics(t)
	if err != nil {
		return err
	}
	for _, metric := range metrics {
		fmt.Println(metric)
	}
	return nil
}

type columnWriter interface {
	Write([]string) error
	Flush()
}

func top500CustomMetrics(c *cli.Context) error {
	api := getAPI()

	metrics, err := api.GetTopAverageMetrics()
	if err != nil {
		return err
	}

	var w columnWriter
	if c.String("format") == "md" {
		w = markdown.NewTableWriter(os.Stdout)
	} else {
		w = csv.NewWriter(os.Stdout)
	}
	if err := w.Write([]string{"Name", "Average per hour", "Max per hour"}); err != nil {
		return errors.New("failed to write output: " + err.Error())
	}

	for _, m := range metrics {
		if err := w.Write([]string{
			m.Name,
			strconv.Itoa(m.AvgPerHour),
			strconv.Itoa(m.MaxPerHour),
		}); err != nil {
			return errors.New("failed to write output: " + err.Error())
		}
	}

	w.Flush()
	return nil
}
