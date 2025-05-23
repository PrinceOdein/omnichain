package core

// import "time"

type Transaction struct {
	Sender    string  `json:"sender"`
	Receiver  string  `json:"receiver"`
	Amount    float64 `json:"amount"`
	Timestamp int64   `json:"timestamp"`
}
