package datadog

import "time"

type Monitor struct {
	ID                int           `json:"id"`
	OrgID             int           `json:"org_id"`
	Tags              []string      `json:"tags"`
	Deleted           interface{}   `json:"deleted"`
	Query             string        `json:"query"`
	Message           string        `json:"message"`
	MatchingDowntimes []interface{} `json:"matching_downtimes"`
	Multi             bool          `json:"multi"`
	Name              string        `json:"name"`
	Created           time.Time     `json:"created"`
	CreatedAt         int64         `json:"created_at"`
	Modified          time.Time     `json:"modified"`
	OverallState      string        `json:"overall_state"`
	Type              string        `json:"type"`
	Creator           struct {
		ID     int    `json:"id"`
		Handle string `json:"handle"`
		Name   string `json:"name"`
		Email  string `json:"email"`
	} `json:"creator"`
	Options struct {
		NotifyAudit bool `json:"notify_audit"`
		Locked      bool `json:"locked"`
		TimeoutH    int  `json:"timeout_h"`
		Silenced    struct {
		} `json:"silenced"`
		Thresholds struct {
			Critical float64 `json:"critical"`
		} `json:"thresholds"`
		RequireFullWindow bool `json:"require_full_window"`
		NotifyNoData      bool `json:"notify_no_data"`
		RenotifyInterval  int  `json:"renotify_interval"`
		NoDataTimeframe   int  `json:"no_data_timeframe"`
	} `json:"options"`
}
