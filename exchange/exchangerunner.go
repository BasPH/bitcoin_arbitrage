package exchange

import (
	"github.com/BasPH/bitcoin_arbitrage/config"
	"github.com/BasPH/bitcoin_arbitrage/exchange/bitfinex"
	"github.com/BasPH/bitcoin_arbitrage/exchange/bitstamp"
	"github.com/BasPH/bitcoin_arbitrage/exchange/kraken"
	"github.com/sirupsen/logrus"
	"time"
)

type Exchange interface {
	Name() string
	Delay() time.Duration
}

type ExchangeRunner struct {
	log      *logrus.Logger
	exchange Exchange

	stop chan struct{}
}

func NewExchangeRunner(c config.Exchange, logger *logrus.Logger) *ExchangeRunner {
	var e Exchange
	switch c.Name {
	case "kraken":
		e = kraken.New(c, logger)
	case "bitstamp":
		e = bitstamp.New(c, logger)
	case "bitfinex":
		e = bitfinex.New(c, logger)
	default:
		return nil
	}

	return &ExchangeRunner{
		log:      logger,
		exchange: e,
		stop:     make(chan struct{}),
	}
}

func (er *ExchangeRunner) scrape() {
	ticker := time.NewTicker(er.exchange.Delay()).C
	for {
		select {
		case <-ticker:
			er.log.Infof("Scrape %v", er.exchange.Name())
		}
	}
}

func (er *ExchangeRunner) Run() {
	go er.scrape()

	for {
		select {
		case <-er.stop:
			er.log.Infof("Stopped exchange %v", er.exchange.Name())
			return
		}
	}
	//ticker := time.NewTicker(er.exchange.Delay()).C
	//
	//for {
	//	select {
	//	case <-ticker:
	//		fmt.Printf("Scrape exchange %v", er.exchange.Name())
	//	}
	//	case <-er.stop:
	//		logrus.Debugf("Stopping ExchangeRunner")
	//		close(er.stop)
	//		return
	//	}
	//}

	//fmt.Println("bla")
}

func (er *ExchangeRunner) Shutdown() {
	<-er.stop
}
