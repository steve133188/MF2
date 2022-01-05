package model

type Chatroom struct {
	RoomID int `json:"room_id" dynamodbav:"room_id"`
	UserID int `json:"user_id" dynamodbav:"user_id"`

	Unread int  `json:"unread" dynamodbav:"unread"`
	IsPin  bool `json:"is_pin" dynamodbav:"is_pin"`

	CustomerID string `json:"customer_id" dynamodbav:"customer_id"`
	Name       string `json:"name" dynamodbav:"name"`
	Phone      string `json:"phone" dynamodbav:"phone"`
	Channel    string `json:"channel" dynamodbav:"channel"`
}
