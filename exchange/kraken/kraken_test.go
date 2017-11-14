package kraken

import (
	"github.com/BasPH/bitcoin_arbitrage/config"
	"testing"
)

func TestKrakenApi(t *testing.T) {
	c, err := config.LoadFile("../../config.toml")
	if err != nil {
		t.Fatal(err)
	}
	k, _ := NewKrakenExchange(c.Kraken, nil)

	k.AssetPairs()
}
