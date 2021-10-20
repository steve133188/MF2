package Model


type User struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Phone string `json:"phone"`
}


type UserCredential struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
