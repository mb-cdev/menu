package validator

import (
	"reflect"
)

var validators map[string]func(value []byte) bool

func init() {
	validators = map[string]func(value []byte) bool{
		"login": IsValidLogin,
	}
}

func IsModelValid(model interface{}) (bool, []error) {
	var errs []error
	t := reflect.ValueOf(model).Elem()

	var i int = 0
	for ; i < t.NumField(); i++ {
		f := t.Field(i)

	}

	if errs == nil {
		return true, nil
	}

	return false, errs
}
