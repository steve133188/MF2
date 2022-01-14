package Model

// type Message struct {
// 	MediaKey  string      `json:"media_key" bson:"media_key"`
// 	Id        interface{} `json:"id" bson:"id"`
// 	Ack       int         `json:"ack" bson:"ack"`
// 	HasMedia  bool        `json:"has_media" bson:"has_media"`
// 	Body      interface{} `json:"body" bson:"body"`
// 	Type      string      `json:"type" bson:"type"`
// 	TimeStamp uint64      `json:"timestamp" bson:"timestamp"`
// 	From      string      `json:"from" bson:"from"`
// 	To        string      `json:"to" bson:"to"`
// 	VCards    interface{} `json:"vcards" bson:"vcards"`
// }

type Chat struct {
	RoomID []string `json:"room_id" bson:"room_id"`
	UserID string   `json:"user_id" bson:"user_id"`
}

// type ClientReq struct {
// 	UserID      string      `json:"user_id"`
// 	ChannelType string      `json:"channel_type"`
// 	Msg         interface{} `json:"message"`
// }

type ClientMsg struct {
	Topic       string      `json:"topic" bson:"topic"`
	UserID      string      `json:"user_id"`
	Phone       string      `json:"phone"`
	ChannelType string      `json:"channel_type"`
	ChatID      string      `json:"chat_id"`
	Url         string      `json:"url"`
	Token       string      `json:"token"`
	Msg         interface{} `json:"message"`
}

type WebhookMsg struct {
	Status bool      `json:"status"`
	Resp   ClientMsg `json:"response"`
}
