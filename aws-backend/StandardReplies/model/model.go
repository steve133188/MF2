package model

type StandardReplies struct {
	ID        string   `json:"id" dynamodbav:"id"`
	Name      string   `json:"name" dynamodbav:"name"`
	Body      []string `json:"body" dynamodbav:"body"`
	Variables []string `json:"variables" dynamodbav:"variables"`
	CreatedAt string   `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt string   `json:"updated_at" dynamodbav:"updated_at"`
	Channels  []string `json:"channels" dynamodbav:"channels"`
}
