// Package repositories contains all the database logic
package repositories

import (
	"database/sql"
	"fmt"

	"api/src/model"
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

func (repository Users) Find(nameOrNick string) ([]model.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%
	lines, error := repository.db.Query("select id, name, nick, email, createdAt from users where name like ? or nick like ?", nameOrNick, nameOrNick)
	if error != nil {
		return nil, error
	}

	defer lines.Close()

	var users []model.User

	for lines.Next() {
		var user model.User

		if error = lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); error != nil {
			return nil, error
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) FindByID(ID uint64) (model.User, error) {
	lines, error := repository.db.Query("select id, name, nick, email, createdAt from users where id = ?", ID)
	if error != nil {
		return model.User{}, error
	}

	defer lines.Close()

	var user model.User

	if lines.Next() {
		if error = lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); error != nil {
			return model.User{}, error
		}
	}

	return user, nil
}

