package Library

import (
	"crypto/md5"
	"encoding/hex"
)

func Hash(data string) string {
	hash := md5.Sum([]byte(data))

	return hex.EncodeToString(hash[:])

}

func HashToken(data string) string {
	token := data + TimeStamp()
	hash := Hash(token)

	return hash
}