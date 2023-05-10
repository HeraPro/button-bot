package main

import (
	"os"
	"tg-bot/internal/clients/bots"
)

var (
	PORT  = os.Getenv("PORT")
	TOKEN = os.Getenv("TOKEN_TELEGRAM")
)

func main() {
	//()

	tgBot := bots.NewBotTelegram(TOKEN)
	tgBot.Listen()
}
