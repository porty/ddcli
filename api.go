package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/porty/ddcli/models"
)

type DatadogAPI struct {
	apiKey string
	appKey string
}

func (dd DatadogAPI) GetDashboards() ([]models.DashboardSummary, error) {
	req, err := dd.newRequest("GET", "/api/v1/dash", nil)
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
		Dashes []models.DashboardSummary `json:"dashes"`
	}{}

	err = json.Unmarshal(b, &dashes)
	if err != nil {
		return nil, errors.New("Failed to unmarshal dashboards: " + err.Error())
	}

	return dashes.Dashes, nil
}

func (dd DatadogAPI) GetDashboard(id string) (*models.Dashboard, error) {
	req, err := dd.newRequest("GET", "/api/v1/dash/"+id, nil)
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
		Dash     models.Dashboard `json:"dash"`
		URL      string           `json:"url"`
		Resource string           `json:"resource"`
	}{}

	err = json.Unmarshal(b, &respObj)
	if err != nil {
		return nil, errors.New("Failed to unmarshal JSON: " + err.Error())
	}

	return &respObj.Dash, nil
}

func (dd DatadogAPI) GetScreenboards() ([]models.ScreenboardSummary, error) {
	req, err := dd.newRequest("GET", "/api/v1/screen", nil)
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
		Screenboards []models.ScreenboardSummary `json:"screenboards"`
	}{}

	err = json.Unmarshal(b, &screens)
	if err != nil {
		return nil, errors.New("Failed to unmarshal screenboard summaries: " + err.Error())
	}

	return screens.Screenboards, nil
}

func (dd DatadogAPI) GetScreenboard(id int) (*models.Screenboard, error) {
	req, err := dd.newRequest("GET", fmt.Sprintf("/api/v1/screen/%d", id), nil)
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

	screenboard := new(models.Screenboard)
	err = json.Unmarshal(b, screenboard)
	if err != nil {
		return nil, errors.New("Failed to unmarshal JSON: " + err.Error())
	}

	return screenboard, nil
}

func (dd DatadogAPI) GetMonitors() ([]models.Monitor, error) {
	req, err := dd.newRequest("GET", "/api/v1/monitor", nil)
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

	monitors := []models.Monitor{}

	err = json.Unmarshal(b, &monitors)
	if err != nil {
		return nil, errors.New("Failed to unmarshal monitors: " + err.Error())
	}

	return monitors, nil
}

func (dd DatadogAPI) newRequest(method string, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, "https://app.datadoghq.com"+endpoint, body)
	if err != nil {
		return nil, err
	}
	values := url.Values{}
	values.Add("api_key", dd.apiKey)
	values.Add("application_key", dd.appKey)
	req.URL.RawQuery = values.Encode()
	return req, nil
}

func printJSON(b []byte) {
	buf := bytes.Buffer{}
	err := json.Indent(&buf, b, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(buf.String())
}
