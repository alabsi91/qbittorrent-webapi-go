package qbittorrent

import (
	"encoding/json"
	"net/url"
	"strconv"
)

/*
# Params
  - "rid" Response ID. If not provided, rid=0 will be assumed. If the given rid is different from the one of last server reply, full_update will be true (see the server reply details for more info)

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-main-data
*/
func (c *Client) GetSyncMainData(rid int) (results SyncMainDataResponse, err error) {

	queryParams := url.Values{}
	queryParams.Add("rid", strconv.Itoa(rid))

	body, err := c.getReq("/api/v2/sync/maindata", &queryParams)
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
  - "hash" Torrent hash
  - "rid" Response ID. If not provided, rid=0 will be assumed. If the given rid is different from the one of last server reply, full_update will be true (see the server reply details for more info)

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-torrent-peers-data
*/
func (c *Client) GetSyncTorrentPeersData(hash string, rid int) (results map[string]interface{}, err error) {

	queryParams := url.Values{}
	queryParams.Add("hash", hash)
	queryParams.Add("rid", strconv.Itoa(rid))

	body, err := c.getReq("/api/v2/sync/torrentPeers", &queryParams)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &results)
	if err != nil {
		return
	}

	return
}
