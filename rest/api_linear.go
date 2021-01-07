package rest

import (
	"fmt"
	"net/http"
)

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

// GetOrders
func (b *ByBit) LinearGetOrders(symbol string, orderStatus string, limit int, page int) (query string, resp []byte, result OrderListResponsePaginated, err error) {
	var cResult OrderListResponsePaginated
	if limit == 0 {
		limit = 20
	}
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["page"] = page
	params["limit"] = limit
	if orderStatus != "" {
		params["order_status"] = orderStatus
	}
	query, resp, err = b.SignedRequest(http.MethodGet, "private/linear/order/list", params, &cResult)
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

// GetActiveOrders
func (b *ByBit) LinearGetActiveOrders(symbol string) (query string, resp []byte, result OrderArrayResponse, err error) {
	var cResult OrderArrayResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	query, resp, err = b.SignedRequest(http.MethodGet, "private/linear/order/search", params, &cResult)
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
func (b *ByBit) LinearCreateOrder(side string, orderType string, price float64,
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
	query, resp, err = b.SignedRequest(http.MethodPost, "private/linear/order/create", params, &cResult)
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
func (b *ByBit) LinearReplaceOrder(symbol string, orderID string, qty int, price float64) (query string, resp []byte, result Order, err error) {
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
	query, resp, err = b.SignedRequest(http.MethodPost, "private/linear/order/replace", params, &cResult)
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
func (b *ByBit) LinearCancelOrder(orderID string, symbol string) (query string, resp []byte, result Order, err error) {
	var cResult OrderResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	if orderID != "" {
		params["order_id"] = orderID
	}
	query, resp, err = b.SignedRequest(http.MethodPost, "private/linear/order/cancel", params, &cResult)
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
func (b *ByBit) LinearCancelAllOrder(symbol string) (query string, resp []byte, result []string, err error) {
	var cResult ResultStringArrayResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	query, resp, err = b.SignedRequest(http.MethodPost, "private/linear/order/cancel-all", params, &cResult)
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
func (b *ByBit) LinearGetStopOrders(symbol string, stopOrderStatus string, limit int, page string) (query string, resp []byte, result StopOrderListResponseResult, err error) {
	var cResult StopOrderListResponse
	if limit == 0 {
		limit = 20
	}
	params := map[string]interface{}{}
	params["symbol"] = symbol
	if stopOrderStatus != "" {
		params["stop_order_status"] = stopOrderStatus
	}
	params["page"] = page
	params["limit"] = limit
	query, resp, err = b.SignedRequest(http.MethodGet, "private/linear/stop-order/list", params, &cResult)
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
func (b *ByBit) LinearGetActiveStopOrders(symbol string) (query string, resp []byte, result StopOrderArrayResponse, err error) {
	var cResult StopOrderArrayResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	query, resp, err = b.SignedRequest(http.MethodGet, "private/linear/order/search", params, &cResult)
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
func (b *ByBit) LinearCreateStopOrder(side string, orderType string, price float64, basePrice float64, stopPx float64,
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
	query, resp, err = b.SignedRequest(http.MethodPost, "private/linear/stop-order/create", params, &cResult)
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
func (b *ByBit) LinearReplaceStopOrder(symbol string, orderID string, qty int, price float64, triggerPrice float64) (query string, resp []byte, result StopOrder, err error) {
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
	query, resp, err = b.SignedRequest(http.MethodPost, "private/linear/stop-order/replace", params, &cResult)
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
func (b *ByBit) LinearCancelStopOrder(orderID string, symbol string) (query string, resp []byte, result StopOrder, err error) {
	var cResult StopOrderResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["stop_order_id"] = orderID
	query, resp, err = b.SignedRequest(http.MethodPost, "private/linear/stop-order/cancel", params, &cResult)
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
func (b *ByBit) LinearCancelAllStopOrders(symbol string) (query string, resp []byte, result []string, err error) {
	var cResult ResultStringArrayResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	query, resp, err = b.SignedRequest(http.MethodPost, "private/linear/stop-order/cancel-all", params, &cResult)
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
