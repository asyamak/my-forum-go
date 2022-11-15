package usecase

import "forum/internal/entity"

type (
	UserRepo interface {
		CreateUser(user entity.UserModel)
		Login()
		DeleteUser(username string) error
		GetUser(username string) (entity.UserModel, error)
		UpdateUser(username, password string) error
		GetAll() ([]entity.UserModel, error)
		ComparePassword(password string) (string, error)
	}
	UserUseCase interface {
		CreateUserandValidate(user entity.UserModel)
		Login(user entity.UserModel) error
	}

	Post interface {
		Create()
		Delete()
		Update()
	}
)
