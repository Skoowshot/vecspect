package domain

type Update struct {
	UpdateId int64    `json:"update_id"`
	Message  *Message `json:"message"`
}

type Updates struct {
	Result []Update `json:"result"`
}

func (u Updates) LastUpdateId(def int64) int64 {
	if len(u.Result) == 0 {
		return def
	}
	return u.Result[len(u.Result)-1].UpdateId
}

type Message struct {
	MessageId int64    `json:"message_id"`
	Chat      Chat     `json:"chat"`
	Date      int      `json:"date"`
	Text      string   `json:"text"`
	FromUser  User     `json:"from"`
	ReplyTo   *Message `json:"reply_to_message"`
}

type Chat struct {
	ChatId int64  `json:"id"`
	Type   string `json:"type"`
}

type User struct {
	Id           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}
