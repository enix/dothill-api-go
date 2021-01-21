/*
 * Copyright (c) 2021 Enix, SAS
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing
 * permissions and limitations under the License.
 *
 * Authors:
 * Paul Laffitte <paul.laffitte@enix.fr>
 * Arthur Chaloin <arthur.chaloin@enix.fr>
 * Alexandre Buisine <alexandre.buisine@enix.fr>
 */

package dothill

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"k8s.io/klog"
)

// Client : Can be used to request the dothill API
type Client struct {
	Username   string
	Password   string
	Addr       string
	HTTPClient http.Client
	Collector  *Collector
	sessionKey string
}

// NewClient : Creates a dothill client by setting up its HTTP client
func NewClient() *Client {
	return &Client{
		HTTPClient: http.Client{
			Timeout: time.Duration(15 * time.Second),
			Transport: &http.Transport{
				// Proxy: http.ProxyURL(proxy),
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		Collector: newCollector(),
	}
}

// Request : Execute the given request with client's configuration
// Deprecated: Use FormattedRequest instead
func (client *Client) Request(endpoint string) (*Response, error) {
	if client.Addr == "" {
		err := errors.New("missing server address")
		return nil, err
	}

	return client.request(&Request{Endpoint: endpoint})
}

// FormattedRequest : Format and execute the given request with client's configuration
func (client *Client) FormattedRequest(endpointFormat string, opts ...interface{}) (*Response, error) {
	endpoint := fmt.Sprintf(endpointFormat, opts...)
	stopTrackAPICall := client.Collector.trackAPICall(endpointFormat)
	resp, err := client.Request(endpoint)
	stopTrackAPICall(err == nil)
	return resp, err
}

func (client *Client) request(req *Request) (*Response, error) {
	isLoginReq := strings.Contains(req.Endpoint, "login")
	if !isLoginReq {
		if len(client.sessionKey) == 0 {
			klog.V(1).Info("no session key stored, authenticating before sending request")
			err := client.Login()
			if err != nil {
				return NewErrorResponse("login failed"), err
			}
		}

		klog.Infof("-> GET %s", req.Endpoint)
	} else {
		klog.Infof("-> GET /login/<hidden>")
	}

	raw, code, err := req.execute(client)
	if code == 401 && !isLoginReq {
		klog.V(1).Info("session key may have expired, trying to re-login")
		err = client.Login()
		if err != nil {
			return NewErrorResponse("re-login failed"), err
		}
		klog.V(1).Info("re-login succeed, re-trying request")
		raw, _, err = req.execute(client)
	}
	if err != nil {
		return NewErrorResponse("request failed"), err
	}

	res, err := NewResponse(raw)
	if err != nil {
		if res != nil {
			return res, err
		}

		return NewErrorResponse("corrupted response"), err
	}

	status := res.Status
	if !isLoginReq {
		klog.Infof("<- [%d %s] %s", status.ReturnCode, status.ResponseType, status.Response)
	} else {
		klog.Infof("<- [%d %s] <hidden>", status.ReturnCode, status.ResponseType)
	}
	if status.ResponseTypeNumeric != 0 {
		return res, fmt.Errorf("Dothill API returned non-zero code %d (%s)", status.ReturnCode, status.Response)
	}

	return res, nil
}
