package dothill

var rawResponse = `
	<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	<RESPONSE VERSION="L100">
	<OBJECT basetype="status" name="status" oid="1">
			<PROPERTY name="response-type" type="string" size="12" draw="false"
	sort="nosort" display-name="Response Type">Success</PROPERTY>
			<PROPERTY name="response-type-numeric" type="uint32" size="12" draw="false"
	sort="nosort" display-name="Response Type">0</PROPERTY>
			<PROPERTY name="response" type="string" size="180" draw="true" sort="nosort"
	display-name="Response">ok</PROPERTY>
			<PROPERTY name="return-code" type="sint32" size="15" draw="false"
	sort="nosort" display-name="Return Code">0</PROPERTY>
			<PROPERTY name="component-id" type="string" size="80" draw="false"
	sort="nosort" display-name="Component ID">vd-1</PROPERTY>
			<PROPERTY name="time-stamp" type="string" size="25" draw="false"
	sort="datetime" display-name="Time">2010-08-10 12:07:18</PROPERTY>
			<PROPERTY name="time-stamp-numeric" type="uint32" size="25" draw="false"
	sort="datetime" display-name="Time">1281442038</PROPERTY>
	</OBJECT>
	</RESPONSE>
`

// Request : Used internally, and can be used to send custom requests (see Client.Request())
type Request struct {
	Endpoint string
	Data     interface{}
}

func (req *Request) execute(client *Client) ([]byte, error) {
	return []byte(rawResponse), nil
}
