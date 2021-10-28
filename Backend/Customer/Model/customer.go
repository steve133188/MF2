package Model

import "time"

// type Customer struct {
// 	Id                 string    `json:"id"`
// 	UserId             string    `json:"user_id"`
// 	Username           string    `json:"username"`
// 	CustomerFirstName  string    `json:"customer_first_name"`
// 	CustomerLastName   string    `json:"customer_last_name"`
// 	Phone              string    `json:"phone"`
// 	Email              string    `json:"email"`
// 	TimeZone           string    `json:"timezone"`
// 	LastUpdatedTime    time.Time `json:"last_updated_time"` //date of updating customer info
// 	AccountCreatedTime time.Time `json:"account_created_time"`
// }

type Customer struct {
	ID           string      `json:"id"`
	Name         string      `json:"name" bson:"name"`
	FirstName    string      `json:"first_name" bson:"first_name"`
	LastName     string      `json:"last_name" bson:"last_name"`
	Phone        string      `json:"phone"`
	Identifier   string      `json:"identifier"`
	Enabled      bool        `json:"enabled"`
	Organization string      `json:"organization"`
	Channel      string      `json:"channel"`
	SourceId     string      `json:"source_id"`
	Source       string      `json:"source"`
	FirstSeen    time.Time   `json:"first_seen" bson:"first_seen"`
	LastSeen     interface{} `json:"last_seen"`
	UpdatedAt    time.Time   `json:"updated_at"`
	TeamAssignee interface{} `json:"team_assignee"`
	Tages        interface{} `json:"tages"`
	AssignedTo   interface{} `json:"assigned_to"`
	MetaFields   interface{} `json:"meta_fields"`
	TeamListing  interface{} `json:"team_listing"`
	IsHandled    interface{} `json:"is_handled"`
}
