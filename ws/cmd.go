package ws

/*
ws.send('{"op":"subscribe","args":["topic","topic.filter"]}');

// 同一个类型的filter有多个时，以'|'分割
// 如订阅BTCUSD一分钟和三分钟的kline
ws.send('{"op":"subscribe","args":["kline.BTCUSD.1m|3m"]}');

// 订阅同一个类型filter的所有数据时请使用'*'
// 如订阅所有产品的所有interval kline
ws.send('{"op":"subscribe","args":["kline.*.*"]}')
*/
type Cmd struct {
	Op   string        `json:"op"`
	Args []interface{} `json:"args"`
}
