package Model

type SystemLog struct {
	ID            string `json:"id"`
	Des           string `json:"description"` // system Activity
	UserId        string `json:"user_id"`
	SystemLogType string `json:"system_log_type"`
	Date          string `json:"date"`
}

type UserLog struct {
	ID          string `json:"id"`
	Des         string `json:"description"` //user Activity
	UserId      string `json:"user_id"`
	UserLogType string `json:"user_log_type"`
	Date        string `json:"date"`
}

type CustomerLog struct {
	ID              string `json:"id" bson:"id"`
	Des             string `json:"description" bson:"description"` // Activity
	CusId           string `json:"customer_id" bson:"customer_id"`
	CustomerLogType string `json:"type" bson:"type"`
	UserID          string `json:"user_id" bson:"user_id"` // system or user_id
	Date            string `json:"date" bson:"date"`
}
