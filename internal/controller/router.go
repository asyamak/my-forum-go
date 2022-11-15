package controller

import (
	"net/http"
)

func SetupRouter(h *handler) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", h.Home)
	router.HandleFunc("/auth/sign-up", h.Auth)
	// router.HandleFunc("/auth", h.Login)
	return router
}
