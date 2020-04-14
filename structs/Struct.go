package structs

type Users struct {
	Id        string `form:"id" json:"id"`
	FirstName string `form:"firstname" json:"firstname"`
	LastName  string `form:"lastname" json:"lastname"`
	Username  string `form:"username" json:"username"`
}

type Response struct {
	ErrNumber int    `json:"errnumber"`
	Status    string `json:"status"`
	Message   string `json:"message"`
	Data      []Users
}

// type ResponseError struct {
// 	Status  int    `json:"status"`
// 	Message string `json:"message"`
// 	Data    string `json:"data"`
// }
