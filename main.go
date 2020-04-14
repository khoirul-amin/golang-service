package main

import (
	"fmt"
	"log"
	"net/http"

	"restapi/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/users", user.ReturnAllUsers).Methods("GET")
	router.HandleFunc("/users", user.InsertUsersMultipart).Methods("POST")
	router.HandleFunc("/users", user.UpdateUsersMultipart).Methods("PUT")
	router.HandleFunc("/users", user.DeleteUsersMultipart).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))

}
