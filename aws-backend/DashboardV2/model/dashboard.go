package model

type Dashboard struct {
	PK        string
	TimeStamp int64
	User      []UserInfo
}

type UserInfo struct {
	ChannelName string
	UserID      int
	Data        UserData
}

type UserData struct {
	UserName            string `json:"user_name" dynamodbav:"user_name"`
	UserRoleName        string `json:"user_role_name" dynamodbav:"user_role_name"`
	UserStatus          string `json:"user_status" dynamodbav:"user_status"`
	LastLogin           int64  `json:"last_login" dynamodbav:"last_login"`
	AssignedContacts    int    `json:"assigned_contacts" dynamodbav:"assigned_contacts"`
	ActiveContacts      int    `json:"active_contacts" dynamodbav:"active_contacts"`
	UnhandledContact    int    `json:"unhandled_contact" dynamodbav:"unhandled_contact"`
	AvgRespTime         int64  `json:"avg_resp_time" dynamodbav:"avg_resp_time"`
	AvgFirstRespTime    int64  `json:"avg_first_resp_time" dynamodbav:"avg_first_resp_time"`
	AllContacts         int    `json:"all_contacts" dynamodbav:"all_contacts"`
	TotalMsgSent        int    `json:"total_msg_sent" dynamodbav:"total_msg_sent"`
	TotalMsgRev         int    `json:"total_msg_rev" dynamodbav:"total_msg_rev"`
	CommunicationNumber int    `json:"communication_number" dynamodbav:"communication_number"`
	Tags                []Tags `json:"tags" dynamodbav:"tags"`
}

type Tags struct {
	Name string `json:"name" dynamodbav:"name"`
	No   int    `json:"no" dynamodbav:"no"`
}
