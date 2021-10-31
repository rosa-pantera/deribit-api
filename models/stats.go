package models

type Stats struct {
	VolumeUSD   float64 `json:"volume_usd"`
	Volume      float64 `json:"volume"`
	PriceChange float64 `json:"price_change"`
	Low         float64 `json:"low"`
	High        float64 `json:"high"`
}
