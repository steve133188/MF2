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

// type Customer struct {
// 	ID           string   `json:"id" bson:"id"`
// 	Name         string   `json:"name" bson:"name"`
// 	FirstName    string   `json:"first_name" bson:"first_name"`
// 	LastName     string   `json:"last_name" bson:"last_name"`
// 	Phone        []string `json:"phone" bson:"phone"`
// 	Organization string   `json:"organization" bson:"organization"`
// 	Channel      []string `json:"channel" bson:"channel"`
// 	ChannelInfo  []string `json:"channel_info" bson:"channel_info"`
// 	Group        string   `json:"group" bson:"group"`
// 	Agent        []string `json:"agent" bson:"agent"`
// 	Tags         []string `json:"tags" bson:"tags"`
// 	Birthday     string   `json:"birthday" bson:"birthday"`
// 	Country      string   `json:"country" bson:"country"`
// 	Address      string   `json:"address" bson:"address"`
// 	CreatedAt    string   `json:"created_at" bson:"created_at"`
// 	UpdatedAt    string   `json:"updated_at" bson:"updated_at"`
// }

type Customer struct {
	ID           string    `json:"id" bson:"id"`
	Name         string    `json:"name" bson:"name"`
	FirstName    string    `json:"first_name" bson:"first_name"`
	LastName     string    `json:"last_name" bson:"last_name"`
	Phone        []string  `json:"phone" bson:"phone"`
	Organization string    `json:"organization" bson:"organization"`
	Channel      []string  `json:"channel" bson:"channel"`
	ChannelInfo  []string  `json:"channel_info" bson:"channel_info"`
	Group        string    `json:"group" bson:"group"`
	Agent        []string  `json:"agent" bson:"agent"`
	Tags         []string  `json:"tags" bson:"tags"`
	Birthday     string    `json:"birthday" bson:"birthday"`
	Country      string    `json:"country" bson:"country"`
	Address      string    `json:"address" bson:"address"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}
type Sort struct {
	Data []string `json:"data"`
}
