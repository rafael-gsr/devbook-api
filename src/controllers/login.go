package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"api/src/authorization"
	"api/src/database"
	"api/src/model"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
)

// Login checks the user email and password
func Login(w http.ResponseWriter, r *http.Request) {
	reqBody, error := io.ReadAll(r.Body)

	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	var user model.User

	if error := json.Unmarshal(reqBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	repository := repositories.NewUserRepository(db)

	PasswordAndID, error := repository.GetPasswordAndID(user.Email)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = security.VerifyPassword(user.Password, PasswordAndID.Password); error != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	token, _ := authorization.GenerateToken(PasswordAndID.ID)

	responses.JSON(w, http.StatusOK, token)
}
