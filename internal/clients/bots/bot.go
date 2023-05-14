package bots

import (
	"net/http"
	"tg-bot/internal/clients/bots/telegram"
)

// Listener to abstract from the implementation of the client, potentially to use business logic for the vk bot
type Listener interface {
	Listen()
}

func NewBotTelegram(token string) Listener {
	bot := &telegram.Bot{
		Token:  token,
		Client: http.DefaultClient,
	}
	return bot
}
