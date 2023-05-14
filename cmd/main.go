package main

import (
	"log"
	"os"
	"tg-bot/internal/clients/bots"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Panic("No token provided")
	}
	tgBot := bots.NewBotTelegram(token)
	tgBot.Listen()
}
