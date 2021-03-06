package tests

import (
	"menu/common/validator"
	"testing"
)

func TestNotEmptyValidator(t *testing.T) {
	emptyString := []byte("")
	emptyString2 := []byte("   ")
	emptyString3 := []byte("\n")
	emptyString4 := []byte("\t")
	emptyString5 := []byte("\n\t")
	fullString := []byte("test\nAnother\n")

	if validator.IsNotEmpty(emptyString2) {
		t.Error("Empty string is not empty! 1")
	}

	if validator.IsNotEmpty(emptyString) {
		t.Error("Empty string is not empty! 2")
	}

	if validator.IsNotEmpty(emptyString3) {
		t.Error("Empty string is not empty! 3")
	}

	if validator.IsNotEmpty(emptyString4) {
		t.Error("Empty string is not empty! 4")
	}

	if validator.IsNotEmpty(emptyString5) {
		t.Error("Empty string is not empty! 5")
	}

	if !validator.IsNotEmpty(fullString) {
		t.Error("full string marked as empty")
	}
}
