package exchange

import (
	"github.com/sirupsen/logrus"
	"github.com/BasPH/bitcoin_arbitrage/config"
	"github.com/BasPH/bitcoin_arbitrage/exchange/kraken"
)

type ExchangeManager struct {
	log       *logrus.Logger
	Exchanges []*Exchange
}

func NewExchangeManager(config *config.Config, logger *logrus.Logger) *ExchangeManager {
	k, err := kraken.NewKrakenExchange(config.Kraken, logger)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info(k.LatestPrice("XXBTZEUR"))

	return &ExchangeManager{
		log: logger,
	}
}
