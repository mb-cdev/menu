package middleware

import (
	"menu/auth/login"
	"menu/auth/model"
	"menu/session/session"
	"net/http"
)

type MiddlewareFunc interface {
	Execute(http.ResponseWriter, *http.Request) bool
}

func Middleware(next func(w http.ResponseWriter, r *http.Request), funcs ...MiddlewareFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, f := range funcs {
			if !f.Execute(w, r) {
				return
			}
		}

		next(w, r)
	}
}

type IsLogged struct {
	MustBeAdmin bool
}

func (il *IsLogged) Execute(w http.ResponseWriter, r *http.Request) bool {
	sessID := r.Header.Get("X-MENU-SESSION-ID")
	s, err := session.Open(sessID)

	if err != nil {
		http.Error(w, "Not Authorized", http.StatusForbidden)
		return false
	}

	if v, ok := s.Get(login.UserObjectSessionKey); ok {
		if u, ok := v.(model.User); ok {
			if !il.MustBeAdmin {
				return true
			} else {
				if u.IsBackendUser {
					return true
				}
			}

		}
	}

	http.Error(w, "Not Authorized", http.StatusForbidden)
	return false
}
