package Model

import "time"

type ChatHistory struct {
	From         time.Time `json:"from"`
	To           time.Time `json:"to"`
	ChannelId    string    `json:"channelId"`
	MemberId     string    `json:"memberId"`
	Platform     string    `json:"platform"`
	GroupId      string    `json:"groupId"`
	AssignmentId string    `json:"assignmentId"`
}
