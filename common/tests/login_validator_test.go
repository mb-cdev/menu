package tests

import (
	"menu/common/validator"
	"testing"
)

func TestLoginValidator(t *testing.T) {
	validLogin := []byte("testlogin1990")
	invalidLogin := []byte("$%#Test_")

	if !validator.IsValidLogin(validLogin) {
		t.Error("Valid login is invalid")
	}

	if validator.IsValidLogin(invalidLogin) {
		t.Error("Invalid login is valid")
	}
}
