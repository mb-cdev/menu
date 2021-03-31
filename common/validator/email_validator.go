package validator

import "regexp"

func IsValidEmail(value []byte) bool {
	r := regexp.MustCompile("[a-zA-Z0-9-_]+@[a-zA-Z0-9-_]+.[a-zA-Z.]+")

	return r.Match(value)
}
