package Model

type BoardCast struct {
	Name        string `json:"name" bson:"name"`
	Period      string `json:"period" bson:"period"`
	Status      string `json:"status" bson:"status"`
	Group       string `json:"group" bson:"group"`
	CreatedBy   string `json:"created_by" bson:"created_by"`
	CreatedDate string `json:"created_date" bson:"created_date"`
	Des         string `json:"description" bson:"description"`
}

type Param struct {
	Param string `json:"param" bson:"param"`
}
