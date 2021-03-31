package tests

import (
	"bytes"
	"crypto/sha512"
	"menu/auth/model"
	"testing"
)

func TestUserSetPassword(t *testing.T) {
	password := "admin1234"

	/**
	 * Test Password
	 **/
	hashFunc := sha512.New()
	hashFunc.Write([]byte(password))

	passwordToCompare := hashFunc.Sum(nil)
	/** end **/

	u := model.User{}
	u.SetPassword(password)

	if bytes.Compare(passwordToCompare, u.Password) != 0 {
		t.Error("Password do not match!")
	}

}
