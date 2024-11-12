package proxmox_client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"net/http"
	"net/url"
)

func (c *Client) CreateSdnZone(sdnZoneCreationBody url.Values) error {

	request, requestCreationError := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/cluster/sdn/zones", c.HostURL), bytes.NewBufferString(sdnZoneCreationBody.Encode()))

	if requestCreationError != nil {
		return requestCreationError
	}

	body, responseError := c.doRequest(request, FORM_URL_ENCODED)
	if responseError != nil {
		return errors.Join(responseError, errors.New(string(body)))
	}

	return nil
}

func (c *Client) GetSdnZone(zone string) (*SdnZoneResponse, error) {

	request, requestCreationError := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/cluster/sdn/zones/%s", c.HostURL, url.PathEscape(zone)), nil)

	if requestCreationError != nil {
		return nil, requestCreationError
	}

	body, responseError := c.doRequest(request, FORM_URL_ENCODED)
	if responseError != nil {
		return nil, errors.Join(responseError, errors.New(string(body)))
	}
	tflog.Debug(c.Context, string(body))
	var zoneCreationResponse = SdnZoneResponse{}

	unmarshallingError := json.Unmarshal(body, &zoneCreationResponse)
	if unmarshallingError != nil {
		return nil, unmarshallingError
	}

	return &zoneCreationResponse, nil
}

func (c *Client) DeleteSdnZone(zone string) error {

	request, requestCreationError := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/cluster/sdn/zones/%s", c.HostURL, url.PathEscape(zone)), nil)

	if requestCreationError != nil {
		return requestCreationError
	}

	body, responseError := c.doRequest(request, FORM_URL_ENCODED)
	if responseError != nil {
		return errors.Join(responseError, errors.New(string(body)))
	}

	return nil
}

func (c *Client) UpdateSdnZone(sdnZoneCreationBody url.Values) error {

	request, requestCreationError := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/cluster/sdn/zones", c.HostURL), bytes.NewBufferString(sdnZoneCreationBody.Encode()))

	if requestCreationError != nil {
		return requestCreationError
	}

	body, responseError := c.doRequest(request, FORM_URL_ENCODED)
	if responseError != nil {
		return errors.Join(responseError, errors.New(string(body)))
	}

	return nil
}
