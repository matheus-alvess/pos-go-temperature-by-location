package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"pos-go-temperature-by-location/internal/core/domain"
)

type WeatherAPIClient struct {
	URL    string
	APIKey string
	Client HTTPClient
}

func NewWeatherAPIClient(url, apiKey string, client HTTPClient) *WeatherAPIClient {
	return &WeatherAPIClient{URL: url, APIKey: apiKey, Client: client}
}

func (w *WeatherAPIClient) GetTemperature(city string) (float64, error) {
	encodedCity := url.QueryEscape(city)
	requestURL := fmt.Sprintf("%s?key=%s&q=%s", w.URL, w.APIKey, encodedCity)
	log.Printf("Requesting URL: %s", requestURL)

	resp, err := w.Client.Get(requestURL)
	if err != nil {
		return 0, fmt.Errorf("failed to get temperature: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to get temperature: status %d, body: %s", resp.StatusCode, body)
	}

	var result domain.WeatherAPIResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return result.Current.TempC, nil
}
