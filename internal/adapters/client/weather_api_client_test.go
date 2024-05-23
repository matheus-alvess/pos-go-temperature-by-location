package client_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"pos-go-temperature-by-location/internal/adapters/client/mock"
	"testing"

	"pos-go-temperature-by-location/internal/adapters/client"

	"github.com/stretchr/testify/assert"
)

func TestGetTemperature_Success(t *testing.T) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"current": {"temp_c": 20}}`)),
	}
	httpClient := mock.NewMockHTTPClient(response, nil)
	weatherClient := client.NewWeatherAPIClient("http://example.com", "api_key", httpClient)

	temperature, err := weatherClient.GetTemperature("São Paulo")

	assert.NoError(t, err)
	assert.Equal(t, 20.0, temperature)
}

func TestGetTemperature_HTTPError(t *testing.T) {
	httpClient := mock.NewMockHTTPClient(nil, errors.New("HTTP error"))
	weatherClient := client.NewWeatherAPIClient("http://example.com", "api_key", httpClient)

	_, err := weatherClient.GetTemperature("São Paulo")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "HTTP error")
}

func TestGetTemperature_InvalidResponseBody(t *testing.T) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"invalid_json}`)),
	}
	httpClient := mock.NewMockHTTPClient(response, nil)
	weatherClient := client.NewWeatherAPIClient("http://example.com", "api_key", httpClient)

	_, err := weatherClient.GetTemperature("São Paulo")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal response")
}

func TestGetTemperature_StatusError(t *testing.T) {
	response := &http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(bytes.NewBufferString("error message")),
	}
	httpClient := mock.NewMockHTTPClient(response, nil)
	weatherClient := client.NewWeatherAPIClient("http://example.com", "api_key", httpClient)

	_, err := weatherClient.GetTemperature("São Paulo")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get temperature")
}
