package login

import "net/http"

func init() {
	http.HandleFunc("/admin/login", HandleBackendLogin)
}

func HandleBackendLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		logInUser(w, r, true)
	}
}
