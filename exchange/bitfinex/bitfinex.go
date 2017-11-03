package bitfinex

import (
	"github.com/bitfinexcom/bitfinex-api-go/v1"
	"github.com/BasPH/bitcoin_arbitrage/config"
	"github.com/sirupsen/logrus"
	"strconv"
)

type bitfinexExchange struct {
	client *bitfinex.Client
	log *logrus.Logger
}

func NewBitfinexExchange(config config.BitfinexExchange, logger *logrus.Logger) (*bitfinexExchange, error) {
	client := bitfinex.NewClient().Auth(config.APIKey, config.APISecret)
	return &bitfinexExchange{
		client: client,
		log: logger,
	}, nil
}

func (bf *bitfinexExchange) LatestPrice(ticker string) float64 {
	tick, err := bf.client.Ticker.Get("btcusd")
	if err != nil {
		bf.log.Error(err)
	}

	latest, err := strconv.ParseFloat(tick.LastPrice, 32)
	if err != nil {
		bf.log.Error(err)
	}
	return latest
}
