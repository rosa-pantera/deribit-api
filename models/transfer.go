package models

type Transfer struct {
	Amount            float64 `json:"amount"`
	CreatedTimestampu int64   `json:"created_timestamp"`
	Currency          string  `json:"currency"`
	Direction         string  `json:"direction"`
	ID                int     `json:"id"`
	OtherSide         string  `json:"other_side"`
	State             string  `json:"state"`
	Type              string  `json:"type"`
	UpdatedTimestamp  uint64  `json:"updated_timestamp"`
}
