package controller

import (
	"html/template"
	"log"
	"net/http"
)

// func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
// 	// h.u.Login()
// }

// func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
// 	token := uuid.NewV4().String()
// 	expiration := time.Now().Add(6)
// 	userModel := entity.UserModel{
// 		Username:     "Asya",
// 		Password:     "qwer1234",
// 		SessionToken: token,
// 	}
// 	http.SetCookie(w, &http.Cookie{Value: token, Path: "/", Name: "session_token", Expires: expiration})
// 	h.u.Login(userModel)
// }

// func (h *handler) Auth(w http.ResponseWriter, r *http.Request) {
// 	// userFirstName := r.FormValue("userFirstName")
// 	// userName := r.FormValue("userName")
// 	// password := r.FormValue("password")
// 	// email := r.FormValue("email")

// 	userModel := entity.UserModel{
// 		Password: "qwer1234",
// 		Username: "Asya",
// 		// UserFirstname: "M",
// 		Email: "m.a_k@mail.ru",
// 		// Cookie:        ,
// 	}
// 	h.u.CreateUserandValidate(userModel)
// }

func (h *handler) CheckMethod(path, method string, r *http.Request) int {
	if r.Method != method {
		return http.StatusMethodNotAllowed
	}
	if r.URL.Path != path {
		return http.StatusNotFound
	}
	return 200
}

type Error struct {
	Message string
	Status  int
}

func (h *handler) ErrorHandler(w http.ResponseWriter, status int) {
	errHandler := h.setError(status)
	w.WriteHeader(status)
	h.Execute(w, "ui/templates/error.html", errHandler)
}

func (h *handler) setError(status int) *Error {
	return &Error{
		Status:  status,
		Message: http.StatusText(status),
	}
}

func (h *handler) Execute(w http.ResponseWriter, parse string, data interface{}) {
	html, err := template.ParseFiles(parse)
	if err != nil {
		log.Println(err.Error())
		h.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	err = html.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		h.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
}
