package telegram

const (
	ENDPOINT              = `https://api.telegram.org/bot%s/%s`
	GET_UPDATES           = "getUpdates"
	SEND_MESSAGE          = "sendMessage"
	ANSWER_CALLBACK_QUERY = "answerCallbackQuery"
	EDIT_MESSAGE_TEXT     = "editMessageText"
)

var (
	updateId = 0
	timeout  = 60
)

var (
	inlineKeyboards = map[string]string{
		"root_layer":      `{"inline_keyboard":[[{"text":"Url Layer","callback_data":"url_layer"},{"text":"Switch Layer","callback_data":"switch_layer"}],[{"text":"Turn Layer","callback_data":"turn_layer"},{"text":"Something?","callback_data":"something_layer"}]]}`,
		"url_layer":       `{"inline_keyboard":[[{"text":"vk", "url":"vk.com"}],[{"text":"<-back", "callback_data":"root_layer"}]]}`,
		"switch_layer":    `{"inline_keyboard":[[{"text":"switch to", "switch_inline_query":"null"}],[{"text":"<-back", "callback_data":"root_layer"}]]}`,
		"turn_layer":      `{"inline_keyboard":[[{"text":"turn off WOW", "callback_data":"turn_off_WOW"},{"text":"turn on WOW", "callback_data":"turn_on_WOW"}],[{"text":"<-back", "callback_data":"root_layer"}]]}`,
		"something_layer": `{"inline_keyboard":[[{"text":"nothing", "callback_data": "nothing"}],[{"text":"<-back", "callback_data":"root_layer"}]]}`,
	}
)
