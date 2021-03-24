package dothill

import (
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/gomega"
)

var client = NewClient()

func init() {
	client.Addr = "http://localhost:8080"
	client.Username = "manage"
	client.Password = "!manage"

	if endpoint := os.Getenv("API_ENDPOINT"); endpoint != "" {
		client.Addr = endpoint
	}
}

func assert(t *testing.T, cond bool, msg string) {
	if !cond {
		t.Errorf(msg)
	} else {
		fmt.Printf("%s: OK\n", msg)
	}
}

func TestLogin(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(client.Login()).To(BeNil())
}

func TestLoginFailed(t *testing.T) {
	var wrongClient = NewClient()
	wrongClient.Addr = client.Addr
	wrongClient.Username = client.Username
	wrongClient.Password = "wrongpassword"

	g := NewGomegaWithT(t)
	g.Expect(wrongClient.Login()).ToNot(BeNil())
}

func TestReLoginFailed(t *testing.T) {
	var wrongClient = NewClient()
	wrongClient.Addr = client.Addr
	wrongClient.Username = client.Username
	wrongClient.Password = client.Password

	g := NewGomegaWithT(t)
	g.Expect(wrongClient.Login()).To(BeNil())

	wrongClient.Password = "wrongpassword"
	wrongClient.sessionKey = "outdated-session-key"

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
