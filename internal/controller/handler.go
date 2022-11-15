package controller

import (
	"forum/internal/usecase"
)

type handler struct {
	u usecase.UserUseCase
	p usecase.Post
}

func NewHandler(u usecase.UserUseCase, p usecase.Post) *handler {
	return &handler{
		u: u,
		p: p,
	}
}
