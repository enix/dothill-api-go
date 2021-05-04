package dothill

import (
	"fmt"
	"os"
	"testing"

	dothill "github.com/Seagate/seagate-exos-x-api"
	"github.com/joho/godotenv"
)

var client *dothill.Client = dothill.NewClient()

func init() {
	settingsfile := ".env"
	err := godotenv.Load(settingsfile)
	if err != nil {
		fmt.Printf("Error loading file (%s), error: %v\n", settingsfile, err)
		return
	}

	client.Addr = os.Getenv("STORAGEIP")
	client.Username = os.Getenv("TEST_USERNAME")
	client.Password = os.Getenv("TEST_PASSWORD")
}

func TestLoginG(t *testing.T) {

	fmt.Printf("init: addr=%s, username=%s, password=%s\n", client.Addr, client.Username, client.Password)

	err := client.Login()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("login was successful\n")
		fmt.Printf("SessionKey: %s\n", client.SessionKey)
	}

}

func TestLoginI(t *testing.T) {

	fmt.Printf("init: addr=%s, username=%s, password=%s\n", client.Addr, client.Username, client.Password)

	err := client.Login("sha256")
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	} else {
		fmt.Printf("login was successful\n")
		fmt.Printf("SessionKey: %s\n", client.SessionKey)
	}

}
