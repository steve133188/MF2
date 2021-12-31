package model

type Dashboard struct {
	PK        string    `json:"PK" dynamodbav:"PK"`
	TimeStamp int64     `json:"timestamp" dynamodbav:"timestamp"`
	Channel   []Channel `json:"channel" dynamodbav:"Channel"`
}

type Channel struct {
	User []UserInfo

	ChannelName string `json:"channel_name" dynamodbav:"channel_name"`

	AvgTotalRespTime         int64 `json:"avg_resp_time" dynamodbav:"avg_resp_time"`
	AvgTotalFirstRespTime    int64 `json:"avg_total_first_resp_time" dynamodbav:"avg_total_first_resp_time"`
	TotalMsgSent             int   `json:"total_msg_sent" dynamodbav:"total_msg_sent"`
	TotalMsgRev              int   `json:"total_msg_rev" dynamodbav:"total_msg_rev"`
	TotalCommunicationNumber int   `json:"total_communication_number" dynamodbav:"total_communication_number"`
}

type UserInfo struct {
	UserID       int    `json:"user_id" dynamodbav:"user_id"`
	UserName     string `json:"user_name" dynamodbav:"user_name"`
	UserRoleName string `json:"user_role_name" dynamodbav:"user_role_name"`
	UserStatus   string `json:"user_status" dynamodbav:"user_status"`
	LastLogin    int64  `json:"last_login" dynamodbav:"last_login"`

	AssignedContacts int    `json:"assigned_contacts" dynamodbav:"assigned_contacts"`
	ActiveContacts   int    `json:"active_contacts" dynamodbav:"active_contacts"`
	UnhandledContact int    `json:"unhandled_contact" dynamodbav:"unhandled_contact"`
	AllContacts      int    `json:"all_contacts" dynamodbav:"all_contacts"`
	Tags             []Tags `json:"tags" dynamodbav:"tags"`

	//ChannelData              []ChannelData ` json:"channel_data" dynamodbav:"channel_data" `
	AvgRespTime     int64 `json:"avg_resp_time" dynamodbav:"avg_resp_time"`
	FirstRespTime   int64 `json:"first_resp_time" dynamodbav:"first_resp_time"`
	LongestRespTime int64 `json:"longest_resp_time" dynamodbav:"longest_resp_time"`

	MsgSent             int `json:"msg_sent" dynamodbav:"msg_sent"`
	MsgRev              int `json:"msg_rev" dynamodbav:"msg_rev"`
	CommunicationNumber int `json:"communication_number" dynamodbav:"communication_number"`
}

type Tags struct {
	Name string `json:"name" dynamodbav:"name"`
	No   int    `json:"no" dynamodbav:"no"`
}

type RespTime struct {
	Longest int64 `json:"longest" dynamodbav:"longest"`
	Average int64 `json:"average" dynamodbav:"average"`
	First   int64 `json:"first" dynamodbav:"first"`
}
