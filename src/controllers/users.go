// Package controllers have all the route handlers
package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"api/src/authorization"
	"api/src/database"
	"api/src/model"
	"api/src/repositories"
	"api/src/responses"

	"github.com/gorilla/mux"
)

// closeDB closes the database connection and treats it error
func closeDB(db *sql.DB) {
	error := db.Close()
	if error != nil {
		log.Print("Error while closing the DB: ", error)
	}
}

// GetUser gets the user inside database
func GetUser(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer closeDB(db)

	repository := repositories.NewUserRepository(db)

	users, error := repository.Find(nameOrNick)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// GetUserByID gets the user inside database
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, error := strconv.ParseUint(parameters["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer closeDB(db)

	repository := repositories.NewUserRepository(db)
	user, error := repository.FindByID(userID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, user)
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

	if error = user.Prepare("create"); error != nil {
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
	params := mux.Vars(r)

	ID, error := strconv.ParseUint(params["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	tokenID, error := authorization.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if ID != tokenID {
		responses.Error(w, http.StatusForbidden, errors.New("the user do not have the permisssions to perform this actions"))
		return
	}

	newBody, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	var user model.User

	if error = json.Unmarshal(newBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("update"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer closeDB(db)

	repository := repositories.NewUserRepository(db)
	error = repository.Update(ID, user)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser gets the user inside database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, error := strconv.ParseUint(params["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	tokenUserID, error := authorization.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if tokenUserID != ID {
		responses.Error(w, http.StatusForbidden, errors.New("the user cannot perform this action"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer closeDB(db)

	repository := repositories.NewUserRepository(db)

	error = repository.Delete(ID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FollowUser links two users as followers
func FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	IDToFollow, error := strconv.ParseUint(params["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	userID, error := authorization.ExtractUserID(r)
	if error != nil {

		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if IDToFollow == userID {
		responses.Error(w, http.StatusBadRequest, errors.New("you cannot follow yourself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer closeDB(db)

	repository := repositories.NewUserRepository(db)
	if error = repository.Follow(userID, IDToFollow); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUser removes the following relation between two users
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	IDToUnfollow, error := strconv.ParseUint(params["userID"], 10, 64)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	userID, error := authorization.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if userID == IDToUnfollow {
		responses.Error(w, http.StatusForbidden, errors.New("you cannot unfollow yourself"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer closeDB(db)

	repository := repositories.NewUserRepository(db)

	error = repository.Unfollow(userID, IDToUnfollow)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UserFollowers(w http.ResponseWriter, r *http.Request) {
	ID, error := authorization.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	defer closeDB(db)

	repository := repositories.NewUserRepository(db)

	followers, error := repository.UserFollowers(ID)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}
