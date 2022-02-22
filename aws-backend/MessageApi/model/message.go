package model

type Message struct {
	RoomID      string        `json:"room_id" dynamodbav:"room_id"`
	Timestamp   string        `json:"timestamp" dynamodbav:"timestamp"`
	Status      string        `json:"status" dynamodbav:"status"`
	MessageType string        `json:"message_type" dynamodbav:"message_type"`
	HasQuoteMsg bool          `json:"hasQuotedMsg" dynamodbav:"hasQuotedMsg"`
	IsMedia     bool          `json:"is_media" dynamodbav:"is_media"`
	MessageId   string        `json:"message_id" dynamodbav:"message_id"`
	SignName    string        `json:"sign_name" dynamodbav:"sign_name"`
	Channel     string        `json:"channel" dynamodbav:"channel"`
	MediaUrl    string        `json:"media_url" dynamodbav:"media_url"`
	Sender      string        `json:"sender" dynamodbav:"sender"`
	Recipient   string        `json:"recipient" dynamodbav:"recipient"`
	Read        bool          `json:"read" dynamodbav:"read"`
	IsForwarded bool          `json:"is_forwarded" dynamodbav:"is_forwarded"`
	FromMe      bool          `json:"from_me" dynamodbav:"from_me"`
	Link        string        `json:"link" dynamodbav:"link"`
	Body        string        `json:"body" dynamodbav:"body"`
	Quote       string        `json:"quote" dynamodbav:"quote"`
	VCard       []interface{} `json:"v_card" dynamodbav:"v_card"`
}
