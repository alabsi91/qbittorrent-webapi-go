package qbittorrent

import (
	"encoding/json"
	"net/url"
	"strconv"
	"strings"
)

/*
This method returns info you usually see in qBt status bar.

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-global-transfer-info
*/
func (c *Client) GetGlobalTransferInfo() (results TransferInfoResponse, err error) {
	body, err := c.getReq("/api/v2/transfer/info", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &results)
	return
}

/*
# Returns
  - The response is `1` if alternative speed limits are enabled, `0` otherwise.

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-alternative-speed-limits-state
*/
func (c *Client) GetAlternativeSpeedLimitsState() (results AlternativeSpeedLimitsStatus, err error) {
	body, err := c.getReq("/api/v2/transfer/speedLimitsMode", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &results)
	return
}

/*
# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#toggle-alternative-speed-limits
*/
func (c *Client) ToggleAlternativeSpeedLimits() (err error) {
	_, err = c.postReq("/api/v2/transfer/toggleSpeedLimitsMode", nil)
	if err != nil {
		return
	}
	return
}

/*
# Returns
  - the value of current global download speed limit in bytes/second; this value will be zero if no limit is applied.

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-global-download-limit
*/
func (c *Client) GetGlobalDownloadLimit() (results int, err error) {
	body, err := c.getReq("/api/v2/transfer/downloadLimit", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &results)
	return
}

/*
# Params
  - "limit" The global download speed limit to set in bytes/second

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-global-download-limit
*/
func (c *Client) SetGlobalDownloadLimit(limit int) (err error) {
	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	_, err = c.postReq("/api/v2/transfer/setDownloadLimit", &params)
	return
}

/*
# Returns
  - the value of current global upload speed limit in bytes/second; this value will be zero if no limit is applied.

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-global-upload-limit
*/
func (c *Client) GetGlobalUploadLimit() (results int, err error) {
	body, err := c.getReq("/api/v2/transfer/uploadLimit", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &results)
	return
}

/*
# Params
  - "limit" The global upload speed limit to set in bytes/second

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-global-upload-limit
*/
func (c *Client) SetGlobalUploadLimit(limit int) (err error) {
	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	_, err = c.postReq("/api/v2/transfer/setUploadLimit", &params)
	return
}

/*
# Params
  - "peers" The list of peers to ban

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#ban-peers
*/
func (c *Client) BanPeers(peers []string) (err error) {
	params := url.Values{}
	params.Add("peers", strings.Join(peers, "|"))
	_, err = c.postReq("/api/v2/transfer/banPeers", &params)
	return
}
