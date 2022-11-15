package controller

import "net/http"

func Controller(mux *http.ServeMux) {
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/auth", Authorisation)
}

func Home(w http.ResponseWriter, r *http.Request)            {}
func Authorisation(w http.ResponseWriter, r *http.Request)   {}
func NewRegistration(w http.ResponseWriter, r *http.Request) {}
func GoogleAuth(w http.ResponseWriter, r *http.Request)      {}
