package models

type InstrumentNotification struct {
	Timestamp      uint64 `json:"timestamp"`
	State          string `json:"state"`
	InstrumentName string `json:"instrument_name"`
}
