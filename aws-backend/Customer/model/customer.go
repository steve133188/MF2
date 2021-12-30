package model

type Customer struct {
	CustomerID int      `json:"customer_id" dynamodbav:"customer_id"`
	Name       string   `json:"customer_name" dynamodbav:"customer_name"`
	Email      string   `json:"email" dynamodbav:"email"`
	FirstName  string   `json:"first_name" dynamodbav:"first_name"`
	LastName   string   `json:"last_name" dynamodbav:"last_name"`
	Phone      string   `json:"phone" dynamodbav:"phone"`
	Channels   []string `json:"channels" dynamodbav:"channels"`
	TeamID     int      `json:"team_id" dynamodbav:"team_id"`
	AgentsID   []int    `json:"agents_id" dynamodbav:"agents_id"`
	TagsID     []int    `json:"tags_id" dynamodbav:"tags_id"`
	Group      string   `json:"customer_group" dynamodbav:"customer_group"`
	Birthday   string   `json:"birthday" dynamodbav:"birthday"`
	Country    string   `json:"country" dynamodbav:"country"`
	Address    string   `json:"address" dynamodbav:"address"`
	Gender     string   `json:"gender" dynamodbav:"gender"`
	CreatedAt  int64    `json:"created_at" dynamodbav:"created_at"`
	UpdateAt   int64    `json:"update_at" dynamodbav:"update_at"`
}

type FullCustomer struct {
	CustomerID int      `json:"customer_id" dynamodbav:"customer_id"`
	Name       string   `json:"customer_name" dynamodbav:"customer_name"`
	Email      string   `json:"email" dynamodbav:"email"`
	FirstName  string   `json:"first_name" dynamodbav:"first_name"`
	LastName   string   `json:"last_name" dynamodbav:"last_name"`
	Phone      string   `json:"phone" dynamodbav:"phone"`
	Channels   []string `json:"channels" dynamodbav:"channels"`
	Team       Team     `json:"team"`
	Agents     []User   `json:"agents"`
	Tags       []Tag    `json:"tags"`
	Group      string   `json:"customer_group" dynamodbav:"customer_group"`
	Birthday   string   `json:"birthday" dynamodbav:"birthday"`
	Country    string   `json:"country" dynamodbav:"country"`
	Address    string   `json:"address" dynamodbav:"address"`
	Gender     string   `json:"gender" dynamodbav:"gender"`
	CreatedAt  string   `json:"created_at" dynamodbav:"created_at"`
	UpdateAt   string   `json:"update_at" dynamodbav:"update_at"`
}

type User struct {
	UserID        int         `json:"user_id" dynamodbav:"user_id"`
	Username      string      `json:"username" dynamodbav:"username"`
	Email         string      `json:"email" dynamodbav:"email"`
	Password      string      `json:"password" dynamodbav:"password"`
	Phone         string      `json:"phone" dynamodbav:"phone"`
	Role          string      `json:"role_name" dynamodbav:"role_name"`
	Leads         int         `json:"leads" dynamodbav:"leads"`
	Status        string      `json:"user_status" dynamodbav:"user_status"`
	TeamID        int         `json:"team_id" dynamodbav:"team_id"`
	Authority     Auth        `json:"authority" dynamodbav:"authority"`
	Channels      interface{} `json:"channels" dynamodbav:"channels"`
	Subscriptions []int       `json:"subscriptions" dynamodbav:"subscriptions"`
	CheckAuth     bool        `json:"check_auth" dynamodbav:"check_auth" default:"false"`
	CreateAt      string      `json:"create_at" dynamodbav:"create_at"`
	LastLogin     string      `json:"last_login" dynamodbav:"last_login"`
}

type Auth struct {
	Dashboard        bool `json:"dashboard" dynamodbav:"dashboard" default:"false"`
	Livechat         bool `json:"livechat" dynamodbav:"livechat" default:"false"`
	Contact          bool `json:"contact" dynamodbav:"contact" default:"false"`
	Broadcast        bool `json:"broadcast" dynamodbav:"broadcast" default:"false"`
	Flowbuilder      bool `json:"flowbuilder" dynamodbav:"flowbuilder" default:"false"`
	Integrations     bool `json:"integrations" dynamodbav:"integrations" default:"false"`
	ProductCatalogue bool `json:"product_catalogue" dynamodbav:"product_catalogue" default:"false"`
	Organization     bool `json:"organization" dynamodbav:"organization" default:"false"`
	Admin            bool `json:"admin" dynamodbav:"admin" default:"false"`
}

type Tag struct {
	TagID    int    `json:"tag_id" dynamodbav:"tag_id"`
	TagName  string `json:"tag_name" dynamodbav:"tag_name"`
	Color    string `json:"color" dynamodbav:"color"`
	CreateAt string `json:"create_at" dynamodbav:"create_at"`
	UpdateAt string `json:"update_at" dynamodbav:"update_at"`
}

type Team struct {
	TeamID     int    `json:"team_id" dynamodbav:"team_id"`
	Type       string `json:"type" dynamodbav:"type"`
	ChildrenID []int  `json:"children_id" dynamodbav:"children_id"`
	ParentID   int    `json:"parent_id" dynamodbav:"parent_id"`
	Name       string `json:"name" dynamodbav:"name"`
}
