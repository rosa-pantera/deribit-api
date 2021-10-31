package models

type DepositAddress struct {
	Address           string `json:"address"`
	CreationTimestamp uint64 `json:"creation_timestamp"`
	Currency          string `json:"currency"`
	Type              string `json:"type"`
}
