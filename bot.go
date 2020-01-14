// functions
// 1. Init socket
// 2. Get market value
// 3. Get user status
// TBD. orderer

package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SetBot() {
	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		log.Fatal(err)
	}

}
