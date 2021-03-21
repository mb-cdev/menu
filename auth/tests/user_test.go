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

	ub.AddAddress().
		Name("Test Address").
		Line1("Os Zgoda 60").
		Postcode("34-300").
		City("Zywiec").
		CountryID(1)

	ub.AddAddress().
		Name("Second Address").
		Line1("Line 1").
		Postcode("12-200").
		City("Polkowice").
		CountryID(1)

	u := ub.Build()
	database.DB.Create(u)
}
