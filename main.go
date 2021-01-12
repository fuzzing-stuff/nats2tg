package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"gopkg.in/yaml.v3"
)

func tgBot() {
	bot, err := tgbotapi.NewBotAPI("token")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	channel := tgbotapi.NewUpdate(0)
	channel.Timeout = 30

	updates, err := bot.GetUpdatesChan(channel)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello World")

		_, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}

type Socks5Type struct {
	URL      string `json:"url"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type NatsConfig struct {
	Server string `json:"server"`
}

type TgConfig struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

type RoutesConfig struct {
	Topic   string `json:"topic"`
	Channel string `json:"channel"`
}

type Config struct {
	Token     string
	TestToken string
	Channel   string
	DataDir   string
	Socks     Socks5Type
	Nats      NatsConfig
	Telegram  TgConfig
	Routes    []RoutesConfig
}

func parseConfig(inputConfig []byte) (Config, error) {

	config := Config{}
	err := yaml.Unmarshal(inputConfig, &config)
	if err != nil {
		log.Fatal("error parsing config file", err)
	}
	return config, err

}

func main() {

	defaultConfigPath := ""
	configFile := flag.String("config", defaultConfigPath, "Path to config file")
	_, err := os.Stat(*configFile)
	if err != nil {
		*configFile = defaultConfigPath
	}

	confYAML, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	parseConfig(confYAML)

	tgBot()

}
