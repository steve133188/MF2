package model

type Message struct {
	RoomID      int    `json:"room_id" dynamodbav:"room_id"`
	TimeStamp   int64  `json:"timestamp" dynamodbav:"timestamp"`
	Receiver    string `json:"receiver" dynamodbav:"receiver"`
	Sender      string `json:"sender" dynamodbav:"sender"`
	MediaUrl    string `json:"media_url" dynamodbav:"media_url"`
	MessageType string `json:"message_type" dynamodbav:"message_type"`
	IsMedia     bool   `json:"is_media" dynamodbav:"is_media"`
	FromMe      bool   `json:"from_me" dynamodbav:"from_me"`
	Body        string `json:"body" dynamodbav:"body"`
}
