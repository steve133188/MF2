package Model

import "time"

// type User struct {
// 	ID         string        `json:"id", bson:"_id"`
// 	Username   string        `json:"username"`
// 	Email      string        `json:"email"`
// 	Password   string        `json:"password"`
// 	Phone      string        `json:"phone"`
// 	Firstname  string        `json:"firstname"`
// 	Lastname   string        `json:"lastname"`
// 	Channels   []interface{} //TODO define the channels datatype
// 	Teams      []string      `json:"teams"`
// 	Role       string        `json:"role"`
// 	Preference interface{}   //TODO define the preference datatype
// 	Date       time.Time     `json:"date"`
// }

type User struct {
	OldId     string      `json:"_id" bson:"old_id"`
	CreatedAt time.Time   `json:"created_at"`
	Services  interface{} `json:"services"`
	UserName  string      `json:"username"`
	Emails    interface{} `json:"emails"`
	Profile   interface{} `json:"profile"`
}

type Services struct {
	Password interface{} `json:"password"`
}

type Password struct {
	Bcrypt string `json:"bcrypt"`
}

type Emails struct {
	Address  string `json:"address"`
	Verified bool   `json:"verified"`
}

type Profile struct {
	Name         string      `json:"name"`
	Phone        string      `json:"phone"`
	Modules      interface{} `json:"modules"`
	Channels     interface{} `json:"channels"`
	Team         string      `json:"team"`
	Organization string      `json:"organization"`
	Enabled      bool        `json:"enabled"`
}
