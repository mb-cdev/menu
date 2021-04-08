package login

import (
	"menu/common/middleware"
	"net/http"
)

func init() {
	http.HandleFunc("/admin/login", middleware.Middleware(HandleBackendLogin, &middleware.IsNotLogged{}))
}

func HandleBackendLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		logInUser(w, r, true)
	}
}
