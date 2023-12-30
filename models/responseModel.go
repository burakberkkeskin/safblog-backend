package models

type Response struct {
	Status string `json:"status"`
	Data   string `json:"Data"`
	Error  string `json:"Error"`
}
