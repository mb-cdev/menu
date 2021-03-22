package validator

import "regexp"

func IsValidLogin(value []byte) bool {
	r := regexp.MustCompile("^[a-zA-Z0-9]+$")

	return r.Match(value)
}
