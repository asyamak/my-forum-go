package controller

import (
	"forum/internal/entity"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	// h.u.Login()
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	token := uuid.NewV4().String()
	expiration := time.Now().Add(6)
	userModel := entity.UserModel{
		Username:     "Asya",
		Password:     "qwer1234",
		SessionToken: token,
	}
	http.SetCookie(w, &http.Cookie{Value: token, Path: "/", Name: "session_token", Expires: expiration})
	h.u.Login(userModel)
}

func (h *handler) Auth(w http.ResponseWriter, r *http.Request) {
	// userFirstName := r.FormValue("userFirstName")
	// userName := r.FormValue("userName")
	// password := r.FormValue("password")
	// email := r.FormValue("email")

	userModel := entity.UserModel{
		Password: "qwer1234",
		Username: "Asya",
		// UserFirstname: "M",
		Email: "m.a_k@mail.ru",
		// Cookie:        ,
	}
	h.u.CreateUserandValidate(userModel)
}
