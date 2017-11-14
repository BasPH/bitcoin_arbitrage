package config

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
)

type Config struct {
	Exchanges []Exchange `toml:"exchange"`
	Limits    struct {
		Eur float64 `toml:"eur,omitempty"`
		Btc float64 `toml:"btc,omitempty"`
	} `toml:"limits,omitempty"`
}

type Exchange struct {
	Name           string `toml:"name"`
	APIKey         string `toml:"api_key"`
	APISecret      string `toml:"api_secret"`
	ScrapeInterval int    `toml:"scrape_interval,omitempty"`
	ClientID       string `toml:"client_id,omitempty"`
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
