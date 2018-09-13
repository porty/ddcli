package datadog

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type API struct {
	apiKey  string
	appKey  string
	baseURL string
}

func New(apiKey string, appKey string) *API {
	return &API{
		apiKey:  apiKey,
		appKey:  appKey,
		baseURL: "https://app.datadoghq.com",
	}
}

func (d API) GetDashboards() ([]DashboardSummary, error) {
	req, err := d.newRequest("GET", "/api/v1/dash", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad status code: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("Bad content type: " + resp.Header.Get("Content-Type"))
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Failed to read response body: " + err.Error())
	}

	dashes := struct {
		Dashes []DashboardSummary `json:"dashes"`
	}{}

	err = json.Unmarshal(b, &dashes)
	if err != nil {
		return nil, errors.New("Failed to unmarshal dashboards: " + err.Error())
	}

	return dashes.Dashes, nil
}

func (d API) GetDashboard(id string) (*Dashboard, error) {
	req, err := d.newRequest("GET", "/api/v1/dash/"+id, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad status code: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("Bad content type: " + resp.Header.Get("Content-Type"))
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Failed to read response body: " + err.Error())
	}

	respObj := struct {
		Dash     Dashboard `json:"dash"`
		URL      string    `json:"url"`
		Resource string    `json:"resource"`
	}{}

	err = json.Unmarshal(b, &respObj)
	if err != nil {
		return nil, errors.New("Failed to unmarshal JSON: " + err.Error())
	}

	return &respObj.Dash, nil
}

func (d API) GetScreenboards() ([]ScreenboardSummary, error) {
	req, err := d.newRequest("GET", "/api/v1/screen", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad status code: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("Bad content type: " + resp.Header.Get("Content-Type"))
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Failed to read response body: " + err.Error())
	}

	screens := struct {
		Screenboards []ScreenboardSummary `json:"screenboards"`
	}{}

	err = json.Unmarshal(b, &screens)
	if err != nil {
		return nil, errors.New("Failed to unmarshal screenboard summaries: " + err.Error())
	}

	return screens.Screenboards, nil
}

func (d API) GetScreenboard(id int) (*Screenboard, error) {
	req, err := d.newRequest("GET", fmt.Sprintf("/api/v1/screen/%d", id), nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad status code: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("Bad content type: " + resp.Header.Get("Content-Type"))
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Failed to read response body: " + err.Error())
	}

	screenboard := new(Screenboard)
	err = json.Unmarshal(b, screenboard)
	if err != nil {
		return nil, errors.New("Failed to unmarshal JSON: " + err.Error())
	}

	return screenboard, nil
}

func (d API) GetMonitors() ([]Monitor, error) {
	req, err := d.newRequest("GET", "/api/v1/monitor", nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad status code: %d", resp.StatusCode)
	}

	if resp.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("Bad content type: " + resp.Header.Get("Content-Type"))
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Failed to read response body: " + err.Error())
	}

	monitors := []Monitor{}

	err = json.Unmarshal(b, &monitors)
	if err != nil {
		return nil, errors.New("Failed to unmarshal monitors: " + err.Error())
	}

	return monitors, nil
}

func (d API) newRequest(method string, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, d.baseURL+endpoint, body)
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	values.Add("api_key", d.apiKey)
	values.Add("application_key", d.appKey)
	req.URL.RawQuery = values.Encode()
	return req, nil
}
