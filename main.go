package main

import (
	"log"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
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

func main() {

	tgBot()

}
