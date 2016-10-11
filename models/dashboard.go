package models

import "time"

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
