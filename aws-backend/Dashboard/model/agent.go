package model

type Agent struct {
	PK        string            `json:"pk" dynamodbav:"PK"`
	TimeStamp int64             `json:"time_stamp" dynamodbav:"timestamp"`
	Agents    map[int]AgentInfo `json:"agents" dynamodbav:"agents"`
}

//AssignedContacts: customer table assignee == user
//ActiveContacts: no. of customer with communication
//UnhandledContact: customer -> user, user unread
type AgentInfo struct {
	UserName         string `json:"user_name" dynamodbav:"user_name"`
	UserRoleName     string `json:"user_role_name" dynamodbav:"user_role_name"`
	UserStatus       string `json:"user_status" dynamodbav:"user_status"`
	LastLogin        int64  `json:"last_login" dynamodbav:"last_login"`
	AssignedContacts int    `json:"assigned_contacts" dynamodbav:"assigned_contacts"`
	ActiveContacts   int    `json:"active_contacts" dynamodbav:"active_contacts"`
	UnhandledContact int    `json:"unhandled_contact" dynamodbav:"unhandled_contact"`
	TotalMsgSent     int    `json:"total_msg_sent" dynamodbav:"total_msg_sent"`
	AvgRespTime      int64  `json:"avg_resp_time" dynamodbav:"avg_resp_time"`
	AvgFirstRespTime int64  `json:"avg_first_resp_time" dynamodbav:"avg_first_resp_time"`
}

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
}
