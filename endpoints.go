package dothill

// ShowDisks : /show/disks API call
func (client *Client) ShowDisks() (*TestModel, *ResponseStatus, error) {
	res := &TestModel{}
	status, err := client.requestAndConvert(res, &Request{Endpoint: "/create/vdisk"})
	return res, status, err
}

// type loginOutput struct{}

// func (client *Client) login() (string, error) {
// 	userpass := fmt.Sprintf("%s_%s", client.Options.Username, client.Options.Password)
// 	err := client.Request()
// 	md5.Sum([]byte(userpass))
// 	return "", nil
// }
