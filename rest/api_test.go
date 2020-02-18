package rest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 开发电报群:
// https://t.me/Bybitapi

func newByBit() *ByBit {
	//baseURL := "https://api.bybit.com/"
	baseURL := "https://api-testnet.bybit.com/"
	b := New(baseURL, "6IASD6KDBdunn5qLpT", "nXjZMUiB3aMiPaQ9EUKYFloYNd0zM39RjRWF")
	return b
}

func TestByBit_GetWalletBalance(t *testing.T) {
	b := newByBit()
	balance, err := b.GetWalletBalance("BTC")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", balance)
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
	//t.Logf("%#v", orders)
	for _, order := range orders {
		if order.ExtFields != nil {
			t.Logf("%#v %v", order, *order.ExtFields)
		} else {
			t.Logf("%#v", order)
		}
	}
}

func TestByBit_CancelOrder(t *testing.T) {
	b := newByBit()
	orderID := "c5b96b82-6a79-4b15-a797-361fe2ca0260"
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
