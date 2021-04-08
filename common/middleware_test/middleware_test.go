package middleware_test

import (
	"encoding/json"
	"io"
	"menu/auth/login"
	"menu/common/middleware"
	"menu/common/response"
	"menu/session/session"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	middleware.Middleware(login.HandleBackendLogin, &middleware.IsNotLogged{})(w, r)
}

type HandlerIsLogged struct {
}

func (h *HandlerIsLogged) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	middleware.Middleware(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<h1>Only for loggedin!</h1>")
	}, &middleware.IsLogged{})(w, r)
}

func TestIsLoggedMiddleware(t *testing.T) {
	s := httptest.NewServer(&HandlerIsLogged{})
	defer s.Close()

	resp, err := s.Client().Get(s.URL + "/test")

	if err != nil {
		t.Error(err.Error())
	}

	if resp.StatusCode != 403 {
		t.Error("Status code other than 403")
	}
}

func TestIsNotLoggedMiddleware(t *testing.T) {
	s := httptest.NewServer(&Handler{})
	defer s.Close()

	// Log in to service
	v := url.Values{}
	v.Set("login", "FirstLast")
	v.Set("password", "Admin1234@#")

	resp, err := s.Client().PostForm(s.URL+"/admin/login", v)

	if err != nil {
		t.Error(err.Error())
	}
	var jsonResp response.JSONResponse
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&jsonResp)

	if !jsonResp.IsSucceed {
		t.Error("Response is false")
	}
	//open session
	sessID := jsonResp.ResponseData["session"]
	sess, err := session.Open(sessID)
	if err != nil {
		t.Error(err.Error())
	}
	defer sess.Drop()

	//try to log in again
	vEncoded := v.Encode()

	bufferB := strings.NewReader(vEncoded)
	req, errR := http.NewRequest(http.MethodPost, s.URL+"/admin/login", bufferB)
	req.Header.Add(session.SessionHeaderKey, sess.ID)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.FormatInt(int64(len(vEncoded)), 10))
	if errR != nil {
		t.Error(errR.Error())
	}

	c := s.Client()
	res, err2 := c.Do(req)

	if err2 != nil {
		t.Error(err2.Error())
	}

	if res.StatusCode != 403 {
		t.Error("Status code different than 403")
	}
}
