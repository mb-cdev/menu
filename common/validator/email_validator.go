package validator

import "regexp"

func IsValidEmail(value string) bool {
	r := regexp.MustCompile("[a-zA-Z0-9-_]+@[a-zA-Z0-9-_]+.[a-zA-Z.]+")

	return r.MatchString(value)
}
