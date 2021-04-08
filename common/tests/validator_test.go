package tests

import (
	"fmt"
	"menu/auth/model"
	"menu/common/validator"
	"testing"
)

func TestValidator(t *testing.T) {
	model, _ := model.NewUserFactory("Test", "User", "TestUser", "Admin1234!", "mcbury1@o2.pl", true)

	fmt.Println(validator.IsModelValid(model))
}
