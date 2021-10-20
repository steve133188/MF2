package Model

import "time"

type SystemLog struct {
	ID     string    `json:"id"`
	Des    string    `json:"description"` // system Activity
	UserId string    `json:"user_id"`
	Date   time.Time `json:"date"`
}

type UserLog struct {
	ID     string    `json:"id"`
	Des    string    `json:"description"` //user Activity
	UserId string    `json:"user_id"`
	Date   time.Time `json:"date"`
}

type CustomerLog struct {
	ID      string    `json:"id"`
	Des     string    `json:"description"` // Activity
	CusId   string    `json:"customer_id"`
	Handler string    `json:"handler"` // system or user_id
	Date    time.Time `json:"date"`
}
