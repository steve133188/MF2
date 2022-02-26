package model

type Customer struct {
	CustomerID  int      `json:"customer_id" dynamodbav:"customer_id"`
	Name        string   `json:"customer_name" dynamodbav:"customer_name"`
	Email       string   `json:"email" dynamodbav:"email"`
	FirstName   string   `json:"first_name" dynamodbav:"first_name"`
	LastName    string   `json:"last_name" dynamodbav:"last_name"`
	Phone       int      `json:"phone" dynamodbav:"phone"`
	CountryCode int      `json:"country_code" dynamodbav:"country_code"`
	Channels    []string `json:"channels" dynamodbav:"channels"`
	TeamID      int      `json:"team_id" dynamodbav:"team_id"`
	AgentsID    []int    `json:"agents_id" dynamodbav:"agents_id"`
	TagsID      []int    `json:"tags_id" dynamodbav:"tags_id"`
	Group       string   `json:"customer_group" dynamodbav:"customer_group"`
	Birthday    string   `json:"birthday" dynamodbav:"birthday"`
	Country     string   `json:"country" dynamodbav:"country"`
	Address     string   `json:"address" dynamodbav:"address"`
	Gender      string   `json:"gender" dynamodbav:"gender"`
	CreatedAt   int64    `json:"created_at" dynamodbav:"created_at"`
	UpdateAt    int64    `json:"update_at" dynamodbav:"update_at"`
	ECMID       string   `json:"ECMID" dynamodbav:"ECMID"`
	HandlerId   int      `json:"handler_id" dynamodbav:"handler_id"`
}

type FullCustomer struct {
	CustomerID  int      `json:"customer_id" dynamodbav:"customer_id"`
	Name        string   `json:"customer_name" dynamodbav:"customer_name"`
	Email       string   `json:"email" dynamodbav:"email"`
	FirstName   string   `json:"first_name" dynamodbav:"first_name"`
	LastName    string   `json:"last_name" dynamodbav:"last_name"`
	Phone       int      `json:"phone" dynamodbav:"phone"`
	CountryCode int      `json:"country_code" dynamodbav:"country_code"`
	Channels    []string `json:"channels" `
	Team        Team     `json:"team"`
	Agents      []User   `json:"agents"`
	Tags        []Tag    `json:"tags"`
	Group       string   `json:"customer_group" dynamodbav:"customer_group"`
	Birthday    string   `json:"birthday" dynamodbav:"birthday"`
	Country     string   `json:"country" dynamodbav:"country"`
	Address     string   `json:"address" dynamodbav:"address"`
	Gender      string   `json:"gender" dynamodbav:"gender"`
	CreatedAt   int64    `json:"created_at" dynamodbav:"created_at"`
	UpdateAt    int64    `json:"update_at" dynamodbav:"update_at"`
	ECMID       string   `json:"ECMID" dynamodbav:"ECMID"`
	HandlerId   int      `json:"handler_id" dynamodbav:"handler_id"`
}

type Tag struct {
	TagID    int    `json:"tag_id" dynamodbav:"tag_id"`
	TagName  string `json:"tag_name" dynamodbav:"tag_name"`
	Color    string `json:"color" dynamodbav:"color"`
	CreateAt int64  `json:"create_at" dynamodbav:"create_at"`
	UpdateAt int64  `json:"update_at" dynamodbav:"update_at"`
}

type Team struct {
	TeamID     int    `json:"org_id" dynamodbav:"org_id"`
	Type       string `json:"type" dynamodbav:"type"`
	ChildrenID []int  `json:"children_id" dynamodbav:"children_id"`
	ParentID   int    `json:"parent_id" dynamodbav:"parent_id"`
	Name       string `json:"name" dynamodbav:"name"`
}
