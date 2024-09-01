package qbittorrent

import (
	"encoding/json"
	"fmt"
)

/*
The response is a string with the application version, e.g. v4.1.3

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized
*/
func (c *Client) GetApplicationVersion() (version string, err error) {
	body, err := c.getReq("/api/v2/app/version", nil)
	if err != nil {
		return
	}

	version = string(body)

	return
}

/*
The response is a string with the WebAPI version, e.g. 2.0

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized
*/
func (c *Client) GetAPIVersion() (version string, err error) {
	body, err := c.getReq("/api/v2/app/webapiVersion", nil)
	if err != nil {
		return
	}

	version = string(body)

	return
}

/*
The response is a [BuildInfo] struct

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized
*/
func (c *Client) GetBuildInfo() (info BuildInfo, err error) {
	body, err := c.getReq("/api/v2/app/buildInfo", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &info)

	return
}

/*
# Http Error Codes:
  - 403 Forbidden, if the client is not authorized
*/
func (c *Client) ShutdownApplication() (err error) {
	_, err = c.postReq("/api/v2/app/shutdown", nil)
	if err != nil {
		return
	}
	return
}

/*
# Http Error Codes:
  - 403 Forbidden, if the client is not authorized
*/
func (c *Client) GetApplicationPreferences() (results ApplicationPreferences, err error) {
	body, err := c.getReq("/api/v2/app/preferences", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &results)

	return
}

/*
# Tip
  - Use [NewPreferences] to build the form then pass its `data` field

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized
*/
func (c *Client) SetApplicationPreferences(prefs map[string]interface{}) (err error) {
	jsonPrefs, err := json.Marshal(prefs)
	if err != nil {
		return
	}

	data := []byte(fmt.Sprintf("json=%s", jsonPrefs))

	_, err = c.postReq("/api/v2/app/setPreferences", &data)
	if err != nil {
		return
	}

	return
}

/*
The response is a string with the default save path, e.g. C:/Users/Dayman/Downloads

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized
*/
func (c *Client) GetDefaultSavePath() (path string, err error) {
	body, err := c.getReq("/api/v2/app/defaultSavePath", nil)
	if err != nil {
		return
	}

	path = string(body)

	return
}
