package Model

import (
	"time"
)

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

//type Password struct {
//	Bcrypt string `json:"bcrypt"`
//}

//type Services struct {
//	Password `json:"password"`
//}

// type Emails struct {
// 	Address  string `json:"address"`
// 	Verified bool   `json:"verified"`
// }

// type Profile struct {
// 	Name         string   `json:"name"`
// 	Phone        string   `json:"phone"`
// 	Modules      []string `json:"modules"`
// 	Channels     []string `json:"channels"`
// 	Team         string   `json:"team"`
// 	Organization string   `json:"organization"`
// 	Enabled      bool     `json:"enabled"`
// }

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Password  string    `json:"password" bson:"password"`
	UserName  string    `json:"username" bson:"username"`
	Email     string    `json:"email" bson:"email"`
	Role      string    `json:"role" bson:"role"`
	Status    string    `json:"status" bson:"status"`
	//Profile   `json:"profile"`
}

// type Token struct {
// 	Password string `json:"password"`
// 	UserName string `json:"username"`
// 	Email    string `json:"email"`
// 	jwt.StandardClaims
// }
