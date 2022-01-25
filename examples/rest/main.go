package main

import (
	"github.com/frankrap/bybit-api/rest"
	"log"
	"net/http"
	"net/url"
)

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

func main() {
	//baseURL := "https://api.bybit.com/"	// 主网络
	baseURL := "https://api-testnet.bybit.com/" // 测试网络
	client := newClient("socks5://127.0.0.1:1080")
	b := rest.New(client,
		baseURL, "rwEwhfC6mDFYIGfcyb", "yfNJSzGapfFwbJyvguAyVXLJSIOCIegBg42Z", true)

	// 获取持仓
	_, _, positions, err := b.GetPositions()
	if err != nil {
		log.Printf("%v", err)
		return
	}

	log.Printf("positions: %#v", positions)

	// 创建委托
	symbol := "BTCUSD"
	side := "Buy"
	orderType := "Limit"
	qty := 10
	price := 35000.0
	timeInForce := "GoodTillCancel"
	_, _, order, err := b.CreateOrder(side, orderType, price, qty, timeInForce, 0, 0, false, false, "", symbol)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Create order: %#v", order)

	// 获取委托单
}
