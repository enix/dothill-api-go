package dothill_test

import (
	"fmt"
	"testing"

	"enix.io/dothill-api-go"
)

func assert(t *testing.T, cond bool, msg string) {
	if !cond {
		t.Errorf(msg)
	} else {
		fmt.Printf("%s:\tOK\n", msg)
	}
}

func TestE2E(t *testing.T) {
	client, _ := dothill.NewClient(&dothill.Options{
		Username: "aze",
		Password: "aze",
	})

	res, status, err := client.ShowDisks()
	assert(t, res != nil, "it should return something")
	assert(t, err == nil, "it should not return an error")
	assert(t, status.ReturnCode == 0, "it should return 0 status code")
}
