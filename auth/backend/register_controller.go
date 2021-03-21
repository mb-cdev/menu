package backend

import (
	"fmt"
	"io"
	"menu/auth/model"
	"menu/common/database"
	"net/http"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		registerUser(w, r)
	}
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	ub := model.NewUserBuilder()

	u := ub.Firstname(r.PostFormValue("firstname")).
		Lastname(r.PostFormValue("lastname")).
		Login(r.PostFormValue("login")).
		Password(r.PostFormValue("password")).
		Build()

	database.DB.Create(u)

	io.WriteString(w, fmt.Sprintf("Added new user ID: %d", u.Model.ID))
}
