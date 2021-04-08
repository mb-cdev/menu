package middleware

import (
	"menu/auth/constans"
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

/**
 * IsLogged
 **/
type IsLogged struct {
	MustBeAdmin bool
}

func (il *IsLogged) Execute(w http.ResponseWriter, r *http.Request) bool {
	sessID := r.Header.Get(session.SessionHeaderKey)

	if sessID == "" {
		http.Error(w, "Not Authorized", http.StatusForbidden)
		return false
	}

	s, err := session.Open(sessID)

	if err != nil {
		http.Error(w, "Not Authorized", http.StatusForbidden)
		return false
	}

	if v, ok := s.Get(constans.UserObjectSessionKey); ok {
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

/**
 * IsNotLogged
 */
type IsNotLogged struct{}

func (inl *IsNotLogged) Execute(w http.ResponseWriter, r *http.Request) bool {
	sessID := r.Header.Get(session.SessionHeaderKey)

	if sessID == "" {
		return true
	}

	s, err := session.Open(sessID)

	if err != nil {
		return true
	}

	if _, ok := s.Get(constans.UserObjectSessionKey); !ok {
		return true
	}

	http.Error(w, "Not Allowed", http.StatusForbidden)
	return false
}
