package Model

// type Organization struct {
// 	ID       string   `json:"id" bson:"id"`
// 	Name     string   `json:"name" bson:"name"`
// 	Team     []string `json:"team" bson:"team"`
// 	Division []string `json:"division" bson:"division"`
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
