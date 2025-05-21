package main

import (
	"log"

	"../../pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

    "github.com/Gromp-13/IronCurtainBot/pkg/telegram"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

}
