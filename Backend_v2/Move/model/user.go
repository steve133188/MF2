package model

type FullUser struct {
	UserID      int    `json:"user_id" dynamodbav:"user_id"`
	Username    string `json:"username" dynamodbav:"username"`
	Email       string `json:"email" dynamodbav:"email"`
	Password    string `json:"password" dynamodbav:"password"`
	Phone       int    `json:"phone" dynamodbav:"phone"`
	CountryCode int    `json:"country_code" dynamodbav:"country_code"`
	Status      string `json:"user_status" dynamodbav:"user_status"`
	TeamID      int    `json:"team_id" dynamodbav:"team_id"`
	Team        Team   `json:"team" dynamodbav:"team"`
	Leads       int    `json:"leads" `
	RoleID      int    `json:"role_id" dynamodbav:"role_id"`
	RoleName    string `json:"role_name" dynamodbav:"role_name"`
	Authority   Auth   `json:"authority" dynamodbav:"authority"`
	Channels    []Chan `json:"channels" dynamodbav:"channels"`
	CheckAuth   bool   `json:"check_auth" dynamodbav:"check_auth" default:"false"`
	CreateAt    int64  `json:"create_at" dynamodbav:"create_at"`
	LastLogin   int64  `json:"last_login" dynamodbav:"last_login"`
	ActivityLog int    `json:"activity_log" dynamodbav:"activity_log"`
	IsBot       bool   `json:"is_bot" dynamodbav:"is_bot" default:"true"`
}

type Node struct {
	NodeIndex int    `json:"node_index" dynamodbav:"node_index"`
	UserId    int    `json:"user_id" dynamodbav:"user_id"`
	ChannelId string `json:"channel_id" dynamodbav:"channel_id"`
	Status    string `json:"status" dynamodbav:"status"`
	Url       string `json:"url" dynamodbav:"url"`
	NodeName  string `json:"node_name" dynamodbav:"node_name"`
	Init      bool   `json:"init" dynamodbav:"init"`
	NodeId    string `json:"node_id" dynamodbav:"node_id"`
}
