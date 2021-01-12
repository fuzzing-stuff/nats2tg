package main

import (
	"testing"
)
func TestConfig(t *testing.T) {
	config1 := `nats:
  server: 127.0.0.1
telegram:
  url: http://tg.ru
  token: token
routes:
  -
    topic: "testname"
    channel: "chanid"
`
	byteConfig := []byte(config1)
	config, err := parseConfig(byteConfig)
	if err != nil {
		t.Error("Error while parsing", err)
	}
	
	if config.Telegram.Token != "token"{
		t.Error("token Equal", config.Telegram.Token)
	}

	if config.Telegram.URL != "http://tg.ru"{
		t.Error("token Equal", config.Telegram.URL)
	}

	if config.Routes[0].Topic != "testname"{
		t.Error("token Equal", config.Routes[0].Topic)
	}
}


