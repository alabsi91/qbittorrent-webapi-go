package qbittorrent

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

/*
# Params
  - "pattern" Pattern to search for (e.g. "Ubuntu 18.04")
  - "plugins" Plugins to use for searching (e.g. "legittorrents"). Also supports `all` and `enabled`
  - "category" Categories to limit your search to (e.g. "legittorrents"). Available categories depend on the specified plugins. Also supports `all`

# Returns
  - "id" ID of the search job

# Http Error Codes
  - 409 User has reached the limit of max Running searches (currently set to 5)
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#start-search
*/
func (c *Client) StartSearch(pattern string, plugins []string, category []string) (results int, err error) {
	form := url.Values{}
	form.Add("pattern", pattern)
	form.Add("plugins", strings.Join(plugins, "|"))
	form.Add("category", strings.Join(category, ","))

	formDataBytes := []byte(form.Encode())
	body, err := c.postReq("/api/v2/search/start", &formDataBytes)
	if err != nil {
		return
	}

	resultStruct := SearchId{}
	err = json.Unmarshal(body, &resultStruct)
	if err != nil {
		return
	}

	results = resultStruct.Id

	return
}

/*
# Params
  - "id" ID of the search job

# Http Error Codes
  - 404 Forbidden, Search job was not found
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#stop-search
*/
func (c *Client) StopSearch(id int) (err error) {
	form := url.Values{}
	form.Add("id", strconv.Itoa(id))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/search/stop", &formDataBytes)
	return
}

/*
# Params
  - "id" (optional) ID of the search job. If not specified, all search jobs are returned

# Http Error Codes
  - 404 Search job was not found
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-search-status
*/
func (c *Client) GetSearchStatus(id *int) (results []SearchStatusResponse, err error) {
	form := url.Values{}
	if id != nil {
		form.Add("id", strconv.Itoa(*id))
	}

	body, err := c.getReq("/api/v2/search/status", &form)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &results)
	return
}

/*
# Params
  - "id" ID of the search job
  - "limit" (optional) max number of results to return. 0 or negative means no limit
  - "offset" (optional) result to start at. A negative number means count backwards (e.g. -2 returns the 2 most recent results)

# Http Error Codes
  - 404 Search job was not found
  - 409 Offset is too large, or too small (e.g. absolute value of negative number is greater than # results)
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-search-results
*/
func (c *Client) GetSearchResults(id int, limit, offset *int) (results SearchResultsResponse, err error) {
	form := url.Values{}
	form.Add("id", strconv.Itoa(id))
	if limit != nil {
		form.Add("limit", strconv.Itoa(*limit))
	}
	if offset != nil {
		form.Add("offset", strconv.Itoa(*offset))
	}

	body, err := c.getReq("/api/v2/search/results", &form)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &results)
	return
}

/*
# Params
  - "id" ID of the search job

# Http Error Codes
  - 404 Search job was not found
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#delete-search
*/
func (c *Client) DeleteSearch(id int) (err error) {
	form := url.Values{}
	form.Add("id", strconv.Itoa(id))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/search/delete", &formDataBytes)
	return
}

/*
# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-search-plugins
*/
func (c *Client) GetSearchPlugins() (results []SearchPluginsResponse, err error) {
	body, err := c.getReq("/api/v2/search/plugins", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &results)
	return
}

/*
# Params
  - "sources" Url or file path of the plugin to install (e.g. "https://raw.githubusercontent.com/qbittorrent/search-plugins/master/nova3/engines/legittorrents.py")

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#install-search-plugin
*/
func (c *Client) InstallSearchPlugin(sources []string) (err error) {
	form := url.Values{}
	form.Add("sources", strings.Join(sources, "|"))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/search/installPlugin", &formDataBytes)
	return
}

/*
# Params
  - "names" Name of the plugin to uninstall (e.g. "legittorrents").

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#uninstall-search-plugin
*/
func (c *Client) UninstallSearchPlugin(names []string) (err error) {
	form := url.Values{}
	form.Add("names", strings.Join(names, "|"))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/search/uninstallPlugin", &formDataBytes)
	return
}

/*
# Params
  - "names" Name of the plugin to enable/disable (e.g. "legittorrents").
  - "enable" Whether the plugins should be enabled

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#enable-search-plugin
*/
func (c *Client) EnableSearchPlugin(names []string, enable bool) (err error) {
	form := url.Values{}
	form.Add("names", strings.Join(names, "|"))
	form.Add("enable", strconv.FormatBool(enable))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/search/enablePlugin", &formDataBytes)
	return
}

/*
# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#update-search-plugins
*/
func (c *Client) UpdateSearchPlugins() (err error) {
	_, err = c.postReq("/api/v2/search/updatePlugins", nil)
	return
}
