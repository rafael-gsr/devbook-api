// Package authorization implements all the authorization details
package authorization

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(ID uint64) (string, error) {
	perms := jwt.MapClaims{}

	perms["authorized"] = true
	perms["exp"] = time.Now().Add(time.Hour * 6).Unix()
	perms["userID"] = ID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perms)

	signSalt := os.Getenv("JWT_SALT")
	return token.SignedString([]byte(signSalt))
}
