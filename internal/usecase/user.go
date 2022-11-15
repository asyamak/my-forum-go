package usecase

import (
	"log"

	"forum/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

type UserUse struct {
	ur UserRepo
}

func newUserUseCase(ur UserRepo) *UserUse {
	return &UserUse{
		ur: ur,
	}
}

func (u *UserUse) CreateUserandValidate(user entity.UserModel) {
	// uid := uuid.NewV4().String()
	if comparePassword(user.Password, user.ConfirmPassword) {
		log.Println("error - password not the same - createuserandvalidate")
		return
	}
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		log.Printf("error - create user - usecase user : %v", err)
	}
	userModel := entity.UserModel{
		// Id:       uid,
		Password: hashedPassword,
		Username: user.Username,
		// UserFirstname: user.UserFirstname,
		Email: user.Email,
		// SessionToken: ,
	}
	u.ur.CreateUser(userModel)
	u.ur.Login()
	// fmt.Println(user)
	// _, err := u.ur.CreateUser(user)
}

func (u *UserUse) Login(user entity.UserModel) error {
	hashpassword, err := u.ur.ComparePassword(user.Password)
	if err != nil {
		if hashpassword == "" {
			return err
		}
		return err
	}
	if CheckPasswordHash(user.Password, hashpassword) {
		log.Print("password are not the same")
		return err
	}
	return nil
}

func comparePassword(password, compare string) bool {
	if password == compare {
		return true
	} else {
		return false
	}
}

// create new hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// checks login in the process of login
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
