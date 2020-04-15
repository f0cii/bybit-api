package ws

import (
	"log"
	"testing"
)

func TestConnect(t *testing.T) {
	cfg := &Configuration{
		Addr:          HostTestnet,
		ApiKey:        "wKuYtkeNdC2PaMKjoy",
		SecretKey:     "5ekcDn3KnKoCRbfvrPImYzVdx7Ri2hhVxkmw",
		AutoReconnect: true,
	}
	b := New(cfg)

	b.Start()

	forever := make(chan struct{})
	<-forever
}

func handleOrderBook(symbol string, data OrderBook) {
	log.Printf("handleOrderBook %v/%v", symbol, data)
}

func handleTrade(symbol string, data []*Trade) {
	log.Printf("handleTrade %v/%v", symbol, data)
}

func handleKLine(symbol string, data KLine) {
	log.Printf("handleKLine %v/%v", symbol, data)
}

func handleInsurance(currency string, data []*Insurance) {
	log.Printf("handleInsurance %v/%v", currency, data)
}

func handleInstrument(symbol string, data []*Instrument) {
	log.Printf("handleInstrument %v/%v", symbol, data)
}

func handlePosition(data []*Position) {
	log.Printf("handlePosition %v", data)
}

func handleExecution(data []*Execution) {
	log.Printf("handleExecution %v", data)
}

func handleOrder(data []*Order) {
	log.Printf("handleOrder %v", data)
}

func TestOrderBookL2(t *testing.T) {
	cfg := &Configuration{
		Addr:          HostTestnet,
		ApiKey:        "wKuYtkeNdC2PaMKjoy",
		SecretKey:     "5ekcDn3KnKoCRbfvrPImYzVdx7Ri2hhVxkmw",
		AutoReconnect: true,
		DebugMode:     true,
	}
	b := New(cfg)

	// 订阅新版25档orderBook
	// b.Subscribe(WSOrderBook25L1 + ".BTCUSD")
	b.Subscribe("orderBookL2_25" + ".BTCUSD")

	b.On(WSOrderBook25L1, handleOrderBook)

	//// 实时交易
	// b.Subscribe(WSTrade) // BTCUSD/ETHUSD/EOSUSD/XRPUSD
	////b.Subscribe("trade.BTCUSD")
	b.On(WSTrade, handleTrade)

	//// K线
	//b.Subscribe(WSKLine + ".BTCUSD.1m")
	//// {"topic":"kline.BTCUSD.1m","data":{"id":0,"symbol":"BTCUSD","open_time":1563777600,"open":10613.5,"high":10613.5,"low":10613.5,"close":10613.5,"volume":130077,"turnover":12.255806170000001,"interval":"1m"}}
	//

	b.On(WSKLine, handleKLine)

	//
	//// 每日保险基金更新
	// b.Subscribe(WSInsurance)
	//// {"topic":"insurance.BTC","data":[{"currency":"BTC","timestamp":"2019-07-21T20:00:00Z","wallet_balance":30494668519}]}

	b.On(WSInsurance, handleInsurance)
	//
	//// 产品最新行情
	//b.Subscribe(WSInstrument + ".BTCUSD")
	//// {"topic":"instrument.BTCUSD","data":[{"symbol":"BTCUSD"}]}
	//// {"topic":"instrument.BTCUSD","data":[{"symbol":"BTCUSD","mark_price":10599.9,"index_price":10599.92}]}

	b.On(WSInstrument, handleInstrument)

	// 私有类

	// 仓位变化
	b.Subscribe(WSPosition)

	// 委托单成交信息
	b.Subscribe(WSExecution)

	// 委托单的更新
	b.Subscribe(WSOrder)

	b.On(WSPosition, handlePosition)
	b.On(WSExecution, handleExecution)
	b.On(WSOrder, handleOrder)

	b.Start()

	forever := make(chan struct{})
	<-forever
}
