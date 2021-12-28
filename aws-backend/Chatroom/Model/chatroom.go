package Model

type ChatRoom struct {
	RoomID     int    `json:"room_id" dynamodbav:"room_id"`
	UserID     int    `json:"user_id" dynamodbav:"user_id"`
	Type       string `json:"type" dynamodbav:"type"`
	Channel    string `json:"channel" dynamodbav:"channel"`
	CustomerID int    `json:"customer_id" dynamodbav:"customer_id"`
}
