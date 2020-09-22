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
	ObjectsMap map[string][]*Object
}

// Object : Typed representation of any XML API object
type Object struct {
	Typ           string     `xml:"basetype,attr"`
	Name          string     `xml:"name,attr"`
	OID           int32      `xml:"oid,attr"`
	Format        string     `xml:"format,attr,omitempty"`
	Objects       []Object   `xml:"OBJECT"`
	Properties    []Property `xml:"PROPERTY"`
	ObjectsMap    map[string][]*Object
	PropertiesMap map[string]*Property
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

	res.ObjectsMap = objectsToMap(res.Objects)

	return res, nil
}

// GetStatus : Creates and returns the final ResponseStatus struct
// from the raw status object in response
func (res *Response) GetStatus() *ResponseStatus {
	statusObject := res.ObjectsMap["status"][0]

	responseTypeNumeric, _ := strconv.Atoi(statusObject.PropertiesMap["response-type-numeric"].Data)
	returnCode, _ := strconv.Atoi(statusObject.PropertiesMap["return-code"].Data)
	timestampNumeric, _ := strconv.Atoi(statusObject.PropertiesMap["time-stamp-numeric"].Data)

	return &ResponseStatus{
		ResponseType:        statusObject.PropertiesMap["response-type"].Data,
		ResponseTypeNumeric: responseTypeNumeric,
		Response:            statusObject.PropertiesMap["response"].Data,
		ReturnCode:          returnCode,
		Time:                time.Unix(int64(timestampNumeric), 0),
	}
}

func objectsToMap(objects []Object) map[string][]*Object {
	objectsMap := make(map[string][]*Object)
	
	for idx := range objects {
		subObject := &objects[idx]
		fillObjectMap(subObject)
		if objectsMap[subObject.Name] == nil {
			objectsMap[subObject.Name] = make([]*Object, 0)
		}
		objectsMap[subObject.Name] = append(objectsMap[subObject.Name], subObject)
	}

	return objectsMap
}

func fillObjectMap(obj *Object) {
	obj.PropertiesMap = make(map[string]*Property)

	for idx := range obj.Properties {
		prop := &obj.Properties[idx]
		prop.Data = strings.TrimSpace(prop.Data)
		obj.PropertiesMap[prop.Name] = prop
	}

	obj.ObjectsMap = objectsToMap(obj.Objects)
}
