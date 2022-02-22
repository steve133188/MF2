package model

type Activity struct {
	TimeStamp  string `json:"timestamp" dynamodbav:"timestamp"`
	Payload    string `json:"payload" dynamodbav:"payload"`
	UserID     int    `json:"user_id" dynamodbav:"user_id"`
	Action     string `json:"action" dynamodbav:"action"`
	Status     string `json:"status" dynamodbav:"status"`
	CustomerID int    `json:"customer_id" dynamodbav:"customer_id"`
	IsSys      bool   `json:"is_sys" dynamodbav:"is_sys"`
	Type       string `json:"type" dynamodbav:"type"`
}
