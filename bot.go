// functions
// 1. Init socket
// 2. Get market value
// 3. Get user status
// TBD. orderer

package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SetBot() {
	key := ""
	if k, ok := os.LookupEnv("BOT_API"); ok {
		key = k
	}

	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	fmt.Println(bot.Self.String())

}
