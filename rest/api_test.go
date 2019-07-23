package rest

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 开发电报群:
// https://t.me/Bybitapi

func newByBit() *ByBit {
	//baseURL := "https://api.bybit.com/"
	baseURL := "https://api-testnet.bybit.com/"
	b := New(baseURL, "YIxOY2RhFkylPudq96", "Bg9G2oFOb3aaIMguD3FOvOJJVBycaoXqXNcI")
	return b
}

func TestByBit_GetSymbols(t *testing.T) {
	//b := newByBit()
	//b.GetSymbols()
}

func TestByBit_CreateOrder(t *testing.T) {
	b := newByBit()
	symbol := "BTCUSD"
	side := "Buy" // Buy Sell
	orderType := "Limit"
	qty := 30
	price := 7000.0
	timeInForce := "GoodTillCancel"
	// {"ret_code":0,"ret_msg":"ok","ext_code":"","result":{"user_id":103061,"symbol":"BTCUSD","side":"Buy","order_type":"Limit","price":"7000","qty":30,"time_in_force":"GoodTillCancel","order_status":"Created","ext_fields":{"cross_status":"PendingNew","xreq_type":"x_create","xreq_offset":148672558},"leaves_qty":30,"leaves_value":"0.00428571","reject_reason":"","cross_seq":-1,"created_at":"2019-07-23T08:54:54.000Z","updated_at":"2019-07-23T08:54:54.000Z","last_exec_time":"0.000000","last_exec_price":0,"order_id":"603c41e0-c9fb-450c-90b6-ea870d5b0180"},"ext_info":null,"time_now":"1563872094.895918","rate_limit_status":98}
	order, err := b.CreateOrder(side, orderType, price, qty, timeInForce, symbol)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", order)
	// Created:创建订单;Rejected:订单被拒绝;New:订单待成交;PartiallyFilled:订单部分成交;Filled:订单全部成交,Cancelled:订单被取消
}

func TestByBit_GetOrders(t *testing.T) {
	b := newByBit()
	symbol := "BTCUSD"
	orders, err := b.GetOrders("", "", 0, 20, "", symbol)
	assert.Nil(t, err)
	t.Logf("%#v", orders)
}

func TestByBit_CancelOrder(t *testing.T) {
	b := newByBit()
	orderID := "cce81de0-5fcd-4367-a627-419a973e4770"
	symbol := "BTCUSD"
	order, err := b.CancelOrder(orderID, symbol)
	assert.Nil(t, err)
	t.Logf("%#v", order)
}

func TestByBit_GetLeverages(t *testing.T) {
	b := newByBit()
	l, err := b.GetLeverages()
	assert.Nil(t, err)
	t.Logf("%#v", l)
}

func TestByBit_SetLeverage(t *testing.T) {
	b := newByBit()
	b.SetLeverage(3, "BTCUSD")
}

func TestByBit_GetPositions(t *testing.T) {
	b := newByBit()
	positions, err := b.GetPositions()
	assert.Nil(t, err)
	t.Logf("%#v", positions)
}

func TestOrderListResult(t *testing.T) {
	j := `{"ret_code":0,"ret_msg":"ok","ext_code":"","result":{"current_page":1,"last_page":1,"data":[{"user_id":103061,"symbol":"BTCUSD","side":"Buy","order_type":"Limit","price":7000,"qty":30,"time_in_force":"GoodTillCancel","order_status":"Cancelled","ext_fields":{"cross_status":"Canceled"},"last_exec_time":"0.000000","last_exec_price":0,"leaves_qty":0,"leaves_value":0,"cum_exec_qty":0,"cum_exec_value":0,"cum_exec_fee":0,"reject_reason":"EC_PerCancelRequest","order_link_id":"","created_at":"2019-07-23T08:12:30.000Z","updated_at":"2019-07-23T08:13:34.000Z","order_id":"f9617bb6-f748-4604-ad7d-a9b7f1433ff8"},{"user_id":103061,"symbol":"BTCUSD","side":"Buy","order_type":"Limit","price":7000,"qty":30,"time_in_force":"GoodTillCancel","order_status":"Cancelled","ext_fields":{"cross_status":"Canceled"},"last_exec_time":"0.000000","last_exec_price":0,"leaves_qty":0,"leaves_value":0,"cum_exec_qty":0,"cum_exec_value":0,"cum_exec_fee":0,"reject_reason":"EC_PerCancelRequest","order_link_id":"","created_at":"2019-07-23T08:07:36.000Z","updated_at":"2019-07-23T08:07:59.000Z","order_id":"43b55540-3bd5-46bf-b89a-4ab0447797f9"},{"user_id":103061,"symbol":"BTCUSD","side":"Buy","order_type":"Limit","price":7000,"qty":30,"time_in_force":"GoodTillCancel","order_status":"Cancelled","ext_fields":{"cross_status":"Canceled"},"last_exec_time":"0.000000","last_exec_price":0,"leaves_qty":0,"leaves_value":0,"cum_exec_qty":0,"cum_exec_value":0,"cum_exec_fee":0,"reject_reason":"EC_PerCancelRequest","order_link_id":"","created_at":"2019-07-23T07:26:55.000Z","updated_at":"2019-07-23T08:02:56.000Z","order_id":"9b4ce60e-bc77-4c3e-8b61-1eeed7e3edb7"},{"user_id":103061,"symbol":"BTCUSD","side":"Buy","order_type":"Limit","price":7000,"qty":30,"time_in_force":"GoodTillCancel","order_status":"New","ext_fields":[],"last_exec_time":"0.000000","last_exec_price":0,"leaves_qty":30,"leaves_value":0.00428571,"cum_exec_qty":0,"cum_exec_value":0,"cum_exec_fee":0,"reject_reason":"","order_link_id":"","created_at":"2019-07-23T06:43:59.000Z","updated_at":"2019-07-23T06:43:59.000Z","order_id":"29d6ac55-44f1-45c4-9c0a-9685e79fcc55"},{"user_id":103061,"symbol":"BTCUSD","side":"Sell","order_type":"Market","price":10482,"qty":1,"time_in_force":"ImmediateOrCancel","order_status":"Filled","ext_fields":{"close_on_trigger":true,"orig_order_type":"BLimit","prior_x_req_price":"5250.5"},"last_exec_time":"1563788170.642047","last_exec_price":10481.5,"leaves_qty":0,"leaves_value":0,"cum_exec_qty":1,"cum_exec_value":9.54e-5,"cum_exec_fee":8.0e-8,"reject_reason":"","order_link_id":"","created_at":"2019-07-22T09:36:10.000Z","updated_at":"2019-07-22T09:36:10.000Z","order_id":"1d8ef634-6ec4-4b50-aff7-850fa882ab85"},{"user_id":103061,"symbol":"BTCUSD","side":"Buy","order_type":"Market","price":10500.5,"qty":1,"time_in_force":"ImmediateOrCancel","order_status":"Filled","ext_fields":[],"last_exec_time":"1563787508.834188","last_exec_price":10500.5,"leaves_qty":0,"leaves_value":0,"cum_exec_qty":1,"cum_exec_value":9.523e-5,"cum_exec_fee":8.0e-8,"reject_reason":"","order_link_id":"","created_at":"2019-07-22T09:25:08.000Z","updated_at":"2019-07-22T09:25:08.000Z","order_id":"02ebb7b2-32ea-4961-9786-c22b474490f5"}]},"ext_info":null,"time_now":"1563869744.317030","rate_limit_status":98}`
	var c OrderListResult
	err := json.Unmarshal([]byte(j), &c)
	assert.Nil(t, err)
}

func TestCancelOrderResult(t *testing.T) {
	//j := `{"ret_code":30037,"ret_msg":"Order already cancelled","ext_code":"","result":null,"ext_info":null,"time_now":"1563869179.061704","rate_limit_status":97}`
	//j := `{"ret_code":0,"ret_msg":"ok","ext_code":"","result":{"user_id":103061,"symbol":"BTCUSD","side":"Buy","order_type":"Limit","price":"7000","qty":30,"time_in_force":"GoodTillCancel","order_status":"New","last_exec_time":"0","last_exec_price":"0","leaves_qty":30,"leaves_value":"0.00428571","cum_exec_qty":0,"cum_exec_value":"0","cum_exec_fee":"0","reject_reason":"","created_at":"2019-07-23T08:07:36.000Z","updated_at":"2019-07-23T08:07:59.000Z","ext_fields":{"cross_status":"PendingCancel","xreq_type":"x_cancel","xreq_offset":148624672},"order_id":"43b55540-3bd5-46bf-b89a-4ab0447797f9"},"ext_info":null,"time_now":"1563869279.239529","rate_limit_status":98}`
	j := `{"ret_code":30037,"ret_msg":"Order already cancelled","ext_code":"","result":null,"ext_info":null,"time_now":"1563869505.166926","rate_limit_status":99}`
	var c CancelOrderResult
	err := json.Unmarshal([]byte(j), &c)
	assert.Nil(t, err)
}
