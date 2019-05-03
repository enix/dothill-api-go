package dothill_test

import (
	"testing"

	"enix.io/dothill-api-go"
)

var client, _ = dothill.NewClient(&dothill.Options{
	Addr:     "http://localhost:8080",
	Username: "aze",
	Password: "aze",
})

func assert(t *testing.T, cond bool, msg string) {
	if !cond {
		t.Errorf(msg)
	}
}

func TestInvalidXML(t *testing.T) {
	_, _, err := client.Request(&dothill.Request{Endpoint: "/invalid/xml"})
	assert(t, err != nil, "it should return an error")
}

func TestStatusCodeNotZero(t *testing.T) {
	_, status, err := client.Request(&dothill.Request{Endpoint: "/status/code/1"})
	assert(t, err != nil, "it should return an error")
	assert(t, status.ReturnCode == 1, "it should return the status code 1 to the user")
}
