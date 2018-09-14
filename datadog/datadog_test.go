package datadog

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetDashboards(t *testing.T) {
	requestCount := 0
	handler := func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		require.Equal(t, "GET", r.Method)
		require.Equal(t, "/api/v1/dash", r.URL.Path)
		require.Equal(t, "api-key", r.URL.Query().Get("api_key"))
		require.Equal(t, "app-key", r.URL.Query().Get("application_key"))

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"dashes": [
			  {
				"read_only": false,
				"resource": "/api/v1/dash/150947",
				"description": "Description 1",
				"title": "Title 1",
				"created": "2016-06-23T04:47:42.419919+00:00",
				"id": "150947",
				"created_by": {
				  "disabled": true,
				  "handle": "email1@example.com",
				  "name": "Example Name",
				  "is_admin": true,
				  "role": null,
				  "access_role": "adm",
				  "verified": true,
				  "email": "email1@example.com",
				  "icon": "https://secure.gravatar.com/avatar/x?s=48&d=retro"
				},
				"modified": "2018-08-30T00:39:37.132905+00:00"
			  },
			  {
				"read_only": true,
				"resource": "/api/v1/dash/56724",
				"description": "Description 2",
				"title": "Title 2",
				"created": "2015-06-25T07:36:02.389983+00:00",
				"id": "56724",
				"created_by": {
				  "disabled": false,
				  "handle": "email1@example.com",
				  "name": "Example Name",
				  "is_admin": false,
				  "role": "Timelord",
				  "access_role": "st",
				  "verified": true,
				  "email": "email1@example.com",
				  "icon": "https://secure.gravatar.com/avatar/x?s=48&d=retro"
				},
				"modified": "2015-06-25T08:07:33.876645+00:00"
			  }
			]
		  }`)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	api := API{
		apiKey:  "api-key",
		appKey:  "app-key",
		baseURL: server.URL,
	}

	summaries, err := api.GetDashboards()
	require.NoError(t, err)
	require.Equal(t, 1, requestCount)
	require.Len(t, summaries, 2)

	expected := []DashboardSummary{
		{
			ID:          "150947",
			ReadOnly:    false,
			Resource:    "/api/v1/dash/150947",
			Description: "Description 1",
			Title:       "Title 1",
			Created:     time.Date(2016, 6, 23, 4, 47, 42, 419919000, time.UTC),
			Modified:    time.Date(2018, 8, 30, 0, 39, 37, 132905000, time.UTC),
		},
		{
			ID:          "56724",
			ReadOnly:    true,
			Resource:    "/api/v1/dash/56724",
			Description: "Description 2",
			Title:       "Title 2",
			Created:     time.Date(2015, 6, 25, 7, 36, 2, 389983000, time.UTC),
			Modified:    time.Date(2015, 06, 25, 8, 7, 33, 876645000, time.UTC),
		},
	}

	// work around comparison of location pointers in time structures
	for i, actual := range summaries {
		require.Equal(t, expected[i].Created.UTC(), actual.Created.UTC())
		require.Equal(t, expected[i].Modified.UTC(), actual.Modified.UTC())
		expected[i].Created = actual.Created
		expected[i].Modified = actual.Modified
	}
	require.Equal(t, expected, summaries)
}

func TestGetMetrics(t *testing.T) {
	requestCount := 0
	handler := func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		require.Equal(t, "GET", r.Method)
		require.Equal(t, "/api/v1/metrics", r.URL.Path)
		require.Equal(t, "api-key", r.URL.Query().Get("api_key"))
		require.Equal(t, "app-key", r.URL.Query().Get("application_key"))
		require.Equal(t, "1545717600", r.URL.Query().Get("from"))

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"metrics": [
			  "system.load.1",
			  "system.load.15",
			  "system.load.5"
			],
			"from": "1467815773"
		  }`)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	api := API{
		apiKey:  "api-key",
		appKey:  "app-key",
		baseURL: server.URL,
	}

	metrics, err := api.GetMetrics(time.Date(2018, 12, 25, 6, 0, 0, 0, time.UTC))
	require.NoError(t, err)
	require.Equal(t, 1, requestCount)

	require.Equal(t, []string{
		"system.load.1",
		"system.load.15",
		"system.load.5",
	}, metrics)
}

func TestGetTopAverageMetrics(t *testing.T) {
	requestCount := 0
	handler := func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		require.Equal(t, "GET", r.Method)
		require.Equal(t, "/api/v1/usage/top_avg_metrics", r.URL.Path)
		require.Equal(t, "api-key", r.URL.Query().Get("api_key"))
		require.Equal(t, "app-key", r.URL.Query().Get("application_key"))
		// this might fail if you get lucky and test it at midnight over a month boundary
		thisMonth := fmt.Sprintf("%d-%02d", time.Now().Year(), time.Now().Month())
		require.Equal(t, thisMonth, r.URL.Query().Get("month"))

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
			"usage": [
			  {
				"metric_category": "custom",
				"metric_name": "custom.metric.1",
				"max_metric_hour": 12786,
				"avg_metric_hour": 7841
			  },
			  {
				"metric_category": "custom",
				"metric_name": "custom.metric.2",
				"max_metric_hour": 10828,
				"avg_metric_hour": 5986
			  }]}`)
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	api := API{
		apiKey:  "api-key",
		appKey:  "app-key",
		baseURL: server.URL,
	}

	usage, err := api.GetTopAverageMetrics()
	require.NoError(t, err)
	require.Equal(t, 1, requestCount)

	expected := []MetricsUsage{
		{
			Category:   "custom",
			Name:       "custom.metric.1",
			MaxPerHour: 12786,
			AvgPerHour: 7841,
		}, {
			Category:   "custom",
			Name:       "custom.metric.2",
			MaxPerHour: 10828,
			AvgPerHour: 5986,
		},
	}

	require.Equal(t, expected, usage)
}
