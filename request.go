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
 * Joe Skazinski <joseph.skazinski@seagate.com>
 */

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
	httpReq.SetBasicAuth(client.Username, client.Password)
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
