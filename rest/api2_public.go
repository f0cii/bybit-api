package rest

import (
	sjson "encoding/json"
	"net/http"
	"time"
)

type FundingResult struct {
	RetCode int     `json:"ret_code"`
	RetMsg  string  `json:"ret_msg"`
	ExtCode string  `json:"ext_code"`
	ExtInfo string  `json:"ext_info"`
	Result  Funding `json:"result"`
	TimeNow string  `json:"time_now"`
}

type Funding struct {
	CurrentPage  int           `json:"current_page"`
	Data         []FundingData `json:"data"`
	FirstPageUrl string        `json:"first_page_url"`
	From         int           `json:"from"`
	LastPage     int           `json:"last_page"`
	LastPageUrl  string        `json:"last_page_url"`
	NextPageUrl  string        `json:"next_page_url"`
	Path         string        `json:"path"`
	PerPage      sjson.Number  `json:"per_page"`
	PrevPageUrl  string        `json:"prev_page_url"`
	To           int           `json:"to"`
	Total        int           `json:"total"`
}

type FundingData struct {
	Id     int          `json:"id"`
	Symbol string       `json:"symbol"`
	Value  sjson.Number `json:"value"`
	Time   time.Time    `json:"time"`
}

// https://api2.bybit.com/funding-rate/list?symbol=BTCUSD&date=&export=false&page=1&limit=20
// To use this you will need to set b.BaseUrl to api2.bybit.com or api2-testnet.bybit.com
func (b *ByBit) GetFunding(symbol string, page int, limit int) (query string, result []FundingData, e error) {
	var ret FundingResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["page"] = page
	params["limit"] = limit  // fixed limit 20
	params["export"] = false // fixed export
	query, _, e = b.PublicRequest(http.MethodGet, "funding-rate/list", params, &ret)
	if e != nil {
		return
	}
	result = ret.Result.Data
	return
}

// for linear: https://api2.bybit.com/linear/funding-rate/list?symbol=BTCUSDT&date=&export=false&page=1
