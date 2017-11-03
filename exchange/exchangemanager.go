package exchange

import (
	"github.com/sirupsen/logrus"
	"github.com/BasPH/bitcoin_arbitrage/config"
	"github.com/BasPH/bitcoin_arbitrage/exchange/kraken"
	"github.com/BasPH/bitcoin_arbitrage/exchange/bitstamp"
	"github.com/BasPH/bitcoin_arbitrage/exchange/bitfinex"
)

type ExchangeManager struct {
	log       *logrus.Logger
	Exchanges []Exchange
}

func NewExchangeManager(config *config.Config, logger *logrus.Logger) *ExchangeManager {
	var exchanges []Exchange

	k, err := kraken.NewKrakenExchange(config.Kraken, logger)
	if err != nil {
		logger.Fatal(err)
	}
	if k != nil {
		exchanges = append(exchanges, k)
	}

	bs, err := bitstamp.NewBitstampExchange(config.Bitstamp, logger)
	if err != nil {
		logger.Fatal(err)
	}
	if bs != nil {
		exchanges = append(exchanges, bs)
	}

	bf, err := bitfinex.NewBitfinexExchange(config.Bitfinex, logger)
	if err != nil {
		logger.Fatal(err)
	}
	if bf != nil {
		exchanges = append(exchanges, bf)
	}

	return &ExchangeManager{
		log: logger,
		Exchanges: exchanges,
	}
}
