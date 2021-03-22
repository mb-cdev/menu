package tests

import (
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

	validator.IsModelValid(&model)
}
