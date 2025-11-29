// Package repositories contains all the database logic
package repositories

import (
	"api/src/model"
	"database/sql"
)

// Users represents the users repository
type Users struct {
	db *sql.DB
}

// NewUserRepository returns the user repository receiving a database connection
func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create inserts the user into the database
func (repository Users) Create(user model.User) (uint64, error) {
	statement, error := repository.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if error != nil {
		return 0, error
	}

	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	lastInsertID, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(lastInsertID), nil
}
