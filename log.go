package qbittorrent

import (
	"encoding/json"
	"net/url"
	"strconv"
)

/*
# Params [GetLogParams] struct
  - "Normal" Include normal messages (default: true)
  - "Info" Include info messages (default: true)
  - "Warning" Include warning messages (default: true)
  - "Critical" Include critical messages (default: true)
  - "LastKnownId" integer Exclude messages with "message id" <= last_known_id (default: -1)

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-log
*/
func (c *Client) GetLog(params *GetLogParams) (results []GetLogResponse, err error) {

	defaultLog := true
	if params == nil {
		params = &GetLogParams{}
	}
	if params.Normal == nil {
		params.Normal = &defaultLog
	}
	if params.Info == nil {
		params.Info = &defaultLog
	}
	if params.Warning == nil {
		params.Warning = &defaultLog
	}
	if params.Critical == nil {
		params.Critical = &defaultLog
	}
	if params.LastKnownId == nil {
		defaultLastKnownId := -1
		params.LastKnownId = &defaultLastKnownId
	}

	queryParams := url.Values{}
	queryParams.Add("normal", strconv.FormatBool(*params.Normal))
	queryParams.Add("info", strconv.FormatBool(*params.Info))
	queryParams.Add("warning", strconv.FormatBool(*params.Warning))
	queryParams.Add("critical", strconv.FormatBool(*params.Critical))
	queryParams.Add("last_known_id", strconv.Itoa(*params.LastKnownId))

	body, err := c.getReq("/api/v2/log/main", &queryParams)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &results)
	if err != nil {
		return
	}

	return
}

/*
# Params
  - "lastKnownId" Exclude messages with "message id" <= last_known_id (default: -1)

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-peer-log
*/
func (c *Client) GetPeerLog(lastKnownId int) (results []GetPeerLogResponse, err error) {

	queryParams := url.Values{}
	queryParams.Add("last_known_id", strconv.Itoa(lastKnownId))

	body, err := c.getReq("/api/v2/log/peers", &queryParams)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &results)
	if err != nil {
		return
	}

	return
}
