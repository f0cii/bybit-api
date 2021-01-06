package rest

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Bybit
type ByBit struct {
	baseURL          string // https://api-testnet.bybit.com/open-api/
	apiKey           string
	secretKey        string
	serverTimeOffset int64
	client           *http.Client
	debugMode        bool
}

// New
func New(httpClient *http.Client, baseURL string, apiKey string, secretKey string, debugMode bool) *ByBit {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	return &ByBit{
		baseURL:   baseURL,
		apiKey:    apiKey,
		secretKey: secretKey,
		client:    httpClient,
		debugMode: debugMode,
	}
}

// SetCorrectServerTime
func (b *ByBit) SetCorrectServerTime() (err error) {
	var timeNow int64
	_, _, timeNow, err = b.GetServerTime()
	if err != nil {
		return
	}
	b.serverTimeOffset = timeNow - time.Now().UnixNano()/1e6
	return
}

// PublicRequest
func (b *ByBit) PublicRequest(method string, apiURL string, params map[string]interface{}, result interface{}) (fullURL string, resp []byte, err error) {
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
	fullURL = b.baseURL + apiURL
	if param != "" {
		fullURL += "?" + param
	}
	if b.debugMode {
		log.Printf("PublicRequest: %v", fullURL)
	}
	var binBody = bytes.NewReader(make([]byte, 0))

	// get a http request
	var request *http.Request
	request, err = http.NewRequest(method, fullURL, binBody)
	if err != nil {
		return
	}

	var response *http.Response
	response, err = b.client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if b.debugMode {
		log.Printf("PublicRequest: %v", string(resp))
	}

	err = json.Unmarshal(resp, result)
	return
}

// SignedRequest
func (b *ByBit) SignedRequest(method string, apiURL string, params map[string]interface{}, result interface{}) (fullURL string, resp []byte, err error) {
	timestamp := time.Now().UnixNano()/1e6 + b.serverTimeOffset

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

	fullURL = b.baseURL + apiURL + "?" + param
	if b.debugMode {
		log.Printf("SignedRequest: %v", fullURL)
	}
	var binBody = bytes.NewReader(make([]byte, 0))

	// get a http request
	var request *http.Request
	request, err = http.NewRequest(method, fullURL, binBody)
	if err != nil {
		return
	}

	var response *http.Response
	response, err = b.client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if b.debugMode {
		log.Printf("SignedRequest: %v", string(resp))
	}
	err = json.Unmarshal(resp, result)
	return
}

// getSigned
func (b *ByBit) getSigned(param string) string {
	sig := hmac.New(sha256.New, []byte(b.secretKey))
	sig.Write([]byte(param))
	signature := hex.EncodeToString(sig.Sum(nil))
	return signature
}

// GetServerTime
func (b *ByBit) GetServerTime() (query string, resp []byte, timeNow int64, err error) {
	params := map[string]interface{}{}
	var ret BaseResult
	query, _, err = b.PublicRequest(http.MethodGet, "v2/public/time", params, &ret)
	if err != nil {
		return
	}
	var t float64
	t, err = strconv.ParseFloat(ret.TimeNow, 64)
	if err != nil {
		return
	}
	timeNow = int64(t * 1000)
	return
}

// GetWalletBalance
func (b *ByBit) GetWalletBalance(coin string) (query string, resp []byte, result Balance, err error) {
	var ret GetBalanceResult
	params := map[string]interface{}{}
	params["coin"] = coin
	query, _, err = b.SignedRequest(http.MethodGet, "v2/private/wallet/balance", params, &ret)
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

// GetPositions
func (b *ByBit) GetPositions() (query string, resp []byte, result []PositionData, err error) {
	var r PositionArrayResponse
	params := map[string]interface{}{}
	query, resp, err = b.SignedRequest(http.MethodGet, "v2/private/position/list", params, &r)
	if err != nil {
		return
	}
	if r.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", r.RetMsg, string(resp))
		return
	}
	result = r.Result
	return
}

// GetPosition
func (b *ByBit) GetPosition(symbol string) (query string, resp []byte, result Position, err error) {
	var r PositionResponse
	params := map[string]interface{}{}
	params["symbol"] = symbol
	query, resp, err = b.SignedRequest(http.MethodGet, "v2/private/position/list", params, &r)
	if err != nil {
		return
	}
	if r.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", r.RetMsg, string(resp))
		return
	}
	result = r.Result
	return
}

// SetLeverage
func (b *ByBit) SetLeverage(leverage int, symbol string) (query string, resp []byte, err error) {
	var r BaseResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["leverage"] = fmt.Sprintf("%v", leverage)
	query, _, err = b.SignedRequest(http.MethodPost, "user/leverage/save", params, &r)
	if err != nil {
		return
	}
	return
}

// WalletRecords
func (b *ByBit) WalletRecords(symbol string, page int, limit int) (query string, resp []byte, result []WalletFundRecord, err error) {
	var r WalletFundRecordResponse
	params := map[string]interface{}{}
	if symbol != "" {
		params["currency"] = symbol
	}
	if page > 0 {
		params["page"] = page
	}
	if limit > 0 {
		params["limit"] = limit
	}
	query, resp, err = b.SignedRequest(http.MethodGet, "open-api/wallet/fund/records", params, &r)
	if err != nil {
		return
	}
	if r.RetCode != 0 {
		err = fmt.Errorf("%v body: [%v]", r.RetMsg, string(resp))
		return
	}
	result = r.Result.Data
	return
}
