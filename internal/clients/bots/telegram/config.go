package telegram

const (
	ENDPOINT              = `https://api.telegram.org/bot%s/%s`
	GET_UPDATES           = "getUpdates"
	SEND_MESSAGE          = "sendMessage"
	ANSWER_CALLBACK_QUERY = "answerCallbackQuery"
	EDIT_MESSAGE_TEXT     = "editMessageText"
	GET_ME                = "getMe"
)

type Request interface {
	getMethod() string
}

var (
	updateId = 0
	timeout  = 60
)

var (
	greetingsText = "This is button-bot for VK internship.\n" +
		"You can test vk bot using this url: in develop mode.\n" +
		"To start conversation, send any data."
	firstLayerText  = "1st layer"
	secondLayerText = "2nd layer"
)

var (
	urlLayerButton = &InlineKeyboardButton{
		Text:         "Url Layer",
		CallbackData: "url_layer",
	}
	switchLayerButton = &InlineKeyboardButton{
		Text:         "Switch Layer",
		CallbackData: "switch_layer",
	}
	turnLayerButton = &InlineKeyboardButton{
		Text:         "Turn Layer",
		CallbackData: "turn_layer",
	}
	somethingLayerButton = &InlineKeyboardButton{
		Text:         "Something Layer",
		CallbackData: "something_layer",
	}
	vkButton = &InlineKeyboardButton{
		Text: "vk",
		Url:  "vk.com",
	}
	backButton = &InlineKeyboardButton{
		Text:         "<-back",
		CallbackData: "root_layer",
	}
	switchToButton = &InlineKeyboardButton{
		Text:              "switchTo",
		SwitchInlineQuery: "null",
	}
	turnOffButton = &InlineKeyboardButton{
		Text:         "turn off WOW",
		CallbackData: "turn_off_WOW",
	}
	turnOnButton = &InlineKeyboardButton{
		Text:         "turn on WOW",
		CallbackData: "turn_on_WOW",
	}
	nothingButton = &InlineKeyboardButton{
		Text:         "nothing",
		CallbackData: "nothing",
	}
)
var (
	rootKeyboard = &InlineKeyboardMarkup{
		InlineKeyboard: [][]*InlineKeyboardButton{
			{urlLayerButton, switchLayerButton},
			{turnLayerButton, somethingLayerButton},
		}}
	urlKeyboard = &InlineKeyboardMarkup{
		InlineKeyboard: [][]*InlineKeyboardButton{
			{vkButton, backButton},
		}}
	switchKeyboard = &InlineKeyboardMarkup{
		InlineKeyboard: [][]*InlineKeyboardButton{
			{switchToButton, backButton},
		}}
	turnKeyboard = &InlineKeyboardMarkup{
		InlineKeyboard: [][]*InlineKeyboardButton{
			{turnOnButton, turnOffButton},
			{backButton},
		}}
	somethingKeyboard = &InlineKeyboardMarkup{
		InlineKeyboard: [][]*InlineKeyboardButton{
			{nothingButton, backButton},
		}}

	inlineKeyboards = map[string]*InlineKeyboardMarkup{
		"root_layer":      rootKeyboard,
		"url_layer":       urlKeyboard,
		"switch_layer":    switchKeyboard,
		"turn_layer":      turnKeyboard,
		"something_layer": somethingKeyboard,
	}
)
