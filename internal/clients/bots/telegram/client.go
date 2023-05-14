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

// GetMe provides information about used bot
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

// Request method implementing transport layer and validates if req is sent successfully.
// method string must represent api endpoint of telegram
// body io.Reader must be json-value in any form that is readable
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

// Send method is a wrapper of Request method, validates if response is bad
// req Request represents DTO for determined endpoint
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

// GetUpdates method represents the telegram api endpoint of same name
// updateId int is offset (param), allows to start getting updates starting from certain one
// timeout int is param that allows you to postpone response if there's no update yet
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

// Listen method is determined for business layer. It is potentially in todos.
func (b Bot) Listen() {
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
						ReplyMarkup: inlineKeyboards["root_layer"],
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

					//add keyboard man
				case "nothing":
					callbackText = "GG"
					tired = false

				default:
					req := &EditMessageTextRequest{
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
