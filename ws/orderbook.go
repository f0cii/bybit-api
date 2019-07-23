package ws

import "time"

// Item stores the amount and price values
type Item struct {
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
}

type OrderBook struct {
	Bids      []Item    `json:"bids"`
	Asks      []Item    `json:"asks"`
	Timestamp time.Time `json:"timestamp"`
}
