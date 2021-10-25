package Model

import (
	"time"
)

// type Message struct {
// 	Id                string    `json:"id"`
// 	SenderUsername    string    `json:"sender_username"`
// 	SenderUserId      string    `json:"sender_user_id"`
// 	SenderUserPhone   string    `json:"sender_user_phone"`
// 	ReceiverUserName  string    `json:"receiver_user_name"`
// 	ReceiverUserId    string    `json:"receiver_user_id"`
// 	ReceiverUserPhone string    `json:"receiver_user_phone"`
// 	MessageType       string    `json:"content_type"`
// 	UpdatedTime       time.Time `json:"updated_time"`
// 	CreatedTime       time.Time `json:"created_time"`
// }
type Message struct {
	// Id             primitive.ObjectID `bson:"_id" json:"id"`
	OldId          string    `bson:"old_id" json:"_id"`
	Conversation   string    `json:"conversation"`
	Sender         string    `json:"sender"`
	Text           string    `json:"text"`
	DateTime       time.Time `json:"datetime"`
	Enabled        bool      `json:"enabled"`
	Status         string    `json:"status"`
	PlainText      string    `json:"plaintext"`
	WhatsapppMsgId string    `json:"whatsappmsgid"`
	SentStatus     string    `json:"sentstatus"`
}
