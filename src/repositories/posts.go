package repositories

import (
	"database/sql"

	"api/src/model"
)

type Posts struct {
	db *sql.DB
}

func CreateNewPostRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// Create a new post inside the DB
func (repository Posts) Create(post model.Post) (uint64, error) {
	statement, error := repository.db.Prepare("insert into posts (title, content, author_id, likes, createdAt) values (?, ?, ?, ?, ?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	created, error := statement.Exec(post.Title, post.Content, post.AuthorID, post.Likes, post.CreatedAt)
	if error != nil {
		return 0, error
	}

	ID, error := created.LastInsertId()
	if error != nil {
		return 0, error
	}

	return uint64(ID), nil
}

