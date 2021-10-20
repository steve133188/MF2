package Model

type Customer struct {
	Id                string `json:"id",omitempty bson:"_id",omitempty`
	UserId            string `json:"userId" bson:"userId"`
	CustomerFirstName string `json:"customerName" bson:"customerName"`
	CustomerLastName  string `json:"customerLastName" bson:"customerLastName"`
	Age               string `json:"age" bson:"age"`
}
