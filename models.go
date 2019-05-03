package dothill

// model : interface to allow generic conversion from raw response to user-object
type model interface {
	fillFromResponse(res *Response)
}

// TestModel : used for internal tests purposes
type TestModel struct {
	Data string
}

func (m *TestModel) fillFromResponse(res *Response) {
	m.Data = res.objectsMap["status"].propertiesMap["response"].Data
}
