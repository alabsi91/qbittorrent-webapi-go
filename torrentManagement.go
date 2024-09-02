package qbittorrent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"
)

/*
GetTorrentList - get list of torrents in the client

# Params
  - "options" (optional) [GetTorrentListOptions]

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-torrent-list
*/
func (c *Client) GetTorrentList(opts *GetTorrentListOptions) (results []TorrentListResponse, err error) {
	if opts == nil {
		opts = &GetTorrentListOptions{}
	}

	queryParams := url.Values{}
	if opts.Filter != "" {
		queryParams.Add("filter", opts.Filter)
	}
	if opts.Category != nil {
		queryParams.Add("category", *opts.Category)
	}
	if opts.Tag != nil {
		queryParams.Add("tag", *opts.Tag)
	}
	if opts.Sort != "" {
		queryParams.Add("sort", opts.Sort)
	}
	if opts.Reverse {
		queryParams.Add("reverse", "true")
	}
	if opts.Limit > 0 {
		queryParams.Add("limit", strconv.Itoa(opts.Limit))
	}
	if opts.Offset > 0 {
		queryParams.Add("offset", strconv.Itoa(opts.Offset))
	}
	if len(opts.Hashes) > 0 {
		queryParams.Add("hashes", strings.Join(opts.Hashes, "|"))
	}

	body, err := c.getReq("/api/v2/torrents/info", &queryParams)
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
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hash" The hash of the torrent you want to get the generic properties of

# Http Error Codes
  - 404 Not Found, if the torrent hash is invalid
  - 403 Forbidden, if the client is not authorized

# NB
  - `-1` is returned if the type of the property is integer but its value is not known.

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-torrent-generic-properties
*/
func (c *Client) GetTorrentGenericProperties(hash string) (results TorrentGenericProperties, err error) {
	queryParams := url.Values{}
	queryParams.Add("hash", hash)

	body, err := c.getReq("/api/v2/torrents/properties", &queryParams)
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
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hash" The hash of the torrent you want to get the generic properties of

# Http Error Codes
  - 404 Not Found, if the torrent hash is invalid
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-torrent-trackers
*/
func (c *Client) GetTorrentTrackers(hash string) (results []TorrentTracker, err error) {
	queryParams := url.Values{}
	queryParams.Add("hash", hash)

	body, err := c.getReq("/api/v2/torrents/trackers", &queryParams)
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
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hash" The hash of the torrent you want to get the webseeds of

# Http Error Codes
  - 404 Not Found, if the torrent hash is invalid
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-torrent-web-seeds
*/
func (c *Client) GetTorrentWebSeeds(hash string) (results []TorrentSeed, err error) {
	queryParams := url.Values{}
	queryParams.Add("hash", hash)

	body, err := c.getReq("/api/v2/torrents/webseeds", &queryParams)
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
Get a list of the files of a torrent

Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hash" The hash of the torrent you want to get the contents of
  - "indexes" (optional) The indexes of the files you want to retrieve.

# Http Error Codes
  - 404 Not Found, if the torrent hash is invalid
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-torrent-contents
*/
func (c *Client) GetTorrentContents(hash string, indexes ...int) (results []TorrentFile, err error) {
	queryParams := url.Values{}
	queryParams.Add("hash", hash)
	if len(indexes) > 0 {
		indexesParam := make([]string, len(indexes))
		for i, index := range indexes {
			indexesParam[i] = strconv.Itoa(index)
		}
		queryParams.Add("indexes", strings.Join(indexesParam, "|"))
	}

	body, err := c.getReq("/api/v2/torrents/files", &queryParams)
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
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hash" The hash of the torrent you want to get the pieces' states of

# Http Error Codes
  - 404 Not Found, if the torrent hash is invalid
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-torrent-pieces-states
*/
func (c *Client) GetTorrentPiecesStates(hash string) (results []TorrentPiecesState, err error) {
	queryParams := url.Values{}
	queryParams.Add("hash", hash)

	body, err := c.getReq("/api/v2/torrents/pieceStates", &queryParams)
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
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hash" The hash of the torrent you want to get the pieces' hashes of

# Http Error Codes
  - 404 Not Found, if the torrent hash is invalid
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-torrent-pieces-hashes
*/
func (c *Client) GetTorrentPiecesHashes(hash string) (results []string, err error) {
	queryParams := url.Values{}
	queryParams.Add("hash", hash)

	body, err := c.getReq("/api/v2/torrents/pieceHashes", &queryParams)
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
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to pause. or set to []string{"all"}, to pause all torrents.

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#pause-torrents
*/
func (c *Client) PauseTorrents(hashes []string) (err error) {
	data := []byte(fmt.Sprintf("hashes=%s", strings.Join(hashes, "|")))
	_, err = c.postReq("/api/v2/torrents/pause", &data)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

#Params
  - "hashes" The hashes of the torrents you want to resume. or set to []string{"all"}, to resume all torrents.

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#resume-torrents
*/
func (c *Client) ResumeTorrents(hashes []string) (err error) {
	data := []byte(fmt.Sprintf("hashes=%s", strings.Join(hashes, "|")))
	_, err = c.postReq("/api/v2/torrents/resume", &data)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to delete. or set to []string{"all"}, to resume all torrents.
  - "deleteFiles" If set to true, the downloaded data will also be deleted, otherwise has no effect.

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#delete-torrents
*/
func (c *Client) DeleteTorrents(hashes []string, deleteFiles bool) (err error) {
	data := []byte(fmt.Sprintf("hashes=%s&deleteFiles=%t", strings.Join(hashes, "|"), deleteFiles))
	_, err = c.postReq("/api/v2/torrents/delete", &data)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to recheck. or set to []string{"all"}, to recheck all torrents.

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#recheck-torrents
*/
func (c *Client) RecheckTorrents(hashes []string) (err error) {
	data := []byte(fmt.Sprintf("hashes=%s", strings.Join(hashes, "|")))
	_, err = c.postReq("/api/v2/torrents/recheck", &data)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to reannounce. or set to []string{"all"}, to reannounce all torrents.

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#reannounce-torrents
*/
func (c *Client) ReannounceTorrents(hashes []string) (err error) {
	data := []byte(fmt.Sprintf("hashes=%s", strings.Join(hashes, "|")))
	_, err = c.postReq("/api/v2/torrents/reannounce", &data)
	return
}

/*
This method can add torrents from a local file or from URLs. http://, https://, magnet: and bc://bt/ links are supported.

# Params
  - "formData" a map of key value pairs, where the key is the name of the form field and the value is the value of the form field

# Example

	torrent := qbittorrent.NewTorrent().
		AddUrl("URL: http://, https://, magnet: and bc://bt/ links are supported").
		AddUrl("Additional URL").
		AddFromFile("pathToTorrentFile.torrent").
		AddFromFile("AdditionalPathToTorrentFile.torrent").
		Category("Movies").
		AutoTMM(true)

	err = client.AddNewTorrent(torrent.Data)
	if err != nil {
		panic(err)
	}

# Http Error Codes
  - 415 Torrent file is not valid
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#add-new-torrent
*/
func (c *Client) AddNewTorrent(formData map[string]string) (err error) {
	// add torrent from files and urls
	if files, ok := formData["torrents"]; ok {
		filePaths := strings.Split(files, "\n")

		var buffer bytes.Buffer
		writer := multipart.NewWriter(&buffer)

		for _, filePath := range filePaths {
			c.addFileToWriter(writer, filePath)
		}

		for key, val := range formData {
			if key != "torrents" {
				writer.WriteField(key, val)
			}
		}

		err = writer.Close()
		if err != nil {
			return
		}

		_, err = c.postMultipart("/api/v2/torrents/add", buffer, writer.FormDataContentType())

		return
	}

	// add torrent from url only
	form := url.Values{}
	for k, v := range formData {
		form.Add(k, v)
	}
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/add", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hash" The hash of the torrent you want to add trackers to.
  - "urls" The URLs of the trackers you want to add.

# Http Error Codes
  - 404 Not Found, if the torrent hash is invalid
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#add-trackers-to-torrent
*/
func (c *Client) AddTrackersToTorrent(hash string, trackers []string) (err error) {
	form := url.Values{}
	form.Add("hash", hash)
	form.Add("urls", strings.Join(trackers, "\n"))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/addTrackers", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hash" The hash of the torrent you want to edit trackers for.
  - "origUrl" The tracker URL you want to edit.
  - "newUrl" The new tracker URL  to replace the "origUrl"

# Http Error Codes
  - 400 newUrl is not a valid URL
  - 404 Torrent hash was not found
  - 409 newUrl already exists for the torrent
  - 409 origUrl was not found
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#edit-trackers
*/
func (c *Client) EditTrackers(hash, origUrl, newUrl string) (err error) {
	form := url.Values{}
	form.Add("hash", hash)
	form.Add("origUrl", origUrl)
	form.Add("newUrl", newUrl)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/editTracker", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

#Params
  - "hash" The hash of the torrent you want to remove trackers from.
  - "urls" The URLs of the trackers you want to remove.

# Http Error Codes
  - 404 Torrent hash was not found
  - 409 All urls were not found
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#remove-trackers
*/
func (c *Client) RemoveTrackers(hash string, urls []string) (err error) {
	form := url.Values{}
	form.Add("hash", hash)
	form.Add("urls", strings.Join(urls, "|"))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/removeTrackers", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to add peers to
  - "peers" The peers to add, Each peer is a colon-separated host:port

# Http Error Codes
  - 400 None of the supplied peers are valid
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#add-peers
*/
func (c *Client) AddPeers(hashes, peers []string) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	form.Add("peers", strings.Join(peers, "|"))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/addPeers", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to increase the priority of. or set to []string{"all"}, to increase the priority of all torrents.

# Http Error Codes
  - 409 Torrent queueing is not enabled
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#increase-torrent-priority
*/
func (c *Client) IncreaseTorrentPriority(hashes []string) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/increasePrio", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to decrease the priority of. or set to []string{"all"}, to decrease the priority of all torrents.

# Http Error Codes
  - 409 Torrent queueing is not enabled
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#decrease-torrent-priority
*/
func (c *Client) DecreaseTorrentPriority(hashes []string) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/decreasePrio", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to set to the maximum priority of. or set to []string{"all"}, to set all torrents to the maximum priority.

# Http Error Codes
  - 409 Torrent queueing is not enabled
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#maximal-torrent-priority
*/
func (c *Client) MaximalTorrentPriority(hashes []string) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/topPrio", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to set to the minimum  priority of. or set to []string{"all"}, to set all torrents to the minimum  priority.

# Http Error Codes
  - 409 Torrent queueing is not enabled
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#minimal-torrent-priority
*/
func (c *Client) MinimalTorrentPriority(hashes []string) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/bottomPrio", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hash" The hash of the torrent you want to set the priority of
  - "ids" values correspond to file position inside the array returned by GetTorrentContents e.g. id=0 for first file, id=1 for second file, etc.
  - "priority" The priority you want to set

# Http Error Codes
  - 400 Priority is invalid
  - 400 At least one file id is not a valid integer
  - 404 Torrent hash was not found
  - 409 Torrent metadata hasn't downloaded yet
  - 409 At least one file id was not found
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-file-priority
*/
func (c *Client) SetFilePriority(hash string, ids []int, priority FilePriority) (err error) {
	strIDs := make([]string, len(ids))
	for i, id := range ids {
		strIDs[i] = strconv.Itoa(id)
	}

	form := url.Values{}
	form.Add("hash", hash)
	form.Add("id", strings.Join(strIDs, "|"))
	form.Add("priority", strconv.Itoa(int(priority)))
	formDataBytes := []byte(form.Encode())

	_, err = c.postReq("/api/v2/torrents/filePrio", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to get the download limit of. or set to []string{"all"}, to get the download limit of all torrents.

Returns:
  - map[string]int where key is the torrent hash and value is the download speed limit in bytes per second; this value will be zero if no limit is applied.

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-torrent-download-limit
*/
func (c *Client) GetTorrentDownloadLimit(hashes []string) (results map[string]int, err error) {
	params := url.Values{}
	params.Add("hashes", strings.Join(hashes, "|"))

	body, err := c.getReq("/api/v2/torrents/downloadLimit", &params)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &results)

	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to set the download limit of
  - "limit" The download speed limit in bytes per second

# Http Error Codes:
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-torrent-download-limit
*/
func (c *Client) SetTorrentDownloadLimit(hashes []string, limit int) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	form.Add("limit", strconv.Itoa(limit))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/setDownloadLimit", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params:
  - "hashes" The hashes of the torrents for which you want to set the share limits. or set to []string{"all"}, to set all torrents share limits.
  - "ratioLimit" The maximum seeding ratio for the torrent. -2 means the global limit should be used, -1 means no limit.
  - "seedingTimeLimit" The maximum seeding time (minutes) for the torrent. -2 means the global limit should be used, -1 means no limit.
  - "inactiveSeedingTimeLimit" The maximum amount of time (minutes) the torrent is allowed to seed while being inactive. -2 means the global limit should be used, -1 means no limit.

# Http Error Codes:
  - 400 Bad Request
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-torrent-share-limit
*/
func (c *Client) SetTorrentShareLimit(hashes []string, ratioLimit float64, seedingTimeLimit, inactiveSeedingTimeLimit int) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	form.Add("ratioLimit", strconv.FormatFloat(ratioLimit, 'f', -1, 64))
	form.Add("seedingTimeLimit", strconv.Itoa(seedingTimeLimit))
	form.Add("inactiveSeedingTimeLimit", strconv.Itoa(inactiveSeedingTimeLimit))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/setShareLimits", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to get the upload limit of. or set to []string{"all"}, to get the upload limit of all torrents.

# Returns
  - map[string]int where key is the torrent hash and value is the upload speed limit in bytes per second; this value will be zero if no limit is applied.

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-torrent-upload-limit
*/
func (c *Client) GetTorrentUploadLimit(hashes []string) (results map[string]int, err error) {
	params := url.Values{}
	params.Add("hashes", strings.Join(hashes, "|"))

	body, err := c.getReq("/api/v2/torrents/uploadLimit", &params)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &results)

	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to set the upload limit of. or set to []string{"all"}, to set the upload limit of all torrents.
  - "limit" The upload speed limit in bytes per second

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-torrent-upload-limit
*/
func (c *Client) SetTorrentUploadLimit(hashes []string, limit int) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	form.Add("limit", strconv.Itoa(limit))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/setUploadLimit", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to set the location of. or set to []string{"all"}, to set the location of all torrents.
  - "location" The new location of the torrent

# Http Error Codes
  - 400 Save path is empty
  - 403 User does not have write access to directory
  - 409 Unable to create save path directory
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-torrent-location
*/
func (c *Client) SetTorrentLocation(hashes []string, location string) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	form.Add("location", location)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/setLocation", &formDataBytes)
	return
}

/*
Changes the name of the torrent (Not the folder name)

Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hash" The hash of the torrent for which you want to set the name
  - "name" The new name of the torrent

# Http Error Codes
  - 404 Torrent hash is invalid
  - 409 Torrent name is empty
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-torrent-name
*/
func (c *Client) SetTorrentName(hash, name string) (err error) {
	form := url.Values{}
	form.Add("hash", hash)
	form.Add("name", name)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/rename", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents you want to set the category of. or set to []string{"all"}, to set the category of all torrents.
  - "category" The new category of the torrent

# Http Error Codes
  - 409 Category name does not exist
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-torrent-category
*/
func (c *Client) SetTorrentCategory(hashes []string, category string) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	form.Add("category", category)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/setCategory", &formDataBytes)
	return
}

/*
Returns all categories in the client

	map["Movies":{"Movies" "/media/Plex/Movies"} "TV Shows":{"TV Shows" "/media/Plex/TV Shows"}]

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-all-categories
*/
func (c *Client) GetAllCategories() (results map[string]Category, err error) {
	body, err := c.getReq("/api/v2/torrents/categories", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &results)
	return
}

/*
Adds a new category to the client

# Params
  - "name" The name of the new category
  - "savePath" The path of the new category

# Http Error Codes
  - 400 Category name is empty
  - 409 Category name is invalid
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#add-new-category
*/
func (c *Client) AddNewCategory(name, savePath string) (err error) {
	form := url.Values{}
	form.Add("name", name)
	form.Add("savePath", savePath)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/createCategory", &formDataBytes)
	return
}

/*
Edits an existing category in the client

# Params
  - "name" The name of the category to edit
  - "savePath" The new path of the category

# Http Error Codes
  - 400 Category name is empty
  - 409 Category editing failed
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#edit-category
*/
func (c *Client) EditCategory(name, savePath string) (err error) {
	form := url.Values{}
	form.Add("name", name)
	form.Add("savePath", savePath)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/editCategory", &formDataBytes)
	return
}

/*
Removes categories from the client

# Params
  - "categories" The names of the categories to remove

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#remove-categories
*/
func (c *Client) RemoveCategories(categories []string) (err error) {
	form := url.Values{}
	form.Add("categories", strings.Join(categories, "\n"))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/removeCategories", &formDataBytes)
	return
}

/*
Adds tags to torrents

Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents for which you want to add tags. or set to []string{"all"}, to add tags to all torrents.
  - "tags" The tags to add

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#add-torrent-tags
*/
func (c *Client) AddTorrentTags(hashes, tags []string) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	form.Add("tags", strings.Join(tags, ","))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/addTags", &formDataBytes)
	return
}

/*
Removes tags from torrents

Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents for which you want to remove tags. or set to []string{"all"}, to remove tags from all torrents.
  - "tags" The tags to remove

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#remove-torrent-tags
*/
func (c *Client) RemoveTorrentTags(hashes, tags []string) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	form.Add("tags", strings.Join(tags, ","))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/removeTags", &formDataBytes)
	return
}

/*
Returns all tags

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-all-tags
*/
func (c *Client) GetAllTags() (results []string, err error) {
	body, err := c.getReq("/api/v2/torrents/tags", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &results)
	return
}

/*
Creates new tags in the client

# Params
  - "tags" The tags to create

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#create-tags
*/
func (c *Client) CreateTags(tags []string) (err error) {
	form := url.Values{}
	form.Add("tags", strings.Join(tags, ","))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/createTags", &formDataBytes)
	return
}

/*
Deletes tags from the client

# Params
  - "tags" The tags to delete

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#delete-tags
*/
func (c *Client) DeleteTags(tags []string) (err error) {
	form := url.Values{}
	form.Add("tags", strings.Join(tags, ","))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/deleteTags", &formDataBytes)
	return
}

/*
Enable/disable auto-management for a list of torrents

Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents for which you want to set the auto-management. or set to []string{"all"}, to set auto-management for all torrents.
  - "enable" Enable or disable auto-management

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-automatic-torrent-management
*/
func (c *Client) SetAutomaticTorrentManagement(hashes []string, enable bool) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	form.Add("enable", strconv.FormatBool(enable))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/setAutoManagement", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents for which you want to toggle sequential download. or set to []string{"all"}, to toggle sequential download for all torrents.

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#toggle-sequential-download
*/
func (c *Client) ToggleSequentialDownload(hashes []string) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/toggleSequentialDownload", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents for which you want to toggle first last piece priority. or set to []string{"all"}, to toggle first last piece priority for all torrents.

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-firstlast-piece-priority
*/
func (c *Client) ToggleFirstLastPiecePriority(hashes []string) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/toggleFirstLastPiecePrio", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents for which you want to set the force start. or set to []string{"all"}, to set force start for all torrents.
  - "forceStart" Enable or disable force start

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-force-start
*/
func (c *Client) SetForceStart(hashes []string, enable bool) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	form.Add("value", strconv.FormatBool(enable))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/setForceStart", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hashes" The hashes of the torrents for which you want to set the super seeding. or set to []string{"all"}, to set super seeding for all torrents.
  - "enable" Enable or disable super seeding

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-super-seeding
*/
func (c *Client) SetSuperSeeding(hashes []string, enable bool) (err error) {
	form := url.Values{}
	form.Add("hashes", strings.Join(hashes, "|"))
	form.Add("value", strconv.FormatBool(enable))
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/setSuperSeeding", &formDataBytes)
	return
}

/*
Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hash" The hash of the torrent
  - "oldPath" (relative to the torrent) The old path
  - "newPath" (relative to the torrent) The new path

# Http Error Codes
  - 400 Missing newPath parameter
  - 409 Invalid newPath or oldPath, or newPath already in use
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#rename-file
*/
func (c *Client) RenameFile(hash, oldPath, newPath string) (err error) {
	form := url.Values{}
	form.Add("hash", hash)
	form.Add("oldPath", oldPath)
	form.Add("newPath", newPath)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/renameFile", &formDataBytes)
	return
}

/*
Renames the folder of a torrent (where its stored)

Requires knowing the torrent hash. You can get it from GetTorrentList

# Params
  - "hash" The hash of the torrent
  - "oldPath" (relative to the torrent) The old path
  - "newPath" (relative to the torrent) The new path

# Http Error Codes
  - 400 Missing newPath parameter
  - 409 Invalid newPath or oldPath, or newPath already in use
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#rename-folder
*/
func (c *Client) RenameFolder(hash, oldPath, newPath string) (err error) {
	form := url.Values{}
	form.Add("hash", hash)
	form.Add("oldPath", oldPath)
	form.Add("newPath", newPath)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/torrents/renameFolder", &formDataBytes)
	return
}
