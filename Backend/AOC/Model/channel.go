package Model

// type Channel struct {
// 	ID                       string `json:"id" bson:"id"`
// 	Name                     string `json:"name" bson:"name"`
// 	Title                    string `json:"titie" bson:"title"`
// 	Enabled                  bool   `json:"enabled" bson:"enabled"`
// 	AuthCode                 string `json:"auth_code" bson:"auth_code"`
// 	ChannelId                string `json:"channel_id" bson:"channel_id"`
// 	Phone                    string `json:"phone" bson:"phone"`
// 	StellaServer             string `json:"server" bson:"server"`
// 	StellaServerAuth         string `json:"server_auth" bson:"server_auth"`
// 	StellaServerAuthUsername string `json:"server_auth_username" bson:"server_auth_username"`
// 	StellaServerAuthPassword string `json:"server_auth_password" bson:"server_auth_password"`
// 	CallbackUrl              string `json:"callback_url" bson:"callback_url"`
// }

type Channel struct {
	ID      string `json:"id" bson:"id"`
	Address string `json:"address" bson:"address"`
}

type ORG struct {
	ID         string   `json:"id" bson:"id"`
	Type       string   `json:"type" bson:"type"`
	ChildrenID []string `json:"children_id" bson:"children_id"`
	ParentID   string   `json:"parent_id" bson:"parent_id"`
	Name       string   `json:"name" bson:"name"`
}

type Tags struct {
	ID      string `json:"id" bson:"id:"`
	Tags    string `json:"tags" bson:"tags"`
	Total   int    `json:"total" bson:"total"`
	Created string `json:"created" bson:"created"`
	Updated string `json:"updated" bson:"updated"`
}

type Roles struct {
	Name string `json:"name" bson:"name"`
	Auth Auth   `json:"auth" bson:"auth"`
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

type StandardReply struct {
	ID			string `json:"id" bson:"id"`
	Name 		string `json:"name" bson:"name"`
	Content		[]string `json:"content" bson:"content"`
	Channel 	[]string `json:"channel" bson:"channel"`
	Team 		string `json:"team" bson:"team"`
	Assignee 	[]string `json:"assignee" bson:"assignee"`
}
