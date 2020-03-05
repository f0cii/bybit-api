package rest

import (
	"errors"
	"net/http"
)

func (b *ByBit) CreateOrderV2(side string, orderType string, price float64,
	qty int, timeInForce string, takeProfit float64, stopLoss float64, reduceOnly bool,
	closeOnTrigger bool, orderLinkID string, symbol string) (result OrderV2, err error) {
	var cResult CreateOrderV2Result
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
	err = b.SignedRequest(http.MethodPost, "v2/private/order/create", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = errors.New(cResult.RetMsg)
		return
	}
	result = cResult.Result
	return
}

// CreateOrder 创建委托单
// symbol: 产品类型, 有效选项:BTCUSD,ETHUSD (BTCUSD ETHUSD)
// side: 方向, 有效选项:Buy, Sell (Buy Sell)
// orderType: Limit/Market
// price: 委托价格, 在没有仓位时，做多的委托价格需高于市价的10%、低于1百万。如有仓位时则需优于强平价。单笔价格增减最小单位为0.5。
// qty: 委托数量, 单笔最大1百万
// timeInForce: 执行策略, 有效选项:GoodTillCancel,ImmediateOrCancel,FillOrKill,PostOnly
// reduceOnly: 只减仓
// symbol: 产品类型, 有效选项:BTCUSD,ETHUSD (BTCUSD ETHUSD)
func (b *ByBit) CreateOrder(side string, orderType string, price float64, qty int, timeInForce string, reduceOnly bool, symbol string) (result Order, err error) {
	var cResult CreateOrderResult
	params := map[string]interface{}{}
	params["side"] = side
	params["symbol"] = symbol
	params["order_type"] = orderType
	params["qty"] = qty
	params["price"] = price
	params["time_in_force"] = timeInForce
	if reduceOnly {
		params["reduce_only"] = true
	}
	err = b.SignedRequest(http.MethodPost, "open-api/order/create", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = errors.New(cResult.RetMsg)
		return
	}
	result = cResult.Result
	return
}

func (b *ByBit) ReplaceOrder(symbol string, orderID string, qty int, price float64) (result Order, err error) {
	var cResult ReplaceOrderResult
	params := map[string]interface{}{}
	params["order_id"] = orderID
	params["symbol"] = symbol
	if qty > 0 {
		params["p_r_qty"] = qty
	}
	if price > 0 {
		params["p_r_price"] = price
	}
	err = b.SignedRequest(http.MethodPost, "open-api/order/replace", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = errors.New(cResult.RetMsg)
		return
	}
	result.OrderID = cResult.Result.OrderID
	return
}

// CreateStopOrder 创建条件委托单
// https://github.com/bybit-exchange/bybit-official-api-docs/blob/master/zh_cn/rest_api.md#open-apistop-ordercreatepost
// symbol: 产品类型, 有效选项:BTCUSD,ETHUSD (BTCUSD ETHUSD)
// side: 方向, 有效选项:Buy, Sell (Buy Sell)
// orderType: Limit/Market
// price: 委托价格, 在没有仓位时，做多的委托价格需高于市价的10%、低于1百万。如有仓位时则需优于强平价。单笔价格增减最小单位为0.5。
// qty: 委托数量, 单笔最大1百万
// basePrice: 当前市价。用于和stop_px值进行比较，确定当前条件委托是看空到stop_px时触发还是看多到stop_px触发。主要是用来标识当前条件单预期的方向
// stopPx: 条件委托下单时市价
// triggerBy: 触发价格类型. 默认为上一笔成交价格
// timeInForce: 执行策略, 有效选项:GoodTillCancel,ImmediateOrCancel,FillOrKill,PostOnly
// reduceOnly: 只减仓
// symbol: 产品类型, 有效选项:BTCUSD,ETHUSD (BTCUSD ETHUSD)
func (b *ByBit) CreateStopOrder(side string, orderType string, price float64, basePrice float64, stopPx float64,
	qty int, triggerBy string, timeInForce string, reduceOnly bool, symbol string) (result Order, err error) {
	var cResult CreateOrderResult
	params := map[string]interface{}{}
	params["side"] = side
	params["symbol"] = symbol
	params["order_type"] = orderType
	params["qty"] = qty
	params["price"] = price
	params["base_price"] = basePrice
	params["stop_px"] = stopPx
	params["time_in_force"] = timeInForce
	if reduceOnly {
		params["reduce_only"] = true
	}
	if triggerBy != "" {
		params["trigger_by"] = triggerBy
	}
	err = b.SignedRequest(http.MethodPost, "open-api/stop-order/create", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = errors.New(cResult.RetMsg)
		return
	}
	result = cResult.Result
	return
}

// GetOrders 查询活动委托
// symbol
// orderID: 订单ID
// orderLinkID: 机构自定义订单ID
// sort: 排序字段，默认按创建时间排序 (created_at cum_exec_qty qty last_exec_price price cum_exec_value cum_exec_fee)
// order: 升序降序， 默认降序 (desc asc)
// page: 页码，默认取第一页数据
// limit: 一页数量，一页默认展示20条数据
func (b *ByBit) GetOrders(sort string, order string, page int,
	limit int, orderStatus string, symbol string) (result []Order, err error) {
	return b.getOrders("", "", sort, order, page, limit, orderStatus, symbol)
}

// getOrders 查询活动委托
// symbol
// orderID: 订单ID
// orderLinkID: 机构自定义订单ID
// sort: 排序字段，默认按创建时间排序 (created_at cum_exec_qty qty last_exec_price price cum_exec_value cum_exec_fee)
// order: 升序降序， 默认降序 (desc asc)
// page: 页码，默认取第一页数据
// limit: 一页数量，一页默认展示20条数据
func (b *ByBit) getOrders(orderID string, orderLinkID string, sort string, order string, page int,
	limit int, orderStatus string, symbol string) (result []Order, err error) {
	var cResult OrderListResult

	if limit == 0 {
		limit = 20
	}

	params := map[string]interface{}{}
	params["symbol"] = symbol
	if orderID != "" {
		params["order_id"] = orderID
	}
	if orderLinkID != "" {
		params["order_link_id"] = orderLinkID
	}
	if sort != "" {
		params["sort"] = sort
	}
	if order != "" {
		params["order"] = order
	}
	params["page"] = page
	params["limit"] = limit
	if orderStatus != "" {
		params["order_status"] = orderStatus
	}
	err = b.SignedRequest(http.MethodGet, "open-api/order/list", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = errors.New(cResult.RetMsg)
		return
	}

	result = cResult.Result.Data
	return
}

// GetStopOrders 查询条件委托单
// orderID: 条件委托单ID
// orderLinkID: 机构自定义订单ID
// order: 排序字段为created_at,升序降序，默认降序 (desc asc )
// page: 页码，默认取第一页数据
// stopOrderStatus 条件单状态: Untriggered: 等待市价触发条件单; Triggered: 市价已触发条件单; Cancelled: 取消; Active: 条件单触发成功且下单成功; Rejected: 条件触发成功但下单失败
// limit: 一页数量，默认一页展示20条数据;最大支持50条每页
func (b *ByBit) GetStopOrders(orderID string, orderLinkID string, stopOrderStatus string, order string,
	page int, limit int, symbol string) (result []Order, err error) {
	var cResult OrderListResult

	if limit == 0 {
		limit = 20
	}

	params := map[string]interface{}{}
	params["symbol"] = symbol
	if orderID != "" {
		params["stop_order_id"] = orderID
	}
	if orderLinkID != "" {
		params["order_link_id"] = orderLinkID
	}
	if stopOrderStatus != "" {
		params["stop_order_status"] = stopOrderStatus
	}
	if order != "" {
		params["order"] = order
	}
	params["page"] = page
	params["limit"] = limit
	err = b.SignedRequest(http.MethodGet, "open-api/stop-order/list", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = errors.New(cResult.RetMsg)
		return
	}

	result = cResult.Result.Data
	return
}

// GetOrderByID
func (b *ByBit) GetOrderByID(orderID string, symbol string) (result Order, err error) {
	var orders []Order
	orders, err = b.getOrders(orderID, "", "", "", 0, 20, "", symbol)
	if err != nil {
		return
	}
	if len(orders) != 1 {
		err = errors.New("not found")
		return
	}
	result = orders[0]
	return
}

// GetOrderByOrderLinkID ...
func (b *ByBit) GetOrderByOrderLinkID(orderLinkID string, symbol string) (result Order, err error) {
	var orders []Order
	orders, err = b.getOrders("", orderLinkID, "", "", 0, 20, "", symbol)
	if err != nil {
		return
	}
	if len(orders) != 1 {
		err = errors.New("not found")
		return
	}
	result = orders[0]
	return
}

// CancelOrder 撤销活动委托单
// orderID: 活动委托单ID, 数据来自创建活动委托单返回的订单唯一ID
// symbol:
func (b *ByBit) CancelOrder(orderID string, symbol string) (result Order, err error) {
	var cResult CancelOrderResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["order_id"] = orderID
	err = b.SignedRequest(http.MethodPost, "open-api/order/cancel", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = errors.New(cResult.RetMsg)
		return
	}

	result = cResult.Result
	return
}

// CancelOrder 撤销活动委托单
// orderID: 活动委托单ID, 数据来自创建活动委托单返回的订单唯一ID
// symbol:
func (b *ByBit) CancelOrderV2(orderID string, orderLinkID string, symbol string) (result OrderV2, err error) {
	var cResult CancelOrderV2Result
	params := map[string]interface{}{}
	params["symbol"] = symbol
	if orderID != "" {
		params["order_id"] = orderID
	}
	if orderLinkID != "" {
		params["order_link_id"] = orderLinkID
	}
	err = b.SignedRequest(http.MethodPost, "v2/private/order/cancel", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = errors.New(cResult.RetMsg)
		return
	}

	result = cResult.Result
	return
}

// CancelAllOrder Cancel All Active Orders
func (b *ByBit) CancelAllOrder(symbol string) (result []OrderV2, err error) {
	var cResult CancelAllOrderV2Result
	params := map[string]interface{}{}
	params["symbol"] = symbol
	err = b.SignedRequest(http.MethodPost, "v2/private/order/cancelAll", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = errors.New(cResult.RetMsg)
		return
	}

	result = cResult.Result
	return
}

// CancelStopOrder 撤销活动条件委托单
// orderID: 活动条件委托单ID, 数据来自创建活动委托单返回的订单唯一ID
// symbol:
func (b *ByBit) CancelStopOrder(orderID string, symbol string) (result Order, err error) {
	var cResult CancelOrderResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["stop_order_id"] = orderID
	err = b.SignedRequest(http.MethodPost, "open-api/stop-order/cancel", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = errors.New(cResult.RetMsg)
		return
	}

	result = cResult.Result
	return
}
