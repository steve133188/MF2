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

type Auth struct {
	Dashboard        bool `json:"dashboard" dynamodbav:"dashboard" default:"false"`
	Livechat         bool `json:"livechat" dynamodbav:"livechat" default:"false"`
	Contact          bool `json:"contact" dynamodbav:"contact" default:"false"`
	Broadcast        bool `json:"broadcast" dynamodbav:"broadcast" default:"false"`
	Flowbuilder      bool `json:"flowbuilder" dynamodbav:"flowbuilder" default:"false"`
	Integrations     bool `json:"integrations" dynamodbav:"integrations" default:"false"`
	ProductCatalogue bool `json:"product_catalogue" dynamodbav:"product_catalogue" default:"false"`
	Organization     bool `json:"organization" dynamodbav:"organization" default:"false"`
	Admin            bool `json:"admin" dynamodbav:"admin" default:"false"`
}

type Chan struct {
	ChannelName string `json:"channel_name" dynamodbav:"channel_name"`
	Url         string `json:"url" dynamodbav:"url"`
}
