package rest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/resty.v1"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type ByBit struct {
	baseURL   string // https://api-testnet.bybit.com/open-api/
	apiKey    string
	secretKey string
	client    *resty.Client
}

func New(baseURL string, apiKey string, secretKey string) *ByBit {
	return &ByBit{baseURL: baseURL, apiKey: apiKey, secretKey: secretKey, client: resty.New()}
}

// GetBalance Get Wallet Balance
// coin: BTC,EOS,XRP,ETH,USDT
func (b *ByBit) GetWalletBalance(coin string) (result Balance, err error) {
	var ret GetBalanceResult
	params := map[string]interface{}{}
	params["coin"] = coin
	err = b.SignedRequest(http.MethodGet, "v2/private/wallet/balance", params, &ret) // v2/private/wallet/balance
	if err != nil {
		return
	}
	switch coin {
	case "BTC":
		result = ret.Result.BTC
	case "ETH":
		result = ret.Result.ETH
	case "EOS":
		result = ret.Result.EOS
	case "XRP":
		result = ret.Result.XRP
	case "USDT":
		result = ret.Result.USDT
	}
	return
}

// CreateOrder 创建委托单
// symbol: 产品类型, 有效选项:BTCUSD,ETHUSD (BTCUSD ETHUSD)
// side: 方向, 有效选项:Buy, Sell (Buy Sell)
// orderType: 产品类型, 有效选项:BTCUSD,ETHUSD (BTCUSD ETHUSD)
// price: 委托价格, 在没有仓位时，做多的委托价格需高于市价的10%、低于1百万。如有仓位时则需优于强平价。单笔价格增减最小单位为0.5。
// qty: 委托数量, 单笔最大1百万
// timeInForce: 执行策略, 有效选项:GoodTillCancel,ImmediateOrCancel,FillOrKill,PostOnly
func (b *ByBit) CreateOrder(side string, orderType string, price float64, qty int, timeInForce string, symbol string) (result Order, err error) {
	var cResult CreateOrderResult
	params := map[string]interface{}{}
	params["side"] = side
	params["symbol"] = symbol
	params["order_type"] = orderType
	params["qty"] = qty
	params["price"] = price
	params["time_in_force"] = timeInForce
	err = b.SignedRequest(http.MethodPost, "open-api/order/create", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = errors.New(cResult.RetMsg)
		return
	}
	result = cResult.Result
	return
}

// GetOrders 查询活动委托
// symbol
// orderID: 订单ID
// orderLinkID: 机构自定义订单ID
// sort: 排序字段，默认按创建时间排序 (created_at cum_exec_qty qty last_exec_price price cum_exec_value cum_exec_fee)
// order: 升序降序， 默认降序 (desc asc)
// page: 页码，默认取第一页数据
// limit: 一页数量，一页默认展示20条数据
func (b *ByBit) GetOrders(sort string, order string, page int,
	limit int, orderStatus string, symbol string) (result []Order, err error) {
	return b.getOrders("", "", sort, order, page, limit, orderStatus, symbol)
}

// getOrders 查询活动委托
// symbol
// orderID: 订单ID
// orderLinkID: 机构自定义订单ID
// sort: 排序字段，默认按创建时间排序 (created_at cum_exec_qty qty last_exec_price price cum_exec_value cum_exec_fee)
// order: 升序降序， 默认降序 (desc asc)
// page: 页码，默认取第一页数据
// limit: 一页数量，一页默认展示20条数据
func (b *ByBit) getOrders(orderID string, orderLinkID string, sort string, order string, page int,
	limit int, orderStatus string, symbol string) (result []Order, err error) {
	var cResult OrderListResult

	if limit == 0 {
		limit = 20
	}

	params := map[string]interface{}{}
	params["symbol"] = symbol
	if orderID != "" {
		params["order_id"] = orderID
	}
	if orderLinkID != "" {
		params["order_link_id"] = orderLinkID
	}
	if sort != "" {
		params["sort"] = sort
	}
	if order != "" {
		params["order"] = order
	}
	params["page"] = page
	params["limit"] = limit
	if orderStatus != "" {
		params["order_status"] = orderStatus
	}
	err = b.SignedRequest(http.MethodGet, "open-api/order/list", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = errors.New(cResult.RetMsg)
		return
	}

	result = cResult.Result.Data
	return
}

// GetOrderByID
func (b *ByBit) GetOrderByID(orderID string, symbol string) (result Order, err error) {
	var orders []Order
	orders, err = b.getOrders(orderID, "", "", "", 0, 20, "", symbol)
	if err != nil {
		return
	}
	if len(orders) != 1 {
		err = errors.New("not found")
		return
	}
	result = orders[0]
	return
}

// GetOrderByOrderLinkID ...
func (b *ByBit) GetOrderByOrderLinkID(orderLinkID string, symbol string) (result Order, err error) {
	var orders []Order
	orders, err = b.getOrders("", orderLinkID, "", "", 0, 20, "", symbol)
	if err != nil {
		return
	}
	if len(orders) != 1 {
		err = errors.New("not found")
		return
	}
	result = orders[0]
	return
}

// CancelOrder 撤销活动委托单
// orderID: 活动委托单ID, 数据来自创建活动委托单返回的订单唯一ID
// symbol:
func (b *ByBit) CancelOrder(orderID string, symbol string) (result Order, err error) {
	var cResult CancelOrderResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["order_id"] = orderID
	err = b.SignedRequest(http.MethodPost, "open-api/order/cancel", params, &cResult)
	if err != nil {
		return
	}
	if cResult.RetCode != 0 {
		err = errors.New(cResult.RetMsg)
		return
	}

	result = cResult.Result
	return
}

// GetLeverages 获取用户杠杆
func (b *ByBit) GetLeverages() (result map[string]LeverageItem, err error) {
	var r GetLeverageResult
	params := map[string]interface{}{}
	err = b.SignedRequest(http.MethodGet, "user/leverage", params, &r)
	if err != nil {
		return
	}
	result = r.Result
	return
}

// SetLeverage 设置杠杆
func (b *ByBit) SetLeverage(leverage int, symbol string) (err error) {
	var r BaseResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["leverage"] = fmt.Sprintf("%v", leverage)
	err = b.SignedRequest(http.MethodPost, "user/leverage", params, &r)
	if err != nil {
		return
	}
	log.Println(r)
	return
}

// GetPositions 获取我的仓位
func (b *ByBit) GetPositions() (result []Position, err error) {
	var r PositionListResult

	params := map[string]interface{}{}
	err = b.SignedRequest(http.MethodGet, "position/list", params, &r)
	if err != nil {
		return
	}
	if r.RetCode != 0 {
		err = errors.New(r.RetMsg)
		return
	}
	result = r.Result
	return
}

func (b *ByBit) PublicRequest(method string, apiURL string) {
	fullURL := b.baseURL + apiURL
	log.Println(fullURL)
	resp, err := b.client.R().Execute(method, fullURL)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	log.Printf("%v", string(resp.Body()))
}

func (b *ByBit) SignedRequest(method string, apiURL string, params map[string]interface{}, result interface{}) error {
	timestamp := time.Now().Unix()*1000 + 1000

	params["api_key"] = b.apiKey
	params["timestamp"] = timestamp

	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var p []string
	for _, k := range keys {
		p = append(p, fmt.Sprintf("%v=%v", k, params[k]))
	}

	param := strings.Join(p, "&")
	signature := b.getSigned(param)
	param += "&sign=" + signature

	fullURL := b.baseURL + apiURL + "?" + param
	r, err := b.client.R().Execute(method, fullURL)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	log.Println(string(r.Body()))
	err = json.Unmarshal(r.Body(), result)
	return err
}

func (b *ByBit) getSigned(param string) string {
	sig := hmac.New(sha256.New, []byte(b.secretKey))
	sig.Write([]byte(param))
	signature := hex.EncodeToString(sig.Sum(nil))
	return signature
}
