package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"restapi/Library"
	"restapi/config"
	"restapi/structs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

// var APPLICATION_NAME = "SimpleApp"

var store = sessions.NewCookieStore([]byte("SIMPLEAPP"))

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

				userData, err := db.Query("SELECT id, first_name,last_name,username,saldo FROM users WHERE id=?",
					&cekloginRes.Id,
				)

				if err != nil {
					log.Print(err)
				}

				tokenLogin := Library.JwtAuthUser(w, r, username, token, *&cekloginRes.Id)
				session, _ := store.Get(r, "UserLogin")
				session.Values["tokenJWT"] = tokenLogin
				sessions.Save(r, w)
				users.Token = tokenLogin

				userData.Next()
				if err := userData.Scan(&users.Id, &users.FirstName, &users.LastName, &users.Username, &users.Saldo); err != nil {
					log.Fatal(err.Error())
				} else {
					arr_user = append(arr_user, users)
				}

				response.ErrNumber = 0
				response.Status = "SUCCESS"
				response.Message = "Login berhasil"
				response.Data = arr_user
				response.RespTime = Library.TimeStamp()
			} else {
				response.ErrNumber = 3
				response.Status = "ERROR"
				response.Message = "Username Atau Password Salah"
				response.RespTime = Library.TimeStamp()
			}
		} else {
			response.ErrNumber = 2
			response.Status = "ERROR"
			response.Message = "Lengkapi Data Terlebih Dahulu"
			response.RespTime = Library.TimeStamp()
			// response.Data = arr_user
		}
	} else {
		response.ErrNumber = 1
		response.Status = "ERROR"
		response.Message = "Header Salah"
		response.RespTime = Library.TimeStamp()
		// response.Data = arr_user
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GoLogout(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "UserLogin")
	session.Options.MaxAge = -1
	sessions.Save(r, w)
	// authorizationHeader := r.Header.Get("Authorization")
	// if authorizationHeader == "" {
	// 	Library.ErrorResponse(w, "Invalid Token", "Lengkapi data terlebih dahulu", 2)
	// } else {
	// 	result := Library.CekAuth(w, r, authorizationHeader)
	// 	tokenUpdate := ""
	// 	if result {
	// 		tokenJwt := Library.MiddlewareJWTAuthorization(w, r, authorizationHeader)
	// 		db := config.Connect()
	// 		defer db.Close()
	// 		_, _ = db.Exec("UPDATE users set token = ? where token = ?",
	// 			tokenUpdate,
	// 			tokenJwt,
	// 		)
	// 		Library.ErrorResponse(w, "Logout", "Logout Berhasil", 0)
	// 	} else {
	// 		Library.ErrorResponse(w, "Invalid Token", "Token anda sudah tidak bisa digunakan, silahkan login kembali", 4)
	// 	}
	// }
}

func CekUserSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Library.CekLocalSession(r)

	result := Library.CekAuth(w, r)
	if result {
		Library.ErrorResponse(w, "Ok sukses", "Mantab jiwa babang", 0)
	} else {
		Library.ErrorResponse(w, "Invalid Token", "Token anda sudah tidak bisa digunakan, silahkan login kembali", 4)
	}
}
