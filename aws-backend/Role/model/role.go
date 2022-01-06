package model

type Role struct {
	RoleID   int      `json:"role_id" dynamodbav:"role_id"`
	RoleName string   `json:"role_name" dynamodbav:"role_name"`
	Auth     Auth     `json:"authority" dynamodbav:"authority"`
	Channel  []string `json:"role_channel" dynamodbav:"role_channel"`
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

type FullRole struct {
	RoleID   int      `json:"role_id" dynamodbav:"role_id"`
	RoleName string   `json:"role_name" dynamodbav:"role_name"`
	Auth     Auth     `json:"authority" dynamodbav:"authority"`
	Channel  []string `json:"role_channel" dynamodbav:"role_channel"`
	Total    int      `json:"total"`
}
