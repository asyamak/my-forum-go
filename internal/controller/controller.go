package controller

import (
	"html/template"
	"log"
	"net/http"
)

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
	Error   error
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
		// Error: errors.New()
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
