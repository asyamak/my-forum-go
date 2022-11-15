package entity

import "time"

type Post struct {
	PostId       int
	PostAuthor   string
	Post         string
	Title        string
	CreationTime time.Time
	Category     string
	Likes        int
	Dislikes     int
}
