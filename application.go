package qbittorrent

import (
	"encoding/json"
	"net/url"
)

/*
The response is a string with the application version, e.g. v4.1.3

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-application-version
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

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-api-version
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

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-build-info
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

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#shutdown-application
*/
func (c *Client) ShutdownApplication() (err error) {
	_, err = c.postReq("/api/v2/app/shutdown", nil)
	if err != nil {
		return
	}
	return
}

/*
Returns the application's settings. The contents may vary depending on which settings are present in qBittorrent.ini.

# Http Error Codes:
- 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-application-preferences
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

# Example

	params := qbittorrent.NewPreferences().
	 Locale("en").
	 SavePath("/home/user/Downloads")

	err = client.SetApplicationPreferences(params.Data)
	if err != nil {
	 panic(err)
	}

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-application-preferences
*/
func (c *Client) SetApplicationPreferences(prefs map[string]interface{}) (err error) {
	jsonPrefs, err := json.Marshal(prefs)
	if err != nil {
		return
	}

	params := url.Values{}
	params.Add("json", string(jsonPrefs))

	_, err = c.postReq("/api/v2/app/setPreferences", &params)
	if err != nil {
		return
	}

	return
}

/*
The response is a string with the default save path, e.g. C:/Users/Dayman/Downloads

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-default-save-path
*/
func (c *Client) GetDefaultSavePath() (path string, err error) {
	body, err := c.getReq("/api/v2/app/defaultSavePath", nil)
	if err != nil {
		return
	}

	path = string(body)

	return
}
