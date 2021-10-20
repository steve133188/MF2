package Services

type UpdateUser interface {
	UpdateOneUser(obj interface{}) (status string,err error)
	UpdateManyUser(obj interface{}) (status string,err error)
}

type Update struct {

}

func NewUpdate() *Update{
	return &Update{}
}

func UpdateOneUser(obj interface{}) (status string,err error) {

}

func UpdateManyUser(obj interface{}) (status string,err error) {

}