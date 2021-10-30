package Model

type Organization struct {
	ID       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Role     string `json:"role" bson:"role"`
	Email    string `json:"email" bson:"email"`
	Phone    string `json:"phone" bson:"phone"`
	Leads    string `json:"leads" bson:"leads"`
	TeamId   string `json:"team_id" bson:"team_id"`
	Division string `json:"division" bson:"division"`
}
