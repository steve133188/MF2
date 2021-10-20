package Model

import "time"

type User struct {
	ID         string        `json:"id", bson:"_id"`
	Username   string        `json:"username"`
	Email      string        `json:"email"`
	Password   string        `json:"password"`
	Phone      string        `json:"phone"`
	Firstname  string        `json:"firstname"`
	Lastname   string        `json:"lastname"`
	Channels   []interface{} //TODO define the channels datatype
	Teams      []string      `json:"teams"`
	Role       string        `json:"role"`
	Preference interface{}   //TODO define the preference datatype
	Date       time.Time     `json:"date"`
}
