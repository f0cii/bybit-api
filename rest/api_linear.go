package rest

import "net/http"

// LinearGetKLine
func (b *ByBit) LinearGetKLine(symbol string, interval string, from int64, limit int) (query string, resp []byte, result []OHLC, err error) {
	var ret GetKlineResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["interval"] = interval
	params["from"] = from
	if limit > 0 {
		params["limit"] = limit
	}
	query, _, err = b.PublicRequest(http.MethodGet, "public/linear/kline", params, &ret)
	if err != nil {
		return
	}
	result = ret.Result
	return
}
