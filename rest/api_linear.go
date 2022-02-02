package rest

import (
	"fmt"
	"net/http"
)

// LinearGetKLine
func (b *ByBit) LinearGetKLine(symbol string, interval string, from int64, limit int) (query string, resp []byte, result []OHLCLinear, err error) {
	var ret GetLinearKlineResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["interval"] = interval
	params["from"] = from
	if limit > 0 {
		params["limit"] = limit
	}
	query, resp, err = b.PublicRequest(http.MethodGet, "public/linear/kline", params, &ret)
	if err != nil {
		return
	}
	result = ret.Result
	return
}

// LinearGetOrders GetOrders
// orderStatus:
// Created - order has been accepted by the system but not yet put through the matching engine
// Rejected - order has been triggered but failed to be placed (e.g. due to insufficient margin)
// New - order has been placed successfully
// PartiallyFilled
// Filled
// Cancelled
// PendingCancel - matching engine has received the cancelation request but it may not be canceled successfully
func (b *ByBit) LinearGetOrders(symbol string, orderStatus string, limit int, page int) (query string, resp []byte, result OrderListResponseResultPaginated, err error) {
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
	result = cResult.Result
	return
}

// LinearGetActiveOrders Query real-time active order information. If only order_id or order_link_id are passed, a single order will be returned; otherwise, returns up to 500 unfilled orders.
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

// LinearGetActiveOrder Query real-time active order information. If only order_id or order_link_id are passed, a single order will be returned; otherwise, returns up to 500 unfilled orders.
func (b *ByBit) LinearGetActiveOrder(symbol string, orderId string, orderLinkId string) (query string, resp []byte, result OrderResponse, err error) {
	var cResult OrderResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	if orderId != "" {
		params["order_id"] = orderId
	}
	if orderLinkId != "" {
		params["order_link_id"] = orderLinkId
	}
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

// LinearCreateOrder CreateOrder
// side: Buy/Sell
// orderType: Limit/Market
// timeInForce: GoodTillCancel/ImmediateOrCancel/FillOrKill/PostOnly
func (b *ByBit) LinearCreateOrder(side string, orderType string, price float64,
	qty float64, timeInForce string, takeProfit float64, stopLoss float64, reduceOnly bool,
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
	params["reduce_only"] = reduceOnly
	params["close_on_trigger"] = closeOnTrigger
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
	// {"ret_code":0,"ret_msg":"OK","ext_code":"","ext_info":"","result":{"order_id":"6f771a91-0f4e-4c01-973d-b58e6390ece0","user_id":443679,"symbol":"BTCUSDT","side":"Buy","order_type":"Limit","price":37927.5,"qty":1,"time_in_force":"GoodTillCancel","order_status":"Created","last_exec_price":0,"cum_exec_qty":0,"cum_exec_value":0,"cum_exec_fee":0,"reduce_only":false,"close_on_trigger":false,"order_link_id":"","created_time":"2022-01-25T02:06:25Z","updated_time":"2022-01-25T02:06:25Z","take_profit":0,"stop_loss":0,"tp_trigger_by":"UNKNOWN","sl_trigger_by":"UNKNOWN","position_idx":1},"time_now":"1643076385.967696","rate_limit_status":99,"rate_limit_reset_ms":1643076385963,"rate_limit":100}
	result = cResult.Result
	return
}

// LinearReplaceOrder Replace order can modify/amend your active orders.
func (b *ByBit) LinearReplaceOrder(symbol string, orderID string, orderLinkId string, qty float64, price float64,
	takeProfit float64, stopLoss float64, tpTriggerBy string, slTriggerBy string) (query string, resp []byte, orderId string, err error) {
	var cResult OrderResponse
	params := map[string]interface{}{}
	if orderID != "" {
		params["order_id"] = orderID
	}
	if orderLinkId != "" {
		params["order_link_id"] = orderLinkId
	}
	params["symbol"] = symbol
	if qty > 0 {
		params["p_r_qty"] = qty
	}
	if price > 0 {
		params["p_r_price"] = price
	}
	if takeProfit > 0 {
		params["take_profit"] = takeProfit
	}
	if stopLoss > 0 {
		params["stop_loss"] = stopLoss
	}
	if tpTriggerBy != "" {
		params["tp_trigger_by"] = tpTriggerBy
	}
	if slTriggerBy != "" {
		params["sl_trigger_by"] = slTriggerBy
	}
	query, resp, err = b.SignedRequest(http.MethodPost, "private/linear/order/replace", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	orderId = cResult.Result.OrderId
	return
}

// LinearCancelOrder
// orderID: Order ID. Required if not passing order_link_id
// orderLinkId: Unique user-set order ID. Required if not passing order_id
func (b *ByBit) LinearCancelOrder(orderID string, orderLinkId string, symbol string) (query string, resp []byte, result Order, err error) {
	var cResult OrderResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	if orderID != "" {
		params["order_id"] = orderID
	}
	if orderLinkId != "" {
		params["order_link_id"] = orderLinkId
	}
	query, resp, err = b.SignedRequest(http.MethodPost, "private/linear/order/cancel", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	// {"ret_code":0,"ret_msg":"OK","ext_code":"","ext_info":"","result":{"order_id":"d328974d-bfe8-484f-a0e9-30159bc78aaf"},"time_now":"1643077335.069762","rate_limit_status":99,"rate_limit_reset_ms":1643077335056,"rate_limit":100}
	result = cResult.Result
	return
}

// LinearCancelAllOrder
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

// LinearGetStopOrders
func (b *ByBit) LinearGetStopOrders(symbol string, stopOrderStatus string, limit int, page int) (query string, resp []byte, result StopOrderListResponseResult, err error) {
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

// LinearGetActiveStopOrders
func (b *ByBit) LinearGetActiveStopOrders(symbol string) (query string, resp []byte, result StopOrderArrayResponse, err error) {
	var cResult StopOrderArrayResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	query, resp, err = b.SignedRequest(http.MethodGet, "private/linear/stop-order/search", params, &cResult)
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
	qty float64, triggerBy string, timeInForce string, closeOnTrigger bool, symbol string, reduceOnly bool) (query string, resp []byte, result StopOrder, err error) {
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
	params["close_on_trigger"] = closeOnTrigger
	params["reduce_only"] = reduceOnly
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
func (b *ByBit) LinearReplaceStopOrder(symbol string, orderID string, qty float64, price float64, triggerPrice float64) (query string, resp []byte, result StopOrder, err error) {
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

// LinearGetPositions
func (b *ByBit) LinearGetPositions() (query string, resp []byte, result []LinearPositionData, err error) {
	var r LinearPositionDataArrayResponse
	params := map[string]interface{}{}
	query, resp, err = b.SignedRequest(http.MethodGet, "private/linear/position/list", params, &r)
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

// LinearGetPosition
func (b *ByBit) LinearGetPosition(symbol string) (query string, resp []byte, result []LinearPosition, err error) {
	var r LinearPositionArrayResponse
	params := map[string]interface{}{}
	if symbol != "" {
		params["symbol"] = symbol
	}
	query, resp, err = b.SignedRequest(http.MethodGet, "private/linear/position/list", params, &r)
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
