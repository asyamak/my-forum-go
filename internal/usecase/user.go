package usecase

import (
	"errors"
	"fmt"
	"forum/internal/entity"
	"log"
	"net/mail"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidNameLength = errors.New("invalid username length")
	ErrUserExist = errors.New("username is already exist")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidUser = errors.New("invalid username")
	ErrInvalidEmail = errors.New("invalid email address")
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
	if user.Password != user.ConfirmPassword {
		log.Println("error - password not the same - createuserandvalidate")
		return fmt.Errorf("usecase: check user :%v",ErrInvalidPassword)
	}
	if len(user.Username) < 4 || len(user.Username) > 40 {
		return fmt.Errorf("usecase: check user :%v",ErrInvalidNameLength)
	}
	for _, w := range user.Username{
		if w < 32 || w > 126{
			return fmt.Errorf("usecase: check user :%v",ErrInvalidUser)
		}
	}
	if _, err := mail.ParseAddress(user.Email); err != nil{
		return fmt.Errorf("usecase: check user :%v",ErrInvalidEmail )
	}
	return nil
}


// isEmailValid checks if the email provided is valid by regex.
func isEmailValid(e string) bool {
    emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    return emailRegex.MatchString(e)
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
