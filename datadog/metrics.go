package datadog

type metricsResponse struct {
	Metrics []string `json:"metrics"`
	From    string   `json:"from"`
}

type MetricsUsage struct {
	Category   string `json:"metric_category"`
	Name       string `json:"metric_name"`
	MaxPerHour int    `json:"max_metric_hour"`
	AvgPerHour int    `json:"avg_metric_hour"`
}

type metricsUsageResponse struct {
	Usage []MetricsUsage `json:"usage"`
}
