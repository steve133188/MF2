package Model

import "time"

type Analysis struct {
	ID           string `json:"id" bson:"id"`
	UserId       string `json:"user_id"`
	Username     string `json:"username"`
	CustomerId   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	Description  string `json:"description"`
	Duration     int    `json:"duration"`

	UpdatedTime time.Time `json:"updated_time"`
	CreatedTime time.Time `json:"created_time"`
}
