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

type ErrorResponse struct {
	ErrNumber int    `json:"errnumber"`
	Status    string `json:"status"`
	Data      string `json:"data"`
	Message   string `json:"message"`
	RespTime  string `json:"respTime"`
}

type Barang struct {
	Id         int    `form:"id" json:"id"`
	NamaProduk string `form:"nama_produk" json:"nama_produk"`
	NamaBarang string `form:"nama_barang" json:"nama_barang"`
	HargaJual  int    `form:"harga_jual" json:"harga_jual"`
	Satuan     string `form:"satuan" json:"satuan"`
	Stok       int    `form:"stock" json:"stock"`
}

type ResponseBarang struct {
	ErrNumber int    `json:"errnumber"`
	Status    string `json:"status"`
	Data      []Barang
	Message   string `json:"message"`
	RespTime  string `json:"respTime"`
}
