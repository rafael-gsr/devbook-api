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

func (repository Users) GetPasswordAndID(email string) (model.User, error) {
	user := model.User{}
	lines, error := repository.db.Query("select id, password from users where email = ?", email)
	if error != nil {
		return user, nil
	}
	defer lines.Close()

	if lines.Next() {
		error = lines.Scan(&user.ID, &user.Password)
		if error != nil {
			return user, error
		}
	}

	return user, nil
}

func (repository Users) Update(ID uint64, body model.User) error {
	statement, error := repository.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	_, error = statement.Exec(body.Name, body.Nick, body.Email, ID)
	if error != nil {
		return error
	}

	return nil
}

func (repository Users) Delete(ID uint64) error {
	statement, error := repository.db.Prepare("delete from users where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	_, error = statement.Exec(ID)
	if error != nil {
		return error
	}

	return nil
}

func (repository Users) Follow(userID, IDToFollow uint64) error {
	statement, error := repository.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) values (?, ?)")
	if error != nil {
		return error
	}
	defer statement.Close()

	_, error = statement.Exec(userID, IDToFollow)
	if error != nil {
		return error
	}

	return nil
}

func (repository Users) Unfollow(userID, IDToUnfollow uint64) error {
	statement, error := repository.db.Prepare("DELETE FROM followers where user_id=? AND follower_id=?")
	if error != nil {
		return error
	}
	defer statement.Close()

	_, error = statement.Exec(userID, IDToUnfollow)
	if error != nil {
		return error
	}

	return nil
}

func (repository Users) UserFollowers(userID uint64) ([]model.User, error) {
	lines, error := repository.db.Query("SELECT u.id, u.name, u.nick, u.email, u.createdAt FROM users u INNER JOIN followers f ON u.id = f.follower_id WHERE f.user_id = ?", userID)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var followers []model.User

	for lines.Next() {
		var user model.User

		if error = lines.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); error != nil {
			return nil, error
		}

		followers = append(followers, user)
	}
	return followers, nil
}

func (repository Users) UserFollowing(userID uint64) ([]model.User, error) {
	lines, error := repository.db.Query("SELECT u.id, u.name, u.nick, u.email, u.createdAt FROM users u INNER JOIN followers f on u.id = f.user_id WHERE f.follower_id = ? ", userID)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var following []model.User

	for lines.Next() {
		var follower model.User

		if error = lines.Scan(&follower.ID, &follower.Name, &follower.Nick, &follower.Email, &follower.CreatedAt); error != nil {
			return nil, error
		}

		following = append(following, follower)
	}

	return following, nil
}
