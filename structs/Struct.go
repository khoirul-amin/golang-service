package structs

type Users struct {
	Id        string `form:"id" json:"id"`
	FirstName string `form:"firstname" json:"firstname"`
	LastName  string `form:"lastname" json:"lastname"`
	Username  string `form:"username" json:"username"`
	Token     string `form:"token" json:"token"`
}

type Produk struct {
	Id     string `form:"id" json:"id"`
	Nama   string `form:"nama_produk" json:"nama_produk"`
	Status string `form:"status" json:"status"`
}

type ResponseProduk struct {
	ErrNumber int    `json:"errnumber"`
	Status    string `json:"status"`
	Data      []Produk
	Message   string `json:"message"`
	RespTime  string `json:"respTime"`
}

type Response struct {
	ErrNumber int    `json:"errnumber"`
	Status    string `json:"status"`
	Data      []Users
	Message   string `json:"message"`
	RespTime  string `json:"respTime"`
}

type CekLogin struct {
	Id string `form:"id" json:"id"`
}

type Login struct {
	Id        string `form:"id" json:"id"`
	FirstName string `form:"firstname" json:"firstname"`
	LastName  string `form:"lastname" json:"lastname"`
	Username  string `form:"username" json:"username"`
	Token     []byte `form:"token" json:"token"`
}

type ResponseLogin struct {
	ErrNumber int    `json:"errnumber"`
	Status    string `json:"status"`
	Data      []Login
	Message   string `json:"message"`
	RespTime  string `json:"respTime"`
}
