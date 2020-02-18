package ws

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chuckpreslar/emission"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"nhooyr.io/websocket"
	"strings"
	"sync"
	"time"
)

const (
	MaxTryTimes = 10
)

// https://github.com/bybit-exchange/bybit-official-api-docs/blob/master/zh_cn/websocket.md

// 测试网地址
// wss://stream-testnet.bybit.com/realtime

// 主网地址
// wss://stream.bybit.com/realtime

const (
	HostReal    = "wss://stream.bybit.com/realtime"
	HostTestnet = "wss://stream-testnet.bybit.com/realtime"
)

const (
	WSOrderBook25L1 = "order_book_25L1" // 新版25档orderBook: order_book_25L1.BTCUSD
	WSKLine         = "kline"           // K线: kline.BTCUSD.1m
	WSTrade         = "trade"           // 实时交易: trade/trade.BTCUSD
	WSInsurance     = "insurance"       // 每日保险基金更新: insurance
	WSInstrument    = "instrument"      // 产品最新行情: instrument

	WSPosition  = "position"  // 仓位变化: position
	WSExecution = "execution" // 委托单成交信息: execution
	WSOrder     = "order"     // 委托单的更新: order

	EventDisconnected = "disconnected" // 连接断开事件
)

var (
	topic_order_book_25L1Prefix = WSOrderBook25L1 + "."
)

type Configuration struct {
	Addr          string `json:"addr"`
	ApiKey        string `json:"api_key"`
	SecretKey     string `json:"secret_key"`
	AutoReconnect bool   `json:"auto_reconnect"`
	DebugMode     bool   `json:"debug_mode"`
}

type ByBitWS struct {
	cfg         *Configuration
	conn        *websocket.Conn
	mu          sync.RWMutex
	isConnected bool

	subscribeCmds   []Cmd
	orderBookLocals map[string]*OrderBookLocal // key: symbol

	emitter *emission.Emitter
}

func New(config *Configuration) *ByBitWS {
	b := &ByBitWS{
		cfg:             config,
		emitter:         emission.NewEmitter(),
		orderBookLocals: make(map[string]*OrderBookLocal),
	}
	b.On(EventDisconnected, b.handleDisconnected)
	return b
}

// setIsConnected sets state for isConnected
func (b *ByBitWS) setIsConnected(state bool) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.isConnected = state
}

// IsConnected returns the WebSocket connection state
func (b *ByBitWS) IsConnected() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.isConnected
}

func (b *ByBitWS) Subscribe(arg string) {
	b.subscribeCmds = append(b.subscribeCmds, Cmd{
		Op:   "subscribe",
		Args: []interface{}{arg},
	})
}

func (b *ByBitWS) SendCmd(cmd Cmd) error {
	data, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	return b.Send(string(data))
}

func (b *ByBitWS) Send(msg string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprintf("send error: %v", r))
		}
	}()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = b.conn.Write(ctx, websocket.MessageText, []byte(msg))
	return
}

func (b *ByBitWS) Start() error {
	b.setIsConnected(false)

	b.conn = nil
	for i := 0; i < MaxTryTimes; i++ {
		c, _, err := b.connect()
		if err != nil {
			log.Println(err)
			time.Sleep(1 * time.Second)
			continue
		}
		b.conn = c
		break
	}
	if b.conn == nil {
		return errors.New("connect fail")
	}

	b.setIsConnected(true)

	if b.cfg.ApiKey != "" && b.cfg.SecretKey != "" {
		b.Auth()
	}

	for _, cmd := range b.subscribeCmds {
		b.SendCmd(cmd)
	}

	cancel := make(chan struct{})

	go func() {
		for {
			select {
			case <-time.After(30 * time.Second):
				b.ping()
			case <-cancel:
				return
			}
		}
	}()

	go func() {
		defer close(cancel)

		for {
			ctx := context.Background()
			messageType, data, err := b.conn.Read(ctx)
			if err != nil {
				log.Printf("Read error: %v", err)
				if strings.Contains(err.Error(), "context deadline exceeded") {
					log.Println("111")
					time.Sleep(3 * time.Second)
					continue
				}
				//if err == context.DeadlineExceeded {
				//	continue
				//}
				log.Printf("%v", err)
				b.Emit(EventDisconnected)
				return
			}

			b.processMessage(messageType, data)
		}
	}()

	return nil
}

func (b *ByBitWS) handleDisconnected() {
	log.Println("handleDisconnected")

	if !b.cfg.AutoReconnect {
		return
	}

	log.Println("close")

	b.Close()
	b.Start()
}

func (b *ByBitWS) connect() (*websocket.Conn, *http.Response, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	c, resp, err := websocket.Dial(ctx, b.cfg.Addr, &websocket.DialOptions{})
	return c, resp, err
}

func (b *ByBitWS) ping() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("ping error: %v", r)
		}
	}()

	if !b.IsConnected() {
		return
	}
	if b.conn == nil {
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := b.conn.Write(ctx, websocket.MessageText, []byte(`{"op":"ping"}`))
	if err != nil {
		log.Printf("ping error: %v", err)
	}
}

func (b *ByBitWS) Auth() error {
	// 单位:毫秒
	expires := time.Now().Unix()*1000 + 10000
	req := fmt.Sprintf("GET/realtime%d", expires)
	sig := hmac.New(sha256.New, []byte(b.cfg.SecretKey))
	sig.Write([]byte(req))
	signature := hex.EncodeToString(sig.Sum(nil))

	cmd := Cmd{
		Op: "auth",
		Args: []interface{}{
			b.cfg.ApiKey,
			//fmt.Sprintf("%v", expires),
			expires,
			signature,
		},
	}
	err := b.SendCmd(cmd)
	return err
}

func (b *ByBitWS) processMessage(messageType websocket.MessageType, data []byte) {
	ret := gjson.ParseBytes(data)

	if b.cfg.DebugMode {
		log.Printf("%v", string(data))
	}

	// 处理心跳包
	retMsg := ret.Get("ret_msg").String()
	if retMsg != "" && retMsg == "pong" {
		return
	}

	if ret.Get("success").Exists() {
		return
	}

	topic := ret.Get("topic").String()
	if topic != "" {
		if strings.HasPrefix(topic, topic_order_book_25L1Prefix) {
			symbol := topic[len(topic_order_book_25L1Prefix):]
			type_ := ret.Get("type").String()
			raw := ret.Get("data").Raw

			switch type_ {
			case "snapshot":
				var data []*OrderBookL2
				err := json.Unmarshal([]byte(raw), &data)
				if err != nil {
					log.Printf("%v", err)
					return
				}
				b.processOrderBookSnapshot(symbol, data...)
			case "delta":
				var delta OrderBookL2Delta
				err := json.Unmarshal([]byte(raw), &delta)
				if err != nil {
					log.Printf("%v", err)
					return
				}
				b.processOrderBookDelta(symbol, &delta)
			}
		} else if strings.HasPrefix(topic, WSTrade) {
			symbol := strings.TrimLeft(topic, WSTrade+".")
			raw := ret.Get("data").Raw
			var data []*Trade
			err := json.Unmarshal([]byte(raw), &data)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			b.processTrade(symbol, data...)
		} else if strings.HasPrefix(topic, WSKLine) {
			// kline.BTCUSD.1m
			topicArray := strings.Split(topic, ".")
			if len(topicArray) != 3 {
				return
			}
			symbol := topicArray[1]
			raw := ret.Get("data").Raw
			var data KLine
			err := json.Unmarshal([]byte(raw), &data)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			b.processKLine(symbol, data)
		} else if strings.HasPrefix(topic, WSInsurance) {
			// insurance.BTC
			topicArray := strings.Split(topic, ".")
			if len(topicArray) != 2 {
				return
			}
			currency := topicArray[1]
			raw := ret.Get("data").Raw
			var data []*Insurance
			err := json.Unmarshal([]byte(raw), &data)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			b.processInsurance(currency, data...)
		} else if strings.HasPrefix(topic, WSInstrument) {
			topicArray := strings.Split(topic, ".")
			if len(topicArray) != 2 {
				return
			}
			symbol := topicArray[1]
			raw := ret.Get("data").Raw
			var data []*Instrument
			err := json.Unmarshal([]byte(raw), &data)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			b.processInstrument(symbol, data...)
		} else if topic == WSPosition {
			raw := ret.Get("data").Raw
			var data []*Position
			err := json.Unmarshal([]byte(raw), &data)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			b.processPosition(data...)
		} else if topic == WSExecution {
			raw := ret.Get("data").Raw
			var data []*Execution
			err := json.Unmarshal([]byte(raw), &data)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			b.processExecution(data...)
		} else if topic == WSOrder {
			raw := ret.Get("data").Raw
			var data []*Order
			err := json.Unmarshal([]byte(raw), &data)
			if err != nil {
				log.Printf("%v", err)
				return
			}
			b.processOrder(data...)
		}
		return
	}
}

func (b *ByBitWS) Close() {
	if b.conn == nil {
		return
	}
	b.conn.Close(websocket.StatusNormalClosure, "")
	b.setIsConnected(false)
}
