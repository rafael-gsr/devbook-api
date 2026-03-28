package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"api/src/authorization"
	"api/src/database"
	"api/src/model"
	"api/src/repositories"
	"api/src/responses"

	"github.com/gorilla/mux"
)

// CreatePost creates a new post and saves it on the DB
func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	ID, error := authorization.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	post := model.Post{AuthorID: ID}

	if error = json.Unmarshal(body, &post); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}
	post.CreatedAt = time.Now()

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer closeDB(db)

	repository := repositories.CreateNewPostRepository(db)

	post.ID, error = repository.Create(post)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	userID, error := authorization.ExtractUserID(r)
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

	repository := repositories.CreateNewPostRepository(db)
	posts, error := repository.Find(userID)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID, error := strconv.ParseUint(params["postID"], 10, 64)
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

	repository := repositories.CreateNewPostRepository(db)

	post, error := repository.FindByID(postID)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	responses.JSON(w, http.StatusOK, post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userID, error := authorization.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)
	postID, error := strconv.ParseUint(params["postID"], 10, 64)
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

	repository := repositories.CreateNewPostRepository(db)
	postToEdit, error := repository.FindByID(postID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if postToEdit.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("it is not possible to edit a posts that don't belongs to you"))
		return
	}

	body, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	var post model.Post
	if error = json.Unmarshal(body, &post); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = post.Prepare(); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = repository.Update(postID, post); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID, error := authorization.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)
	postID, error := strconv.ParseUint(params["postID"], 10, 64)
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

	repository := repositories.CreateNewPostRepository(db)

	postToBeDeleted, error := repository.FindByID(postID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if postToBeDeleted.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("it is not possible to delete a post that don't belongs to you"))
		return
	}

	error = repository.Delete(postID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, error := strconv.ParseUint(params["postID"], 10, 64)
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

	repository := repositories.CreateNewPostRepository(db)

	post, error := repository.FindByID(postID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if post.ID == 0 {
		responses.Error(w, http.StatusBadRequest, errors.New("post not found"))
		return
	}

	post.Likes++
	fmt.Println(post.Likes)

	error = repository.Update(postID, post)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
