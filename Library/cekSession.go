package Library

import (
	"log"
	"net/http"
	"restapi/config"

	"github.com/gorilla/sessions"
	// "github.com/kataras/go-sessions"
)

var store = sessions.NewCookieStore([]byte("SIMPLEAPP"))

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

func CekAuth(w http.ResponseWriter, r *http.Request) bool {
	authorizationHeader := r.Header.Get("Authorization")
	session, _ := store.Get(r, "UserLogin")

	if len(session.Values) == 0 {
		return false
	} else {
		if authorizationHeader == "" {
			return false
		} else {
			tokenJwt := MiddlewareJWTAuthorization(w, r, authorizationHeader)

			if tokenJwt != "<nil>" {
				return CekSession(w, r, tokenJwt)
			} else {
				return false
			}
		}
		// return true
	}
}

func CekLocalSession(r *http.Request) {
	session, _ := store.Get(r, "UserLogin")
	if len(session.Values) == 0 {
		log.Print("Kosong")
	} else {

		log.Print(session.Values)
	}
}
