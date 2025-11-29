// Package controllers have all the route handlers
package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"api/src/database"
	"api/src/model"
	"api/src/repositories"
	"api/src/responses"
)

// GetUser gets the user inside database
func GetUser(w http.ResponseWriter, r *http.Request) {
	if _, error := w.Write([]byte("get")); error != nil {
		log.Fatal(error)
	}
}

// GetUserByID gets the user inside database
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	if _, error := w.Write([]byte("get by id")); error != nil {
		log.Fatal(error)
	}
}

// CreateUser inserts the user inside database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user model.User

	error = json.Unmarshal(requestBody, &user)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	repository := repositories.NewUserRepository(db)
	ID, error := repository.Create(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	user.ID = ID

	responses.JSON(w, http.StatusCreated, user)
}

// PutUser gets the user inside database
func PutUser(w http.ResponseWriter, r *http.Request) {
	if _, error := w.Write([]byte("put")); error != nil {
		log.Fatal(error)
	}
}

// DeleteUser gets the user inside database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if _, error := w.Write([]byte("get")); error != nil {
		log.Fatal(error)
	}
}
