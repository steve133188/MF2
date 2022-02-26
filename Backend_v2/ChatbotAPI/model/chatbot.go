package model

type Flow struct {
	FlowID string
}

type Option struct {
	OptionID string
}

type Action struct {
	ActionID string
}

type ChatListItem struct {
	Stage int `json:"stage"`
	parent string `json:"parent"`
	FlowId string `json:"flow_id"`
}