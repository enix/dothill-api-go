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
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

// Response : Typed representation of any XML API response
type Response struct {
	Version string
	Status  ResponseStatus
	Objects map[string][]Object
}

// Object : Typed representation of any XML API object
type Object struct {
	ID         int32
	Type       string
	Name       string
	Format     string
	Objects    map[string][]Object
	Properties map[string]Property
}

// Property : Typed representation of any XML API property
type Property struct {
	Name        string `xml:"name,attr"`
	Type        string `xml:"type,attr"`
	Size        int32  `xml:"size,attr"`
	Draw        bool   `xml:"draw,attr"`
	Sort        string `xml:"sort,attr"`
	DisplayName string `xml:"display-name,attr"`
	Data        string `xml:",chardata"`
}

// ResponseStatus : Final representation of the "status" object in every API response
type ResponseStatus struct {
	ResponseType        string
	ResponseTypeNumeric int
	Response            string
	ReturnCode          int
	Time                time.Time
}

type rawObject struct {
	BaseType   string      `xml:"basetype,attr"`
	Name       string      `xml:"name,attr"`
	OID        int32       `xml:"oid,attr"`
	Format     string      `xml:"format,attr,omitempty"`
	Objects    []rawObject `xml:"OBJECT"`
	Properties []Property  `xml:"PROPERTY"`
}

// NewResponse : Unmarshals the raw data into a typed response object
// and generate a hash map from fields for optimization
func NewResponse(data []byte) (*Response, error) {
	type rawResponse struct {
		Version string      `xml:"VERSION,attr"`
		Objects []rawObject `xml:"OBJECT"`
	}

	res := &rawResponse{}
	err := xml.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	objects := objectsToMap(res.Objects)
	response := Response{
		Version: res.Version,
		Objects: objects,
	}

	response.computeStatus()
	return &response, nil
}

// NewErrorResponse : Creates a response with an error status
func NewErrorResponse(err string) *Response {
	return &Response{Status: *NewErrorStatus(err)}
}

// NewErrorStatus : Creates an error status
func NewErrorStatus(err string) *ResponseStatus {
	return &ResponseStatus{
		ResponseType: "Error",
		Response:     err,
		Time:         time.Now(),
	}
}

func (res *Response) computeStatus() {
	statusObjects, statusObjectsExists := res.Objects["status"]
	if !statusObjectsExists || len(statusObjects) == 0 {
		return
	}

	statusObject := statusObjects[0]
	responseTypeNumeric, _ := strconv.Atoi(statusObject.Properties["response-type-numeric"].Data)
	returnCode, _ := strconv.Atoi(statusObject.Properties["return-code"].Data)
	timestampNumeric, _ := strconv.Atoi(statusObject.Properties["time-stamp-numeric"].Data)

	res.Status = ResponseStatus{
		ResponseType:        statusObject.Properties["response-type"].Data,
		ResponseTypeNumeric: responseTypeNumeric,
		Response:            statusObject.Properties["response"].Data,
		ReturnCode:          returnCode,
		Time:                time.Unix(int64(timestampNumeric), 0),
	}

	delete(res.Objects, "status")
}

func objectsToMap(objects []rawObject) map[string][]Object {
	objectsMap := make(map[string][]Object)

	for idx := range objects {
		subObject := &Object{}
		rawSubObject := &objects[idx]
		fillObjectMap(rawSubObject, subObject)
		if objectsMap[rawSubObject.Name] == nil {
			objectsMap[rawSubObject.Name] = make([]Object, 0)
		}
		objectsMap[rawSubObject.Name] = append(objectsMap[rawSubObject.Name], *subObject)
	}

	return objectsMap
}

func fillObjectMap(src *rawObject, dest *Object) {
	dest.ID = src.OID
	dest.Type = src.BaseType
	dest.Name = src.Name
	dest.Format = src.Format
	dest.Objects = objectsToMap(src.Objects)
	dest.Properties = make(map[string]Property)

	for idx := range src.Properties {
		prop := src.Properties[idx]
		prop.Data = strings.TrimSpace(prop.Data)
		dest.Properties[prop.Name] = prop
	}
}
