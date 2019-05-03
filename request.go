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

func (req *Request) execute(client *Client) ([]byte, error) {
	httpClient := &http.Client{}
	url := fmt.Sprintf("%s/api%s", client.Options.Addr, req.Endpoint)
	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("sessionKey", client.sessionKey)
	res, err := httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("API returned unexpected HTTP status %d", res.StatusCode)
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
