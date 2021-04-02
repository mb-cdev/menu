package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"menu/auth/login"
	"menu/common/response"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

type BackendHandler struct{}

func (b *BackendHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	login.HandleBackendLogin(w, r)
}

func TestLoginBackendUser(t *testing.T) {
	s := httptest.NewServer(&BackendHandler{})
	defer s.Close()

	v := url.Values{}

	v.Set("login", "FirstLast")
	v.Set("password", "Admin1234@#")

	resp, err := s.Client().PostForm(s.URL+"/admin/login", v)

	if err != nil {
		log.Default().Fatal(err.Error())
	}

	io.Copy(os.Stdout, resp.Body)
	jr := response.JSONResponse{}

	d := json.NewDecoder(resp.Body)
	d.Decode(&jr)

	for i, v := range jr.ResponseData {
		fmt.Printf("%s %s\n", i, v)
	}
}
