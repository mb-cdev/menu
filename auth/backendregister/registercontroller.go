package backendregister

import "net/http"

func init() {
	http.HandleFunc("/admin/register", handleBackendRegister)
}

func handleBackendRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		registerNewUser(w, r)
	}
}
