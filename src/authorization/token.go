// Package authorization implements all the authorization details
package authorization

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"api/src/config"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(ID uint64) (string, error) {
	perms := jwt.MapClaims{}

	perms["authorized"] = true
	perms["exp"] = time.Now().Add(time.Hour * 6).Unix()
	perms["userID"] = ID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perms)

	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken validate if the request token is valid
func ValidateToken(r *http.Request) error {
	token := extractToken(r)
	parsedToken, error := jwt.Parse(token, getSecretSalt)

	if error != nil {
		return error
	}

	if _, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return nil
	}

	return errors.New("invalid token")
}

// extractToken extract the token from headers (bearer)
func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	splittedToken := strings.Split(token, " ")

	if len(splittedToken) == 2 {
		return splittedToken[1]
	}

	return ""
}

func getSecretSalt(token *jwt.Token) (any, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method '%v'", token.Header["alg"])
	}

	return []byte(config.SecretKey), nil
}
