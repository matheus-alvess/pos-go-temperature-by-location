package main

import (
	"fmt"
	"log"
	"os"
	"pos-go-temperature-by-location/internal/adapters/api"
	"pos-go-temperature-by-location/internal/adapters/client"
	"pos-go-temperature-by-location/internal/core/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	viaCEPURL := os.Getenv("VIA_CEP_URL")
	weatherAPIURL := os.Getenv("WEATHER_API_URL")
	weatherAPIKey := os.Getenv("WEATHER_API_KEY")

	log.Printf("ViaCEP URL: %s", viaCEPURL)
	log.Printf("WeatherAPI URL: %s", weatherAPIURL)
	log.Printf("WeatherAPI Key: %s", weatherAPIKey)

	httpClient := &client.RealHTTPClient{}

	viaCEPClient := client.NewViaCEPClient(viaCEPURL, httpClient)
	weatherAPIClient := client.NewWeatherAPIClient(weatherAPIURL, weatherAPIKey, httpClient)

	weatherService := services.NewWeatherService(viaCEPClient, weatherAPIClient)
	weatherHandler := api.NewWeatherHandler(weatherService)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/weather/:cep", weatherHandler.GetWeather)

	fmt.Println("Server is running on port 8080...")
	router.Run(":8080")
}
