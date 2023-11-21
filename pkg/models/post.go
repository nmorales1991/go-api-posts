package models

import "time"

type Post struct {
	ID          string        `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Status      bool          `json:"status"`
	Date        time.Time     `json:"date" bson:"date"`
	Comments    []PostComment `json:"comments"`
}

type PostComment struct {
	ID      int32     `json:"id"`
	Comment string    `json:"comment"`
	Date    time.Time `json:"date" bson:"date"`
	Status  bool      `json:"status"`
}
