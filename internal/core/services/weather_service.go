package services

import (
	"math"
	"pos-go-temperature-by-location/internal/core/domain"
	"pos-go-temperature-by-location/internal/ports"
)

type WeatherHandler struct {
	cityInfoPort ports.CityInfoPort
	weatherPort  ports.WeatherPort
}

func NewWeatherService(cityInfoPort ports.CityInfoPort, weatherPort ports.WeatherPort) *WeatherHandler {
	return &WeatherHandler{
		cityInfoPort: cityInfoPort,
		weatherPort:  weatherPort,
	}
}

func (s *WeatherHandler) GetWeather(cep string) (*domain.Temperature, error) {
	city, err := s.cityInfoPort.GetCityFromCEP(cep)
	if err != nil {
		return nil, err
	}

	tempC, err := s.weatherPort.GetTemperature(city)
	if err != nil {
		return nil, err
	}

	tempF := tempC*1.8 + 32
	tempK := tempC + 273.15

	return &domain.Temperature{
		TempC: math.Round(tempC*10) / 10,
		TempF: math.Round(tempF*10) / 10,
		TempK: math.Round(tempK*10) / 10,
	}, nil
}
