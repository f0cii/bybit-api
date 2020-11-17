package rest

import (
	sjson "encoding/json"
	"time"
)

type BaseResult struct {
	RetCode          int         `json:"ret_code"`
	RetMsg           string      `json:"ret_msg"`
	ExtCode          string      `json:"ext_code"`
	ExtInfo          string      `json:"ext_info"`
	Result           interface{} `json:"result"`
	TimeNow          string      `json:"time_now"`
	RateLimitStatus  int         `json:"rate_limit_status"`
	RateLimitResetMs int64       `json:"rate_limit_reset_ms"`
	RateLimit        int         `json:"rate_limit"`
}

type Item struct {
	Price float64 `json:"price,string"`
	Size  float64 `json:"size"`
}

type OrderBook struct {
	Asks []Item    `json:"asks"`
	Bids []Item    `json:"bids"`
	Time time.Time `json:"time"`
}

type RawItem struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
	Size   float64 `json:"size"`
	Side   string  `json:"side"` // Buy/Sell
}

type GetOrderBookResult struct {
	RetCode int       `json:"ret_code"`
	RetMsg  string    `json:"ret_msg"`
	ExtCode string    `json:"ext_code"`
	ExtInfo string    `json:"ext_info"`
	Result  []RawItem `json:"result"`
	TimeNow string    `json:"time_now"`
}

type OHLC struct {
	Symbol   string  `json:"symbol"`
	Interval string  `json:"interval"`
	OpenTime int64   `json:"open_time"`
	Open     float64 `json:"open,string"`
	High     float64 `json:"high,string"`
	Low      float64 `json:"low,string"`
	Close    float64 `json:"close,string"`
	Volume   float64 `json:"volume,string"`
	Turnover float64 `json:"turnover,string"`
}

type GetKlineResult struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  []OHLC `json:"result"`
	TimeNow string `json:"time_now"`
}

type Ticker struct {
	Symbol               string    `json:"symbol"`
	BidPrice             float64   `json:"bid_price,string"`
	AskPrice             float64   `json:"ask_price,string"`
	LastPrice            float64   `json:"last_price,string"`
	LastTickDirection    string    `json:"last_tick_direction"`
	PrevPrice24H         float64   `json:"prev_price_24h,string"`
	Price24HPcnt         float64   `json:"price_24h_pcnt,string"`
	HighPrice24H         float64   `json:"high_price_24h,string"`
	LowPrice24H          float64   `json:"low_price_24h,string"`
	PrevPrice1H          float64   `json:"prev_price_1h,string"`
	Price1HPcnt          float64   `json:"price_1h_pcnt,string"`
	MarkPrice            float64   `json:"mark_price,string"`
	IndexPrice           float64   `json:"index_price,string"`
	OpenInterest         float64   `json:"open_interest"`
	OpenValue            float64   `json:"open_value,string"`
	TotalTurnover        float64   `json:"total_turnover,string"`
	Turnover24H          float64   `json:"turnover_24h,string"`
	TotalVolume          float64   `json:"total_volume"`
	Volume24H            float64   `json:"volume_24h"`
	FundingRate          float64   `json:"funding_rate,string"`
	PredictedFundingRate float64   `json:"predicted_funding_rate,string"`
	NextFundingTime      time.Time `json:"next_funding_time"`
	CountdownHour        int       `json:"countdown_hour"`
}

type GetTickersResult struct {
	RetCode int      `json:"ret_code"`
	RetMsg  string   `json:"ret_msg"`
	ExtCode string   `json:"ext_code"`
	ExtInfo string   `json:"ext_info"`
	Result  []Ticker `json:"result"`
	TimeNow string   `json:"time_now"`
}

type TradingRecord struct {
	ID     int       `json:"id"`
	Symbol string    `json:"symbol"`
	Price  float64   `json:"price"`
	Qty    int       `json:"qty"`
	Side   string    `json:"side"`
	Time   time.Time `json:"time"`
}

type GetTradingRecordsResult struct {
	RetCode int             `json:"ret_code"`
	RetMsg  string          `json:"ret_msg"`
	ExtCode string          `json:"ext_code"`
	ExtInfo string          `json:"ext_info"`
	Result  []TradingRecord `json:"result"`
	TimeNow string          `json:"time_now"`
}

type LeverageFilter struct {
	MinLeverage  int     `json:"min_leverage"`
	MaxLeverage  int     `json:"max_leverage"`
	LeverageStep float64 `json:"leverage_step,string"`
}

type PriceFilter struct {
	MinPrice float64 `json:"min_price,string"`
	MaxPrice float64 `json:"max_price,string"`
	TickSize float64 `json:"tick_size,string"`
}

type LotSizeFilter struct {
	MaxTradingQty int `json:"max_trading_qty"`
	MinTradingQty int `json:"min_trading_qty"`
	QtyStep       int `json:"qty_step"`
}

type SymbolInfo struct {
	Name           string         `json:"name"`
	BaseCurrency   string         `json:"base_currency"`
	QuoteCurrency  string         `json:"quote_currency"`
	PriceScale     int            `json:"price_scale"`
	TakerFee       float64        `json:"taker_fee,string"`
	MakerFee       float64        `json:"maker_fee,string"`
	LeverageFilter LeverageFilter `json:"leverage_filter"`
	PriceFilter    PriceFilter    `json:"price_filter"`
	LotSizeFilter  LotSizeFilter  `json:"lot_size_filter"`
}

type GetSymbolsResult struct {
	RetCode int          `json:"ret_code"`
	RetMsg  string       `json:"ret_msg"`
	ExtCode string       `json:"ext_code"`
	ExtInfo string       `json:"ext_info"`
	Result  []SymbolInfo `json:"result"`
	TimeNow string       `json:"time_now"`
}

type Balance struct {
	Equity           float64 `json:"equity"`
	AvailableBalance float64 `json:"available_balance"`
	UsedMargin       float64 `json:"used_margin"`
	OrderMargin      float64 `json:"order_margin"`
	PositionMargin   float64 `json:"position_margin"`
	OccClosingFee    float64 `json:"occ_closing_fee"`
	OccFundingFee    float64 `json:"occ_funding_fee"`
	WalletBalance    float64 `json:"wallet_balance"`
	RealisedPnl      float64 `json:"realised_pnl"`
	UnrealisedPnl    float64 `json:"unrealised_pnl"`
	CumRealisedPnl   float64 `json:"cum_realised_pnl"`
	GivenCash        float64 `json:"given_cash"`
	ServiceCash      float64 `json:"service_cash"`
}

type GetBalanceResult struct {
	RetCode          int                  `json:"ret_code"`
	RetMsg           string               `json:"ret_msg"`
	ExtCode          string               `json:"ext_code"`
	ExtInfo          string               `json:"ext_info"`
	Result           GetBalanceResultData `json:"result"`
	TimeNow          string               `json:"time_now"`
	RateLimitStatus  int                  `json:"rate_limit_status"`
	RateLimitResetMs int64                `json:"rate_limit_reset_ms"`
	RateLimit        int                  `json:"rate_limit"`
}

type GetBalanceResultData struct {
	BTC  Balance `json:"BTC"`
	ETH  Balance `json:"ETH"`
	EOS  Balance `json:"EOS"`
	XRP  Balance `json:"XRP"`
	USDT Balance `json:"USDT"`
}

type PositionResponse struct {
	BaseResult
	Result Position `json:"result"`
}

type PositionArrayResponse struct {
	BaseResult
	Result []Position `json:"result"`
}

type Position struct {
	ID                  int       `json:"id"`
	UserID              int       `json:"user_id"`
	RiskID              int       `json:"risk_id"`
	Symbol              string    `json:"symbol"`
	Size                float64   `json:"size"`
	Side                string    `json:"side"`
	EntryPrice          float64   `json:"entry_price,string"`
	LiqPrice            float64   `json:"liq_price,string"`
	BustPrice           float64   `json:"bust_price,string"`
	TakeProfit          float64   `json:"take_profit,string"`
	StopLoss            float64   `json:"stop_loss,string"`
	TrailingStop        float64   `json:"trailing_stop,string"`
	PositionValue       float64   `json:"position_value,string"`
	Leverage            float64   `json:"leverage,string"`
	PositionStatus      string    `json:"position_status"`
	AutoAddMargin       float64   `json:"auto_add_margin"`
	OrderMargin         float64   `json:"order_margin,string"`
	PositionMargin      float64   `json:"position_margin,string"`
	OccClosingFee       float64   `json:"occ_closing_fee,string"`
	OccFundingFee       float64   `json:"occ_funding_fee,string"`
	WalletBalance       float64   `json:"wallet_balance,string"`
	CumRealisedPnl      float64   `json:"cum_realised_pnl,string"`
	CumCommission       float64   `json:"cum_commission,string"`
	RealisedPnl         float64   `json:"realised_pnl,string"`
	DeleverageIndicator float64   `json:"deleverage_indicator"`
	OcCalcData          string    `json:"oc_calc_data"`
	CrossSeq            float64   `json:"cross_seq"`
	PositionSeq         float64   `json:"position_seq"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	UnrealisedPnl       float64   `json:"unrealised_pnl"`
}

type Order struct {
	UserID        int          `json:"user_id"`
	OrderID       string       `json:"order_id"`
	Symbol        string       `json:"symbol"`
	Side          string       `json:"side"`
	OrderType     string       `json:"order_type"`
	Price         sjson.Number `json:"price"`
	Qty           float64      `json:"qty"`
	TimeInForce   string       `json:"time_in_force"`
	OrderStatus   string       `json:"order_status"`
	LastExecTime  sjson.Number `json:"last_exec_time"`
	LastExecPrice sjson.Number `json:"last_exec_price"`
	LeavesQty     float64      `json:"leaves_qty"`
	CumExecQty    float64      `json:"cum_exec_qty"`
	CumExecValue  sjson.Number `json:"cum_exec_value"`
	CumExecFee    sjson.Number `json:"cum_exec_fee"`
	RejectReason  string       `json:"reject_reason"`
	OrderLinkID   string       `json:"order_link_id"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

type OrderListResponse struct {
	BaseResult
	Result OrderListResponseResult `json:"result"`
}

type OrderListResponseResult struct {
	Data   []Order `json:"data"`
	Cursor string  `json:"cursor"`
}

type OrderResponse struct {
	BaseResult
	Result Order `json:"result"`
}

type OrderArrayResponse struct {
	BaseResult
	Result []Order `json:"result"`
}

type StopOrder struct {
	StopOrderId       string       `json:"StopOrderId"`
	UserID            int64        `json:"user_id"`
	Symbol            string       `json:"symbol"`
	Side              string       `json:"side"`
	OrderType         string       `json:"order_type"`
	Price             sjson.Number `json:"price"`
	Qty               float64      `json:"qty"`
	TimeInForce       string       `json:"time_in_force"`
	CreateType        string       `json:"create_type"`
	CancelType        string       `json:"cancel_type"`
	OrderStatus       string       `json:"order_status"`
	LeavesQty         float64      `json:"leaves_qty"`
	LeavesValue       string       `json:"leaves_value"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
	CrossStatus       string       `json:"cross_status"`
	CrossSeq          float64      `json:"cross_seq"`
	StopOrderType     string       `json:"stop_order_type"`
	TriggerBy         string       `json:"trigger_by"`
	BasePrice         sjson.Number `json:"base_price"`
	ExpectedDirection string       `json:"expected_direction"`
}

type StopOrderListResponse struct {
	BaseResult
	Result StopOrderListResponseResult `json:"result"`
}

type StopOrderListResponseResult struct {
	Data   []StopOrder `json:"data"`
	Cursor string      `json:"cursor"`
}

type StopOrderResponse struct {
	BaseResult
	Result StopOrder `json:"result"`
}

type StopOrderArrayResponse struct {
	BaseResult
	Result []StopOrder `json:"result"`
}
