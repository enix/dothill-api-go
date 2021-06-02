package dothill

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	. "github.com/onsi/gomega"
)

var client = NewClient()

func init() {
	var exists bool
	settingsfile := ".env"

	// Note, any defined environment variable is used over the ones defined in .env
	if _, err := os.Stat(settingsfile); err == nil {
		fmt.Printf("Testing setup: Loading (%s)\n", settingsfile)
		err := godotenv.Load(settingsfile)
		if err != nil {
			fmt.Printf("Error loading file (%s), error: %v\n", settingsfile, err)
			return
		}
	}

	client.Addr, exists = os.LookupEnv("TEST_STORAGEIP")
	if exists {
		fmt.Printf("Testing setup: %s=%s\n", "TEST_STORAGEIP", client.Addr)
	}

	client.Username, exists = os.LookupEnv("TEST_USERNAME")
	if exists {
		fmt.Printf("Testing setup: %s=%s\n", "TEST_USERNAME", client.Username)
	}

	client.Password, exists = os.LookupEnv("TEST_PASSWORD")
	if exists {
		fmt.Printf("Testing setup: %s=%s\n", "TEST_PASSWORD", client.Password)
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
	g := NewWithT(t)
	g.Expect(client.Login()).To(BeNil())
}

func TestLoginFailed(t *testing.T) {
	var wrongClient = NewClient()
	wrongClient.Addr = client.Addr
	wrongClient.Username = client.Username
	wrongClient.Password = "wrongpassword"

	g := NewWithT(t)
	g.Expect(wrongClient.Login()).ToNot(BeNil())
}

func TestReLoginFailed(t *testing.T) {
	var wrongClient = NewClient()
	wrongClient.Addr = client.Addr
	wrongClient.Username = client.Username
	wrongClient.Password = client.Password

	g := NewWithT(t)
	g.Expect(wrongClient.Login()).To(BeNil())

	wrongClient.Password = "wrongpassword"
	wrongClient.sessionKey = "outdated-session-key"

	_, status, err := wrongClient.Request("/bad/request")
	g.Expect(err).NotTo(BeNil())
	g.Expect(err).To(MatchError("Dothill API returned non-zero code 2 (Authentication Unsuccessful)"))
	g.Expect(status.ResponseType).To(Equal("Error"))
	// This test returns one of three different values based on the  API version.
	g.Expect(status.Response).Should(BeElementOf([]string{"re-login failed", "request failed", "Invalid sessionkey"}))
}

func TestInvalidXML(t *testing.T) {
	g := NewGomegaWithT(t)
	res, err := NewResponse([]byte(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<RESPONSE VERSION="L100">
	<OBJECT basetype="status" name="status" oid="1">
		<PROPERTY name="response-type" type="string" size="12" draw="false" sort="nosort" display-name="Response Type">
			Success
		</PROPERTY>`))

	g.Expect(err).NotTo(BeNil())
	g.Expect(err).To(MatchError(&xml.SyntaxError{
		Msg:  "unexpected EOF",
		Line: 6,
	}))
	g.Expect(res).To(BeNil())
}

func TestBadRequest(t *testing.T) {
	g := NewGomegaWithT(t)
	response, status, err := client.Request("/bad/request")

	g.Expect(response).To(BeNil())
	g.Expect(status.ResponseType).To(Equal("Error"))
	g.Expect(status.Response).To(Equal("request failed"))
	g.Expect(status.Time).To(BeTemporally("~", time.Now(), time.Second))
	g.Expect(err).NotTo(BeNil())
	g.Expect(err).To(MatchError("API returned unexpected HTTP status 400"))
}
