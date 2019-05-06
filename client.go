package dothill

import (
	"fmt"
)

// Options : Client configuration, including authentication credentials
type Options struct {
	Username string
	Password string
	Addr     string
}

// Client : Can be used to request the dothill API
type Client struct {
	Options    *Options
	sessionKey string
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

func (client *Client) requestAndConvert(model model, req *Request) (*ResponseStatus, error) {
	res, status, err := client.Request(req)
	if err != nil {
		return status, err
	}
	model.fillFromResponse(res)
	return status, nil
}
