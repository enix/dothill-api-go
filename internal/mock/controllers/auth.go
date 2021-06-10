package controllers

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomHexa(n int) string {
	var hexa = []rune("abcdef0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = hexa[rand.Intn(len(hexa))]
	}
	return string(b)
}

type AuthController struct {
	Tokens   map[string]struct{}
	username string
	password string
}

func NewAuthController() *AuthController {
	username := os.Getenv("TEST_USERNAME")
	password := os.Getenv("TEST_PASSWORD")

	if username == "" {
		panic("missing TEST_USERNAME environment variable")
	}
	if password == "" {
		panic("missing TEST_PASSWORD environment variable")
	}

	fmt.Printf("Starting Auth Controller with username=%q and password=%q\n", username, password)

	return &AuthController{
		Tokens:   map[string]struct{}{},
		username: username,
		password: password,
	}
}

func (a AuthController) Login(c *gin.Context) {
	userpass := fmt.Sprintf("%s_%s", a.username, a.password)
	hash := md5.Sum([]byte(userpass))

	if c.Param("hash") == fmt.Sprintf("%x", hash) {
		token := randomHexa(32)
		a.Tokens[token] = struct{}{}
		c.Data(http.StatusOK, "", []byte(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<RESPONSE VERSION="L100">
	<OBJECT basetype="status" name="status" oid="1">
		<PROPERTY name="response-type" type="string" size="12" draw="false" sort="nosort" display-name="Response Type">Success</PROPERTY>
		<PROPERTY name="response-type-numeric" type="uint32" size="12" draw="false" sort="nosort" display-name="Response Type">0</PROPERTY>
		<PROPERTY name="response" type="string" size="180" draw="true" sort="nosort" display-name="Response">`+token+`</PROPERTY>
		<PROPERTY name="return-code" type="sint32" size="15" draw="false" sort="nosort" display-name="Return Code">1</PROPERTY>
		<PROPERTY name="component-id" type="string" size="80" draw="false" sort="nosort" display-name="Component ID"></PROPERTY>
		<PROPERTY name="time-stamp" type="string" size="25" draw="false" sort="datetime" display-name="Time">2021-06-02 09:51:26</PROPERTY>
		<PROPERTY name="time-stamp-numeric" type="uint32" size="25" draw="false" sort="datetime" display-name="Time">1622627486</PROPERTY>
	</OBJECT>
</RESPONSE>`))
		return
	}

	c.Data(http.StatusOK, "", []byte(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<RESPONSE VERSION="L100">
	<OBJECT basetype="status" name="status" oid="1">
		<PROPERTY name="response-type" type="string" size="12" draw="false" sort="nosort" display-name="Response Type">Error</PROPERTY>
		<PROPERTY name="response-type-numeric" type="uint32" size="12" draw="false" sort="nosort" display-name="Response Type">1</PROPERTY>
		<PROPERTY name="response" type="string" size="180" draw="true" sort="nosort" display-name="Response">Authentication Unsuccessful</PROPERTY>
		<PROPERTY name="return-code" type="sint32" size="15" draw="false" sort="nosort" display-name="Return Code">2</PROPERTY>
		<PROPERTY name="component-id" type="string" size="80" draw="false" sort="nosort" display-name="Component ID"></PROPERTY>
		<PROPERTY name="time-stamp" type="string" size="25" draw="false" sort="datetime" display-name="Time">2021-06-02 16:39:24</PROPERTY>
		<PROPERTY name="time-stamp-numeric" type="uint32" size="25" draw="false" sort="datetime" display-name="Time">1622651964</PROPERTY>
	</OBJECT>
</RESPONSE>`))
	return
}
