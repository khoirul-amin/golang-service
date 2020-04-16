package Library

import (
	"log"
	"net/http"
	"restapi/config"
)

func CekSession(w http.ResponseWriter, r *http.Request, tokenJwt string) bool {
	// session := sessions.Start(w, r)
	// sessionLogin := session.GetString("userToken")

	// var cekloginRes structs.CekLogin
	db := config.Connect()
	defer db.Close()

	cekUser, err := db.Query("SELECT id FROM users WHERE token = ?",
		tokenJwt,
	)
	if err != nil {
		log.Print(err)
	}

	if cekUser.Next() {
		return true
	} else {
		return false
	}
}

func CekAuth(w http.ResponseWriter, r *http.Request, authorizationHeader string) bool {
	tokenJwt := MiddlewareJWTAuthorization(w, r, authorizationHeader)

	if tokenJwt != "<nil>" {
		return CekSession(w, r, tokenJwt)
	} else {
		return false
	}
}
