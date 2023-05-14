package telegram

//some helper functions for template

func (r GetUpdatesRequest) getMethod() string {
	return GET_UPDATES
}
func (r SendMessageRequest) getMethod() string {
	return SEND_MESSAGE
}

func (r AnswerCallbackQueryRequest) getMethod() string {
	return ANSWER_CALLBACK_QUERY
}

func (r EditMessageTextRequest) getMethod() string {
	return EDIT_MESSAGE_TEXT
}

func (r GetMeRequest) getMethod() string {
	return GET_ME
}
