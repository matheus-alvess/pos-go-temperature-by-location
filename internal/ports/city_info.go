package ports

type CityInfoPort interface {
	GetCityFromCEP(cep string) (string, error)
}
