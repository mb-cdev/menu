package tests

import (
	"menu/auth/login"
	"menu/session/session"
	"net/http"
	"testing"
)

type HandleLogin struct{}

func (h *HandleLogin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	login.HandleBackendLogin(w, r)

}

func TestLogin(t *testing.T) {
	/*s := httptest.NewServer(&HandleLogin{})
	defer s.Close()

	v := url.Values{}
	v.Set("login", "FirstLast")
	v.Set("password", "Admin1234@#")
	resp, _ := s.Client().PostForm(s.URL+"/admin/login", v)

	d := json.NewDecoder(resp.Body)

	ds := response.JSONResponse{}
	d.Decode(&ds)
	fmt.Printf("%#v", ds)*/
}

func TestSesions(t *testing.T) {
	ss := []*session.Session{}
	for i := 0; i < 3; i++ {
		ss = append(ss, session.New())
	}

	session.SessionCache.Range(func(k interface{}, v interface{}) bool {
		exists := false
		for _, val := range ss {
			if v.(*session.Session).ID == val.ID {
				exists = true
				break
			}
		}

		if !exists {
			t.Error("Sessions not exists in memory!")
		}

		return exists
	})

	for _, v := range ss {
		v.Drop()
	}
}
