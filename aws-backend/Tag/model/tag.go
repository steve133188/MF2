package model

type Tag struct {
	TagID     int    `json:"tag_id" dynamodbav:"tag_id"`
	TagName   string `json:"tag_name" dynamodbav:"tag_name"`
	CreatedAt string `json:"create_at" dynamodbav:"create_at"`
	UpdateAt  string `json:"update_at" dynamodbav:"update_at"`
}
