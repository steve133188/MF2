package Model

type Channel struct {
	ID                       string `json:"id" bson:"id"`
	Name                     string `json:"name" bson:"name"`
	Title                    string `json:"titie" bson:"title"`
	Enabled                  bool   `json:"enabled" bson:"enabled"`
	AuthCode                 string `json:"auth_code" bson:"auth_code"`
	ChannelId                string `json:"channel_id" bson:"channel_id"`
	Phone                    string `json:"phone" bson:"phone"`
	StellaServer             string `json:"server" bson:"server"`
	StellaServerAuth         string `json:"server_auth" bson:"server_auth"`
	StellaServerAuthUsername string `json:"server_auth_username" bson:"server_auth_username"`
	StellaServerAuthPassword string `json:"server_auth_password" bson:"server_auth_password"`
	CallbackUrl              string `json:"callback_url" bson:"callback_url"`
}
