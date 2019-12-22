package main

import (
	"github.com/frankrap/bybit-api/rest"
	"log"
)

func main() {
	//baseURL := "https://api.bybit.com/"	// 主网络
	baseURL := "https://api-testnet.bybit.com/" // 测试网络
	b := rest.New(baseURL, "YIxOY2RhFkylPudq96", "Bg9G2oFOb3aaIMguD3FOvOJJVBycaoXqXNcI")

	// 获取持仓
	positions, err := b.GetPositions()
	if err != nil {
		log.Printf("%v", err)
		return
	}

	log.Printf("positions: %#v", positions)

	// 创建委托
	symbol := "BTCUSD"
	side := "Buy"
	orderType := "Limit"
	qty := 30
	price := 7000.0
	timeInForce := "GoodTillCancel"
	order, err := b.CreateOrder(side, orderType, price, qty, timeInForce, symbol)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Create order: %#v", order)

	// 获取委托单
}
