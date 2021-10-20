package Services

type AuthServices interface{
	Register() (res string, err error)
	DeleteAccount() (res string, err error)
	Recovery() (res string, err error)
	Authorized() (jwt []byte , err error)
	Freeze (id string , token []byte) (res string, err error)
}

type authServices struct {}

func NewAuthServices()* authServices  {
	return &authServices{}
}

func (a * authServices) Register() (res string, err error){
		return
}
func (a * authServices) DeleteAccount() (res string, err error){
	return
}
func (a * authServices) Recovery() (res string, err error){
	return
}
func (a * authServices) Authorized() (res string, err error){
	return
}

func (a * authServices) Freeze(id string , token []byte) (res string, err error){
	return
}