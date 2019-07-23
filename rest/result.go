package rest

import (
	"time"
)

type BaseResult struct {
	RetCode         int         `json:"ret_code"`
	RetMsg          string      `json:"ret_msg"`
	ExtCode         string      `json:"ext_code"`
	Result          interface{} `json:"result"`
	TimeNow         string      `json:"time_now"`
	RateLimitStatus int         `json:"rate_limit_status"`
}

type CreateOrderResult struct {
	RetCode         int    `json:"ret_code"`
	RetMsg          string `json:"ret_msg"`
	ExtCode         string `json:"ext_code"`
	Result          Order0 `json:"result"`
	TimeNow         string `json:"time_now"`
	RateLimitStatus int    `json:"rate_limit_status"`
}

type CancelOrderResult struct {
	RetCode         int    `json:"ret_code"`
	RetMsg          string `json:"ret_msg"`
	ExtCode         string `json:"ext_code"`
	Result          Order2 `json:"result"`
	TimeNow         string `json:"time_now"`
	RateLimitStatus int    `json:"rate_limit_status"`
}

type OrderListResultData struct {
	Data        []Order1 `json:"data"`
	CurrentPage int      `json:"current_page"`
	LastPage    int      `json:"last_page"`
}

type OrderListResult struct {
	RetCode         int                 `json:"ret_code"`
	RetMsg          string              `json:"ret_msg"`
	ExtCode         string              `json:"ext_code"`
	Result          OrderListResultData `json:"result"`
	TimeNow         string              `json:"time_now"`
	RateLimitStatus int                 `json:"rate_limit_status"`
}

// Order ...
type Order struct {
	OrderID       string    `json:"order_id"`
	UserID        int       `json:"user_id"`
	Symbol        string    `json:"symbol"`
	Side          string    `json:"side"`
	OrderType     string    `json:"order_type"`
	Price         float64   `json:"price"`
	Qty           float64   `json:"qty"`
	TimeInForce   string    `json:"time_in_force"`
	OrderStatus   string    `json:"order_status"`
	LastExecTime  string    `json:"last_exec_time"`
	LastExecPrice float64   `json:"last_exec_price"`
	LeavesQty     float64   `json:"leaves_qty"`
	CumExecQty    float64   `json:"cum_exec_qty"`
	CumExecValue  float64   `json:"cum_exec_value"`
	CumExecFee    float64   `json:"cum_exec_fee"`
	RejectReason  string    `json:"reject_reason"`
	OrderLinkID   string    `json:"order_link_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Create order
type Order0 struct {
	OrderID       string    `json:"order_id"`
	UserID        int       `json:"user_id"`
	Symbol        string    `json:"symbol"`
	Side          string    `json:"side"`
	OrderType     string    `json:"order_type"`
	Price         float64   `json:"price,string"`
	Qty           float64   `json:"qty"`
	TimeInForce   string    `json:"time_in_force"`
	OrderStatus   string    `json:"order_status"`
	LastExecTime  string    `json:"last_exec_time"`
	LastExecPrice float64   `json:"last_exec_price"`
	LeavesQty     float64   `json:"leaves_qty"`
	CumExecQty    float64   `json:"cum_exec_qty"`
	CumExecValue  float64   `json:"cum_exec_value"`
	CumExecFee    float64   `json:"cum_exec_fee"`
	RejectReason  string    `json:"reject_reason"`
	OrderLinkID   string    `json:"order_link_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Get order
type Order1 struct {
	OrderID       string    `json:"order_id"`
	UserID        int       `json:"user_id"`
	Symbol        string    `json:"symbol"`
	Side          string    `json:"side"`
	OrderType     string    `json:"order_type"`
	Price         float64   `json:"price"`
	Qty           float64   `json:"qty"`
	TimeInForce   string    `json:"time_in_force"`
	OrderStatus   string    `json:"order_status"`
	LastExecTime  string    `json:"last_exec_time"`
	LastExecPrice float64   `json:"last_exec_price"`
	LeavesQty     float64   `json:"leaves_qty"`
	CumExecQty    float64   `json:"cum_exec_qty"`
	CumExecValue  float64   `json:"cum_exec_value"`
	CumExecFee    float64   `json:"cum_exec_fee"`
	RejectReason  string    `json:"reject_reason"`
	OrderLinkID   string    `json:"order_link_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Cancel order
type Order2 struct {
	OrderID       string    `json:"order_id"`
	UserID        int       `json:"user_id"`
	Symbol        string    `json:"symbol"`
	Side          string    `json:"side"`
	OrderType     string    `json:"order_type"`
	Price         float64   `json:"price,string"`
	Qty           float64   `json:"qty"`
	TimeInForce   string    `json:"time_in_force"`
	OrderStatus   string    `json:"order_status"`
	LastExecTime  string    `json:"last_exec_time"`
	LastExecPrice float64   `json:"last_exec_price,string"`
	LeavesQty     float64   `json:"leaves_qty"`
	CumExecQty    float64   `json:"cum_exec_qty"`
	CumExecValue  float64   `json:"cum_exec_value,string"`
	CumExecFee    float64   `json:"cum_exec_fee,string"`
	RejectReason  string    `json:"reject_reason"`
	OrderLinkID   string    `json:"order_link_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type GetLeverageResult struct {
	RetCode         int                     `json:"ret_code"`
	RetMsg          string                  `json:"ret_msg"`
	ExtCode         string                  `json:"ext_code"`
	Result          map[string]LeverageItem `json:"result"`
	TimeNow         string                  `json:"time_now"`
	RateLimitStatus int                     `json:"rate_limit_status"`
}

type LeverageItem struct {
	Leverage int `json:"leverage"`
}

type Position struct {
	ID                  int           `json:"id"`
	UserID              int           `json:"user_id"`
	RiskID              int           `json:"risk_id"`
	Symbol              string        `json:"symbol"`
	Size                int           `json:"size"`
	Side                string        `json:"side"`
	EntryPrice          float64       `json:"entry_price"`
	LiqPrice            int           `json:"liq_price"`
	BustPrice           int           `json:"bust_price"`
	TakeProfit          int           `json:"take_profit"`
	StopLoss            int           `json:"stop_loss"`
	TrailingStop        int           `json:"trailing_stop"`
	PositionValue       float64       `json:"position_value"`
	Leverage            int           `json:"leverage"`
	PositionStatus      string        `json:"position_status"`
	AutoAddMargin       int           `json:"auto_add_margin"`
	OrderMargin         float64       `json:"order_margin"`
	PositionMargin      float64       `json:"position_margin"`
	OccClosingFee       float64       `json:"occ_closing_fee"`
	OccFundingFee       float64       `json:"occ_funding_fee"`
	ExtFields           []interface{} `json:"ext_fields"`
	WalletBalance       float64       `json:"wallet_balance"`
	CumRealisedPnl      float64       `json:"cum_realised_pnl"`
	CumCommission       int           `json:"cum_commission"`
	RealisedPnl         float64       `json:"realised_pnl"`
	DeleverageIndicator int           `json:"deleverage_indicator"`
	OcCalcData          string        `json:"oc_calc_data"`
	CrossSeq            int           `json:"cross_seq"`
	PositionSeq         int           `json:"position_seq"`
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           time.Time     `json:"updated_at"`
	UnrealisedPnl       float64       `json:"unrealised_pnl"`
}

//type PositionListResultData struct {
//	Data        []Position `json:"data"`
//	CurrentPage int      `json:"current_page"`
//	LastPage    int      `json:"last_page"`
//}

type PositionListResult struct {
	BaseResult
	ExtInfo interface{} `json:"ext_info"`
	Result  []Position  `json:"result"`
}
