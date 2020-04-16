package Library

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"restapi/structs"
	"strconv"
	"time"
)

func TimeStamp() string {
	var now = time.Now()

	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	var timeNow string = strconv.Itoa(year) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(day) + " " + strconv.Itoa(hour) + ":" + strconv.Itoa(minute) + ":" + strconv.Itoa(second)

	return timeNow

}

func Hash(data string) string {
	hash := md5.Sum([]byte(data))

	return hex.EncodeToString(hash[:])

}

func HashToken(data string) string {
	token := data + TimeStamp()
	hash := Hash(token)

	return hash
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func ErrorResponse(w http.ResponseWriter, data string, Message string, ErrNo int) {
	var response structs.ErrorResponse

	if ErrNo == 0 {
		response.Status = "SUCCESS"
	} else {
		response.Status = "ERROR"
	}
	response.ErrNumber = ErrNo
	response.Data = data
	response.Message = Message
	response.RespTime = TimeStamp()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
