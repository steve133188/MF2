package model

type Chatroom struct {
	Channel     string `json:"channel" dynamodbav:"channel" `
	RoomID      string `json:"room_id" dynamodbav:"room_id"`
	UserID      int    `json:"user_id" dynamodbav:"user_id"`
	CustomerID  int    `json:"customer_id" dynamodbav:"customer_id"`
	Unread      int    `json:"unread" dynamodbav:"unread"`
	IsPin       bool   `json:"is_pin" dynamodbav:"is_pin"`
	Phone       string `json:"phone" dynamodbav:"phone"`
	Name        string `json:"name" dynamodbav:"name"`
	CountryCode int    `json:"country_code" dynamodbav:"country_code"`
	Avatar      string `json:"avatar" dynamodbav:"avatar"`
	LastMsgTime string `json:"last_msg_time" dynamodbav:"last_msg_time"`
	BotOn       bool   `json:"bot_on" dynamodbav:"bot_on"`
	TeamID      int    `json:"team_id" dynamodbav:"team_id"`
}
