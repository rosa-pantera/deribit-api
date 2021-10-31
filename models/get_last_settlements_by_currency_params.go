package models

type GetLastSettlementsByCurrencyParams struct {
	Currency             string `json:"currency"`
	Type                 string `json:"type,omitempty"`
	Count                int    `json:"count,omitempty"`
	Continuation         string `json:"continuation,omitempty"`
	SearchStartTimestamp uint   `json:"search_start_timestamp,omitempty"`
}
