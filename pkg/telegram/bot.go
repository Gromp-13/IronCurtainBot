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

	updates, err := b.initUpdatesChannel()
	if err != nil{
		return err
	}

	b.handleUpdates(updates)

        return nil
}


func handleUpdates(updates tgbotapi.UpdatesChannel){
	for update := range updates {                         if update.Message == nil {
        continue
      }
       
        b.handleMassage(update.Message)

   }
}


func (b *Bot) handleMassage(message *tgbotapi.Message) {

	log.Printf("[%s] %s", message.From.UserName, message.Text)

        msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)                           

        b.bot.Send(msg)
}

func (b *Bot) initUpdatesChannel() (tgbotapi.initUpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)                    u.Timeout = 60                                                                              return b.bot.GetUpdatesChan(u)
}
