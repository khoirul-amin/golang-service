package structs

type Users struct {
	Id        string `form:"id" json:"id"`
	FirstName string `form:"firstname" json:"firstname"`
	LastName  string `form:"lastname" json:"lastname"`
	Username  string `form:"username" json:"username"`
	Token     string `form:"token" json:"token"`
}

type Response struct {
	ErrNumber int    `json:"errnumber"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	Data      []Users
}

type CekLogin struct {
	Id string `form:"id" json:"id"`
}
