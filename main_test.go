package dothill

import (
	"fmt"
	"testing"
)

var client = &Client{
	Addr:     "http://mock:8080",
	Username: "manage",
	Password: "!manage",
}

func assert(t *testing.T, cond bool, msg string) {
	if !cond {
		t.Errorf(msg)
	} else {
		fmt.Printf("%s: OK\n", msg)
	}
}

func TestLogin(t *testing.T) {
	assert(t, client.Login() == nil, "login should succeed")
}

func TestInvalidURL(t *testing.T) {
	_, err := client.Request("/trololol")
	assert(t, err != nil, "it should return an error")
}

func TestInvalidXML(t *testing.T) {
	_, err := client.Request("/invalid/xml")
	assert(t, err != nil, "it should return an error")
}

func TestStatusCodeNotZero(t *testing.T) {
	response, err := client.Request("/status/code/1")
	assert(t, err != nil, "it should return an error")
	assert(t, response.Status.ResponseTypeNumeric == 1, "it should return the status code 1 to the user")
}
