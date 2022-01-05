package model

type Livechat struct {
	ChannelContact []ChannelContact

	AllContacts    []int `json:"all_contacts" dynamodbav:"all_contacts"`
	ActiveContacts []int `json:"active_contacts" dynamodbav:"active_contacts"`

	TotalMsgSent []int `json:"total_msg_sent" dynamodbav:"total_msg_sent"`
	TotalMsgRev  []int `json:"total_msg_rev" dynamodbav:"total_msg_rev"`

	AvgTotalRespTime      int64   `json:"avg_resp_time" dynamodbav:"avg_resp_time"`
	AvgTotalFirstRespTime int64   `json:"avg_total_first_resp_time" dynamodbav:"avg_total_first_resp_time"`
	AvgLongestRespTime    int64   `json:"longest_resp_time" dynamodbav:"longest_resp_time"`
	AvgRespTime           []int64 `json:"avg_resp_time" dynamodbav:"avg_resp_time"`

	NewAddedContacts   []int `json:"new_added_contacts" dynamodbav:"new_added_contacts"`
	CommunicationHours []int `json:"communication_hours" dynamodbav:"communication_hours"`

	Yaxis []int64 `json:"yaxis" dynamodbav:"yaxis"`

	Tags []Tags `json:"tags" dynamodbav:"tags"`
}

type ChannelContact struct {
	ChannelName         string
	ChannelTotalContact int
}

type Agents struct {
	Agent []Agent

	AgentsNo         []int `json:"agents_no" dynamodbav:"agents_no"`
	Connected        []int `json:"connected" dynamodbav:"connected"`
	Disconnected     []int `json:"disconnected" dynamodbav:"disconnected"`
	AllContacts      []int `json:"all_contacts" dynamodbav:"all_contacts"`
	NewAddedContacts []int `json:"new_added_contacts" dynamodbav:"new_added_contacts"`

	TotalAssignedContacts  int `json:"total_assigned_contacts" dynamodbav:"total_assigned_contacts"`
	TotalActiveContacts    int `json:"total_active_contacts" dynamodbav:"total_active_contacts"`
	TotalDeliveredContacts int ` json:"total_delivered_contacts" dynamodbav:"total_delivered_contacts"`
	TotalUnhandledContact  int `json:"total_unhandled_contact" dynamodbav:"total_unhandled_contact"`

	TotalMsgSent int `json:"total_msg_sent" dynamodbav:"total_msg_sent"`
	TotalMsgRev  int `json:"total_msg_rev" dynamodbav:"total_msg_rev"`

	AvgTotalRespTime      int64 `json:"avg_resp_time" dynamodbav:"avg_resp_time"`
	AvgTotalFirstRespTime int64 `json:"avg_total_first_resp_time" dynamodbav:"avg_total_first_resp_time"`
	AvgLongestRespTime    int64 `json:"longest_resp_time" dynamodbav:"longest_resp_time"`

	DataCollected int `json:"data_collected" dynamodbav:"data_collected"`
}

type Agent struct {
	UserID       int    `json:"user_id" dynamodbav:"user_id"`
	UserName     string `json:"user_name" dynamodbav:"user_name"`
	UserRoleName string `json:"user_role_name" dynamodbav:"user_role_name"`
	TeamID       int    `json:"team_id" dynamodbav:"team_id"`
	UserStatus   string `json:"user_status" dynamodbav:"user_status"`

	AssignedContacts  int `json:"assigned_contacts" dynamodbav:"assigned_contacts"`
	ActiveContacts    int `json:"active_contacts" dynamodbav:"active_contacts"`
	DeliveredContacts int ` json:"delivered_contacts" dynamodbav:"delivered_contacts"`
	UnhandledContact  int `json:"unhandled_contact" dynamodbav:"unhandled_contact"`

	NewAddedContacts int `json:"new_added_contacts" dynamodbav:"new_added_contacts"`
	AllContacts      int `json:"all_contacts" dynamodbav:"all_contacts"`

	//ChannelData              []ChannelData ` json:"channel_data" dynamodbav:"channel_data" `
	AvgRespTime     int64 `json:"avg_resp_time" dynamodbav:"avg_resp_time"`
	FirstRespTime   int64 `json:"first_resp_time" dynamodbav:"first_resp_time"`
	LongestRespTime int64 `json:"longest_resp_time" dynamodbav:"longest_resp_time"`

	MsgSent int `json:"msg_sent" dynamodbav:"msg_sent"`
	MsgRev  int `json:"msg_rev" dynamodbav:"msg_rev"`
}
