package model

type Customer struct {
	CustomerID int      `json:"customer_id" dynamodbav:"customer_id"`
	Name       string   `json:"customer_name" dynamodbav:"customer_name"`
	Email      string   `json:"email" dynamodbav:"email"`
	FirstName  string   `json:"first_name" dynamodbav:"first_name"`
	LastName   string   `json:"last_name" dynamodbav:"last_name"`
	Phone      string   `json:"phone" dynamodbav:"phone"`
	Channels   []string `json:"channels" dynamodbav:"channels"`
	TeamID     int      `json:"team_id" dynamodbav:"team_id"`
	AgentsID   []int    `json:"agents_id" dynamodbav:"agents_id"`
	TagsID     []int    `json:"tags_id" dynamodbav:"tags_id"`
	Group      string   `json:"customer_group" dynamodbav:"customer_group"`
	Birthday   string   `json:"birthday" dynamodbav:"birthday"`
	Country    string   `json:"country" dynamodbav:"country"`
	Address    string   `json:"address" dynamodbav:"address"`
	Gender     string   `json:"gender" dynamodbav:"gender"`
	CreatedAt  int64    `json:"created_at" dynamodbav:"created_at"`
	UpdateAt   int64    `json:"update_at" dynamodbav:"update_at"`
}
