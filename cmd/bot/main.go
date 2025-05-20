package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/Gromp-13/IronCurtainBot/tree/main/pkg/telegram"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil{
		log.Fatal(err)
	}

}
