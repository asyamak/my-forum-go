package usecase

import (
	"fmt"
	"forum/internal/entity"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserUse struct {
	repository UserRepo
}

func newUserUseCase(ur UserRepo) *UserUse {
	return &UserUse{
		repository: ur,
	}
}

func (u *UserUse) CreateUserandValidate(user entity.UserModel) error {
	if err := checkUser(user); err != nil {
		return err
	}
	var err error
	user.Password, err = generateHashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("Usecase: cannot generate hash: %v", err)
	}
	return u.repository.CreateUser(user)
}

// func (u *UserUse) Login(user entity.UserModel) error {
// 	hashpassword, err := u.ur.ComparePassword(user.Password)
// 	if err != nil {
// 		if hashpassword == "" {
// 			return err
// 		}
// 		return err
// 	}
// 	if CheckPasswordHash(user.Password, hashpassword) {
// 		log.Print("password are not the same")
// 		return err
// 	}
// 	return nil
// }

func checkUser(user entity.UserModel) error {
	if !comparePassword(user.Password, user.ConfirmPassword) {
		log.Println("error - password not the same - createuserandvalidate")
		return fmt.Errorf("Invalid password")
	}
	if len(user.Username) < 4 || len(user.Username) > 40 {
	}

	return nil
}

// if it will not be used more than in one function checkUser put in it
func comparePassword(password, compare string) bool {
	if password == compare {
		return true
	} else {
		return false
	}
}

// create new hashed password upon signup
func generateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checks login in the process of login
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
