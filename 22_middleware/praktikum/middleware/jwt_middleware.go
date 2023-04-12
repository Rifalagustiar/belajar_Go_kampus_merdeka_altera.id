package middleware

import (
	"time"

	"crud/constants"

	"github.com/dgrijalva/jwt-go"
)

func Auth(id int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = id
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.SECRET_JWT))
}
