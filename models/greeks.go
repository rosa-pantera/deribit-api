package models

type Greeks struct {
	Vega  float64 `json:"vega"`
	Theta float64 `json:"theta"`
	Rho   float64 `json:"rho"`
	Gamma float64 `json:"gamma"`
	Delta float64 `json:"delta"`
}
