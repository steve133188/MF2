package model

type Organization struct {
	OrgID      int    `json:"org_id" dynamodbav:"org_id"`
	Type       string `json:"type" dynamodbav:"type"`
	ChildrenID []int  `json:"children_id" dynamodbav:"children_id"`
	ParentID   int    `json:"parent_id" dynamodbav:"parent_id"`
	Name       string `json:"name" dynamodbav:"name"`
}

type OrgStruct struct {
	OrgID    int         `json:"org_id" dynamodbav:"org_id"`
	Type     string      `json:"type" dynamodbav:"type"`
	ParentID int         `json:"parent_id" dynamodbav:"parent_id"`
	Name     string      `json:"name" dynamodbav:"name"`
	Children []OrgStruct `json:"children"`
}

type User struct {
	UserID int `json:"user_id" dynamodbav:"user_id"`
	TeamID int `json:"team_id" dynamodbav:"team_id"`
}
