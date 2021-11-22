package Model

import "time"

type Customer struct {
	ID           string   `json:"id" bson:"id"`
	Name         string   `json:"name" bson:"name"`
	FirstName    string   `json:"first_name" bson:"first_name"`
	LastName     string   `json:"last_name" bson:"last_name"`
	Phones       []string `json:"phones" bson:"phones"`
	Organization string   `json:"organization" bson:"organization"`
	Channels     []string `json:"channels" bson:"channels"`

	Group     string    `json:"group" bson:"group"`
	TeamID    string    `json:"team_id" bson:"team_id"`
	Agents    []string  `json:"agents" bson:"agents"`
	Tags      []string  `json:"tags" bson:"tags"`
	Birthday  string    `json:"birthday" bson:"birthday"`
	Country   string    `json:"country" bson:"country"`
	Address   string    `json:"address" bson:"address"`
	Email     string    `json:"email" bson:"email"`
	Gender    string    `json:"gender" bson:"gender"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
type Sort struct {
	Data []string `json:"data"`
}
