package rest_test

import (
	"github.com/frankrap/bybit-api/rest"
	"testing"
)

func newByBit2() *rest.ByBit {
	baseURL := "https://api2-testnet.bybit.com/"
	apiKey := "6IASD6KDBdunn5qLpT"
	secretKey := "nXjZMUiB3aMiPaQ9EUKYFloYNd0zM39RjRWF"
	b := rest.New(nil, baseURL, apiKey, secretKey, true)
	return b
}

func TestByBit_GetFunding(t *testing.T) {
	b := newByBit2()
	_, _, funding, e := b.GetFunding("BTCUSD", 1, 200)
	if e != nil {
		t.Error(e)
		return
	}
	t.Logf("Funding: %v", funding)
}
