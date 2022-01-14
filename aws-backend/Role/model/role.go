package model

type Role struct {
	RoleID   int    `json:"role_id" dynamodbav:"role_id"`
	RoleName string `json:"role_name" dynamodbav:"role_name"`
	Auth     Auth   `json:"authority" dynamodbav:"authority"`
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
	Whatsapp         bool `json:"whatsapp" dynamodbav:"whatsapp"`
	WABA             bool `json:"waba" dynamodbav:"waba"`
	Messager         bool `json:"messager" dynamodbav:"messager"`
	WeChat           bool `json:"wechat" dynamodbav:"wechat"`
}

type FullRole struct {
	RoleID   int      `json:"role_id" dynamodbav:"role_id"`
	RoleName string   `json:"role_name" dynamodbav:"role_name"`
	Auth     Auth     `json:"authority" dynamodbav:"authority"`
	Channel  []string `json:"role_channel" dynamodbav:"role_channel"`
	Total    int      `json:"total"`
}

type User struct {
	UserID      int    `json:"user_id" dynamodbav:"user_id"`
	Username    string `json:"username" dynamodbav:"username"`
	Email       string `json:"email" dynamodbav:"email"`
	Password    string `json:"password" dynamodbav:"password"`
	Phone       int    `json:"phone" dynamodbav:"phone"`
	CountryCode int    `json:"country_code" dynamodbav:"country_code"`
	RoleID      int    `json:"role_id" dynamodbav:"role_id"`
	Status      string `json:"user_status" dynamodbav:"user_status"`
	TeamID      int    `json:"team_id" dynamodbav:"team_id"`
	Channels    []Chan `json:"channels" dynamodbav:"channels"`
	CheckAuth   bool   `json:"check_auth" dynamodbav:"check_auth" default:"false"`
	CreateAt    int64  `json:"create_at" dynamodbav:"create_at"`
	LastLogin   int64  `json:"last_login" dynamodbav:"last_login"`
	ActivityLog int    `json:"activity_log" dynamodbav:"activity_log"`
}

type Chan struct {
	ChannelName string `json:"channel_name" dynamodbav:"channel_name"`
	ChannelUrl  string `json:"channel_url" dynamodbav:"channel_url"`
}
