package model

type Dashboard struct {
	PK        string     `json:"PK" dynamodbav:"PK"`
	TimeStamp int64      `json:"timestamp" dynamodbav:"timestamp"`
	Channel   []Channel  `json:"channel" dynamodbav:"Channel"`
	User      []UserInfo `json:"user" dynamodbav:"user"`
	Tags      []Tags     `json:"tags" dynamodbav:"tags"`

	AvgTotalRespTime      int64 `json:"avg_resp_time" dynamodbav:"avg_resp_time"`
	AvgTotalFirstRespTime int64 `json:"avg_total_first_resp_time" dynamodbav:"avg_total_first_resp_time"`
	AvgLongestRespTime    int64 `json:"longest_resp_time" dynamodbav:"longest_resp_time"`

	TotalMsgSent       int   `json:"total_msg_sent" dynamodbav:"total_msg_sent"`
	TotalMsgRev        int   `json:"total_msg_rev" dynamodbav:"total_msg_rev"`
	CommunicationHours []int `json:"communication_hours" dynamodbav:"communication_hours"`
	//TotalCommunicationNumber int   `json:"total_communication_number" dynamodbav:"total_communication_number"`

	AllContacts            int `json:"all_contacts" dynamodbav:"all_contacts"`
	NewAddedContacts       int `json:"new_added_contacts" dynamodbav:"new_added_contacts"`
	TotalAssignedContacts  int `json:"total_assigned_contacts" dynamodbav:"total_assigned_contacts"`
	TotalActiveContacts    int `json:"total_active_contacts" dynamodbav:"total_active_contacts"`
	TotalDeliveredContacts int ` json:"total_delivered_contacts" dynamodbav:"total_delivered_contacts"`
	TotalUnhandledContact  int `json:"total_unhandled_contact" dynamodbav:"total_unhandled_contact"`
}

type Channel struct {
	User []UserInfo `json:"user" dynamodbav:"user"`

	ChannelName string `json:"channel_name" dynamodbav:"channel_name"`

	AvgTotalRespTime      int64 `json:"avg_resp_time" dynamodbav:"avg_resp_time"`
	AvgTotalFirstRespTime int64 `json:"avg_total_first_resp_time" dynamodbav:"avg_total_first_resp_time"`
	AvgLongestRespTime    int64 `json:"longest_resp_time" dynamodbav:"longest_resp_time"`

	TotalMsgSent       int   `json:"total_msg_sent" dynamodbav:"total_msg_sent"`
	TotalMsgRev        int   `json:"total_msg_rev" dynamodbav:"total_msg_rev"`
	CommunicationHours []int `json:"communication_hours" dynamodbav:"communication_hours"`
	//TotalCommunicationNumber int   `json:"total_communication_number" dynamodbav:"total_communication_number"`

	AllContacts            int `json:"all_contacts" dynamodbav:"all_contacts"`
	NewAddedContacts       int `json:"new_added_contacts" dynamodbav:"new_added_contacts"`
	TotalAssignedContacts  int `json:"total_assigned_contacts" dynamodbav:"total_assigned_contacts"`
	TotalActiveContacts    int `json:"total_active_contacts" dynamodbav:"total_active_contacts"`
	TotalDeliveredContacts int ` json:"total_delivered_contacts" dynamodbav:"total_delivered_contacts"`
	TotalUnhandledContact  int `json:"total_unhandled_contact" dynamodbav:"total_unhandled_contact"`
}

type UserInfo struct {
	UserID       int    `json:"user_id" dynamodbav:"user_id"`
	UserName     string `json:"user_name" dynamodbav:"user_name"`
	TeamID       int    `json:"team_id" dynamodbav:"team_id"`
	UserRoleName string `json:"user_role_name" dynamodbav:"user_role_name"`
	UserStatus   string `json:"user_status" dynamodbav:"user_status"`
	LastLogin    int64  `json:"last_login" dynamodbav:"last_login"`

	AssignedContacts  int `json:"assigned_contacts" dynamodbav:"assigned_contacts"`
	ActiveContacts    int `json:"active_contacts" dynamodbav:"active_contacts"`
	DeliveredContacts int ` json:"delivered_contacts" dynamodbav:"delivered_contacts"`
	UnhandledContact  int `json:"unhandled_contact" dynamodbav:"unhandled_contact"`

	NewAddedContacts int    `json:"new_added_contacts" dynamodbav:"new_added_contacts"`
	AllContacts      int    `json:"all_contacts" dynamodbav:"all_contacts"`
	Tags             []Tags `json:"tags" dynamodbav:"tags"`

	//ChannelData              []ChannelData ` json:"channel_data" dynamodbav:"channel_data" `
	AvgRespTime     int64 `json:"avg_resp_time" dynamodbav:"avg_resp_time"`
	FirstRespTime   int64 `json:"first_resp_time" dynamodbav:"first_resp_time"`
	LongestRespTime int64 `json:"longest_resp_time" dynamodbav:"longest_resp_time"`

	MsgSent int `json:"msg_sent" dynamodbav:"msg_sent"`
	MsgRev  int `json:"msg_rev" dynamodbav:"msg_rev"`
	//CommunicationNumber int `json:"communication_number" dynamodbav:"communication_number"`
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
