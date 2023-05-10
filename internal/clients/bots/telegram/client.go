package telegram

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Bot struct {
	Token  string
	Client *http.Client
	//buffer?
}

func (b Bot) Request(method, body string) (*ApiResponse, error) {
	rawResp, err := b.Client.Post(method, "application/json", strings.NewReader(body))
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

func (b Bot) GetUpdates(updateId, timeout int) ([]*Update, error) {
	method := fmt.Sprintf(ENDPOINT, b.Token, GET_UPDATES)
	body := fmt.Sprintf(`{"timeout": %d, "offset": %d}`, timeout, updateId)
	apiResp, err := b.Request(method, body)
	if err != nil {
		return nil, err
	} else if !apiResp.Ok {
		resp := fmt.Sprintf("Error code: %d Description: %s", apiResp.ErrorCode, apiResp.Description)
		return nil, errors.New(resp)
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

	//these two are jokes
	var (
		tired        bool
		callbackText string = "WOW"
	)
	for true {
		updates, err := b.GetUpdates(updateId, timeout)
		if err != nil {
			log.Println(err)
			err = nil
		}

		for _, update := range updates {
			updateId = update.UpdateID + 1
			if update.Message != nil {
				method := fmt.Sprintf(ENDPOINT, b.Token, SEND_MESSAGE)
				if update.Message.Entities != nil && update.Message.Entities[0].Type == "bot_command" && update.Message.Text == "/start" {
					body := fmt.Sprintf(`{"chat_id":%d,"text":"Hi"}`, update.Message.Chat.ID)
					_, err = b.Request(method, body)
					if err != nil {
						log.Println(err)
						err = nil
					}
				} else {
					//todo reply
					body := fmt.Sprintf(`{"chat_id":%d,"text":"1st layer","reply_markup":%s}`,
						update.Message.Chat.ID,
						inlineKeyboards["root_layer"])
					_, err = b.Request(method, body)
					if err != nil {
						log.Println(err)
						err = nil
					}
				}
			} else if update.CallbackQuery != nil {
				method := fmt.Sprintf(ENDPOINT, b.Token, ANSWER_CALLBACK_QUERY)
				var body string
				if !tired {
					body = fmt.Sprintf(`{"callback_query_id": "%s","text": "%s"}`, update.CallbackQuery.ID, callbackText)
				} else {
					body = fmt.Sprintf(`{"callback_query_id": "%s"}`, update.CallbackQuery.ID)

				}
				_, err = b.Request(method, body)
				if err != nil {
					log.Println(err)
					err = nil
				}
				method = fmt.Sprintf(ENDPOINT, b.Token, EDIT_MESSAGE_TEXT)

				//switch, data, url, url_login

				//methods := map[string]string{}

				switch update.CallbackQuery.Data {
				case "root_layer":
					body = fmt.Sprintf(`{"chat_id":%d, "message_id":%d, "text":%s,"reply_markup":%s}`,
						update.CallbackQuery.Message.Chat.ID,
						update.CallbackQuery.Message.MessageID,
						`"1st layer"`,
						inlineKeyboards["root_layer"])

				case "turn_off_WOW":
					if callbackText != "GG" {
						tired = true
					}
					body = fmt.Sprintf(`{"chat_id":%d,"message_id":%d,"text":%s,"reply_markup":%s}`,
						update.CallbackQuery.Message.Chat.ID,
						update.CallbackQuery.Message.MessageID,
						`"2nd layer"`,
						inlineKeyboards["turn_layer"])

				case "turn_on_WOW":
					tired = false
					body = fmt.Sprintf(`{"chat_id":%d,"message_id":%d,"text":%s,"reply_markup":%s}`,
						update.CallbackQuery.Message.Chat.ID,
						update.CallbackQuery.Message.MessageID,
						`"2nd layer"`,
						inlineKeyboards["turn_layer"])

				case "nothing":
					callbackText = "GG"
					tired = false
					body = fmt.Sprintf(`{"chat_id":%d,"message_id":%d,"text":%s,"reply_markup":%s}`,
						update.CallbackQuery.Message.Chat.ID,
						update.CallbackQuery.Message.MessageID,
						`"2nd layer"`,
						inlineKeyboards["something_layer"])

				default:
					body = fmt.Sprintf(`{"chat_id":%d,"message_id":%d,"text":%s,"reply_markup":%s}`,
						update.CallbackQuery.Message.Chat.ID,
						update.CallbackQuery.Message.MessageID,
						`"2nd layer"`,
						inlineKeyboards[update.CallbackQuery.Data])
				}

				_, err = b.Request(method, body)

				if err != nil {
					log.Println(err)
				}

			}
		}
	}
}
