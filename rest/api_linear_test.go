package rest

import (
	"testing"
	"time"
)

func TestLinear_GetOrderBook(t *testing.T) {
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

func TestByBit_LinearGetKLine(t *testing.T) {
	b := newByBit()
	from := time.Now().Add(-1 * time.Hour).Unix()
	_, _, ohlcs, err := b.LinearGetKLine("BTCUSDT", "1", from, 10)
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range ohlcs {
		t.Logf("%#v", v)
	}
}

func TestByBit_LinearCreateOrder(t *testing.T) {
	b := newByBit()
	_, _, order, err := b.LinearCreateOrder("Buy", "Limit", 35000, 1, "GoodTillCancel", 0,
		0, false, false, "", "BTCUSDT")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", order)
}

func TestByBit_LinearGetOrders(t *testing.T) {
	b := newByBit()
	_, _, orders, err := b.LinearGetOrders("BTCUSDT", "Created,New", 10, 0)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", orders.Data)
}

func TestByBit_LinearCancelOrder(t *testing.T) {
	b := newByBit()
	_, _, ret, err := b.LinearCancelOrder("d328974d-bfe8-484f-a0e9-30159bc78aaf", "", "BTCUSDT")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", ret)
}

func TestByBit_LinearCancelAllOrder(t *testing.T) {
	b := newByBit()
	_, _, ret, err := b.LinearCancelAllOrder("BTCUSDT")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", ret)
}

func TestByBit_LinearGetPositions(t *testing.T) {
	b := newByBit()
	_, _, ret, err := b.LinearGetPositions() // BTCUSDT
	if err != nil {
		t.Error(err)
		return
	}
	// t.Logf("%#v", ret)
	for _, v := range ret {
		if !v.IsValid || v.Data.Size == 0 {
			continue
		}
		t.Logf("%#v", v)
	}
}

func TestByBit_LinearGetPosition(t *testing.T) {
	b := newByBit()
	_, _, ret, err := b.LinearGetPosition("BTCUSDT")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", ret)
}
