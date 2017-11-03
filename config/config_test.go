package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	test := []byte(`
[kraken]
api_key = "test_key"
api_secret = "test_secret"
`)

	conf, err := Load(test)
	if err != nil {
		t.Error(err)
	}

	if conf.Kraken.APIKey != "test_key" {
		t.Error("Kraken API key not parsed correctly.")
	}
	if conf.Kraken.APISecret != "test_secret" {
		t.Error("Kraken API secret not parsed correctly.")
	}
}

func TestLoadFile(t *testing.T) {
	conf, err := LoadFile("testdata/example_config.toml")
	if err != nil {
		t.Error(err)
	}

	if conf.Kraken.APIKey != "test_key" {
		t.Error("Kraken API key not parsed correctly.")
	}
	if conf.Kraken.APISecret != "test_secret" {
		t.Error("Kraken API secret not parsed correctly.")
	}
}
