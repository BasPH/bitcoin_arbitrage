package bitstamp

import (
	"github.com/BasPH/bitcoin_arbitrage/config"
	"github.com/sirupsen/logrus"
	"github.com/ajph/bitstamp-go"
)

type bitstampExchange struct {
	log *logrus.Logger
}

// https://www.bitstamp.net/api/
func NewBitstampExchange(config config.BitstampExchange, logger *logrus.Logger) (*bitstampExchange, error) {
	bitstamp.SetAuth(config.ClientId, config.APIKey, config.APISecret)

	return &bitstampExchange{
		log: logger,
	}, nil
}

func (be *bitstampExchange) LatestPrice(ticker string) float64 {
	return 321
}
