package handlers

import (
	"fmt"
	"reflect"
)

// JSONresponse  ...
type JSONresponse struct {
	Result   map[string]string `json:"data"`
	Metadata map[string]string `json:"metadata"`
}

// JWT ...
type JWT struct{}

// Validator ...
type Validator interface {
	validate(args ...func(i interface{}) error) error
}

type getter interface {
	Get(s string) (reflect.Value, error)
}

// ValidatorGetter ...
type ValidatorGetter interface {
	Validator
	getter
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Get ...
func (c *credentials) Get(key string) (reflect.Value, error) {
	return Get(c, key)
}

func (c *credentials) validate(args ...func(i interface{}) error) error {
	return nil
}

// rangeFuncs ranges over multiple functions
func rangeFuncs(i interface{}, args ...func(interface{}) error) error {
	for _, f := range args {
		err := f(i)
		if err != nil {
			return fmt.Errorf("function call failed: %v", err)
		}
	}

	return nil
}

// Get ...
// uses reflection to get a map value out of an empty interface
func Get(i interface{}, key string) (reflect.Value, error) {
	val := reflect.ValueOf(i)

	// Dereference pointer
	rv := val.Elem()

	// Lookup field by name
	fieldname := rv.FieldByName(key)

	if !fieldname.IsValid() {
		return reflect.Value{}, fmt.Errorf("not a field name: %s", key)
	}

	return fieldname, nil
}
