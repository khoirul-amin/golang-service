package main

import (
	"fmt"
	"log"
	"net/http"

	"restapi/controller"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/users", controller.ReturnAllUsers).Methods("GET")
	router.HandleFunc("/users", controller.InsertUsersMultipart).Methods("POST")
	router.HandleFunc("/users", controller.UpdateUsersMultipart).Methods("PUT")
	router.HandleFunc("/users", controller.DeleteUsersMultipart).Methods("DELETE")

	//LoginApi
	router.HandleFunc("/users/login", controller.GetLogin).Methods("POST")
	router.HandleFunc("/users/cek-session", controller.CekUserSession).Methods("GET")
	router.HandleFunc("/users/logout", controller.GoLogout).Methods("GET")

	//Get Produk
	router.HandleFunc("/cekproduk", controller.GetProduk).Methods("POST")

	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))

}
