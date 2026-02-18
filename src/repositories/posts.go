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

// FindByID gets one single post
func (repository Posts) FindByID(ID uint64) (model.Post, error) {
	var post model.Post

	query, error := repository.db.Query("select p.id, p.title, p.content, p.author_id, p.createdAt, u.nick from posts p inner join users u on u.id = p.author_id where p.id = ?", ID)
	if error != nil {
		return post, error
	}
	defer query.Close()

	if query.Next() {
		if error := query.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.CreatedAt, &post.AuthorNick); error != nil {
			return post, error
		}
	}

	return post, nil
}

// Find all the posts
func (repository *Posts) Find(userID uint64) ([]model.Post, error) {
	lines, error := repository.db.Query("select p.title, p.content, p.likes, p.createdAt, u.nick from posts p inner join users u on p.author_id = u.id where u.id = ?", userID)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var posts []model.Post

	for lines.Next() {
		var singlePost model.Post

		error = lines.Scan(&singlePost.Title, &singlePost.Content, &singlePost.Likes, &singlePost.CreatedAt, &singlePost.AuthorNick)
		if error != nil {
			return nil, error
		}

		posts = append(posts, singlePost)
	}

	return posts, nil
}

// Update changes the posts properites
func (repository *Posts) Update(postID uint64, fieldToUpdate model.Post) error {
	statement, error := repository.db.Prepare("update posts set title = ?,  content = ? where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	_, error = statement.Exec(fieldToUpdate.Title, fieldToUpdate.Content, postID)
	if error != nil {
		return error
	}

	return nil
}

// Delete the post
func (repository *Posts) Delete(postID uint64) error {
	statement, error := repository.db.Prepare("delete from posts where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	_, error = statement.Exec(postID)
	if error != nil {
		return error
	}

	return nil
}
