package dothill

import (
	"crypto/md5"
	"fmt"
)

// Login : Must be called before any other route, authentitcate to the API
func (client *Client) Login() error {
	userpass := fmt.Sprintf("%s_%s", client.Username, client.Password)
	hash := md5.Sum([]byte(userpass))
	res, _, err := client.Request(fmt.Sprintf("/login/%x", hash))

	if err != nil {
		return err
	}

	client.sessionKey = res.ObjectsMap["status"].PropertiesMap["response"].Data
	return nil
}

// TestCall : test call for mock API
// func (client *Client) TestCall() (*TestModel, *ResponseStatus, error) {
// 	res := &TestModel{}
// 	status, err := client.requestAndConvert(res, "/create/vdisk/level/r5/disks/2.6,2.7,2.8/vd-1")
// 	return res, status, err
// }

// CreateVolume : creates a volume with the given name, capacity in the given pool
func (client *Client) CreateVolume(name, size, pool string) (*Response, *ResponseStatus, error) {
	return client.Request(fmt.Sprintf("/create/volume/pool/\"%s\"/size/%s/tier-affinity/no-affinity/\"%s\"", pool, size, name))
}

// MapVolume : map a volume to host + LUN
func (client *Client) MapVolume(name, host string, lun int) (*Response, *ResponseStatus, error) {
	return client.Request(fmt.Sprintf("/map/volume/access/rw/lun/%d/host/%s/\"%s\"", lun, host, name))
}

// UnmapVolume : unmap a volume from host
func (client *Client) UnmapVolume(name, host string) (*Response, *ResponseStatus, error) {
	return client.Request(fmt.Sprintf("/unmap/volume/host/\"%s\"/\"%s\"", host, name))
}

// DeleteVolume : deletes a volume
func (client *Client) DeleteVolume(name string) (*Response, *ResponseStatus, error) {
	return client.Request(fmt.Sprintf("/delete/volumes/\"%s\"", name))
}

// ShowHostMaps : list the volume mappings for given host
func (client *Client) ShowHostMaps(host string) ([]Volume, *ResponseStatus, error) {
	res, status, err := client.Request(fmt.Sprintf("/show/host-maps/\"%s\"", host))
	hostView := res.ObjectsMap["host-view"]
	mappings := make([]Volume, 0, len(hostView.Objects))

	for _, object := range hostView.Objects {
		if object.Name == "volume-view" {
			vol := Volume{}
			vol.fillFromObject(&object)
			mappings = append(mappings, vol)
		}
	}

	return mappings, status, err
}
