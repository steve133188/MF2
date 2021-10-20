package Model

import "time"

type SystemLog struct {
	ID            string    `json:"id"`
	Des           string    `json:"description"` // system Activity
	UserId        string    `json:"user_id"`
	SystemLogType string    `json:"system_log_type"`
	Date          time.Time `json:"date"`
}

type UserLog struct {
	ID          string    `json:"id"`
	Des         string    `json:"description"` //user Activity
	UserId      string    `json:"user_id"`
	UserLogType string    `json:"user_log_type"`
	Date        time.Time `json:"date"`
}

type CustomerLog struct {
	ID              string    `json:"id"`
	Des             string    `json:"description"` // Activity
	CusId           string    `json:"customer_id"`
	CustomerLogType string    `json:"customer_log_type"`
	Handler         string    `json:"handler"` // system or user_id
	Date            time.Time `json:"date"`
}
