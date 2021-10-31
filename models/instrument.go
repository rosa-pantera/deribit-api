package models

type Instrument struct {
	TickSize                 float64 `json:"tick_size"`
	TakerCommission          float64 `json:"taker_commission"`
	Strike                   float64 `json:"strike"`
	SettlementPeriod         string  `json:"settlement_period"`
	QuoteCurrency            string  `json:"quote_currency"`
	OptionType               string  `json:"option_type"`
	MinTradeAmount           float64 `json:"min_trade_amount"`
	MaxLiquidationCommission float64 `json:"max_liquidation_commission"`
	MaxLeverage              float64 `json:"max_leverage"`
	MakerCommission          float64 `json:"maker_commission"`
	Kind                     string  `json:"kind"`
	IsActive                 bool    `json:"is_active"`
	InstrumentName           string  `json:"instrument_name"`
	ExpirationTimestamp      uint64  `json:"expiration_timestamp"`
	CreationTimestamp        uint64  `json:"creation_timestamp"`
	ContractSize             float64 `json:"contract_size"`
	BlockTradeCommission     float64 `json:"block_trade_commission"`
	BaseCurrency             string  `json:"base_currency"`
}
