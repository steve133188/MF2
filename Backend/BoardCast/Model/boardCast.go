package Model

import "time"

// type BoardCast struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// 	Des  string `json:"description"`

// 	UserId       string `json:"user_id"`
// 	Username     string `json:"username"`
// 	CustomerId   string `json:"customer_id"`
// 	CustomerName string `json:"customer_name"`
// 	Message      string `json:"message"`

// 	UpdatedTime time.Time `json:"updated_time"`
// 	CreatedTime time.Time `json:"created_time"`
// }

type BoardCast struct {
	OldId        string      `json:"_id" bson:"old_id"`
	Name         string      `json:"name"`
	Message      interface{} `json:"message"`
	Tag          interface{} `json:"tag"`
	Flow         interface{} `json:"flow"`
	Enabled      bool        `json:"enabled"`
	Organization string      `json:"organization"`
	ScheduledAt  time.Time   `json:"scheduled_at"`
	Channels     interface{} `json:"channels"`
	PlainText    string      `json:"plain_text"`
	Status       string      `json:"status"`
	Delivered    int16       `json:"delivered"`
	Failed       int16       `json:"failed"`
	Recipient    int16       `json:"recipient"`
	TeamId       string      `json:"team_id"`
}

type Message struct {
	Content string `json:"content"`
}

type Tags struct {
	Conditions  interface{} `json:"conditions"`
	ConditionOn interface{} `json:"condition_on"`
}

type Flow struct {
	Conditions  interface{} `json:"conditions"`
	ConditionOn interface{} `json:"condition_on"`
}

type Conditions struct {
	As  interface{} `json:"as"`
	Tag interface{} `json:"tag"`
	ID  string      `json:"id"`
}

type As struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type Tag struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type ConditionOn struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}
