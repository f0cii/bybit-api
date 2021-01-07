package rest

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"
)

// GetOrderBook
func (b *ByBit) GetOrderBook(symbol string) (query string, resp []byte, result OrderBook, err error) {
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
func (b *ByBit) GetKLine(symbol string, interval string, from int64, limit int) (query string, resp []byte, result []OHLC, err error) {
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

// GetOrders
func (b *ByBit) GetOrders(symbol string, orderStatus string, direction string, limit int, cursor string) (query string, resp []byte, result OrderListResponseResult, err error) {
	var cResult OrderListResponse
	if limit == 0 {
		limit = 20
	}
	params := map[string]interface{}{}
	params["symbol"] = symbol
	if orderStatus != "" {
		params["order_status"] = orderStatus
	}
	if direction != "" {
		params["direction"] = direction
	}
	params["limit"] = limit
	if cursor != "" {
		params["cursor"] = cursor
	}
	query, resp, err = b.SignedRequest(http.MethodGet, "v2/private/order/list", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	result = cResult.Result
	return
}

// GetActiveOrders
func (b *ByBit) GetActiveOrders(symbol string) (query string, resp []byte, result OrderArrayResponse, err error) {
	var cResult OrderArrayResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	query, resp, err = b.SignedRequest(http.MethodGet, "v2/private/order", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	result = cResult
	return
}

// CreateOrder
func (b *ByBit) CreateOrder(side string, orderType string, price float64,
	qty int, timeInForce string, takeProfit float64, stopLoss float64, reduceOnly bool,
	closeOnTrigger bool, orderLinkID string, symbol string) (query string, resp []byte, result Order, err error) {
	var cResult OrderResponse
	params := map[string]interface{}{}
	params["side"] = side
	params["symbol"] = symbol
	params["order_type"] = orderType
	params["qty"] = qty
	if price > 0 {
		params["price"] = price
	}
	params["time_in_force"] = timeInForce
	if takeProfit > 0 {
		params["take_profit"] = takeProfit
	}
	if stopLoss > 0 {
		params["stop_loss"] = stopLoss
	}
	if reduceOnly {
		params["reduce_only"] = true
	}
	if closeOnTrigger {
		params["close_on_trigger"] = true
	}
	if orderLinkID != "" {
		params["order_link_id"] = orderLinkID
	}
	query, resp, err = b.SignedRequest(http.MethodPost, "v2/private/order/create", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	result = cResult.Result
	return
}

// ReplaceOrder
func (b *ByBit) ReplaceOrder(symbol string, orderID string, qty int, price float64) (query string, resp []byte, result Order, err error) {
	var cResult OrderResponse
	params := map[string]interface{}{}
	params["order_id"] = orderID
	params["symbol"] = symbol
	if qty > 0 {
		params["p_r_qty"] = qty
	}
	if price > 0 {
		params["p_r_price"] = price
	}
	query, resp, err = b.SignedRequest(http.MethodPost, "v2/private/order/replace", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	result.OrderId = cResult.Result.OrderId
	return
}

// CancelOrder
func (b *ByBit) CancelOrder(orderID string, symbol string) (query string, resp []byte, result Order, err error) {
	var cResult OrderResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	if orderID != "" {
		params["order_id"] = orderID
	}
	query, resp, err = b.SignedRequest(http.MethodPost, "v2/private/order/cancel", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	result = cResult.Result
	return
}

// CancelAllOrder
func (b *ByBit) CancelAllOrder(symbol string) (query string, resp []byte, result []Order, err error) {
	var cResult OrderArrayResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	query, resp, err = b.SignedRequest(http.MethodPost, "v2/private/order/cancelAll", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	result = cResult.Result
	return
}

// GetStopOrders
func (b *ByBit) GetStopOrders(symbol string, stopOrderStatus string, direction string, limit int, cursor string) (query string, resp []byte, result StopOrderListResponseResult, err error) {
	var cResult StopOrderListResponse
	if limit == 0 {
		limit = 20
	}
	params := map[string]interface{}{}
	params["symbol"] = symbol
	if stopOrderStatus != "" {
		params["stop_order_status"] = stopOrderStatus
	}
	if direction != "" {
		params["direction"] = direction
	}
	params["limit"] = limit
	if cursor != "" {
		params["cursor"] = cursor
	}
	query, resp, err = b.SignedRequest(http.MethodGet, "v2/private/stop-order/list", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	result = cResult.Result
	return
}

// GetActiveStopOrders
func (b *ByBit) GetActiveStopOrders(symbol string) (query string, resp []byte, result StopOrderArrayResponse, err error) {
	var cResult StopOrderArrayResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	query, resp, err = b.SignedRequest(http.MethodGet, "v2/private/stop-order", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	result = cResult
	return
}

// CreateStopOrder
func (b *ByBit) CreateStopOrder(side string, orderType string, price float64, basePrice float64, stopPx float64,
	qty int, triggerBy string, timeInForce string, closeOnTrigger bool, symbol string) (query string, resp []byte, result StopOrder, err error) {
	var cResult StopOrderResponse
	params := map[string]interface{}{}
	params["side"] = side
	params["symbol"] = symbol
	params["order_type"] = orderType
	params["qty"] = qty
	if price > 0 {
		params["price"] = price
	}
	params["base_price"] = basePrice
	params["stop_px"] = stopPx
	params["time_in_force"] = timeInForce
	if closeOnTrigger {
		params["close_on_trigger"] = true
	}
	if triggerBy != "" {
		params["trigger_by"] = triggerBy
	}
	query, resp, err = b.SignedRequest(http.MethodPost, "v2/private/stop-order/create", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	result = cResult.Result
	return
}

// ReplaceStopOrder
func (b *ByBit) ReplaceStopOrder(symbol string, orderID string, qty int, price float64, triggerPrice float64) (query string, resp []byte, result StopOrder, err error) {
	var cResult StopOrderResponse
	params := map[string]interface{}{}
	params["stop_order_id"] = orderID
	params["symbol"] = symbol
	if qty > 0 {
		params["p_r_qty"] = qty
	}
	if price > 0 {
		params["p_r_price"] = price
	}
	if triggerPrice > 0 {
		params["p_r_trigger_price"] = triggerPrice
	}
	query, resp, err = b.SignedRequest(http.MethodPost, "v2/private/stop-order/replace", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	result.StopOrderId = cResult.Result.StopOrderId
	return
}

// CancelStopOrder
func (b *ByBit) CancelStopOrder(orderID string, symbol string) (query string, resp []byte, result StopOrder, err error) {
	var cResult StopOrderResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["stop_order_id"] = orderID
	query, resp, err = b.SignedRequest(http.MethodPost, "v2/private/stop-order/cancel", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	result.StopOrderId = cResult.Result.StopOrderId
	return
}

// CancelAllStopOrders
func (b *ByBit) CancelAllStopOrders(symbol string) (query string, resp []byte, result []StopOrder, err error) {
	var cResult StopOrderArrayResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	query, resp, err = b.SignedRequest(http.MethodPost, "v2/private/stop-order/cancelAll", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	result = cResult.Result
	return
}
