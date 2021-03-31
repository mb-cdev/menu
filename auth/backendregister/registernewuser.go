package backendregister

import (
	"menu/auth/model"
	"menu/common/validator"
	"net/http"
)

func registerNewUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	u := model.NewUserFactory(
		r.FormValue("firstname"),
		r.FormValue("lastname"),
		r.FormValue("login"),
		r.FormValue("password"),
		true,
	)
	v, err := validator.IsModelValid(u)

	if err != nil {

	}
}
