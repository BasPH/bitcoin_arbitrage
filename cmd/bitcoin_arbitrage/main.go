package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/BasPH/bitcoin_arbitrage/config"
	"github.com/BasPH/bitcoin_arbitrage/exchange"
	"os"
)

var (
	debug = kingpin.Flag("debug", "Debug mode.").Short('d').Default("false").Bool()
	conf  = kingpin.Flag("config", "Configuration file path.").Default("config.toml").String()

	log *logrus.Logger
)

func init() {
	log = logrus.New()
	log.Formatter = logrus.Formatter(&logrus.JSONFormatter{})
	log.Out = os.Stdout
	log.Level = logrus.InfoLevel
}

func main() {
	kingpin.Parse()

	c, err := config.LoadFile(*conf)
	if err != nil {
		log.Fatal(err)
	}

	em := exchange.NewExchangeManager(c, log)
	log.Infof("Num exchanges connected: %d", len(em.Exchanges))

	//ticker, err := api.Ticker(krakenapi.XXBTZEUR)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println(ticker.XXBTZEUR.OpeningPrice)
}
