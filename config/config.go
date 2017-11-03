package config

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
)

type Config struct {
	Kraken KrakenExchange `toml:"kraken,omitempty"`
	Bitstamp BitstampExchange `toml:"bitstamp,omitempty"`
	Bitfinex BitfinexExchange `toml:"bitfinex,omitempty"`
}

type KrakenExchange struct {
	APIKey string `toml:"api_key"`
	APISecret string `toml:"api_secret"`
}

type BitstampExchange struct {
	ClientId string `toml:"client_id"`
	APIKey string `toml:"api_key"`
	APISecret string `toml:"api_secret"`
}

type BitfinexExchange struct {
	APIKey string `toml:"api_key"`
	APISecret string `toml:"api_secret"`
}

func Load(b []byte) (*Config, error) {
	config := &Config{}
	err := toml.Unmarshal(b, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func LoadFile(path string) (*Config, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config, err := Load(content)
	if err != nil {
		return nil, err
	}

	return config, nil
}
