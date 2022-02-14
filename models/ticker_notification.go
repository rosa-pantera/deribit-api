package models

type TickerNotification struct {
	UnderlyingPrice float64 `json:"underlying_price"`
	UnderlyingIndex string  `json:"underlying_index"`
	Timestamp       uint64  `json:"timestamp"`
	Stats           Stats   `json:"stats"`
	State           string  `json:"state"`
	SettlementPrice float64 `json:"settlement_price"`
	OpenInterest    float64 `json:"open_interest"`
	MinPrice        float64 `json:"min_price"`
	MaxPrice        float64 `json:"max_price"`
	MarkPrice       float64 `json:"mark_price"`
	MarkIV          float64 `json:"mark_iv"`
	LastPrice       float64 `json:"last_price"`
	InterestRate    float64 `json:"interest_rate"`
	InstrumentName  string  `json:"instrument_name"`
	IndexPrice      float64 `json:"index_price"`
	Greeks          Greeks  `json:"greeks"`
	Funding8H       float64 `json:"funding_8h"`
	// EstimatedDeliveryPrice float64 `json:"estimated_delivery_price"`
	CurrentFunding float64 `json:"current_funding"`
	BidIV          float64 `json:"bid_iv"`
	BestBidPrice   float64 `json:"best_bid_price"`
	BestBidAmount  float64 `json:"best_bid_amount"`
	BestAskPrice   float64 `json:"best_ask_price"`
	BestAskAmount  float64 `json:"best_ask_amount"`
	AskIV          float64 `json:"ask_iv"`
}
