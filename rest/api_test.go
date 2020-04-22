package rest

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

// 开发电报群:
// https://t.me/Bybitapi

func newByBit() *ByBit {
	//baseURL := "https://api.bybit.com/"
	baseURL := "https://api-testnet.bybit.com/"
	apiKey := "6IASD6KDBdunn5qLpT"
	secretKey := "nXjZMUiB3aMiPaQ9EUKYFloYNd0zM39RjRWF"
	b := New(nil, baseURL, apiKey, secretKey, true)
	err := b.SetCorrectServerTime()
	if err != nil {
		log.Printf("%v", err)
	}
	return b
}

func TestByBit_GetServerTime(t *testing.T) {
	b := newByBit()
	timeNow, err := b.GetServerTime()
	if err != nil {
		t.Error(err)
		return
	}
	now := time.Now().UnixNano() / 1e6
	t.Logf("timeNow: %v Now: %v Diff: %v",
		timeNow,
		now,
		now-timeNow)
}

func TestByBit_SetCorrectServerTime(t *testing.T) {
	b := newByBit()
	err := b.SetCorrectServerTime()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestByBit_GetOrderBook(t *testing.T) {
	b := newByBit()
	ob, err := b.GetOrderBook("BTCUSD")
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
	ohlcs, err := b.GetKLine(
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

func TestByBit_GetTickers(t *testing.T) {
	b := newByBit()
	tickers, err := b.GetTickers()
	if err != nil {
		t.Error()
		return
	}
	for _, v := range tickers {
		t.Logf("%#v", v)
	}
}

func TestByBit_GetTradingRecords(t *testing.T) {
	b := newByBit()
	records, err := b.GetTradingRecords("BTCUSD", 0, 0)
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range records {
		t.Logf("%#v", v)
	}
}

func TestByBit_GetSymbols(t *testing.T) {
	b := newByBit()
	symbols, err := b.GetSymbols()
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range symbols {
		t.Logf("%#v", v)
	}
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

func TestByBit_CreateOrderV2(t *testing.T) {
	b := newByBit()
	symbol := "BTCUSD"
	side := "Buy" // Buy Sell
	orderType := "Limit"
	qty := 30
	price := 5000.0
	timeInForce := "GoodTillCancel"
	// {"ret_code":0,"ret_msg":"ok","ext_code":"","result":{"user_id":103061,"symbol":"BTCUSD","side":"Buy","order_type":"Limit","price":"7000","qty":30,"time_in_force":"GoodTillCancel","order_status":"Created","ext_fields":{"cross_status":"PendingNew","xreq_type":"x_create","xreq_offset":148672558},"leaves_qty":30,"leaves_value":"0.00428571","reject_reason":"","cross_seq":-1,"created_at":"2019-07-23T08:54:54.000Z","updated_at":"2019-07-23T08:54:54.000Z","last_exec_time":"0.000000","last_exec_price":0,"order_id":"603c41e0-c9fb-450c-90b6-ea870d5b0180"},"ext_info":null,"time_now":"1563872094.895918","rate_limit_status":98}
	order, err := b.CreateOrderV2(
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

func TestByBit_CreateOrder(t *testing.T) {
	b := newByBit()
	symbol := "BTCUSD"
	side := "Buy" // Buy Sell
	orderType := "Limit"
	qty := 30
	price := 7000.0
	timeInForce := "GoodTillCancel"
	// {"ret_code":0,"ret_msg":"ok","ext_code":"","result":{"user_id":103061,"symbol":"BTCUSD","side":"Buy","order_type":"Limit","price":"7000","qty":30,"time_in_force":"GoodTillCancel","order_status":"Created","ext_fields":{"cross_status":"PendingNew","xreq_type":"x_create","xreq_offset":148672558},"leaves_qty":30,"leaves_value":"0.00428571","reject_reason":"","cross_seq":-1,"created_at":"2019-07-23T08:54:54.000Z","updated_at":"2019-07-23T08:54:54.000Z","last_exec_time":"0.000000","last_exec_price":0,"order_id":"603c41e0-c9fb-450c-90b6-ea870d5b0180"},"ext_info":null,"time_now":"1563872094.895918","rate_limit_status":98}
	order, err := b.CreateOrder(side, orderType, price, qty, timeInForce, false, symbol)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", order)
	// Created:创建订单;Rejected:订单被拒绝;New:订单待成交;PartiallyFilled:订单部分成交;Filled:订单全部成交,Cancelled:订单被取消
}

func TestByBit_CreateOrder2(t *testing.T) {
	b := newByBit()
	symbol := "BTCUSD"
	side := "Sell" // Buy Sell
	orderType := "Limit"
	qty := 30
	price := 9000.0
	timeInForce := "GoodTillCancel"
	// {"ret_code":0,"ret_msg":"ok","ext_code":"","result":{"user_id":103061,"symbol":"BTCUSD","side":"Buy","order_type":"Limit","price":"7000","qty":30,"time_in_force":"GoodTillCancel","order_status":"Created","ext_fields":{"cross_status":"PendingNew","xreq_type":"x_create","xreq_offset":148672558},"leaves_qty":30,"leaves_value":"0.00428571","reject_reason":"","cross_seq":-1,"created_at":"2019-07-23T08:54:54.000Z","updated_at":"2019-07-23T08:54:54.000Z","last_exec_time":"0.000000","last_exec_price":0,"order_id":"603c41e0-c9fb-450c-90b6-ea870d5b0180"},"ext_info":null,"time_now":"1563872094.895918","rate_limit_status":98}
	order, err := b.CreateOrder(side, orderType, price, qty, timeInForce, true, symbol)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", order)
	// Created:创建订单;Rejected:订单被拒绝;New:订单待成交;PartiallyFilled:订单部分成交;Filled:订单全部成交,Cancelled:订单被取消
}

func TestByBit_CreateStopOrder(t *testing.T) {
	b := newByBit()
	symbol := "BTCUSD"
	side := "Buy" // Buy Sell
	orderType := "Limit"
	qty := 30
	price := 10000.0
	basePrice := 10000.0 // 触发价
	stopPx := 0.0
	triggerBy := ""
	timeInForce := "GoodTillCancel"
	// {"ret_code":0,"ret_msg":"ok","ext_code":"","result":{"user_id":103061,"symbol":"BTCUSD","side":"Buy","order_type":"Limit","price":"7000","qty":30,"time_in_force":"GoodTillCancel","order_status":"Created","ext_fields":{"cross_status":"PendingNew","xreq_type":"x_create","xreq_offset":148672558},"leaves_qty":30,"leaves_value":"0.00428571","reject_reason":"","cross_seq":-1,"created_at":"2019-07-23T08:54:54.000Z","updated_at":"2019-07-23T08:54:54.000Z","last_exec_time":"0.000000","last_exec_price":0,"order_id":"603c41e0-c9fb-450c-90b6-ea870d5b0180"},"ext_info":null,"time_now":"1563872094.895918","rate_limit_status":98}
	order, err := b.CreateStopOrder(side, orderType, price, stopPx, basePrice, qty, triggerBy, timeInForce, false, symbol)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", order)
}

func TestByBit_GetOrders(t *testing.T) {
	b := newByBit()
	symbol := "BTCUSD"
	orders, err := b.GetOrders("", "", 1, 20, "New", symbol)
	assert.Nil(t, err)
	//t.Logf("%#v", orders)
	for _, order := range orders {
		if order.ExtFields != nil {
			t.Logf("%#v %v", order, *order.ExtFields)
			t.Logf("ReduceOnly: %v", order.ExtFields.ReduceOnly)
		} else {
			t.Logf("%#v", order)
		}
	}
}

func TestByBit_GetOrder(t *testing.T) {
	b := newByBit()
	order, err := b.GetOrderByID(
		"9d468e94-14b2-4d2e-88b9-590adaee3549",
		"",
		"BTCUSD")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", order)
}

func TestByBit_GetStopOrders(t *testing.T) {
	b := newByBit()
	symbol := "BTCUSD"
	orders, err := b.GetStopOrders("", "", "", "", 0, 10, symbol)
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

func TestByBit_CancelOrderV2(t *testing.T) {
	b := newByBit()
	orderID := "02f0e920-9bc9-4d87-a010-95923a2c430e"
	symbol := "BTCUSD"
	order, err := b.CancelOrderV2(orderID, "", symbol)
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
	order, err := b.CancelStopOrder(orderID, symbol)
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

func TestByBit_GetPosition(t *testing.T) {
	b := newByBit()
	position, err := b.GetPosition("BTCUSD")
	assert.Nil(t, err)
	t.Logf("%#v", position)
}
