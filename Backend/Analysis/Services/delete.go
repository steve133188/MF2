package Services

type DeleteLog interface {
	DeleteOneLog(obj interface{}) (status string,err error)
	DeleteManyLog(obj interface{}) (status string,err error)
	DeleteAllLog(obj interface{}) (status string,err error)
}

type Delete struct {

}

func NewUDelete() *Delete{
	return &Delete{}
}

func DeleteOneLog(obj interface{}) (status string,err error){}

func DeleteManyLog(obj interface{}) (status string,err error){}

func DeleteAllLog(obj interface{}) (status string,err error){}
