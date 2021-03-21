package main

import (
	authbackend "menu/auth/backend"
	"net/http"
)

func main() {
	http.HandleFunc("/register", authbackend.RegisterController)
	http.ListenAndServe(":80", nil)
}
