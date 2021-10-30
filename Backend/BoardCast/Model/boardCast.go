package Model

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
	Id          string   `json:"id" bson:"id"`
	Name        string   `json:"name" bson:"name"`
	Period      []string `json:"period" bson:"period"`
	Group       string   `json:"group" bson:"group"`
	CreatedBy   string   `json:"created_by" bson:"created_by"`
	CreatedDate string   `json:"created_date" bson:"created_date"`
	Des         string   `json:"description" bson:"description"`
}

// type Message struct {
// 	Content string `json:"content"`
// }

// type Tags struct {
// 	Conditions  interface{} `json:"conditions"`
// 	ConditionOn interface{} `json:"condition_on"`
// }

// type Flow struct {
// 	Conditions  interface{} `json:"conditions"`
// 	ConditionOn interface{} `json:"condition_on"`
// }

// type Conditions struct {
// 	As  interface{} `json:"as"`
// 	Tag interface{} `json:"tag"`
// 	ID  string      `json:"id"`
// }

// type As struct {
// 	Value string `json:"value"`
// 	Name  string `json:"name"`
// }

// type Tag struct {
// 	Value string `json:"value"`
// 	Name  string `json:"name"`
// }

// type ConditionOn struct {
// 	Value string `json:"value"`
// 	Name  string `json:"name"`
// }
