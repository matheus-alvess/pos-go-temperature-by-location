package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"pos-go-temperature-by-location/internal/adapters/client"
	"pos-go-temperature-by-location/internal/core/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type mockCityInfoPort struct{}

func (m *mockCityInfoPort) GetCityFromCEP(cep string) (string, error) {
	if cep == "00000000" {
		return "", fmt.Errorf("CEP not found")
	}
	return "MockCity", nil
}

type mockWeatherPort struct{}

func (m *mockWeatherPort) GetTemperature(city string) (float64, error) {
	return 25.0, nil
}

func TestGetWeatherInvalidCEP(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	weatherService := services.NewWeatherService(&mockCityInfoPort{}, &mockWeatherPort{})
	weatherHandler := NewWeatherHandler(weatherService)

	router.GET("/weather/:cep", weatherHandler.GetWeather)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/weather/123", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Contains(t, w.Body.String(), "invalid zipcode")
}

func TestGetWeatherNotFoundCEP(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	weatherService := services.NewWeatherService(&mockCityInfoPort{}, &mockWeatherPort{})
	weatherHandler := NewWeatherHandler(weatherService)

	router.GET("/weather/:cep", weatherHandler.GetWeather)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/weather/00000000", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "can not find zipcode")
}

func TestGetWeatherSuccess(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	viaCEPURL := os.Getenv("VIA_CEP_URL")
	weatherAPIURL := os.Getenv("WEATHER_API_URL")
	weatherAPIKey := os.Getenv("WEATHER_API_KEY")

	httpClient := &client.RealHTTPClient{}
	viaCEPClient := client.NewViaCEPClient(viaCEPURL, httpClient)
	weatherAPIClient := client.NewWeatherAPIClient(weatherAPIURL, weatherAPIKey, httpClient)

	weatherService := services.NewWeatherService(viaCEPClient, weatherAPIClient)
	weatherHandler := NewWeatherHandler(weatherService)

	router.GET("/weather/:cep", weatherHandler.GetWeather)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/weather/01001000", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
