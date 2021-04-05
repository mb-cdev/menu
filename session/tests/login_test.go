package tests

import (
	"encoding/json"
	"fmt"
	"menu/auth/login"
	"menu/common/response"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type HandleLogin struct{}

func (h *HandleLogin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	login.HandleBackendLogin(w, r)

}

func TestLogin(t *testing.T) {
	s := httptest.NewServer(&HandleLogin{})
	defer s.Close()

	v := url.Values{}
	v.Set("login", "FirstLast")
	v.Set("password", "Admin1234@#")
	resp, _ := s.Client().PostForm(s.URL+"/admin/login", v)

	d := json.NewDecoder(resp.Body)

	ds := response.JSONResponse{}
	d.Decode(&ds)
	fmt.Printf("%#v", ds)
}
