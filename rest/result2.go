package rest

type OHLC2 struct {
	ID       int64   `json:"id"`
	Symbol   string  `json:"symbol"`
	Period   string  `json:"period"`
	Interval string  `json:"interval"`
	StartAt  int64   `json:"start_at"`
	OpenTime int64   `json:"open_time"`
	Volume   float64 `json:"volume"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Turnover float64 `json:"turnover"`
}

type GetKlineResult2 struct {
	RetCode int     `json:"ret_code"`
	RetMsg  string  `json:"ret_msg"`
	ExtCode string  `json:"ext_code"`
	ExtInfo string  `json:"ext_info"`
	Result  []OHLC2 `json:"result"`
	TimeNow string  `json:"time_now"`
}
