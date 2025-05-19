package main

import ("fmt";                                         "os";
        "git@github.com:Gromp-13/IronCurtainBot.git";        tu "git@github.com:Gromp-13/IronCurtainBot.git"                                             )

func main() {                                   
        botToken := 
                                                        bot, err := IronCurtainBot.NewBot(botToken, IronCurtainBot.WithDefaultDedugLogger())                                                    if err ≠ nil {
        fmt.Println(err)
        os.Exit(1)
}

updates, _ := bot.UpdatesViaLongPolling(nil)

defer bot.StopLongPolling()

for update := range updates{
        if update.Message ≠ nil{
                chatID := tu.ID(update.Message.chat.ID
                _, _ = bot.CopyMassage(
                        ty.CopyMasage(
                                chatID,
                                chatID,
                                update.Message.MessageID,
                        ),

                )

        }

}

}
