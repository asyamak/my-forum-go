package controller

import (
	"fmt"
	"forum/internal/entity"
	"log"
	"net/http"
)

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	// h.u.Login()
}

// func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
// 	token := uuid.NewV4().String()
// 	expiration := time.Now().Add(6)
// 	userModel := entity.UserModel{
// 		Username: "Asya",
// 		Password: "qwer1234",
// 		Token:    token,
// 	}
// 	http.SetCookie(w, &http.Cookie{Value: token, Path: "/", Name: "session_token", Expires: expiration})
// 	// h.u.Login(userModel)
// }

// sign-up handler
func (h *handler) Auth(w http.ResponseWriter, r *http.Request) {
	// user := entity.UserModel{
	// 	Username:        r.PostFormValue("username"),
	// 	Email:           r.PostFormValue("email"),
	// 	Password:        r.PostFormValue("password"),
	// 	ConfirmPassword: r.PostFormValue("confirm_password"),
	// }
	userModel := entity.UserModel{
		Password:        "qwer1234",
		ConfirmPassword: "qwer1234",
		Username:        "asya",
		Email:           "m.a_k@mail.ru",
	}
	err := h.u.CreateUserandValidate(userModel)
	if err != nil {
		fmt.Errorf("error create user: %v", err)
	}
	log.Println("success auth handler")
}
