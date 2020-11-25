package rest

import (
	"testing"
	"time"
)

func TestGetKLine2(t *testing.T) {
	b := newByBit()
	from := time.Now().Add(-1 * time.Hour).Unix()
	_, ohlcs, err := b.GetKLine2(
		"BTCUSDT",
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
