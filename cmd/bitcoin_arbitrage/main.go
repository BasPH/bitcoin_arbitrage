package main

import (
	"github.com/BasPH/bitcoin_arbitrage/config"
	"github.com/BasPH/bitcoin_arbitrage/engine"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	debug  = kingpin.Flag("debug", "Debug mode.").Short('d').Default("false").Bool()
	dryrun = kingpin.Flag("dryrun", "Dryrun mode - run the engine but don't make any transactions.").Default("false").Bool()
	conf   = kingpin.Flag("config", "Configuration file path.").Default("config.toml").String()

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
	if *debug {
		log.Level = logrus.DebugLevel
	}

	c, err := config.LoadFile(*conf)
	if err != nil {
		log.Fatal(err)
	}

	engine := engine.New(c, log)
	log.Infof("Initialized engine with %d exchanges", len(engine.Exchanges))
	engine.Run()
}
