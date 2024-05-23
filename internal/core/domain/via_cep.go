package domain

type ViaCepAPIResponse struct {
	Locality string `json:"localidade"`
	Error    bool   `json:"erro,omitempty"`
}
