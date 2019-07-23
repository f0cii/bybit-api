package ws

import (
	"strconv"
	"time"
)

type OrderBookL2 struct {
	ID     int64   `json:"id"`
	Price  float64 `json:"price,string"`
	Side   string  `json:"side"`
	Size   int64   `json:"size"`
	Symbol string  `json:"symbol"`
}

type OrderBookL2Delta struct {
	Delete []*OrderBookL2 `json:"delete"`
	Update []*OrderBookL2 `json:"update"`
	Insert []*OrderBookL2 `json:"insert"`
}

func (o *OrderBookL2) Key() string {
	return strconv.FormatInt(o.ID, 10)
}

type Trade struct {
	Timestamp     time.Time `json:"timestamp"`
	Symbol        string    `json:"symbol"`
	Side          string    `json:"side"`
	Size          int       `json:"size"`
	Price         float64   `json:"price"`
	TickDirection string    `json:"tick_direction"`
	TradeID       string    `json:"trade_id"`
	CrossSeq      int       `json:"cross_seq"`
}

type KLine struct {
	ID       int64   `json:"id"`        // 563
	Symbol   string  `json:"symbol"`    // BTCUSD
	OpenTime int64   `json:"open_time"` // 1539918000
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Volume   float64 `json:"volume"`
	Turnover float64 `json:"turnover"` // 0.0013844
	Interval string  `json:"interval"` // 1m
}

type Insurance struct {
	Currency      string    `json:"currency"`
	Timestamp     time.Time `json:"timestamp"`
	WalletBalance int64     `json:"wallet_balance"`
}

type Instrument struct {
	Symbol     string  `json:"symbol"`
	MarkPrice  float64 `json:"mark_price"`
	IndexPrice float64 `json:"index_price"`
}

type Order struct {
	OrderID      string    `json:"order_id"`
	OrderLinkID  string    `json:"order_link_id"`
	Symbol       string    `json:"symbol"`
	Side         string    `json:"side"`       // Buy/Sell
	OrderType    string    `json:"order_type"` // Market/Limit
	Price        float64   `json:"price"`
	Qty          float64   `json:"qty"`
	TimeInForce  string    `json:"time_in_force"` // GoodTillCancel/ImmediateOrCancel/FillOrKill/PostOnly
	OrderStatus  string    `json:"order_status"`  // Filled
	LeavesQty    float64   `json:"leaves_qty"`
	CumExecQty   float64   `json:"cum_exec_qty"`
	CumExecValue float64   `json:"cum_exec_value"`
	CumExecFee   float64   `json:"cum_exec_fee"`
	Timestamp    time.Time `json:"timestamp"`
}

type Execution struct {
	Symbol      string    `json:"symbol"`
	Side        string    `json:"side"`
	OrderID     string    `json:"order_id"`
	ExecID      string    `json:"exec_id"`
	OrderLinkID string    `json:"order_link_id"`
	Price       float64   `json:"price"`
	ExecQty     float64   `json:"exec_qty"`
	ExecFee     float64   `json:"exec_fee"`
	LeavesQty   float64   `json:"leaves_qty"`
	IsMaker     bool      `json:"is_maker"`
	TradeTime   time.Time `json:"trade_time"`
}

type Position struct {
	Symbol         string  `json:"symbol"`
	Side           string  `json:"side"` // None
	Size           float64 `json:"size"`
	EntryPrice     float64 `json:"entry_price"`
	LiqPrice       float64 `json:"liq_price"`
	BustPrice      float64 `json:"bust_price"`
	TakeProfit     float64 `json:"take_profit"`
	StopLoss       float64 `json:"stop_loss"`
	TrailingStop   float64 `json:"trailing_stop"`
	PositionValue  float64 `json:"position_value"`
	Leverage       float64 `json:"leverage"`
	PositionStatus string  `json:"position_status"` // Normal
	AutoAddMargin  float64 `json:"auto_add_margin"`
	CrossSeq       float64 `json:"cross_seq"`
	PositionSeq    float64 `json:"position_seq"`
}
