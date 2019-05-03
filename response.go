package dothill

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"
)

// Response : Typed representation of any XML API response
type Response struct {
	Version    string   `xml:"VERSION,attr"`
	Objects    []Object `xml:"OBJECT"`
	objectsMap map[string]*Object
}

// Object : Typed representation of any XML API object
type Object struct {
	Typ           string     `xml:"basetype,attr"`
	Name          string     `xml:"name,attr"`
	OID           int32      `xml:"oid,attr"`
	Format        string     `xml:"format,attr,omitempty"`
	Properties    []Property `xml:"PROPERTY"`
	propertiesMap map[string]*Property
}

// Property : Typed representation of any XML API property
type Property struct {
	Name        string `xml:"name,attr"`
	Typ         string `xml:"type,attr"`
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

// NewResponse : Unmarshals the raw data into a typed response object
// and generate a hash map from fields for optimization
func NewResponse(data []byte) (*Response, error) {
	res := &Response{}
	err := xml.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	res.objectsMap = make(map[string]*Object)
	for idx := range res.Objects {
		obj := &res.Objects[idx]
		obj.propertiesMap = make(map[string]*Property)
		res.objectsMap[obj.Name] = obj
		for idx2 := range obj.Properties {
			prop := &obj.Properties[idx2]
			prop.Data = strings.TrimSpace(prop.Data)
			obj.propertiesMap[prop.Name] = prop
		}
	}

	return res, nil
}

// GetStatus : Creates and returns the final ResponseStatus struct
// from the raw status object in response
func (res *Response) GetStatus() *ResponseStatus {
	statusObject := res.objectsMap["status"]
	responseTypeNumeric, _ := strconv.Atoi(statusObject.propertiesMap["response-type-numeric"].Data)
	returnCode, _ := strconv.Atoi(statusObject.propertiesMap["return-code"].Data)
	timestampNumeric, _ := strconv.Atoi(statusObject.propertiesMap["time-stamp-numeric"].Data)

	return &ResponseStatus{
		ResponseType:        statusObject.propertiesMap["response-type"].Data,
		ResponseTypeNumeric: responseTypeNumeric,
		Response:            statusObject.propertiesMap["response"].Data,
		ReturnCode:          returnCode,
		Time:                time.Unix(int64(timestampNumeric), 0),
	}
}
