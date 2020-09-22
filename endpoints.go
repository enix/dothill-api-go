package dothill

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// Login : Called automatically, may be called manually if credentials changed
func (client *Client) Login() error {
	userpass := fmt.Sprintf("%s_%s", client.Username, client.Password)
	hash := md5.Sum([]byte(userpass))
	res, _, err := client.Request(fmt.Sprintf("/login/%x", hash))

	if err != nil {
		return err
	}

	client.sessionKey = res.ObjectsMap["status"][0].PropertiesMap["response"].Data
	return nil
}

// CreateVolume : creates a volume with the given name, capacity in the given pool
func (client *Client) CreateVolume(name, size, pool string) (*Response, *ResponseStatus, error) {
	return client.Request(fmt.Sprintf("/create/volume/pool/\"%s\"/size/%s/tier-affinity/no-affinity/\"%s\"", pool, size, name))
}

// CreateHost : creates a host
func (client *Client) CreateHost(name, iqn string) (*Response, *ResponseStatus, error) {
	return client.Request(fmt.Sprintf("/create/host/id/\"%s\"/\"%s\"", iqn, name))
}

// MapVolume : map a volume to host + LUN
func (client *Client) MapVolume(name, host, access string, lun int) (*Response, *ResponseStatus, error) {
	return client.Request(fmt.Sprintf("/map/volume/access/%s/lun/%d/host/%s/\"%s\"", access, lun, host, name))
}

// ShowVolumes : get informations about volumes
func (client *Client) ShowVolumes(volumes ...string) (*Response, *ResponseStatus, error) {
	return client.Request(fmt.Sprintf("/show/volumes/\"%s\"", strings.Join(volumes, ",")))
}

// UnmapVolume : unmap a volume from host
func (client *Client) UnmapVolume(name, host string) (*Response, *ResponseStatus, error) {
	var url string

	if host == "" {
		url = fmt.Sprintf("/unmap/volume/\"%s\"", name)
	} else {
		url = fmt.Sprintf("/unmap/volume/host/\"%s\"/\"%s\"", host, name)
	}

	return client.Request(url)
}

// ExpandVolume : extend a volume if there is enough space on the vdisk
func (client *Client) ExpandVolume(name, size string) (*Response, *ResponseStatus, error) {
	return client.Request(fmt.Sprintf("/expand/volume/size/\"%s\"/\"%s\"", size, name))
}

// DeleteVolume : deletes a volume
func (client *Client) DeleteVolume(name string) (*Response, *ResponseStatus, error) {
	return client.Request(fmt.Sprintf("/delete/volumes/\"%s\"", name))
}

// DeleteHost : deletes a hotst by its ID or nickname
func (client *Client) DeleteHost(name string) (*Response, *ResponseStatus, error) {
	return client.Request(fmt.Sprintf("/delete/host/\"%s\"", name))
}

// ShowHostMaps : list the volume mappings for given host
// If host is an empty string, mapping for all hosts is shown
func (client *Client) ShowHostMaps(host string) ([]Volume, *ResponseStatus, error) {
	if len(host) > 0 {
		host = fmt.Sprintf("\"%s\"", host)
	}
	res, status, err := client.Request(fmt.Sprintf("/show/host-maps/%s", host))
	if err != nil {
		return nil, status, err
	}

	mappings := make([]Volume, 0)
	for _, rootObj := range res.Objects {
		if rootObj.Name != "host-view" {
			continue
		}

		for _, object := range rootObj.Objects {
			if object.Name == "volume-view" {
				vol := Volume{}
				vol.fillFromObject(&object)
				mappings = append(mappings, vol)
			}
		}
	}

	return mappings, status, err
}
