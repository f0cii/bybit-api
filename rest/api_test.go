package rest

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"
)

// https://t.me/Bybitapi

// HttpProxy  = "http://127.0.0.1:6152"
// SocksProxy = "socks5://127.0.0.1:6153"
func newClient(proxyURL string) *http.Client {
	if proxyURL == "" {
		return nil
	}
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyURL)
	}

	httpTransport := &http.Transport{
		Proxy: proxy,
	}

	httpClient := &http.Client{
		Transport: httpTransport,
	}
	return httpClient
}

func newByBit() *ByBit {
	//baseURL := "https://api.bybit.com/"
	baseURL := "https://api-testnet.bybit.com/"
	//baseURL := "https://api.bytick.com/"
	apiKey := "rwEwhfC6mDFYIGfcyb"
	secretKey := "yfNJSzGapfFwbJyvguAyVXLJSIOCIegBg42Z"
	client := newClient("socks5://127.0.0.1:1080")
	b := New(client, baseURL, apiKey, secretKey, true)
	err := b.SetCorrectServerTime()
	if err != nil {
		log.Printf("%v", err)
	}
	return b
}

func TestByBit_GetServerTime(t *testing.T) {
	b := newByBit()
	_, _, timeNow, err := b.GetServerTime()
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

func TestByBit_GetTickers(t *testing.T) {
	b := newByBit()
	_, _, tickers, err := b.GetTickers()
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
	_, _, records, err := b.GetTradingRecords("BTCUSD", 0, 0)
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
	_, _, symbols, err := b.GetSymbols()
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
	//_, _, balance, err := b.GetWalletBalance("BTC")
	_, _, balance, err := b.GetWalletBalance("USDT")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v", balance)
}

func TestByBit_SetLeverage(t *testing.T) {
	b := newByBit()
	_, _, _ = b.SetLeverage(3, "BTCUSD")
}

func TestByBit_GetPositions(t *testing.T) {
	b := newByBit()
	_, _, positions, err := b.GetPositions()
	assert.Nil(t, err)
	t.Logf("%#v", positions)
}

func TestByBit_GetPosition(t *testing.T) {
	b := newByBit()
	_, _, position, err := b.GetPosition("BTCUSD")
	assert.Nil(t, err)
	t.Logf("%#v", position)
}
