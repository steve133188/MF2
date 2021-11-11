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

type Division struct {
	ID        string   `json:"id" bson:"id"`
	Name      string   `json:"name" bson:"name"`
	Team      []string `json:"team" bson:"team"`
	CreatedAt string   `json:"created_at" bson:"created"`
}

type EditTeam struct {
	DivName string `json:"div_name" bson:"div_name"`
	Old     string `json:"old" bson:"old"`
	New     string `json:"new" bson:"new"`
}

type Role struct {
	ID               string `json:"id" bson:"id"`
	Name             string `json:"name" bson:"name"`
	Dashboard        bool   `json:"dashboard" bson:"dashboard"`
	LiveChat         bool   `json:"livechat" bson:"livechat"`
	Contact          bool   `json:"contact" bson:"contact"`
	Boardcast        bool   `json:"boardcast" bson:"boardcast"`
	FlowBuilder      bool   `json:"flowbuilder" bson:"flowbuilder"`
	Integrations     bool   `json:"integration" bson:"integration"`
	ProductCatalogue bool   `json:"product_catalogue" bson:"product_catalogue"`
	Organization     bool   `json:"organization" bson:"organization"`
	Admin            bool   `json:"admin" bson:"admin"`
}

type Tags struct {
	ID      string `json:"id" bson:"id:"`
	Tags    string `json:"tags" bson:"tags"`
	Total   int    `json:"total" bson:"total"`
	Created string `json:"created" bson:"created"`
	Updated string `json:"updated" bson:"updated"`
}

type Group struct {
	Name string `json:"name" bson:"name"`
}

type EditGroup struct {
	Old string `json:"old" bson:"old"`
	New string
}
