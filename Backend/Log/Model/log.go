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
	Activity     string `json:"activity" bson:"activity"` // Activity
	CustomerName string `json:"customer_name" bson:"customer_name"`
	UserName     string `json:"user_name" bson:"user_name"` // system or user_id
	Date         string `json:"date" bson:"date"`
}
