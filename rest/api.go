package rest

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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
