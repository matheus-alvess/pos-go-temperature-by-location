package api

import (
	"fmt"
	"net/http"
	"pos-go-temperature-by-location/internal/core/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	weatherService *services.WeatherService
}

func NewWeatherHandler(weatherService *services.WeatherService) *WeatherHandler {
	return &WeatherHandler{weatherService: weatherService}
}

func (h *WeatherHandler) GetWeather(c *gin.Context) {
	cep := c.Param("cep")

	if len(cep) != 8 || !isValidCEP(cep) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid zipcode"})
		return
	}

	weather, err := h.weatherService.GetWeather(cep)
	if err != nil {
		if err.Error() == "CEP not found" {
			c.JSON(http.StatusNotFound, gin.H{"message": "can not find zipcode"})
		} else {
			fmt.Println("failed to get temperature", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get temperature"})
		}
		return
	}

	c.JSON(http.StatusOK, weather)
}

func isValidCEP(cep string) bool {
	_, err := strconv.Atoi(cep)
	return err == nil
}
