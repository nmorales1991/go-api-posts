package models

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID          string        `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string        `json:"title" validate:"required"`
	Description string        `json:"description" validate:"required"`
	Status      bool          `json:"status"`
	Date        time.Time     `json:"date" bson:"date" validate:"required"`
	Comments    []PostComment `json:"comments" validate:"required,dive"`
}

type PostComment struct {
	ID      string    `json:"id"`
	Comment string    `json:"comment" validate:"required"`
	Date    time.Time `json:"date" bson:"date" validate:"required"`
	Status  bool      `json:"status"`
}

func NewPost() Post {
	return Post{
		Status: true,
		Comments: []PostComment{
			NewPostComment(),
		},
	}
}

func NewPostComment() PostComment {
	return PostComment{
		ID:     uuid.New().String(),
		Status: true,
	}
}
