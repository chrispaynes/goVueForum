package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"
)

// JSONresponse represents ...
type JSONresponse struct {
	Result   Result   `json:"result"`
	Metadata Metadata `json:"metadata"`
	Error    string   `json:"error,omitempty"`
}

// Result represents...
type Result struct {
	Data map[string]interface{} `json:"data"`
}

// Metadata represents...
type Metadata struct {
	ResponseTime string `json:"responseTime"`
	RedirectURL  string `json:"redirectURL,omitempty"`
}

// JWT represents...
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

func writeJSONresponse(w http.ResponseWriter, h http.Header, t time.Time, err error, m JSONresponse) {
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Origin", h.Get("Origin"))

	code := 200
	elapsed := time.Since(t)

	jsonR := &JSONresponse{
		Metadata: Metadata{
			ResponseTime: fmt.Sprintf("%s", (elapsed * 1000)),
		},
	}

	if err != nil {
		jsonR.Error = err.Error()
		code = 400
	} else {
		jsonR.Result = m.Result
	}

	resp, _ := json.Marshal(jsonR)

	w.WriteHeader(code)
	w.Write(resp)
}
