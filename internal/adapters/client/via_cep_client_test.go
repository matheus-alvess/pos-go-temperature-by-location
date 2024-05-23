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

func TestGetCityFromCEP_Success(t *testing.T) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"localidade": "São Paulo"}`)),
	}
	httpClient := mock.NewMockHTTPClient(response, nil)
	cepClient := client.NewViaCEPClient("http://example.com", httpClient)

	city, err := cepClient.GetCityFromCEP("01000-000")

	assert.NoError(t, err)
	assert.Equal(t, "São Paulo", city)
}

func TestGetCityFromCEP_HTTPError(t *testing.T) {
	httpClient := mock.NewMockHTTPClient(nil, errors.New("HTTP error"))
	cepClient := client.NewViaCEPClient("http://example.com", httpClient)

	_, err := cepClient.GetCityFromCEP("01000-000")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "HTTP error")
}

func TestGetCityFromCEP_InvalidResponseBody(t *testing.T) {
	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`{"invalid_json}`)),
	}
	httpClient := mock.NewMockHTTPClient(response, nil)
	cepClient := client.NewViaCEPClient("http://example.com", httpClient)

	_, err := cepClient.GetCityFromCEP("01000-000")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

func TestGetCityFromCEP_StatusError(t *testing.T) {
	response := &http.Response{
		StatusCode: http.StatusNotFound,
		Body:       ioutil.NopCloser(bytes.NewBufferString("Not Found")),
	}
	httpClient := mock.NewMockHTTPClient(response, nil)
	cepClient := client.NewViaCEPClient("http://example.com", httpClient)

	_, err := cepClient.GetCityFromCEP("01000-000")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get city from CEP")
}
