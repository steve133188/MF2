package model

type LiveChat struct {
	PK        string           `json:"pk" dynamodbav:"PK"`
	TimeStamp int64            `json:"time_stamp" dynamodbav:"timestamp"`
	Users     map[int]UserInfo `json:"users" dynamodbav:"users"`
}

//AllContacts --> Customer table, number of customers with assignee
//TotalMsgSent --> Message table, number of items with from_me == true
//TotalMsgRev --> Message table, number of items with from_me == false
//AvgRespTime --> message table
//Tags --> customer table, scane number of each tags
type UserInfo struct {
	TemplateMsg         tempMsg  `json:"template_msg" dynamodbav:"template_msg"`
	Channels            chann    `json:"channels" dynamodbav:"channels"`
	AllContacts         int      `json:"all_contacts" dynamodbav:"all_contacts"`
	TotalMsgSent        int      `json:"total_msg_sent" dynamodbav:"total_msg_sent"`
	TotalMsgRev         int      `json:"total_msg_rev" dynamodbav:"total_msg_rev"`
	RespTime            RespTime `json:"resp_time" dynamodbav:"resp_time"`
	CommunicationNumber int      `json:"communication_number" dynamodbav:"communication_number"`
	Tags                []Tags   `json:"tags" dynamodbav:"tags"`
}

type Tags struct {
	Name string `json:"name" dynamodbav:"name"`
	No   int    `json:"no" dynamodbav:"no"`
}

type tempMsg struct {
	Quote int `json:"quote" dynamodbav:"quote"`
	Sent  int `json:"sent" dynamodbav:"sent"`
}

//chatroom table, channel
type chann struct {
	WhatsApp int `json:"whats_app" dynamodbav:"whats_app"`
	WABA     int `json:"waba" dynamodbav:"waba"`
	WeChat   int `json:"we_chat" dynamodbav:"we_chat"`
}

//min as unit
type RespTime struct {
	Longest int `json:"longest" dynamodbav:"longest"`
	Average int `json:"average" dynamodbav:"average"`
	First   int `json:"first" dynamodbav:"first"`
}
