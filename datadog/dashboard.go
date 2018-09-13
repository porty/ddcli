package datadog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"time"
)

type DashboardSummary struct {
	ID          string    `json:"id"`
	ReadOnly    bool      `json:"read_only"`
	Resource    string    `json:"resource"`
	Description string    `json:"description"`
	Title       string    `json:"title"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
}

type Dashboard struct {
	ID          int       `json:"id"`
	ReadOnly    bool      `json:"read_only"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Title       string    `json:"title"`
	Modified    time.Time `json:"modified"`
	Graphs      []struct {
		Definition struct {
			Viz      string `json:"viz"`
			Requests []struct {
				Q    string `json:"q"`
				Type string `json:"type"`
			} `json:"requests"`
		} `json:"definition"`
		Title string `json:"title"`
	} `json:"graphs"`
	TemplateVariables []struct {
		Default string `json:"default"`
		Prefix  string `json:"prefix"`
		Name    string `json:"name"`
	} `json:"template_variables"`
}

func (d DashboardSummary) Download(api *API, outputDir string) error {
	dash, err := api.GetDashboard(d.ID)
	if err != nil {
		return fmt.Errorf("failed to get dashboard #%s: %s", d.ID, err.Error())
	}
	dest := path.Join(outputDir, "dashboards", d.ID+".json")
	b, err := json.MarshalIndent(dash, "", "  ")
	if err != nil {
		return errors.New("failed to JSON marshal dashboard: " + err.Error())
	}
	if b[len(b)-1] != '\n' {
		b = append(b, '\n')
	}
	if err = ioutil.WriteFile(dest, b, 0664); err != nil {
		return fmt.Errorf("failed to write to file '%s': %s", dest, err.Error())
	}
	return nil
}
