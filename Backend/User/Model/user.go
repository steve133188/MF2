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
	Leads       string   `json:"leads" bson:"leads"`
	TeamID      string   `json:"team_id" bson:"team_id"`
	LastLogin   string   `json:"last_login" bson:"last_login"`
	Authority   Auth     `json:"authority" bson:"authority"`
	Channels    []string `json:"channels" bson:"channels"`
	ChannelInfo Info     `json:"channel_info" bson:"channel_info"`
	Phone       string   `json:"phone" bson:"phone"`
}

type Roles struct {
	Phone string `json:"phone" bson:"phone"`
	Role  string `json:"role" bson:"role"`
	Auth  Auth   `json:"auth" bson:"auth"`
}

type Auth struct {
	Dashboard        bool `json:"dashboard" bson:"dashboard" default:"false"`
	Livechat         bool `json:"livechat" bson:"livechat" default:"false"`
	Contact          bool `json:"contact" bson:"contact" default:"false"`
	Boardcast        bool `json:"boardcast" bson:"boardcast" default:"false"`
	Flowbuilder      bool `json:"flowbuilder" bson:"flowbuilder" default:"false"`
	Integrations     bool `json:"integrations" bson:"integrations" default:"false"`
	ProductCatalogue bool `json:"product_catalogue" bson:"product_catalogue" default:"false"`
	Organization     bool `json:"organization" bson:"organization" default:"false"`
	Admin            bool `json:"admin" bson:"admin" default:"false"`
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
