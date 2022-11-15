package repository

import (
	"database/sql"
	"forum/internal/entity"
)

type Post struct {
	db *sql.DB
}

func NewPost(db *sql.DB) *Post {
	return &Post{
		db,
	}
}

func (p *Post) Create() {
}

func (p *Post) Delete() {
}

func (p *Post) Update() {
}

func (p *Post) GetAllPosts() ([]entity.Post, error) {
	posts := []entity.Post{}
	query := `SELECT * FROM posts;`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		post := entity.Post{}
		err = rows.Scan(&post.PostId, &post.PostAuthor, &post.Post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	rows.Close()
	return posts, nil
}
