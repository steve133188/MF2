package model

type User struct {
	UserID        int    `json:"user_id" dynamodbav:"user_id"`
	Username      string `json:"username" dynamodbav:"username"`
	Email         string `json:"email" dynamodbav:"email"`
	Password      string `json:"password" dynamodbav:"password"`
	Phone         string `json:"phone" dynamodbav:"phone"`
	RoleID        int    `json:"role_id" dynamodbav:"role_id"`
	Leads         int    `json:"leads" dynamodbav:"leads"`
	Status        string `json:"user_status" dynamodbav:"user_status"`
	TeamID        int    `json:"team_id" dynamodbav:"team_id"`
	Channels      []Chan `json:"channels" dynamodbav:"channels"`
	Subscriptions []int  `json:"subscriptions" dynamodbav:"subscriptions"`
	CheckAuth     bool   `json:"check_auth" dynamodbav:"check_auth" default:"false"`
	CreateAt      string `json:"create_at" dynamodbav:"create_at"`
	LastLogin     string `json:"last_login" dynamodbav:"last_login"`
}

type Chan struct {
	ChannelName string `json:"channel_name" dynamodbav:"channel_name"`
	ChannelUrl  string `json:"channel_url" dynamodbav:"channel_url"`
}
