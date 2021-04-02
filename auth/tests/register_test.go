package tests

import (
	"encoding/json"
	"log"
	"menu/auth/register"
	"menu/common/response"
	"net/http"
	"net/http/httptest"
	nt "net/http/httptest"
	"net/url"
	"testing"
)

/**
 * Backend handler
 **/
type HandleBackend struct{}

func (h HandleBackend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	register.HandleBackendRegister(w, r)
}

/**
 * Frontend handler
 */
type HandleFrontend struct{}

func (h HandleFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	register.HandleFrontendRegister(w, r)
}

func TestMain(t *testing.T) {
}

func TestRegisterBackendUser(t *testing.T) {
	s := nt.NewServer(HandleBackend{})
	defer s.Close()

	v := url.Values{}
	v.Set("firstname", "First")
	v.Set("lastname", "Last")
	v.Set("login", "FirstLast")
	v.Set("password", "Admin1234@#")
	v.Set("email", "test@test.pl")

	resp, err := s.Client().PostForm(s.URL+"/admin/register", v)

	if err != nil {
		log.Default().Fatal(err)
	}
	d := json.NewDecoder(resp.Body)

	jrDecoded := response.JSONResponse{}

	d.Decode(&jrDecoded)

	if !jrDecoded.IsSucceed {
		t.Error("Error while creating new BACKEND user: ", jrDecoded.Errs)
	}
}

func TestRegisterFrontendUser(t *testing.T) {
	s := httptest.NewServer(HandleFrontend{})
	defer s.Close()
	v := url.Values{}
	v.Set("firstname", "FirstFront")
	v.Set("lastname", "LastFront")
	v.Set("login", "FirstFrontLastFront")
	v.Set("password", "Front123@#")
	v.Set("email", "test@front.com")

	resp, err := s.Client().PostForm(s.URL+"/register", v)

	if err != nil {
		t.Error(err.Error())
	}

	jrDecoded := response.JSONResponse{}
	j := json.NewDecoder(resp.Body)

	j.Decode(&jrDecoded)

	if jrDecoded.IsSucceed == false {
		t.Error("Error while creating new FRONTEND user: ", jrDecoded.Errs)
	}
}
