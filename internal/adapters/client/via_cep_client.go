package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"pos-go-temperature-by-location/internal/core/domain"
)

type ViaCEPClient struct {
	URL    string
	Client HTTPClient
}

func NewViaCEPClient(url string, client HTTPClient) *ViaCEPClient {
	return &ViaCEPClient{URL: url, Client: client}
}

func (v *ViaCEPClient) GetCityFromCEP(cep string) (string, error) {
	requestURL := fmt.Sprintf("%s%s/json/", v.URL, cep)
	log.Printf("Requesting URL: %s", requestURL)

	resp, err := v.Client.Get(requestURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get city from CEP: status %d, body: %s", resp.StatusCode, body)
	}

	var response domain.ViaCepAPIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	if response.Error {
		return "", fmt.Errorf("CEP not found")
	}

	return response.Locality, nil
}
