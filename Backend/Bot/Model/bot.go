package Model

import "time"

// type BotMessages struct {
// 	ID          string `json:"id"`
// 	BotName     string `json:"bot_name"`
// 	MessageDes  string `json:"description"`
// 	MessageBody string `json:"message"`

// 	UpdatedTime time.Time `json:"updated_time"`
// 	CreatedTime time.Time `json:"created_time"`
// }

type BotBody struct {
	ID            string    `json:"id" bson:"_id"`
	BotName       string    `json:"botname"`
	Organization  string    `json:"organization"`
	Folder        string    `json:"folder"`
	Activated     bool      `json:"actived"`
	Enabled       bool      `json:"enabled"`
	CreatedOn     time.Time `json:"created_on"`
	UpdatedOn     time.Time `json:"updated_on"`
	SenderId      string    `json:"sender_id"` //sender phone number (externalId in stella)
	Stages        interface{}
	Connections   interface{}
	Keyword       string `json:"keyword"`
	ChannelIcon   string `json:"channel_icon"`
	ChatsStarted  int    `json:"chats_starts"`
	ChatsFinished int    `json:"chats_finished"`
	Ctr           int    `json:"ctr"`
	StagesCtr     interface{}
	LastMessage   time.Time `json:"last_message"`
	TeamData      interface{}
	TeamId        string `json:"teamid"`
	TeamName      string `json:"teamname"`
}

type Stages struct {
	Name      string `json:"name"`
	Id        string `json:"id"`
	IsDefault bool   `json:"isdefault"`
	Postion   interface{}
	Actions   interface{}
}

type Actions struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	IsBot bool   `json:"isbot"`
	Data  interface{}
	Bot   string `json:"bot"`
}

type Data struct {
	Text       string `json:"text"`
	Attachment interface{}
	PlainText  string `json:"plain_text"`
}

type Connection struct {
	Source string `json:"source"`
	Target string `json:"target"`
}
