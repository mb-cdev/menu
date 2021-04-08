package login

import (
	"menu/auth/model"
	"menu/common/database"
	"menu/common/response"
	"menu/session/session"
	"net/http"
)

const UserObjectSessionKey string = "User"

func logInUser(w http.ResponseWriter, r *http.Request, is_backend_user bool) {
	jsonResponse := response.JSONResponse{
		IsSucceed: true,
		Errs:      nil,
	}

	u := model.User{
		Login: r.FormValue("login"),
	}
	u.SetPassword(r.FormValue("password"))

	var user model.User

	res := database.
		DB.
		Model(&u).
		Where(database.DB.
			Where("login = ?", u.Login).
			Where("password = ?", u.Password).
			Where("is_backend_user = ?", is_backend_user)).First(&user)

	if res.Error != nil {
		jsonResponse.IsSucceed = false
		jsonResponse.AddError(res.Error)
		jsonResponse.WriteJSONResponse(w)
		return
	}

	jsonResponse.ResponseData = make(map[string]string)

	/**
	 * Start new session
	 */
	sess := session.New()
	sess.Set(UserObjectSessionKey, user)
	jsonResponse.ResponseData["session"] = sess.ID
	jsonResponse.WriteJSONResponse(w)
}
