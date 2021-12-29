package model

type LiveChat struct {
	PK        string
	TimeStamp int64
	Users     map[int]UserInfo
}

//AllContacts --> Customer table, number of customers with assignee
//TotalMsgSent --> Message table, number of items with from_me == true
//TotalMsgRev --> Message table, number of items with from_me == false
//AvgRespTime --> message table
//Tags --> customer table, scane number of each tags
type UserInfo struct {
	TemplateMsg         tempMsg
	Channels            chann
	AllContacts         int
	TotalMsgSent        int
	TotalMsgRev         int
	RespTime            RespTime
	CommunicationNumber int
	Tags                map[string]int
}

type tempMsg struct {
	Quote int
	Sent  int
}

//chatroom table, channel
type chann struct {
	WhatsApp int
	WABA     int
	WeChat   int
}

//min as unit
type RespTime struct {
	Longest int
	Average int
	First   int
}
