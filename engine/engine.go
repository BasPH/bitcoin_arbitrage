package engine

import (
	"github.com/BasPH/bitcoin_arbitrage/config"
	"github.com/BasPH/bitcoin_arbitrage/exchange"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Engine struct {
	log       *logrus.Logger
	Exchanges []*exchange.ExchangeRunner
	prices    [][]*exchange.Exchange

	//stop chan struct{}
}

func New(c *config.Config, logger *logrus.Logger) *Engine {
	engine := &Engine{log: logger}
	var exchanges []*exchange.ExchangeRunner

	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, c := range c.Exchanges {
		wg.Add(1)
		go func(ce config.Exchange) {
			defer wg.Done()
			mu.Lock()
			exchanges = append(exchanges, exchange.NewExchangeRunner(ce, logger))
			mu.Unlock()
		}(c)
	}
	wg.Wait()

	engine.Exchanges = exchanges
	engine.log.Info("All exchanges initialized!")
	return engine
}

func (e *Engine) Run() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	for _, ex := range e.Exchanges {
		go ex.Run()
	}

	for {
		select {
		case <-signalChan:
			e.log.Info("Stopped engine")
			return
		}
	}
}
