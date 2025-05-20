package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type Bot struct{
	bot *tgbotapi.BotAPI 
}

func NewBot(bot *tgbotapi.BotAPI) *Bot {
	return &Bot{bot: bot}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)                    u.Timeout = 60                                                                              updates := b.bot.GetUpdatesChan(u) 
	if err != nil{
		return err
	}

        for update := range updates {                         if update.Message != nil { 

                        log.Printf("[%s] %s",
update.Message.From.UserName, update.Message.T
ext)

                        msg := tgbotapi.NewMes
sage(update.Message.Chat.ID, update.Message.Te
xt)                                                                   msg.ReplyToMessageID = update.Message.MessageID                                                                                          b.bot.Send(msg)
}

}
return nil
}
