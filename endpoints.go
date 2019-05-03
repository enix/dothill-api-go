package dothill

// ShowDisksOutput : Typed Go representation of the /show/disks API call response
type ShowDisksOutput struct {
	Data string
}

// ShowDisks : /show/disks API call
func (client *Client) ShowDisks() (*ShowDisksOutput, *ResponseStatus, error) {
	res, status, err := client.Request(&Request{
		Endpoint: "/show/disks",
	})
	if err != nil {
		return nil, status, err
	}

	out := &ShowDisksOutput{
		Data: res.objectsMap["status"].propertiesMap["response"].Data,
	}

	return out, status, nil
}

// type loginOutput struct{}

// func (client *Client) login() (string, error) {
// 	userpass := fmt.Sprintf("%s_%s", client.Options.Username, client.Options.Password)
// 	err := client.Request()
// 	md5.Sum([]byte(userpass))
// 	return "", nil
// }
