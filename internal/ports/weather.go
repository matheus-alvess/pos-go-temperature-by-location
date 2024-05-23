package ports

type WeatherPort interface {
	GetTemperature(city string) (float64, error)
}
