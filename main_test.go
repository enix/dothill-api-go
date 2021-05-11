package dothill

import (
	"fmt"
	"os"
	"testing"

	dothill "github.com/enix/dothill-api-go"

	"github.com/joho/godotenv"
	. "github.com/onsi/gomega"
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

func assert(t *testing.T, cond bool, msg string) {
	if !cond {
		t.Errorf(msg)
	} else {
		fmt.Printf("%s: OK\n", msg)
	}
}

func TestLoginG(t *testing.T) {
	g := NewWithT(t)
	g.Expect(client.Login()).To(BeNil())
}
func TestLoginI(t *testing.T) {
	g := NewWithT(t)
	g.Expect(client.Login("sha256")).To(BeNil())
}

func TestLoginFailed(t *testing.T) {
	var wrongClient = dothill.NewClient()
	wrongClient.Addr = client.Addr
	wrongClient.Username = client.Username
	wrongClient.Password = "wrongpassword"

	g := NewGomegaWithT(t)
	g.Expect(wrongClient.Login()).ToNot(BeNil())
}

func TestReLoginFailed(t *testing.T) {
	var wrongClient = dothill.NewClient()
	wrongClient.Addr = client.Addr
	wrongClient.Username = client.Username
	wrongClient.Password = client.Password

	g := NewGomegaWithT(t)
	g.Expect(wrongClient.Login()).To(BeNil())

	wrongClient.Password = "wrongpassword"
	wrongClient.SessionKey = "outdated-session-key"

	_, status, err := wrongClient.Request("/status/code/1")
	g.Expect(err).NotTo(BeNil())
	g.Expect(status.ResponseType).To(Equal("Error"))
	g.Expect(status.Response).To(Equal("re-login failed"))
}

func TestInvalidURL(t *testing.T) {
	g := NewGomegaWithT(t)
	_, status, err := client.Request("/trololol")

	g.Expect(err).NotTo(BeNil())
	g.Expect(status.ResponseType).To(Equal("Error"))
	g.Expect(status.Response).To(Equal("request failed"))
}

func TestInvalidXML(t *testing.T) {
	g := NewGomegaWithT(t)
	_, status, err := client.Request("/invalid/xml")

	g.Expect(err).NotTo(BeNil())
	g.Expect(status.ResponseType).To(Equal("Error"))
	g.Expect(status.Response).To(Equal("corrupted response"))
}

func TestStatusCodeNotZero(t *testing.T) {
	g := NewGomegaWithT(t)
	_, status, err := client.Request("/status/code/1")

	g.Expect(err).NotTo(BeNil())
	g.Expect(status.ResponseTypeNumeric).To(Equal(1))
}
