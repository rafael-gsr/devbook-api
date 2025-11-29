// Package config contains the application configs definitions
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// MySQLConnection is the string that connects with the database
	MySQLConnection = ""

	// Port is the application run port
	Port = ""
)

// LoadEnv initialize the env variables
func LoadEnv() {
	var error error

	if error = godotenv.Load(".env"); error != nil {
		log.Fatal(error)
	}

	Port = os.Getenv("API_PORT")

	if len(Port) == 0 {
		Port = ":9000"
	}

	MySQLConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}
