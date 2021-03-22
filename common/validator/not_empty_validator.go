package validator

import "strings"

func IsNotEmpty(value string) bool {
	return len(strings.Trim(value, "\n\t ")) > 0
}
