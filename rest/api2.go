package rest

import (
	"net/http"
)

// To use this functions will need to set b.BaseUrl to api2.bybit.com or api2-testnet.bybit.com

// GetFunding
// https://api2.bybit.com/funding-rate/list?symbol=BTCUSD&date=&export=false&page=1&limit=20
func (b *ByBit) GetFunding(symbol string, page int, limit int) (query string, resp []byte, result []FundingData, e error) {
	var ret FundingResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["page"] = page
	params["limit"] = limit  // fixed limit 20
	params["export"] = false // fixed export
	query, resp, e = b.PublicRequest(http.MethodGet, "funding-rate/list", params, &ret)
	if e != nil {
		return
	}
	result = ret.Result.Data
	return
}

// GetPriceIndex
// https://api2.bybit.com/api/price/index?symbol=BTCUSD&resolution=1&from=1605087277&to=1605173738
func (b *ByBit) GetPriceIndex(symbol string, resolution string, from int64, to int64) (query string, resp []byte, result []IndexOHLC, e error) {
	var ret IndexOHLCResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["resolution"] = resolution
	params["from"] = from
	params["to"] = to
	query, resp, e = b.PublicRequest(http.MethodGet, "api/price/index", params, &ret)
	if e != nil {
		return
	}
	result = ret.Result
	return
}

// GetPremiumIndex
// https://api2.bybit.com/api/premium-index-price/index?symbol=BTCUSD&from=1605087277&resolution=1&to=1605173738
func (b *ByBit) GetPremiumIndex(symbol string, resolution string, from int64, to int64) (query string, resp []byte, result []IndexOHLC, e error) {
	var ret IndexOHLCResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["resolution"] = resolution
	params["from"] = from
	params["to"] = to
	query, resp, e = b.PublicRequest(http.MethodGet, "api/premium-index-price/index", params, &ret)
	if e != nil {
		return
	}
	result = ret.Result
	return
}

// LinearGetFunding
// https://api2.bybit.com/linear/funding-rate/list?symbol=BTCUSDT&date=&export=false&page=1
func (b *ByBit) LinearGetFunding(symbol string, page int, limit int) (query string, resp []byte, result []FundingData, e error) {
	var ret FundingResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["page"] = page
	params["limit"] = limit  // fixed limit 20
	params["export"] = false // fixed export
	query, resp, e = b.PublicRequest(http.MethodGet, "linear/funding-rate/list", params, &ret)
	if e != nil {
		return
	}
	result = ret.Result.Data
	return
}

// LinearGetPriceIndex
// https://api2.bybit.com/api/linear/public/kline/price?symbol=BTCUSDT&from=1607360460&to=1610006520&resolution=30
func (b *ByBit) LinearGetPriceIndex(symbol string, resolution string, from int64, to int64) (query string, resp []byte, result []IndexOHLC, e error) {
	var ret IndexOHLCResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["resolution"] = resolution
	params["from"] = from
	params["to"] = to
	query, resp, e = b.PublicRequest(http.MethodGet, "api/linear/public/kline/price", params, &ret)
	if e != nil {
		return
	}
	result = ret.Result
	return
}

// LinearGetPremiumIndex
// https://api2.bybit.com/api/linear/public/kline/premium-price?symbol=BTCUSDT&from=1607364960&to=1610011020&resolution=30
func (b *ByBit) LinearGetPremiumIndex(symbol string, resolution string, from int64, to int64) (query string, resp []byte, result []IndexOHLC, e error) {
	var ret IndexOHLCResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["resolution"] = resolution
	params["from"] = from
	params["to"] = to
	query, resp, e = b.PublicRequest(http.MethodGet, "api/linear/public/kline/premium-price", params, &ret)
	if e != nil {
		return
	}
	result = ret.Result
	return
}
