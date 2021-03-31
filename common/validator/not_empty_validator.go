package validator

import (
	"bytes"
)

func IsNotEmpty(value []byte) bool {
	return len(bytes.Trim(value, "\n\t ")) > 0
}
