package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	test := []byte(`
[[exchange]]
name = "kraken"
api_key = "test_key"
api_secret = "test_secret"
scrape_interval = 5
`)

	conf, err := Load(test)
	if err != nil {
		t.Error(err)
	}

	if len(conf.Exchanges) != 1 {
		t.Error("Number of exchanges is not 1")
	}

	if conf.Exchanges[0].Name != "kraken" {
		t.Error("Expected name \"kraken\"")
	}
}

func TestLoadFile(t *testing.T) {
	conf, err := LoadFile("testdata/example_config.toml")
	if err != nil {
		t.Error(err)
	}

	if len(conf.Exchanges) != 4 {
		t.Errorf("Number of exchanges is not 4, counted %v", len(conf.Exchanges))
	}

	if conf.Exchanges[0].Name != "kraken" {
		t.Error("Expected name \"kraken\"")
	}
}
