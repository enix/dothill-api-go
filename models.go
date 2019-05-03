package dothill

// model : interface to allow generic conversion from raw response to user-object
type model interface {
	FillFromResponse(res *Response)
}

type TestModel struct {
	Data string
}

func (m *TestModel) FillFromResponse(res *Response) {
	m.Data = res.objectsMap["status"].propertiesMap["response"].Data
}
