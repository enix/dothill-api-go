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
	Objects []Object
}

// Object : Typed representation of any XML API object
type Object struct {
	ID         int32
	Type       string
	Name       string
	Format     string
	Objects    []Object
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

	objects := objectsPropertiesToMap(res.Objects)
	response := Response{
		Version: res.Version,
		Objects: objects,
	}

	response.computeStatus()
	return &response, nil
}

func (res *Response) computeStatus() {
	var statusObject *Object
	objects := []Object{}
	for _, object := range res.Objects {
		if object.Type == "status" {
			statusObject = &object
		} else {
			objects = append(objects, object)
		}
	}
	if statusObject == nil {
		return
	}

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

	res.Objects = objects
}

func objectsPropertiesToMap(rawObjects []rawObject) []Object {
	objects := []Object{}

	for idx := range rawObjects {
		rawSubObject := &rawObjects[idx]
		objects = append(objects, *objectFromRawObject(rawSubObject))
	}

	return objects
}

func objectFromRawObject(src *rawObject) *Object {
	object := &Object{
		ID:         src.OID,
		Type:       src.BaseType,
		Name:       src.Name,
		Format:     src.Format,
		Properties: make(map[string]Property),
		Objects:    objectsPropertiesToMap(src.Objects),
	}

	for idx := range src.Properties {
		prop := src.Properties[idx]
		prop.Data = strings.TrimSpace(prop.Data)
		object.Properties[prop.Name] = prop
	}

	return object
}
