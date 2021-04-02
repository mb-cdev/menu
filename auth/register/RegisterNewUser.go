package register

import (
	"menu/auth/model"
	"menu/common/database"
	jr "menu/common/response"
	"menu/common/validator"
	"net/http"
)

func registerNewUser(w http.ResponseWriter, r *http.Request, is_backend_user bool) bool {
	r.ParseForm()

	jr := jr.JSONResponse{
		IsSucceed: true,
		Errs:      nil,
	}

	u, err := model.NewUserFactory(
		r.FormValue("firstname"),
		r.FormValue("lastname"),
		r.FormValue("login"),
		r.FormValue("password"),
		r.FormValue("email"),
		is_backend_user,
	)

	var errs []error = nil
	if err == nil {
		_, errs = validator.IsModelValid(u)
	}
	if errs != nil || err != nil {
		jr.IsSucceed = false
		jr.AddError(err)
		jr.AddErrors(errs)
		jr.WriteJSONResponse(w)
		return false
	}

	if res := database.DB.Create(&u); res.Error != nil {
		jr.IsSucceed = false
		jr.AddError(res.Error)
		jr.WriteJSONResponse(w)
		return false
	}
	jr.WriteJSONResponse(w)

	return true
}
