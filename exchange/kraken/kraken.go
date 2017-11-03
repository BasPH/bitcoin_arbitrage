package kraken

import (
	"github.com/beldur/kraken-go-api-client"
	"github.com/BasPH/bitcoin_arbitrage/config"
	"github.com/sirupsen/logrus"
)

type krakenExchange struct {
	api *krakenapi.KrakenApi
	log *logrus.Logger
}

func NewKrakenExchange(config config.KrakenExchange, logger *logrus.Logger) (*krakenExchange, error) {
	api := krakenapi.New(config.APIKey, config.APISecret)
	tr, err := api.Time()
	if err != nil {
		return nil, err
	}
	logger.Info(tr.Rfc1123)

	return &krakenExchange{
		api: api,
		log: logger,
	}, nil
}

func (ke *krakenExchange) LatestPrice(ticker string) float32 {
	return 123
}
