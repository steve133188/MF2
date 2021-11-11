package Model

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
	CreatedAt string `json:"created_at" bson:"created_at"`
	Password  string `json:"password" bson:"password"`
	UserName  string `json:"username" bson:"username"`
	Email     string `json:"email" bson:"email"`
	Role      string `json:"role" bson:"role"`
	Status    string `json:"status" bson:"status"`
	// Interface    string      `json:"interface" bson:"interface"`
	// AssignTo     string      `json:"assign_to" bson:"assign_to"`
	Leads        string      `json:"leads" bson:"leads"`
	Team         string      `json:"team" bson:"team"`
	DivisionName string      `json:"division_name" bson:"division_name"`
	LastLogin    string      `json:"last_login" bson:"last_login"`
	Authority    interface{} `json:"authority" bson:"authority"`
	Channels     []string    `json:"channels" bson:"channels"`
	ChannelInfo  Info        `json:"channel_info" bson:"channel_info"`
	Phone        string      `json:"phone" bson:"phone"`
}

type Div struct {
	Division string `json:"division" bson:"division"`
	Team     string `json:"team" bson:"team"`
}

type Param struct {
	Param string `json:"param" bson:"param"`
}

type Info struct {
	Phone     string `json:"phone" bson:"phone"`
	Address   string `json:"address" bson:"address"`
	ChannelId string `json:"channel_id" bson:"channel"`
}

// type Token struct {
// 	Password string `json:"password"`
// 	UserName string `json:"username"`
// 	Email    string `json:"email"`
// 	jwt.StandardClaims
// }
