package rest

import (
	"fmt"
	"net/http"
)

// getOrders 查询活动委托
func (b *ByBit) GetOrders(symbol string, orderStatus string, direction string, limit int, cursor string) (result OrderListResponseResult, err error) {
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
	var resp []byte
	resp, err = b.SignedRequest(http.MethodGet, "v2/private/order/list", params, &cResult)
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

func (b *ByBit) CreateOrder(side string, orderType string, price float64,
	qty int, timeInForce string, takeProfit float64, stopLoss float64, reduceOnly bool,
	closeOnTrigger bool, orderLinkID string, symbol string) (result Order, err error) {
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
	var resp []byte
	resp, err = b.SignedRequest(http.MethodPost, "v2/private/order/create", params, &cResult)
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
func (b *ByBit) ReplaceOrder(symbol string, orderID string, qty int, price float64) (result Order, err error) {
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
	var resp []byte
	resp, err = b.SignedRequest(http.MethodPost, "v2/private/order/replace", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", cResult.RetMsg, string(resp))
		return
	}
	result.OrderID = cResult.Result.OrderID
	return
}

// CancelOrder 撤销活动委托单
func (b *ByBit) CancelOrder(orderID string, symbol string) (result Order, err error) {
	var cResult OrderResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	if orderID != "" {
		params["order_id"] = orderID
	}
	var resp []byte
	resp, err = b.SignedRequest(http.MethodPost, "v2/private/order/cancel", params, &cResult)
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

// CancelAllOrder Cancel All Active Orders
func (b *ByBit) CancelAllOrder(symbol string) (result []Order, err error) {
	var cResult OrderArrayResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	var resp []byte
	resp, err = b.SignedRequest(http.MethodPost, "v2/private/order/cancelAll", params, &cResult)
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

// getStopOrders 查询条件委托单
func (b *ByBit) GetStopOrders(symbol string, stopOrderStatus string, direction string, limit int, cursor string) (result StopOrderListResponseResult, err error) {
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
	var resp []byte
	resp, err = b.SignedRequest(http.MethodGet, "v2/private/stop-order/list", params, &cResult)
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

// CreateStopOrder 创建条件委托单
func (b *ByBit) CreateStopOrder(side string, orderType string, price float64, basePrice float64, stopPx float64,
	qty int, triggerBy string, timeInForce string, closeOnTrigger bool, symbol string) (result StopOrder, err error) {
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
	var resp []byte
	resp, err = b.SignedRequest(http.MethodPost, "v2/private/stop-order/create", params, &cResult)
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
func (b *ByBit) ReplaceStopOrder(symbol string, orderID string, qty int, price float64, triggerPrice float64) (result StopOrder, err error) {
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
	var resp []byte
	resp, err = b.SignedRequest(http.MethodPost, "v2/private/stop-order/replace", params, &cResult)
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

// CancelStopOrder 撤销活动条件委托单
func (b *ByBit) CancelStopOrder(orderID string, symbol string) (result StopOrder, err error) {
	var cResult StopOrderResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["stop_order_id"] = orderID
	var resp []byte
	resp, err = b.SignedRequest(http.MethodPost, "v2/private/stop-order/cancel", params, &cResult)
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

// CancelAllStopOrders 撤消全部条件委托单
func (b *ByBit) CancelAllStopOrders(symbol string) (result []StopOrder, err error) {
	var cResult StopOrderArrayResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	var resp []byte
	resp, err = b.SignedRequest(http.MethodPost, "v2/private/stop-order/cancelAll", params, &cResult)
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
