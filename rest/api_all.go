package rest

import (
	"fmt"
	"net/http"
	"strconv"
)

// GetServerTime
func (b *ByBit) GetServerTime() (query string, resp []byte, timeNow int64, err error) {
	params := map[string]interface{}{}
	var ret BaseResult
	query, resp, err = b.PublicRequest(http.MethodGet, "v2/public/time", params, &ret)
	if err != nil {
		return
	}
	var t float64
	t, err = strconv.ParseFloat(ret.TimeNow, 64)
	if err != nil {
		return
	}
	timeNow = int64(t * 1000)
	return
}

// GetWalletBalance
func (b *ByBit) GetWalletBalance(coin string) (query string, resp []byte, result Balance, err error) {
	var ret GetBalanceResult
	params := map[string]interface{}{}
	params["coin"] = coin
	query, resp, err = b.SignedRequest(http.MethodGet, "v2/private/wallet/balance", params, &ret)
	if err != nil {
		return
	}
	switch coin {
	case "BTC":
		result = ret.Result.BTC
	case "ETH":
		result = ret.Result.ETH
	case "EOS":
		result = ret.Result.EOS
	case "XRP":
		result = ret.Result.XRP
	case "USDT":
		result = ret.Result.USDT
	}
	return
}

// GetPositions
func (b *ByBit) GetPositions() (query string, resp []byte, result []PositionData, err error) {
	var r PositionArrayResponse
	params := map[string]interface{}{}
	query, resp, err = b.SignedRequest(http.MethodGet, "v2/private/position/list", params, &r)
	if err != nil {
		return
	}
	if r.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", r.RetMsg, string(resp))
		return
	}
	result = r.Result
	return
}

// GetPosition
func (b *ByBit) GetPosition(symbol string) (query string, resp []byte, result Position, err error) {
	var r PositionResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	query, resp, err = b.SignedRequest(http.MethodGet, "v2/private/position/list", params, &r)
	if err != nil {
		return
	}
	if r.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", r.RetMsg, string(resp))
		return
	}
	result = r.Result
	return
}

// SetLeverage
func (b *ByBit) SetLeverage(leverage int, symbol string) (query string, resp []byte, err error) {
	var r BaseResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["leverage"] = fmt.Sprintf("%v", leverage)
	query, resp, err = b.SignedRequest(http.MethodPost, "user/leverage/save", params, &r)
	if err != nil {
		return
	}
	return
}

// WalletRecords
func (b *ByBit) WalletRecords(symbol string, page int, limit int) (query string, resp []byte, result []WalletFundRecord, err error) {
	var r WalletFundRecordResponse
	params := map[string]interface{}{}
	if symbol != "" {
		params["currency"] = symbol
	}
	if page > 0 {
		params["page"] = page
	}
	if limit > 0 {
		params["limit"] = limit
	}
	query, resp, err = b.SignedRequest(http.MethodGet, "open-api/wallet/fund/records", params, &r)
	if err != nil {
		return
	}
	if r.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", r.RetMsg, string(resp))
		return
	}
	result = r.Result.Data
	return
}

// GetTickers
func (b *ByBit) GetTickers() (query string, resp []byte, result []Ticker, err error) {
	var ret GetTickersResult
	params := map[string]interface{}{}
	query, resp, err = b.PublicRequest(http.MethodGet, "v2/public/tickers", params, &ret)
	if err != nil {
		return
	}
	result = ret.Result
	return
}

// GetTradingRecords
func (b *ByBit) GetTradingRecords(symbol string, from int64, limit int) (query string, resp []byte, result []TradingRecord, err error) {
	var ret GetTradingRecordsResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	if from > 0 {
		params["from"] = from
	}
	if limit > 0 {
		params["limit"] = limit
	}
	query, resp, err = b.PublicRequest(http.MethodGet, "v2/public/trading-records", params, &ret)
	if err != nil {
		return
	}
	result = ret.Result
	return
}

// GetSymbols
func (b *ByBit) GetSymbols() (query string, resp []byte, result []SymbolInfo, err error) {
	var ret GetSymbolsResult
	params := map[string]interface{}{}
	query, resp, err = b.PublicRequest(http.MethodGet, "v2/public/symbols", params, &ret)
	if err != nil {
		return
	}
	result = ret.Result
	return
}

// GetOpenInterest, limit max 200
func (b *ByBit) GetOpenInterest(symbol string, period string, limit int) (query string, resp []byte, result []OpenInterest, err error) {
	var ret GetOpenInterestResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["period"] = period
	if limit > 0 {
		params["limit"] = limit
	}
	query, resp, err = b.PublicRequest(http.MethodGet, "v2/public/open-interest", params, &ret)
	if err != nil {
		return
	}
	result = ret.Result
	return
}

// GetAccountRatio, limit max 200
func (b *ByBit) GetAccountRatio(symbol string, period string, limit int) (query string, resp []byte, result []AccountRatio, err error) {
	var ret GetAccountRatioResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["period"] = period
	if limit > 0 {
		params["limit"] = limit
	}
	query, resp, err = b.PublicRequest(http.MethodGet, "v2/public/account-ratio", params, &ret)
	if err != nil {
		return
	}
	result = ret.Result
	return
}
