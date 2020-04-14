package user

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi/Library"
	"restapi/config"
	"restapi/structs"

	_ "github.com/go-sql-driver/mysql"
)

func ReturnAllUsers(w http.ResponseWriter, r *http.Request) {
	var users structs.Users
	var arr_user []structs.Users
	var response structs.Response
	// var errorresponse structs.ResponseError

	ua := r.Header.Get("User-Agent")
	db := config.Connect()
	defer db.Close()

	if ua == "123" {
		rows, err := db.Query("Select id,first_name,last_name,username from users")
		if err != nil {
			log.Print(err)
		}

		for rows.Next() {
			if err := rows.Scan(&users.Id, &users.FirstName, &users.LastName, &users.Username); err != nil {
				log.Fatal(err.Error())

			} else {
				arr_user = append(arr_user, users)
			}
		}

		response.ErrNumber = 0
		response.Status = "SUCCESS"
		response.Message = "Daftar User"
		response.Data = arr_user
	} else {
		response.ErrNumber = 1
		response.Status = "ERROR"
		response.Message = "Header Salah"
		// response.Data = arr_user
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func InsertUsersMultipart(w http.ResponseWriter, r *http.Request) {
	// var users structs.Users
	// var arr_user []structs.Users
	var response structs.Response

	ua := r.Header.Get("User-Agent")
	db := config.Connect()
	defer db.Close()

	if ua == "123" {
		err := r.ParseMultipartForm(4096)
		if err != nil {
			panic(err)
		}

		first_name := r.FormValue("first_name")
		last_name := r.FormValue("last_name")
		username := r.FormValue("username")
		updated_at := ""
		created_at := Library.TimeStamp()

		if first_name != "" && last_name != "" && username != "" && r.FormValue("password") != "" {

			password := Library.Hash(r.FormValue("password"))

			_, err = db.Exec("INSERT INTO users (first_name, last_name, username, password, created_at, updated_at) values (?,?,?,?,?,?)",
				first_name,
				last_name,
				username,
				password,
				created_at,
				updated_at,
			)

			if err != nil {
				log.Print(err)
			}

			response.ErrNumber = 0
			response.Status = "SUCCESS"
			response.Message = "Success Add"
			// log.Print("Insert data to database")
		} else {
			response.ErrNumber = 2
			response.Status = "ERROR"
			response.Message = "Lengkapi Data Terlebih Dahulu"
			// response.Data = arr_user
		}
	} else {
		response.ErrNumber = 1
		response.Status = "ERROR"
		response.Message = "Header Salah"
		// response.Data = arr_user
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateUsersMultipart(w http.ResponseWriter, r *http.Request) {

	var response structs.Response
	ua := r.Header.Get("User-Agent")
	db := config.Connect()
	defer db.Close()

	if ua == "123" {
		err := r.ParseMultipartForm(4096)
		if err != nil {
			panic(err)
		}

		first_name := r.FormValue("first_name")
		last_name := r.FormValue("last_name")
		username := r.FormValue("username")
		id := r.FormValue("user_id")
		updated_at := Library.TimeStamp()

		if first_name != "" && last_name != "" && username != "" && id != "" {
			_, err = db.Exec("UPDATE users set first_name = ?, last_name = ?, username = ?, updated_at = ? where id = ?",
				first_name,
				last_name,
				username,
				updated_at,
				id,
			)

			if err != nil {
				log.Print(err)
			}
			response.ErrNumber = 0
			response.Status = "SUCCESS"
			response.Message = "Success Update Data"
			// log.Print("Update data to database")
		} else {
			response.ErrNumber = 2
			response.Status = "ERROR"
			response.Message = "Lengkapi Data Terlebih Dahulu"
			// response.Data = arr_user
		}
	} else {
		response.ErrNumber = 1
		response.Status = "ERROR"
		response.Message = "Header Salah"
		// response.Data = arr_user
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func DeleteUsersMultipart(w http.ResponseWriter, r *http.Request) {

	var response structs.Response
	ua := r.Header.Get("User-Agent")
	db := config.Connect()
	defer db.Close()

	if ua == "123" {
		err := r.ParseMultipartForm(4096)
		if err != nil {
			panic(err)
		}

		id := r.FormValue("user_id")
		if id != "" {
			_, err = db.Exec("DELETE from users where id = ?",
				id,
			)

			if err != nil {
				log.Print(err)
			}

			response.ErrNumber = 0
			response.Status = "SUCCESS"
			response.Message = "Success Delete Data"
			// log.Print("Delete data to database")

		} else {
			response.ErrNumber = 2
			response.Status = "ERROR"
			response.Message = "Lengkapi Data Terlebih Dahulu"
			// response.Data = arr_user
		}
	} else {
		response.ErrNumber = 1
		response.Status = "ERROR"
		response.Message = "Header Salah"
		// response.Data = arr_user
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
