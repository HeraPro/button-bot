package telegram

import (
	"encoding/json"
)

type ApiResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result,omitempty"`
	ErrorCode   int             `json:"error_code,omitempty"`
	Description string          `json:"description,omitempty"`
	// Parameters (read tg api spec)
}

type GetMeRequest struct {
}

type GetUpdatesRequest struct {
	Offset         int      `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

type SendMessageRequest struct {
	//may be string in case "@channelusername"
	ChatID                   int              `json:"chat_id"`
	MessageThreadId          int              `json:"message_thread_id,omitempty"`
	Text                     string           `json:"text"`
	ParseMode                string           `json:"parse_mode,omitempty"`
	Entities                 []*MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool             `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool             `json:"disable_notification,omitempty"`
	ProtectContent           bool             `json:"protect_content,omitempty"`
	ReplyToMessageID         int              `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool             `json:"allow_sending_without_reply,omitempty"`
	//InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, ForceReply
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type AnswerCallbackQueryRequest struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	Url             string `json:"url,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
}

type EditMessageTextRequest struct {
	// string in case @channelusername
	ChatID                int                   `json:"chat_id,omitempty"`
	MessageID             int                   `json:"message_id,omitempty"`
	InlineMessageID       string                `json:"inline_message_id,omitempty"`
	Text                  string                `json:"text"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	Entities              []*MessageEntity      `json:"entities,omitempty"`
	DisableWebPagePreview bool                  `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type Update struct {
	UpdateID          int            `json:"update_id"`
	Message           *Message       `json:"message,omitempty"`
	EditedMessage     *Message       `json:"edited_message,omitempty"`
	ChannelPost       *Message       `json:"channel_post,omitempty"`
	EditedChannelPost *Message       `json:"edited_channel_post,omitempty"`
	InlineQuery       *Message       `json:"inline_query,omitempty"`
	CallbackQuery     *CallbackQuery `json:"callback_query,omitempty"`
	//ChosenInlineResult *InlineQuery   `json:"chosen_inline_result,omitempty"`
	//ShippingQuery      *ShippingQuery     `json:"shipping_query,omitempty"`
	//PreCheckoutQuery   *PreCheckoutQuery  `json:"pre_checkout_query,omitempty"`
	//Poll               *Poll              `json:"poll,omitempty"`
	//PollAnswer         *PollAnswer        `json:"poll_answer,omitempty"`
	//MyChatMember       *ChatMemberUpdated `json:"my_chat_member,omitempty"`
	//ChatMember         *ChatMemberUpdated `json:"chat_member,omitempty"`
	//ChatJoinRequest    *ChatJoinRequest   `json:"chat_join_request"`
}

type CallbackQuery struct {
	ID      string   `json:"id"`
	From    *User    `json:"from"`
	Message *Message `json:"message,omitempty"`
	Data    string   `json:"data,omitempty"`
	//InlineMessageID string   `json:"inlineMessageID,omitempty"`
	//ChatInstance    string   `json:"chatInstance"`
	//GameShortName   string   `json:"gameShortName,omitempty"`
}

type Message struct {
	MessageID   int                   `json:"message_id"`
	From        *User                 `json:"from,omitempty"`
	Chat        *Chat                 `json:"chat"`
	Date        int                   `json:"date"`
	Text        string                `json:"text,omitempty"`
	Entities    []*MessageEntity      `json:"entities,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"replyMarkup,omitempty"`
	//MessageThreadID      int                   `json:"message_thread_id,omitempty"`
	//SenderChat           *Chat                 `json:"sender_chat,omitempty"`
	//ForwardFrom          *User                 `json:"forward_from,omitempty"`
	//ForwardFromChat      *Chat                 `json:"forward_from_chat,omitempty"`
	//ForwardFromMessageID int                   `json:"forward_from_message_id,omitempty"`
	//ForwardSignature     string                `json:"forward_signature,omitempty"`
	//ForwardSenderName    string                `json:"forward_sender_name,omitempty"`
	//ForwardDate          int                   `json:"forward_date,omitempty"`
	//IsTopicMessage       bool                  `json:"is_topic_message,omitempty"`
	//IsAutomaticForward   bool                  `json:"is_automatic_forward,omitempty"`
	//ReplyToMessage       *Message              `json:"reply_to_message,omitempty"`
	//ViaBot               *User                 `json:"via_bot,omitempty"`
	//EditDate             int                   `json:"edit_date,omitempty"`
	//HasProtectedContent  bool                  `json:"has_protected_content,omitempty"`
	//MediaGroupId         string                `json:"media_group_id,omitempty"`
	//AuthorSignature      string                `json:"author_signature,omitempty"`
	//Animation                     *Animation                     `json:"animation,omitempty"`
	//Audio                         *Audio                         `json:"audio"`
	//Document                      *Document                      `json:"document,omitempty"`
	//Photo                         []*PhotoSize                   `json:"photo,omitempty"`
	//Sticker                       *Sticker                       `json:"sticker,omitempty"`
	//Video                         *Video                         `json:"video,omitempty"`
	//VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	//Voice                         *Voice                         `json:"voice,omitempty"`
	//Caption         string           `json:"caption,omitempty"`
	//CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	//HasMediaSpoiler bool             `json:"has_media_spoiler,omitempty"`
	//Contact                       *Contact                       `json:"contact,omitempty"`
	//Dice                          *Dice                          `json:"dice,omitempty"`
	//Game                          *Game                          `json:"game,omitempty"`
	//Poll                          *Poll                          `json:"poll,omitempty"`
	//Venue                         *Venue                         `json:"venue,omitempty"`
	//Location                      *Location                      `json:"location,omitempty"`
	//NewChatMembers []*User `json:"new_chat_members,omitempty"`
	//LeftChatMember *User   `json:"left_chat_member,omitempty"`
	//NewChatTitle   string  `json:"new_chat_title,omitempty"`
	//NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo,omitempty"`
	//DeleteChatPhoto       bool `json:"delete_chat_photo,omitempty"`
	//GroupChatCreated      bool `json:"group_chat_created,omitempty"`
	//SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`
	//ChannelChatCreated    bool `json:"channel_chat_created,omitempty"`
	//MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	//MigrateToChatID   int      `json:"migrate_to_chat_id,omitempty"`
	//MigrateFromChatID int      `json:"migrate_from_chat_id,omitempty"`
	//PinnedMessage     *Message `json:"pinned_message,omitempty"`
	//Invoice                       *Invoice                       `json:"invoice,omitempty"`
	//SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	//UserShared                    *UserShared                    `json:"user_shared,omitempty"`
	//ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`
	//ConnectedWebsite string `json:"connected_website,omitempty"`
	//WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`
	//PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	//ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	//ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`
	//ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	//ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`
	//ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`
	//GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`
	//GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`
	//VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	//VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`
	//VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`
	//VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	//WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	Url           string `json:"url,omitempty"`
	User          *User  `json:"user,omitempty"`
	Language      string `json:"language,omitempty"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

type User struct {
	ID                       int    `json:"id"`
	IsBot                    bool   `json:"is_bot"`
	FirstName                string `json:"first_name,omitempty"`
	LastName                 string `json:"last_name,omitempty"`
	Username                 string `json:"username"`
	LanguageCode             string `json:"language_code"`
	IsPremium                bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu    bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups            bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupsMessages bool   `json:"can_read_all_groups_messages,omitempty"`
	SupportsInlineQueries    bool   `json:"supports_inline_queries,omitempty"`
}

//todo add other attributes
type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type InlineKeyboardButton struct {
	Text                         string `json:"text"`
	Url                          string `json:"url,omitempty"`
	CallbackData                 string `json:"callback_data,omitempty"`
	SwitchInlineQuery            string `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`
	//LoginUrl                     *LoginUrl `json:"login_url,omitempty"`
	//WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	//SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	//CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	//Pay                          bool                         `json:"pay,omitempty"`
}
