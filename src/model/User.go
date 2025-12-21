// Package model contains all database models of the application
package model

import (
	"errors"
	"strings"
	"time"

	"api/src/security"

	"github.com/badoux/checkmail"
)

// User represents the social app user
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

func (user *User) Prepare(step string) error {
	if error := user.validateIsEmpty(step); error != nil {
		return error
	}

	if error := user.format(step); error != nil {
		return error
	}

	return nil
}

// Validate checks the values of the struct and says if is fullfilled
func (user *User) validateIsEmpty(step string) error {
	if user.Name == "" {
		return errors.New("the name field is required and could not be empty")
	}

	if user.Nick == "" {
		return errors.New("the nick field is required and could not be empty")
	}

	if user.Email == "" {
		return errors.New("the email field is required and could not be empty")
	}

	if error := checkmail.ValidateFormat(user.Email); error != nil {
		return errors.New("User email format is not valid")
	}

	if user.Password == "" && step == "create" {
		return errors.New("the password field is required and could not be empty")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "create" {
		hashedPassword, error := security.Hash(user.Password)
		if error != nil {
			return error
		}

		user.Password = string(hashedPassword)
	}

	return nil
}
