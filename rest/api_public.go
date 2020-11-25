package rest

import (
	"net/http"
	"sort"
	"strconv"
	"time"
)

// GetServerTime Get server time.
func (b *ByBit) GetServerTime() (query string, timeNow int64, err error) {
	params := map[string]interface{}{}
	var ret BaseResult
	query, _, err = b.PublicRequest(http.MethodGet, "v2/public/time", params, &ret)
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

// GetOrderBook Get the orderbook
// 正反向合约通用
func (b *ByBit) GetOrderBook(symbol string) (query string, result OrderBook, err error) {
	var ret GetOrderBookResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	query, _, err = b.PublicRequest(http.MethodGet, "v2/public/orderBook/L2", params, &ret)
	if err != nil {
		return
	}

	for _, v := range ret.Result {
		if v.Side == "Sell" {
			result.Asks = append(result.Asks, Item{
				Price: v.Price,
				Size:  v.Size,
			})
		} else if v.Side == "Buy" {
			result.Bids = append(result.Bids, Item{
				Price: v.Price,
				Size:  v.Size,
			})
		}
	}

	sort.Slice(result.Asks, func(i, j int) bool {
		return result.Asks[i].Price < result.Asks[j].Price
	})

	sort.Slice(result.Bids, func(i, j int) bool {
		return result.Bids[i].Price > result.Bids[j].Price
	})

	var timeNow float64
	timeNow, err = strconv.ParseFloat(ret.TimeNow, 64) // 1582011750.433202
	if err != nil {
		return
	}
	result.Time = time.Unix(0, int64(timeNow*1e9))
	return
}

// GetKLine
// https://bybit-exchange.github.io/docs/inverse/#t-httprequest-2
// interval: 1 3 5 15 30 60 120 240 360 720 "D" "M" "W" "Y"
// from: From timestamp in seconds
// limit: Limit for data size per page, max size is 200. Default as showing 200 pieces of data per page
func (b *ByBit) GetKLine(symbol string, interval string, from int64, limit int) (query string, result []OHLC, err error) {
	var ret GetKlineResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["interval"] = interval
	params["from"] = from
	if limit > 0 {
		params["limit"] = limit
	}
	query, _, err = b.PublicRequest(http.MethodGet, "v2/public/kline/list", params, &ret)
	if err != nil {
		return
	}
	result = ret.Result
	return
}

func (b *ByBit) GetTickers() (query string, result []Ticker, err error) {
	// https://api-testnet.bybit.com/v2/public/tickers
	var ret GetTickersResult
	params := map[string]interface{}{}
	query, _, err = b.PublicRequest(http.MethodGet, "v2/public/tickers", params, &ret)
	if err != nil {
		return
	}
	result = ret.Result
	return
}

// from: From ID. Default: return latest data
// limit: Number of results. Default 500; max 1000
func (b *ByBit) GetTradingRecords(symbol string, from int64, limit int) (query string, result []TradingRecord, err error) {
	var ret GetTradingRecordsResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	if from > 0 {
		params["from"] = from
	}
	if limit > 0 {
		params["limit"] = limit
	}
	query, _, err = b.PublicRequest(http.MethodGet, "v2/public/trading-records", params, &ret)
	if err != nil {
		return
	}
	result = ret.Result
	return
}

func (b *ByBit) GetSymbols() (query string, result []SymbolInfo, err error) {
	var ret GetSymbolsResult
	params := map[string]interface{}{}
	query, _, err = b.PublicRequest(http.MethodGet, "v2/public/symbols", params, &ret)
	if err != nil {
		return
	}
	result = ret.Result
	return
}
