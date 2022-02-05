package models

type UnsubscribeResponse []string

type UnsubscribeAllResponse struct {
	// Result should be "ok"
	Result string `json:"result"`
}
