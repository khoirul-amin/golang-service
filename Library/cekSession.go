package Library

import (
	"net/http"

	"github.com/kataras/go-sessions"
)

func CekSession(w http.ResponseWriter, r *http.Request, tokenJwt string) bool {
	session := sessions.Start(w, r)
	sessionLogin := session.GetString("userToken")

	if sessionLogin == tokenJwt {
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
