package register

import (
	"net/http"
)

func init() {
	http.HandleFunc("/register", HandleFrontendRegister)
	http.HandleFunc("/admin/register", HandleBackendRegister)
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
