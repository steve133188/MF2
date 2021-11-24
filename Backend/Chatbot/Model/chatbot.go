package Model

type ChatBotReply struct {
	ID string `json:"id" bson:"id"`

	ChildrenID []string `json:"children_id" bson:"children_id"`
	ParentID   string   `json:"parent_id" bson:"parent_id"`

	DataType   string `json:"data_type" bson:"data_type"`
	InputData  string `json:"input_data" bson:"input_data"`
	OutputData string `json:"output_data" bson:"output_data"`
}
