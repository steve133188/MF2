package model

import "time"

type Agent struct {
	PK        string
	TimeStamp int64
	Agents    map[int]AgentInfo
}

//AssignedContacts: customer table assignee == user
//ActiveContacts: no of customer with communication
//DeliveredContacts: message table -> user send to customer
//UnhandledContact: customer -> user, user unread
type AgentInfo struct {
	UserName          string
	UserRoleName      string
	UserStatus        string
	LastLogin         time.Time
	AssignedContacts  int
	ActiveContacts    int
	DeliveredContacts int
	UnhandledContact  int
	TotalMsgSent      int
	AvgRespTime       int
	AvgFirstRespTime  int
}
