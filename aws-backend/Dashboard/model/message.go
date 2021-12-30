package model

type Message struct {
	RoomID      int    `json:"room_id" dynamodbav:"room_id"`
	TimeStamp   string `json:"timestamp" dynamodbav:"timestamp"`
	Receiver    string `json:"receiver" dynamodbav:"receiver"`
	Sender      string `json:"sender" dynamodbav:"sender"`
	MediaUrl    string `json:"media_url" dynamodbav:"media_url"`
	MessageType string `json:"message_type" dynamodbav:"message_type"`
	IsMedia     bool   `json:"is_media" dynamodbav:"is_media"`
	FromMe      bool   `json:"from_me" dynamodbav:"from_me"`
	Body        string `json:"body" dynamodbav:"body"`
}

type Tag struct {
	TagID     int    `json:"tag_id" dynamodbav:"tag_id"`
	TagName   string `json:"tag_name" dynamodbav:"tag_name"`
	CreatedAt int64  `json:"create_at" dynamodbav:"create_at"`
	UpdateAt  int64  `json:"update_at" dynamodbav:"update_at"`
}
