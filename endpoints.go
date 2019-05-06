package dothill

import (
	"crypto/md5"
	"fmt"
)

// Login : Must be called before any other route, authentitcate to the API
func (client *Client) Login() error {
	userpass := fmt.Sprintf("%s_%s", client.Username, client.Password)
	hash := md5.Sum([]byte(userpass))
	res, _, err := client.Request(&Request{Endpoint: fmt.Sprintf("/login/%x", hash)})

	if err != nil {
		return err
	}

	client.sessionKey = res.objectsMap["status"].propertiesMap["response"].Data
	return nil
}

// TestCall : test call for mock API
func (client *Client) TestCall() (*TestModel, *ResponseStatus, error) {
	res := &TestModel{}
	status, err := client.requestAndConvert(res, &Request{Endpoint: "/create/vdisk/level/r5/disks/2.6,2.7,2.8/vd-1"})
	return res, status, err
}
