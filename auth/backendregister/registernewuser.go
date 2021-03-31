package backendregister

import (
	"log"
	"menu/auth/model"
	jr "menu/common/response"
	"menu/common/validator"
	"net/http"
	"time"
)

func registerNewUser(w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()

	jr := jr.JSONResponse{
		IsSucceed: true,
		Errs:      nil,
	}

	u := model.NewUserFactory(
		r.FormValue("firstname"),
		r.FormValue("lastname"),
		r.FormValue("login"),
		r.FormValue("password"),
		r.FormValue("email"),
		true,
	)
	_, errs := validator.IsModelValid(u)

	if errs != nil {
		jr.IsSucceed = false
		jr.AddErrors(errs)
		jr.WriteJSONResponse(w)
		return false
	}

	now := time.Now()
	jr.ResponseData = struct {
		Name  string
		Login string
		Time  int64
	}{
		u.Firstname,
		u.Login,
		now.Unix(),
	}
	log.Default().Println(u)
	jr.WriteJSONResponse(w)

	return true
}