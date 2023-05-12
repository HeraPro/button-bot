package main

import (
	"os"
	"tg-bot/internal/clients/bots"
)

func main() {
	tgBot := bots.NewBotTelegram(os.Getenv("TELEGRAM_TOKEN"))
	tgBot.Listen()
}
