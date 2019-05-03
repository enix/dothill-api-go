package dothill

import (
	"errors"
	"fmt"
)

// Options : Client configuration, including authentication credentials
type Options struct {
	Username string
	Password string
	Addr     string
	Port     int16
}

// Client : Can be used to request the dothill API
type Client struct {
	Options    *Options
	sessionKey string
}

// NewClient : Create a client from the given options
func NewClient(options *Options) (*Client, error) {
	if len(options.Username) < 1 {
		return nil, errors.New("please provide a username")
	}

	if len(options.Password) < 1 {
		return nil, errors.New("please provide a password")
	}

	client := &Client{}
	// sessionKey, err := client.login()
	// if err != nil {
	// 	return nil, err
	// }

	// client.sessionKey = sessionKey
	return client, nil
}

// Request : Execute the given request with client's configuration
func (client *Client) Request(req *Request) (*Response, *ResponseStatus, error) {
	raw, err := req.execute(client)
	if err != nil {
		return nil, nil, err
	}

	res, err := NewResponse(raw)
	if err != nil {
		if res != nil {
			return res, res.GetStatus(), err
		}

		return nil, nil, err
	}

	status := res.GetStatus()
	if status.ReturnCode != 0 {
		return res, status, fmt.Errorf("API returned non-zero code %d", status.ReturnCode)
	}

	return res, status, nil
}
