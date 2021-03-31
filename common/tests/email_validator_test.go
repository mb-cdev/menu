package tests

import (
	"menu/common/validator"
	"testing"
)

func TestEmailValidator(t *testing.T) {
	validEmail := []byte("mcbury1@o2.pl")
	invalidEmail1 := []byte("mcbury@o2")
	invalidEmail2 := []byte("domain.com")
	invalidEmail3 := []byte("some string")

	if !validator.IsValidEmail(validEmail) {
		t.Error("Valid email is invalid!")
	}

	if validator.IsValidEmail(invalidEmail1) {
		t.Error("Invalid email is valid! 1")
	}

	if validator.IsValidEmail(invalidEmail2) {
		t.Error("Invalid email is valid! 2")
	}

	if validator.IsValidEmail(invalidEmail3) {
		t.Error("Invalid email is valid! 3")
	}
}
