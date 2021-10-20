package Model

import "time"

type Bot struct {
	ID           string `json:"id"`
	BotName      string `json:"bot_name"`
	Des          string `json:"description"`
	UserId       string `json:"user_id"`
	Username     string `json:"username"`
	UserDes      string `json:"user_description"`
	CustomerId   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	CustomerDes  string `json:"customer_description"`

	UpdatedTime time.Time `json:"updated_time"`
	CreatedTime time.Time `json:"created_time"`
}

type BotBody struct {
	BotName       string    `json:"botname"`
	Organization  string    `json:"organization"`
	Folder        string    `json:"folder"`
	Activated     bool      `json:"actived"`
	Enabled       bool      `json:"enabled"`
	CreatedOn     time.Time `json:"created_on"`
	UpdatedOn     time.Time `json:"updated_on"`
	SenderId      string    `json:"sender_id"` //sender phone number (externalId in stella)
	Stages        interface{}
	Connections   struct{}
	Keyword       string `json:"keyword"`
	ChannelIcon   string `json:"channel_icon"`
	ChatsStarted  int    `json:"chats_starts"`
	ChatsFinished int    `json:"chats_finished"`
	Ctr           int    `json:"ctr"`
	StagesCtr     struct{}
	LastMessage   time.Time `json:"last_message"`
	TeamData      struct{}
	TeamId        string `json:"team_id"`
	TeamName      string `json:"team_name"`
}

type Stages struct {
	Name      string `json:"name"`
	Id        string `json:"id"`
	IsDefault bool   `json:"is_default"`
	Postion   struct{}
	Actions   struct {
		Name  string `json:"actions_name"`
		Id    string `json:"actions_id"`
		IsBot bool   `json:"is_bot"`
		Bot   string `json:"bot"`
		Data  struct{}
	}
}
