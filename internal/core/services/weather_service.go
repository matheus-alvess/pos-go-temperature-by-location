package services

import (
	"pos-go-temperature-by-location/internal/core/domain"
	"pos-go-temperature-by-location/internal/ports"
)

type WeatherService struct {
	cityInfoPort ports.CityInfoPort
	weatherPort  ports.WeatherPort
}

func NewWeatherService(cityInfoPort ports.CityInfoPort, weatherPort ports.WeatherPort) *WeatherService {
	return &WeatherService{
		cityInfoPort: cityInfoPort,
		weatherPort:  weatherPort,
	}
}

func (s *WeatherService) GetWeather(cep string) (*domain.Temperature, error) {
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
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}, nil
}
