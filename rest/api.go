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

	for _, v := range r.Result {
		result = append(result, b.convertPositionV1(v))
	}
	return
}

func (b *ByBit) convertPositionV1(position PositionV1) (result Position) {
	result.ID = position.ID
	result.UserID = position.UserID
	result.RiskID = position.RiskID
	result.Symbol = position.Symbol
	result.Size = position.Size
	result.Side = position.Side
	result.EntryPrice = position.EntryPrice
	result.LiqPrice = position.LiqPrice
	result.BustPrice = position.BustPrice
	result.TakeProfit = position.TakeProfit
	result.StopLoss = position.StopLoss
	result.TrailingStop = position.TrailingStop
	result.PositionValue = position.PositionValue
	result.Leverage = position.Leverage
	result.PositionStatus = position.PositionStatus
	result.AutoAddMargin = position.AutoAddMargin
	result.OrderMargin = position.OrderMargin
	result.PositionMargin = position.PositionMargin
	result.OccClosingFee = position.OccClosingFee
	result.OccFundingFee = position.OccFundingFee
	result.ExtFields = position.ExtFields
	result.WalletBalance = position.WalletBalance
	result.CumRealisedPnl = position.CumRealisedPnl
	result.CumCommission = position.CumCommission
	result.RealisedPnl = position.RealisedPnl
	result.DeleverageIndicator = position.DeleverageIndicator
	result.OcCalcData = position.OcCalcData
	result.CrossSeq = position.CrossSeq
	result.PositionSeq = position.PositionSeq
	result.CreatedAt = position.CreatedAt
	result.UpdatedAt = position.UpdatedAt
	result.UnrealisedPnl = position.UnrealisedPnl
	return
}

// GetPosition 获取我的仓位
func (b *ByBit) GetPosition(symbol string) (result Position, err error) {
	var r GetPositionResult

	params := map[string]interface{}{}
	params["symbol"] = symbol
	err = b.SignedRequest(http.MethodGet, "v2/private/position/list", params, &r)
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

func (b *ByBit) PublicRequest(method string, apiURL string, params map[string]interface{}, result interface{}) error {
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
	fullURL := b.baseURL + apiURL
	if param != "" {
		fullURL += "?" + param
	}
	//log.Println(fullURL)
	r, err := b.client.R().Execute(method, fullURL)
	if err != nil {
		log.Printf("%v", err)
		return err
	}
	//log.Printf("%v", string(r.Body()))
	err = json.Unmarshal(r.Body(), result)
	return err
}

func (b *ByBit) SignedRequest(method string, apiURL string, params map[string]interface{}, result interface{}) error {
	timestamp := time.Now().UnixNano() / 1e6

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
