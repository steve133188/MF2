package Model

import "time"

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
	SenderNo   string    `json:"sender_no" bson:"sender_no"`
	ReceiverNo string    `json:"receiver_no" bson:"receiver_no"`
	Message    string    `json:"message" bson:"message"`
	SentTime   time.Time `json:"sent_time" bson:"sent_time"`
	CustomerID string    `json:"customer_id" bson:"customer_id"`
}
