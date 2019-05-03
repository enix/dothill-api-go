package dothill_test

import (
	"fmt"
	"testing"

	"enix.io/dothill-api-go"
)

func TestE2E(t *testing.T) {
	client, _ := dothill.NewClient(&dothill.Options{
		Username: "aze",
		Password: "aze",
	})

	fmt.Println(client.ShowDisks())
}
