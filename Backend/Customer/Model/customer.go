package Model

import "time"

type Customer struct {
	Id                 string    `json:"id"`
	UserId             string    `json:"user_d"`
	Username           string    `json:"username"`
	CustomerFirstName  string    `json:"customer_first_name"`
	CustomerLastName   string    `json:"customer_last_name"`
	Phone              string    `json:"phone"`
	Email              string    `json:"email"`
	TimeZone           string    `json:"timezone"`
	LastUpdatedTime    time.Time `json:"last_updated_time"` //date of updating customer info
	AccountCreatedTime time.Time `json:"account_created_time"`
}
