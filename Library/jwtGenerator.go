package Library

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// type M map[string]interface{}

type MyClaims struct {
	jwt.StandardClaims
	Id       string `json:"Id"`
	Username string `json:"Username"`
	Token    string `json:"token"`
	// Id    string `json:"id"`
}

var APPLICATION_NAME = "SimpleApp"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

func JwtAuthUser(w http.ResponseWriter, r *http.Request, username string, token string, id string) string {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Id:       id,
		Username: username,
		Token:    token,
	}

	jwtToken := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	signedToken, err := jwtToken.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusBadRequest)
		// return
		log.Print(w, err.Error(), http.StatusBadRequest)
	}

	return signedToken
}

func MiddlewareJWTAuthorization(w http.ResponseWriter, r *http.Request, authorizationHeader string) string {
	// tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	tokenString := authorizationHeader

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return JWT_SIGNATURE_KEY, nil
	})
	if err != nil {

	}
	claims, ok := token.Claims.(jwt.MapClaims)
	// log.Print(claims)
	if !ok || !token.Valid {
		// log.Print(w, err.Error(), http.StatusBadRequest)
		// fmt.Print("")
	}
	// ctx := context.WithValue(context.Background(), "userInfo", claims)
	// r = r.WithContext(ctx)
	// userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userInfo := claims
	// log.Print(userInfo)
	message := fmt.Sprintf("%v", userInfo["token"])
	return message
	// json.NewEncoder(w).Encode(message)
}
