package bitfinex

import (
	"github.com/BasPH/bitcoin_arbitrage/config"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
	"github.com/sirupsen/logrus"
	"time"
)

const name = "Bitfinex"

type bitfinexExchange struct {
	client *bitfinex.Client
	log    *logrus.Logger
	delay  time.Duration
}

func (bf *bitfinexExchange) Name() string {
	return name
}

func (bf *bitfinexExchange) Delay() time.Duration {
	return bf.delay
}

func New(c config.Exchange, logger *logrus.Logger) *bitfinexExchange {
	logger.Infof("Initializing %v", name)
	client := bitfinex.NewClient().Auth(c.APIKey, c.APISecret)
	return &bitfinexExchange{
		client: client,
		log:    logger,
		delay:  time.Duration(time.Second * time.Duration(c.ScrapeInterval)),
	}
}
