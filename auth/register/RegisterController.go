package register

import (
	"menu/common/middleware"
	"net/http"
)

func init() {
	inl := &middleware.IsNotLogged{}
	http.HandleFunc("/register", middleware.Middleware(HandleFrontendRegister, inl))
	http.HandleFunc("/admin/register", middleware.Middleware(HandleBackendRegister, inl))
}

func HandleBackendRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		registerNewUser(w, r, true)
	}
}

func HandleFrontendRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		registerNewUser(w, r, false)
	}
}
