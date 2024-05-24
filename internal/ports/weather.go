package ports

import "pos-go-temperature-by-location/internal/core/domain"

type WeatherPort interface {
	GetTemperature(city string) (float64, error)
}

//go:generate mockgen -source=./weather.go -destination=./mock/weather_service_mock.go -package=mocks
type WeatherService interface {
	GetWeather(cep string) (*domain.Temperature, error)
}
