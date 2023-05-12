package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Bot struct {
	Token  string
	Client *http.Client
	//buffer?
}

func (b Bot) GetMe() (*User, error) {
	req := &GetMeRequest{}
	apiResp, err := b.Send(req)
	if err != nil {
		return nil, err
	}
	var bot *User
	err = json.Unmarshal(apiResp.Result, &bot)
	if err != nil {
		return nil, err
	}
	return bot, nil
}

func (b Bot) Request(method string, body io.Reader) (*ApiResponse, error) {
	rawResp, err := b.Client.Post(method, "application/json", body)
	if err != nil {
		return nil, err
	}

	var apiResp *ApiResponse
	err = json.NewDecoder(rawResp.Body).Decode(&apiResp)
	if err != nil {
		return nil, err
	}

	return apiResp, err
}

func (b Bot) Send(req Request) (*ApiResponse, error) {
	//what if i put this in config
	method := fmt.Sprintf(ENDPOINT, b.Token, req.getMethod())
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	apiResp, err := b.Request(method, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	} else if !apiResp.Ok {
		resp := fmt.Sprintf("Error code: %d Description: %s", apiResp.ErrorCode, apiResp.Description)
		return nil, errors.New(resp)
	}
	return apiResp, nil
}

func (b Bot) GetUpdates(updateId, timeout int) ([]*Update, error) {
	req := &GetUpdatesRequest{
		Offset:  updateId,
		Timeout: timeout,
	}
	apiResp, err := b.Send(req)
	if err != nil {
		return nil, err
	}

	var updates []*Update
	err = json.Unmarshal(apiResp.Result, &updates)
	if err != nil {
		return nil, err
	}

	return updates, nil
}

func (b Bot) Listen() {
	//ctx
	bot, err := b.GetMe()
	if err != nil {
		log.Panic(err)
	}
	log.Println(fmt.Sprintf("telegram bot username: %s", bot.Username))
	//these two are jokes
	var (
		tired        bool
		callbackText string = "WOW"
	)
	for true {
		updates, err := b.GetUpdates(updateId, timeout)
		if err != nil {
			log.Println(err)
		}

		for _, update := range updates {
			updateId = update.UpdateID + 1
			if update.Message != nil {
				log.Println(fmt.Sprintf(
					`{"username": "%s", "firstname": "%s", "text": "%s"}`,
					update.Message.From.Username,
					update.Message.From.FirstName,
					update.Message.Text))
				if update.Message.Entities != nil && update.Message.Entities[0].Type == "bot_command" && update.Message.Text == "/start" {
					req := &SendMessageRequest{
						ChatID: update.Message.Chat.ID,
						Text:   greetingsText,
					}
					_, err = b.Send(req)
					if err != nil {
						log.Println(err)
					}
				} else {
					req := &SendMessageRequest{
						ChatID:      update.Message.Chat.ID,
						Text:        firstLayerText,
						ReplyMarkup: inlineKeyboards["root_layer"],
					}
					_, err = b.Send(req)
					if err != nil {
						log.Println(err)
					}
				}
			} else if update.CallbackQuery != nil {
				log.Println(fmt.Sprintf(
					`{"username": "%s", "firstname": "%s", "callback_date": "%s"}`,
					update.CallbackQuery.From.Username,
					update.CallbackQuery.From.FirstName,
					update.CallbackQuery.Data))
				req := &AnswerCallbackQueryRequest{
					CallbackQueryID: update.CallbackQuery.ID,
				}
				if !tired {
					req.Text = callbackText
				}
				_, err = b.Send(req)
				if err != nil {
					log.Println(err)
				}

				switch update.CallbackQuery.Data {
				case "root_layer":
					req := &EditMessageTextRequest{
						ChatID:      update.CallbackQuery.Message.Chat.ID,
						MessageID:   update.CallbackQuery.Message.MessageID,
						Text:        firstLayerText,
						ReplyMarkup: rootKeyboard,
					}
					_, err = b.Send(req)
					if err != nil {
						log.Println(err)
					}

				case "turn_off_WOW":
					if callbackText != "GG" {
						tired = true
					} else {
						callbackText = "NOT_WOW"
					}

				case "turn_on_WOW":
					if callbackText != "WOW" {
						callbackText = "WOW"
					}
					tired = false

				case "nothing":
					callbackText = "GG"
					tired = false

				default:
					req := EditMessageTextRequest{
						ChatID:      update.CallbackQuery.Message.Chat.ID,
						MessageID:   update.CallbackQuery.Message.MessageID,
						Text:        secondLayerText,
						ReplyMarkup: inlineKeyboards[update.CallbackQuery.Data],
					}
					_, err = b.Send(req)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}
}
