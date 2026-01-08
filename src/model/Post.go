package model

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt"`
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("the title of the post cannot be empty")
	}

	if post.Content == "" {
		return errors.New("the content of the post cannot be empty")
	}

	return nil
}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)

	post.Content = strings.TrimSpace(post.Content)
}

func (post *Post) Prepare() error {
	if error := post.validate(); error != nil {
		return error
	}

	post.format()

	return nil
}
