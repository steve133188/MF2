package model

type Message struct {
	RoomID      string `json:"room_id"`
	Timestamp   string `json:"timestamp"`
	Status      string `json:"status"`
	MessageType string `json:"message_type"`
	HasQuoteMsg bool   `json:"hasQuotedMsg"`
	IsMedia     bool   `json:"is_media"`
	MessageId   string `json:"message_id"`
	SignName    string `json:"sign_name"`
	Channel     string `json:"channel"`
	MediaUrl    string `json:"media_url"`
	Sender      string `json:"sender"`
	Recipient   string `json:"recipient"`
	Read        bool   `json:"read"`
	IsForwarded bool   `json:"is_forwarded"`
	FromMe      bool   `json:"from_me"`
	Link        string `json:"link"`
	Body        string `json:"body"`
	Quote       string `json:"quote"`
	//VCard       []interface{} `json:"v_card"`
}
