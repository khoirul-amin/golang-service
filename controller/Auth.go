package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi/Library"
	"restapi/config"
	"restapi/structs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/go-sessions"
)

func GetLogin(w http.ResponseWriter, r *http.Request) {
	var users structs.Users
	var cekloginRes structs.CekLogin
	var arr_user []structs.Users
	var response structs.Response

	ua := r.Header.Get("User-Agent")
	db := config.Connect()
	defer db.Close()

	if ua == "123" {
		username := r.FormValue("username")
		if username != "" && r.FormValue("password") != "" {
			password := Library.Hash(r.FormValue("password"))

			cekUser, err := db.Query("SELECT id FROM users WHERE username = ? AND password = ?",
				username, password,
			)
			if err != nil {
				log.Print(err)
			}

			if cekUser.Next() {
				cekUser.Scan(&cekloginRes.Id)

				token := Library.HashToken(password)
				updated_at := Library.TimeStamp()
				_, err = db.Exec("UPDATE users SET token = ?, updated_at = ? WHERE id = ?",
					token,
					updated_at,
					&cekloginRes.Id,
				)
				if err != nil {
					log.Print(err)
				}

				userData, err := db.Query("SELECT id, first_name,last_name,username,token FROM users WHERE id=?",
					&cekloginRes.Id,
				)

				if err != nil {
					log.Print(err)
				}

				userData.Next()
				if err := userData.Scan(&users.Id, &users.FirstName, &users.LastName, &users.Username, &users.Token); err != nil {
					log.Fatal(err.Error())
				} else {
					arr_user = append(arr_user, users)
				}

				session := sessions.Start(w, r)

				session.Set("isLogin", "Login")
				session.Set("userName", username)
				session.Set("userToken", token)
				response.ErrNumber = 0
				response.Status = "SUCCESS"
				response.Message = "Success Login"
				response.Data = arr_user
			} else {
				response.ErrNumber = 3
				response.Status = "ERROR"
				response.Message = "Username Atau Password Salah"
			}
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

func GoLogout(w http.ResponseWriter, r *http.Request) {
	var response structs.Response
	ua := r.Header.Get("User-Agent")
	if ua == "123" {
		session := sessions.Start(w, r)
		session.Clear()
		sessions.Destroy(w, r)

		response.ErrNumber = 0
		response.Status = "SUCCESS"
		response.Message = "Logout Success"
		// log.Print("Delete data to database")
	} else {
		response.ErrNumber = 1
		response.Status = "ERROR"
		response.Message = "Header Salah"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CekUserSession(w http.ResponseWriter, r *http.Request) {
	session := Library.CekSession(w, r)

	log.Print(session)
}
