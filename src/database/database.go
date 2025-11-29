// Package database connects the application to the database
package database

import (
	"database/sql"

	"api/src/config"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

// Connect connects to the database
func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.MySQLConnection)

	if error != nil {
		return nil, error
	}

	if error = db.Ping(); error != nil {
		return nil, error
	}

	return db, nil
}
