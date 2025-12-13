// Package security implements all the security features
package security

import "golang.org/x/crypto/bcrypt"

// Hash receives an value and returns a hashed byte array
func Hash(value string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
