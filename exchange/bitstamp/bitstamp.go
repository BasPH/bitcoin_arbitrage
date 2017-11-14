package bitstamp

import (
	"github.com/BasPH/bitcoin_arbitrage/config"
	"github.com/sirupsen/logrus"
	"github.com/ajph/bitstamp-go"
	"time"
)

const name = "Bitstamp"

type bitstampExchange struct {
	log *logrus.Logger
	delay time.Duration
}

func (bf *bitstampExchange) Name() string {
	return name
}

func (bf *bitstampExchange) Delay() time.Duration {
	return bf.delay
}

func New(c config.Exchange, logger *logrus.Logger) *bitstampExchange {
	logger.Infof("Initializing %v", name)
	bitstamp.SetAuth(c.ClientID, c.APIKey, c.APISecret)
	return &bitstampExchange{
		log:    logger,
		delay:  time.Duration(time.Second * time.Duration(c.ScrapeInterval)),
	}
}
