package repositories

import "database/sql"

// Login is the interface that will receive the database connection
type Login struct {
	db *sql.DB
}

func CreateLoginRepository(db *sql.DB) *Login {
	return &Login{db}
}

// FindByEmail find user by
func (login *Login) FindByEmail(email string) (*Users, error) {
	return nil, nil
}
