package dothill

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request : Used internally, and can be used to send custom requests (see Client.Request())
type Request struct {
	Endpoint string
	Data     interface{}
}

func (req *Request) execute(client *Client) ([]byte, int, error) {
	url := fmt.Sprintf("%s/api%s", client.Addr, req.Endpoint)
	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}

	httpReq.Header.Set("sessionKey", client.sessionKey)
	res, err := client.HTTPClient.Do(httpReq)
	if err != nil {
		return nil, 0, err
	}

	if res.StatusCode >= 400 {
		return nil, res.StatusCode, fmt.Errorf("API returned unexpected HTTP status %d", res.StatusCode)
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	return data, res.StatusCode, err
}
