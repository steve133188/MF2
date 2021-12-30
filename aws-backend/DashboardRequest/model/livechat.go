package model

type LiveChat struct {
	PK        string           `json:"pk,omitempty" dynamodbav:"pk"`
	TimeStamp int64            `json:"time_stamp,omitempty" dynamodbav:"time_stamp"`
	Users     map[int]UserInfo `json:"users,omitempty" dynamodbav:"users"`
}

//AllContacts --> Customer table, number of customers with assignee
//TotalMsgSent --> Message table, number of items with from_me == true
//TotalMsgRev --> Message table, number of items with from_me == false
//AvgRespTime --> message table
//Tags --> customer table, scane number of each tags
type UserInfo struct {
	TemplateMsg         tempMsg        `json:"template_msg" dynamodbav:"template_msg"`
	Channels            chann          `json:"channels" dynamodbav:"channels"`
	AllContacts         int            `json:"all_contacts,omitempty" dynamodbav:"all_contacts"`
	TotalMsgSent        int            `json:"total_msg_sent,omitempty" dynamodbav:"total_msg_sent"`
	TotalMsgRev         int            `json:"total_msg_rev,omitempty" dynamodbav:"total_msg_rev"`
	RespTime            RespTime       `json:"resp_time" dynamodbav:"resp_time"`
	CommunicationNumber int            `json:"communication_number,omitempty" dynamodbav:"communication_number"`
	Tags                map[string]int `json:"tags,omitempty" dynamodbav:"tags"`
}

type tempMsg struct {
	Quote int `json:"quote,omitempty" dynamodbav:"quote"`
	Sent  int `json:"sent,omitempty" dynamodbav:"sent"`
}

//chatroom table, channel
type chann struct {
	WhatsApp int `json:"whats_app,omitempty" dynamodbav:"whats_app"`
	WABA     int `json:"waba,omitempty" dynamodbav:"waba"`
	WeChat   int `json:"we_chat,omitempty" dynamodbav:"we_chat"`
}

//min as unit
type RespTime struct {
	Longest int `json:"longest,omitempty" dynamodbav:"longest"`
	Average int `json:"average,omitempty" dynamodbav:"average"`
	First   int `json:"first,omitempty" dynamodbav:"first"`
}
