package authentication

import (
	"devbook-api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func TokenGenerator(userID uint64) (string, error) {

	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, permissions)

	return token.SignedString(config.SecretKey)
}
