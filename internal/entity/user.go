package entity

import "time"

type UserModel struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Email           string `json:"email"`
	Post            int    `json:"user_posts"`

	Token          string        `json:"token"`
	ExpirationTime time.Duration `json:"expiration"`
}
