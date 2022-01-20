package rest

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestByBit_GetOrderBook(t *testing.T) {
	b := newByBit()
	_, _, ob, err := b.GetOrderBook("BTCUSD")
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range ob.Asks {
		t.Logf("Ask: %#v", v)
	}
	for _, v := range ob.Bids {
		t.Logf("Bid: %#v", v)
	}
	t.Logf("%v", ob.Time)
}

func TestByBit_GetOrderBook2(t *testing.T) {
	b := newByBit()
	_, _, ob, err := b.GetOrderBook("BTCUSDT")
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range ob.Asks {
		t.Logf("Ask: %#v", v)
	}
	for _, v := range ob.Bids {
		t.Logf("Bid: %#v", v)
	}
	t.Logf("%v", ob.Time)
}

func TestByBit_GetKLine(t *testing.T) {
	b := newByBit()
	from := time.Now().Add(-1 * time.Hour).Unix()
	_, _, ohlcs, err := b.GetKLine(
		"BTCUSD",
		"1",
		from,
		0,
	)
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range ohlcs {
		t.Logf("%#v", v)
	}
}

func TestByBit_GetOrders(t *testing.T) {
	b := newByBit()
	symbol := "BTCUSD"
	_, _, orders, err := b.GetOrders(symbol, "", "next", 20, "")
	assert.Nil(t, err)
	//t.Logf("%#v", orders)
	for _, order := range orders.Data {
		t.Logf("%#v", order)
	}
}

func TestByBit_CreateOrder(t *testing.T) {
	b := newByBit()
	symbol := "BTCUSD"
	side := "Buy" // Buy Sell
	orderType := "Limit"
	qty := 30
	price := 5000.0
	timeInForce := "GoodTillCancel"
	_, _, order, err := b.CreateOrder(
		side,
		orderType,
		price,
		qty,
		timeInForce,
		0,
		0,
		false,
		false,
		"",
		symbol)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", order)
}

func TestByBit_CancelOrder(t *testing.T) {
	b := newByBit()
	orderID := "c5b96b82-6a79-4b15-a797-361fe2ca0260"
	symbol := "BTCUSD"
	_, _, order, err := b.CancelOrder(orderID, symbol)
	assert.Nil(t, err)
	t.Logf("%#v", order)
}

func TestByBit_GetStopOrders(t *testing.T) {
	b := newByBit()
	symbol := "BTCUSD"
	status := "Untriggered,Triggered,Active"
	_, _, result, err := b.GetStopOrders(symbol, status, "next", 20, "")
	assert.Nil(t, err)
	//t.Logf("%#v", orders)
	for _, order := range result.Data {
		//if order.ExtFields != nil {
		//	t.Logf("%#v %v", order, *order.ExtFields)
		//} else {
		t.Logf("CreatedAt: %v %#v", order.CreatedAt.Local(), order)
		//}
	}
}

func TestByBit_CreateStopOrder(t *testing.T) {
	b := newByBit()
	symbol := "BTCUSD"
	side := "Buy" // Buy Sell
	orderType := "Limit"
	qty := 30
	price := 10000.0
	basePrice := 7100.0
	stopPx := 10000.0
	triggerBy := ""
	timeInForce := "GoodTillCancel"
	// {"ret_code":0,"ret_msg":"ok","ext_code":"","result":{"user_id":103061,"symbol":"BTCUSD","side":"Buy","order_type":"Limit","price":"7000","qty":30,"time_in_force":"GoodTillCancel","order_status":"Created","ext_fields":{"cross_status":"PendingNew","xreq_type":"x_create","xreq_offset":148672558},"leaves_qty":30,"leaves_value":"0.00428571","reject_reason":"","cross_seq":-1,"created_at":"2019-07-23T08:54:54.000Z","updated_at":"2019-07-23T08:54:54.000Z","last_exec_time":"0.000000","last_exec_price":0,"order_id":"603c41e0-c9fb-450c-90b6-ea870d5b0180"},"ext_info":null,"time_now":"1563872094.895918","rate_limit_status":98}
	_, _, order, err := b.CreateStopOrder(side,
		orderType, price, basePrice, stopPx, qty, triggerBy, timeInForce, true, symbol)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", order)
}

func TestByBit_CancelStopOrder(t *testing.T) {
	b := newByBit()
	orderID := "c6e535a9-6900-4b64-b983-3b220f6f41f8"
	symbol := "BTCUSD"
	_, _, order, err := b.CancelStopOrder(orderID, symbol)
	assert.Nil(t, err)
	t.Logf("%#v", order)
}

func TestByBit_CancelAllStopOrders(t *testing.T) {
	b := newByBit()
	symbol := "BTCUSD"
	_, _, orders, err := b.CancelAllStopOrders(symbol)
	assert.Nil(t, err)
	t.Logf("%#v", orders)
}
