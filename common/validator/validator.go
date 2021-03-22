package validator

import (
	"errors"
	"fmt"
	"reflect"
)

var validators map[string]func(value []byte) bool

func init() {
	validators = map[string]func(value []byte) bool{
		"login": IsValidLogin,
	}
}

func IsModelValid(model interface{}) (bool, []error) {
	var errs []error = nil
	v := reflect.ValueOf(model).Elem()

	var i int = 0
	for ; i < v.NumField(); i++ {
		f := v.Field(i)
		t := v.Type().Field(i)

		tag := t.Tag.Get("validator")

		if len(tag) > 0 {
			if fu, ok := validators[tag]; ok {
				if !fu([]byte(f.String())) {
					errs = append(errs, errors.New(fmt.Sprintf("%v is invalid!", t.Name)))
				}
			}
		}

	}

	if errs == nil {
		return true, nil
	}

	return false, errs
}
