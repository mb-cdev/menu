package tests

import (
	"bytes"
	"crypto/sha512"
	"menu/auth/model"
	"menu/common/database"
	"testing"
)

func TestUserSetPassword(t *testing.T) {
	password := "admin1234"

	/**
	 * Test Password
	 **/
	hashFunc := sha512.New512_256()
	hashFunc.Write([]byte(password))

	passwordToCompare := hashFunc.Sum(nil)
	/** end **/

	u := model.User{}
	u.SetPassword(password)

	if bytes.Compare(passwordToCompare, u.Password) != 0 {
		t.Error("Password do not match!")
	}

}

func TestUserBuilder(t *testing.T) {

	database.DB.Exec("DELETE FROM addresses;DELETE FROM users")

	ub := model.NewUserBuilder().
		Firstname("John").
		Lastname("Smith").
		Login("jSmith").
		Password("admin1234")

	u := ub.Build()
	database.DB.Create(u)
}
