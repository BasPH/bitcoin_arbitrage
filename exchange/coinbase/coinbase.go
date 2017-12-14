package coinbase

import (
	"github.com/BasPH/bitcoin_arbitrage/config"
	"github.com/fabioberger/coinbase-go"
	"github.com/sirupsen/logrus"
	"time"
)

const name = "Coinbase"

type coinbaseExchange struct {
	client coinbase.Client
	log    *logrus.Logger
	delay  time.Duration
}

func (cb *coinbaseExchange) Name() string {
	return name
}

func (cb *coinbaseExchange) Delay() time.Duration {
	return cb.delay
}

func New(c config.Exchange, logger *logrus.Logger) *coinbaseExchange {
	logger.Infof("Initializing %v", name)
	client := coinbase.ApiKeyClient(c.APIKey, c.APISecret)
	return &coinbaseExchange{
		client: client,
		log:    logger,
		delay:  time.Duration(time.Second * time.Duration(c.ScrapeInterval)),
	}
}

func (cb *coinbaseExchange) Scrape() {
	cb.log.Info("Scrape %v", name)
}
