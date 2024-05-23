package domain

type WeatherAPIResponse struct {
	Current Current `json:"current"`
}

type Current struct {
	TempC float64 `json:"temp_c"`
}

type Temperature struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}
