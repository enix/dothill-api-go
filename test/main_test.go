package dothill_test

import (
	"fmt"
	"testing"

	"enix.io/dothill-api-go"
)

var client = &dothill.Client{
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
	_, _, err := client.Request("/trololol")
	assert(t, err != nil, "it should return an error")
}

func TestInvalidXML(t *testing.T) {
	_, _, err := client.Request("/invalid/xml")
	assert(t, err != nil, "it should return an error")
}

func TestStatusCodeNotZero(t *testing.T) {
	_, status, err := client.Request("/status/code/1")
	assert(t, err != nil, "it should return an error")
	assert(t, status.ReturnCode == 1, "it should return the status code 1 to the user")
}

// func TestValidCall(t *testing.T) {
// 	res, status, err := client.TestCall()
// 	assert(t, err == nil, "it should not return an error")
// 	assert(t, status.ReturnCode == 0, "it should return status code 0")
// 	assert(t, res.Data == "Command completed successfully. (vd-1) - The vdisk was created.", "it should return the correct message")
// }
