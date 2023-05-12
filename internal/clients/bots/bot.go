package bots

import (
	"net/http"
	"tg-bot/internal/clients/bots/telegram"
	"tg-bot/internal/clients/bots/vk"
)

type Listener interface {
	//Listen(ctx context.Context)
	Listen()
}

func NewBotTelegram(token string) Listener {
	bot := &telegram.Bot{
		Token:  token,
		Client: http.DefaultClient,
	}
	return bot
}

//token?
func NewBotVK() Listener {
	bot := &vk.Bot{}
	return bot
}
