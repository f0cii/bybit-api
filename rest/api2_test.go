package rest_test

import (
	"github.com/frankrap/bybit-api/rest"
	"log"
	"testing"
)

func newByBit2() *rest.ByBit {
	baseURL := "https://api2-testnet.bybit.com/"
	apiKey := "6IASD6KDBdunn5qLpT"
	secretKey := "nXjZMUiB3aMiPaQ9EUKYFloYNd0zM39RjRWF"
	b := rest.New(nil, baseURL, apiKey, secretKey, true)
	err := b.SetCorrectServerTime()
	if err != nil {
		log.Printf("%v", err)
	}
	return b
}

func TestByBit_GetFunding(t *testing.T) {
	b := newByBit2()
	funding, e := b.GetFunding("BTCUSD", 1)
	if e != nil {
		t.Error(e)
		return
	}
	t.Logf("Funding: %v", funding)
}
