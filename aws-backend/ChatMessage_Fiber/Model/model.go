package Model

type Message struct {
	RoomId    string ` json:"room_id" dynamodbav:"room_id"`
	Timestamp string `json:"timestamp" dynamodbav:"timestamp"`
	Status    string `json:"status" dynamodbav:"status"`

	MessageId   string `json:"message_id" dynamodbav:"message_id"`
	MessageType string `json:"message_type" dynamodbav:"message_type"`
	Channel     string `json:"channel" dynamodbav:"channel"`
	SignName    string `json:"sign_name" dynamodbav:"sign_name"`

	Sender    string `json:"sender" dynamodbav:"sender"`
	Recipient string `json:"recipient" dynamodbav:"recipient"`

	Link     string `json:"link" dynamodbav:"link"`
	Body     string `json:"body" dynamodbav:"body"`
	MediaUrl string `json:"media_url" dynamodbav:"media_url"`
	Quote    string `json:"quote" dynamodbav:"quote"`

	HasQuotedMsg bool `json:"hasQuotedMsg" dynamodbav:"hasQuotedMsg"`
	IsMedia      bool `json:"is_media" dynamodbav:"is_media"`
	Read         bool `json:"read" dynamodbav:"read"`
	IsForwarded  bool `json:"is_forwarded" dynamodbav:"is_forwarded"`
	FromMe       bool `json:"from_me" dynamodbav:"from_me"`

	VCard interface{} `json:"v_card" dynamodbav:"v_card"`
}
