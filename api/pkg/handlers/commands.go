package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
)

// Command represents a POST or PUT command
type Command struct {
	Payload ValidatorGetter
	Issuer  interface{}
	writer  io.WriteCloser
}

// NewCommand creates a new command
func NewCommand(arg io.ReadCloser) (*Command, error) {
	c := &Command{
		Payload: &credentials{},
	}

	err := marshalPayload(c, arg)

	if err != nil {
		return nil, fmt.Errorf("could not marshal payload: %v", err)
	}

	return c, nil
}

// should accept a validatorFunc or interface that returns an error
// if needed validate that it has data (JSON)
// validate the command issue
// ensure the command has everything it needs to execute the command
// false -> return error with unmet requirements
// true ->
// check the issued at date and the issuer
// see if JWT hasn't expired
// look at the JWT and see if the hash of the JWT matches that of one current in the session DB
func (c *Command) validate(args ...func(i interface{}) error) error {
	return rangeFuncs(*c, args...)
}

// should accept a prepareFunc or interface
// check for SQL injection stuff
// trim spaces
// escape special characters
// strip slashes
// check white lists
func (c *Command) prepare(args ...func(i interface{}) error) error {
	log.Info("preparing the command")
	return rangeFuncs(c, args...)
}

// should receive an executor func or interface
// if not executor is defined then use the Command's default writer
// open connection to receiving entity
// return success or error JSON
func (c *Command) execute(args ...func(i interface{}) error) error {
	return rangeFuncs(c, args...)
}

// Register registers a new user to the forum
func Register(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func marshalPayload(i interface{}, rc io.ReadCloser) error {
	body, err := ioutil.ReadAll(rc)

	if err != nil {
		return fmt.Errorf("could not read from body %s", err)
	}

	defer rc.Close()

	var prop interface{}

	if v, ok := i.(Command); ok {
		v.Payload = &credentials{}
		prop = v.Payload
	}

	json.Unmarshal(body, prop)

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	return nil
}

func writeJSONresponse(w http.ResponseWriter, h http.Header, t time.Time, err error, m ...map[string]string) {
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	w.Header().Set("Access-Control-Allow-Origin", h.Get("Origin"))

	code := 200
	elapsed := time.Since(t)

	jsonR := &JSONresponse{
		Metadata: map[string]string{
			"responseTime": fmt.Sprintf("%s", (elapsed * 1000)),
			"redirectURL":  "",
		},
	}

	if err != nil {
		jsonR.Result = map[string]string{
			"error": err.Error(),
		}
		code = 400
	} else {
		jsonR.Result = m[0]
	}

	resp, _ := json.Marshal(jsonR)

	w.WriteHeader(code)
	w.Write(resp)
}

// Login logs a user into the forum
func Login(w http.ResponseWriter, req *http.Request) {
	start := time.Now()

	c, err := NewCommand(req.Body)
	if err != nil {
		writeJSONresponse(w, req.Header, start, err)
		return
	}

	// verify that they're registered
	// // check if user ID is in DB (using Query)
	// // true -> write JWT -> store JWT in redis session store -> return JWT to resp writer -> resp writer to JSON
	err = c.validate(hasUsernameAndPassword)

	if err != nil {
		writeJSONresponse(w, req.Header, start, err)
		return
	}

	// // trim func
	// // sanitize func
	// add a session token in the sessionDB (redis)
	// fetch user info from DB
	// create a JWT
	// return a JWT
	// err = c.prepare(pf1, pf2)
	// if err != nil {
	// 	log.Fatalf("failed to prepare command: %s", err.Error())
	// }

	// err = c.execute(ef1, ef2)
	// if err != nil {
	// 	log.Fatalf("failed to execute command: %s", err.Error())
	// }

	result := map[string]string{
		"jwt":     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.ffI2NH89HmR6Swa7aIAsjSv65vftOBOCa_nuxRevF0E",
		"user_id": "1234",
	}

	writeJSONresponse(w, req.Header, start, nil, result)
}

// Logout logs a user out of the forum
func Logout(w http.ResponseWriter, req *http.Request) {
	// create a new command
	// receive user name and password via POST
	// verify that they're registered
	// remove the session token in the sessionDB (redis)
	// fetch user info from DB
	// create a JWT
	// return a JWT
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

// validate has username and password and not empty strings
func hasUsernameAndPassword(i interface{}) error {
	value, ok := i.(Command)

	if !ok {
		return fmt.Errorf("failed to type assert type: %t into type *Command", i)
	}

	user, err := value.Payload.Get("Username")
	if err != nil {
		return fmt.Errorf("could not find 'Username' in Payload: %v", err)
	}

	pass, err := value.Payload.Get("Password")
	if err != nil {
		return fmt.Errorf("could not find 'Password' in Payload: %v", err)
	}

	if reflect.DeepEqual("", pass) {
		return fmt.Errorf("'Password' cannot be empty: %v", err)
	}

	if reflect.DeepEqual("", user) {
		return fmt.Errorf("'User' cannot be empty: %v", err)
	}

	return nil
}
