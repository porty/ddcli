package datadog

import "time"

type ScreenboardSummary struct {
	ID       int       `json:"id"`
	ReadOnly bool      `json:"read_only"`
	Resource string    `json:"resource"`
	Created  time.Time `json:"created"`
	Title    string    `json:"title"`
	Modified time.Time `json:"modified"`
}

type Screenboard struct {
	ID                int           `json:"id"`
	BoardTitle        string        `json:"board_title"`
	ReadOnly          bool          `json:"read_only"`
	BoardBgtype       string        `json:"board_bgtype"`
	Created           time.Time     `json:"created"`
	OriginalTitle     string        `json:"original_title"`
	Modified          time.Time     `json:"modified"`
	Height            int           `json:"height"`
	Width             string        `json:"width"`
	TemplateVariables []interface{} `json:"template_variables"`
	Shared            bool          `json:"shared"`
	TitleEdited       bool          `json:"title_edited"`
	Widgets           []struct {
		BoardID    int    `json:"board_id"`
		TitleSize  int    `json:"title_size"`
		Title      bool   `json:"title"`
		TitleAlign string `json:"title_align"`
		TitleText  string `json:"title_text"`
		Height     int    `json:"height"`
		TileDef    struct {
			Viz      string `json:"viz"`
			Requests []struct {
				Q                  string        `json:"q"`
				Aggregator         string        `json:"aggregator,omitempty"`
				ConditionalFormats []interface{} `json:"conditional_formats"`
				Type               string        `json:"type"`
				Style              struct {
					Type string `json:"type"`
				} `json:"style"`
			} `json:"requests"`
		} `json:"tile_def"`
		Width      int    `json:"width"`
		Timeframe  string `json:"timeframe"`
		Y          int    `json:"y"`
		X          int    `json:"x"`
		LegendSize string `json:"legend_size"`
		Type       string `json:"type"`
		Legend     bool   `json:"legend"`
	} `json:"widgets"`
}
