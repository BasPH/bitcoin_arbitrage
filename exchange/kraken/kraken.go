package kraken

import (
	"github.com/beldur/kraken-go-api-client"
	"github.com/BasPH/bitcoin_arbitrage/config"
	"github.com/sirupsen/logrus"
	"time"
)

const name = "Kraken"

type krakenExchange struct {
	api   *krakenapi.KrakenApi
	log   *logrus.Logger
	delay time.Duration
}

func (ke *krakenExchange) Name() string {
	return name
}

func (ke *krakenExchange) Delay() time.Duration {
	return ke.delay
}

func New(c config.Exchange, logger *logrus.Logger) *krakenExchange {
	logger.Infof("Initializing %v", name)
	api := krakenapi.New(c.APIKey, c.APISecret)
	return &krakenExchange{
		api:   api,
		log:   logger,
		delay: time.Duration(time.Second * time.Duration(c.ScrapeInterval)),
	}
}
