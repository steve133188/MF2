package Model

import "time"

type BoardCast struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Des  string `json:"description"`

	UserId       string `json:"user_id"`
	Username     string `json:"username"`
	CustomerId   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	Message      string `json:"message"`

	UpdatedTime time.Time `json:"updated_time"`
	CreatedTime time.Time `json:"created_time"`
}
