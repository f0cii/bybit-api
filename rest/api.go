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
	"strings"
	"time"
)

type ByBit struct {
	baseURL          string // https://api-testnet.bybit.com/open-api/
	apiKey           string
	secretKey        string
	serverTimeOffset int64 // 时间偏差(ms)
	client           *http.Client
	debugMode        bool
}

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

// SetCorrectServerTime 校正服务器时间
func (b *ByBit) SetCorrectServerTime() (err error) {
	var timeNow int64
	_, timeNow, err = b.GetServerTime()
	if err != nil {
		return
	}
	b.serverTimeOffset = timeNow - time.Now().UnixNano()/1e6
	return
}

// GetBalance Get Wallet Balance
// coin: BTC,EOS,XRP,ETH,USDT
func (b *ByBit) GetWalletBalance(coin string) (query string, result Balance, err error) {
	var ret GetBalanceResult
	params := map[string]interface{}{}
	params["coin"] = coin
	query, _, err = b.SignedRequest(http.MethodGet, "v2/private/wallet/balance", params, &ret) // v2/private/wallet/balance
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

// GetPositions 获取我的仓位
func (b *ByBit) GetPositions() (query string, result []PositionData, err error) {
	var r PositionArrayResponse

	params := map[string]interface{}{}
	var resp []byte
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

// GetPosition 获取我的仓位
func (b *ByBit) GetPosition(symbol string) (query string, result Position, err error) {
	var r PositionResponse

	params := map[string]interface{}{}
	params["symbol"] = symbol
	var resp []byte
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

// SetLeverage 设置杠杆
func (b *ByBit) SetLeverage(leverage int, symbol string) (query string, err error) {
	var r BaseResult
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["leverage"] = fmt.Sprintf("%v", leverage)
	query, _, err = b.SignedRequest(http.MethodPost, "user/leverage/save", params, &r)
	if err != nil {
		return
	}
	log.Println(r)
	return
}

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

func (b *ByBit) getSigned(param string) string {
	sig := hmac.New(sha256.New, []byte(b.secretKey))
	sig.Write([]byte(param))
	signature := hex.EncodeToString(sig.Sum(nil))
	return signature
}
