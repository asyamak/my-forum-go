package entity

type Comments struct {
	Id       int
	Author   string
	PostId   int
	Content  string
	Likes    int
	Dislikes int
}
