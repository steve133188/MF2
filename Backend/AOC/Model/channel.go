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

// type Organization struct {
// 	ID       string `json:"id" bson:"id"`
// 	Name     string `json:"name" bson:"name"`
// 	Role     string `json:"role" bson:"role"`
// 	Email    string `json:"email" bson:"email"`
// 	Phone    string `json:"phone" bson:"phone"`
// 	Leads    string `json:"leads" bson:"leads"`
// 	TeamId   string `json:"team_id" bson:"team_id"`
// 	Division string `json:"division" bson:"division"`
// }
type Team struct {
	ID        string   `json:"id" bson:"id"`
	Name      string   `json:"name" bson:"name"`
	Division  string   `json:"division" bson:"division"`
	Num       int      `json:"num" bson:"num"`
	UserName  []string `json:"user_name" bson:"user_name"`
	CreatedAt string   `json:"created_at" bson:"created"`
}

type Division struct {
	ID        string   `json:"id" bson:"id"`
	Name      string   `json:"name" bson:"name"`
	Team      []string `json:"team" bson:"team"`
	CreatedAt string   `json:"created_at" bson:"created"`
}

type Admin struct {
	ID          string `json:"id" bson:"id"`
	AdminName   string `json:"username"`
	Description string `json:"description"`
	LastAction  string `json:"last_action"`

	TargetUserId        string `json:"target_user_id"`
	TargetUsername      string `json:"target_username"`
	TargetUserPhone     string `json:"target_user_phone"`
	TargetCustomerId    string `json:"target_customer_id"`
	TargetCustomerName  string `json:"target_custmer_name"`
	TargetCustomerPhone string `json:"target_customer_phone"`

	UpdatedTime string `json:"updated_time"`
	CreatedTime string `json:"created_time"`
}
