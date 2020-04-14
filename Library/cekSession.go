package Library

import (
	"net/http"

	"github.com/kataras/go-sessions"
)

func CekSession(w http.ResponseWriter, r *http.Request) bool {
	session := sessions.Start(w, r)
	sessionLogin := session.GetString("userToken")

	if sessionLogin == "" {
		return false
	} else {
		return true
	}
}
