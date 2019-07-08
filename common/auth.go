package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go_songs/models"

	jwt "github.com/dgrijalva/jwt-go"
)

var mykey = []byte("Encoding account")

//GenerateJWT generate a token
func GenerateJWT(Acc models.Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": Acc.Username,
		"password": Acc.Password,
		"exp":      time.Now().Add(time.Minute * 20).Unix(),
	})
	tokenString, error := token.SignedString(mykey)
	if error != nil {
		fmt.Println(error)
		return "", error
	}
	return tokenString, error
}

//Validate
func Validate(w http.ResponseWriter, r *http.Request, endpoint http.HandlerFunc) {
	if r.URL.Path != "/signin" && r.URL.Path != "/signup" {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Error occurs when parsing token")
				}
				return mykey, nil
			})
			if err != nil {
				json.NewEncoder(w).Encode(err)
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Unauthorized")
		}
	} else {
		endpoint(w, r)
	}
}
