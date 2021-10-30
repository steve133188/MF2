package Model

import "time"

type Admin struct {
	ID          string `json:"id" bson:"id"`
	AdminName   string `json:"username"`
	Description string `json:"description"`
	LastAction  string `json:"last_action"`

	TargetUserId        string `json:"target_user_id"`
	TargetUsername      string `json:"target_username"`
	TargetUserPhone     string `json:"target_user_phone"`
	TargetCustomerId    string `json:"target_customer_id"`
	TargetCustomerName  string `json:"target_custmer_name"`
	TargetCustomerPhone string `json:"target_customer_phone"`

	UpdatedTime time.Time `json:"updated_time"`
	CreatedTime time.Time `json:"created_time"`
}
