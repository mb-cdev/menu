package validator

import (
	"bytes"
	"regexp"
)

func IsPasswordValid(value []byte) bool {
	value = bytes.Trim(value, "\n\t ")
	m1, _ := regexp.Match("[a-z]+", value)
	m2, _ := regexp.Match("[A-Z]+", value)
	m3, _ := regexp.Match("[-_!+@#$%^&*()]+", value)
	m4 := len(value) > 6

	if !m1 || !m2 || !m3 || !m4 {
		return false
	}

	return true
}
