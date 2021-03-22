package tests

import (
	"fmt"
	"menu/common/validator"
	"testing"
)

func TestValidator(t *testing.T) {
	model := struct {
		login string `validator:"login"`
		age   uint8
	}{
		"mcbury_",
		12,
	}

	fmt.Println(validator.IsModelValid(&model))
}
