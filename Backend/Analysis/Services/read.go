package Services

import "mf-log-servies/Model"

type ReadUser interface {
	FindOne( id string) Model.User
	FindMany( condition interface{} ) []Model.User
	FindAll() []Model.User
}

type Read struct {}

func NewRead() *Read{
	return &Read{}
}

func (r*Read) FindOne(id string) Model.User {
	return Model.User{}
}

func (r*Read) FindMany(c interface{}) []Model.User {
	return [] Model.User{}
}

func (r*Read) FindAll() []Model.User {
	return [] Model.User{}
}