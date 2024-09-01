package qbittorrent

import (
	"strconv"
	"strings"
)

type NewTorrentOptions struct {
	Data map[string]string
}

func NewTorrent() *NewTorrentOptions {
	return &NewTorrentOptions{
		Data: make(map[string]string),
	}
}

// Add a torrent from a URL: http://, https://, magnet: and bc://bt/ links are supported.
func (o *NewTorrentOptions) AddUrl(uri string) *NewTorrentOptions {
	if urls, ok := o.Data["urls"]; ok {
		o.Data["urls"] = urls + "\n" + uri
	} else {
		o.Data["urls"] = uri
	}

	return o
}

// Download folder
func (o *NewTorrentOptions) SavePath(path string) *NewTorrentOptions {
	o.Data["savepath"] = path
	return o
}

// Cookie sent to download the .torrent file
func (o *NewTorrentOptions) Cookie(v string) *NewTorrentOptions {
	o.Data["cookie"] = v
	return o
}

// Category for the torrent
func (o *NewTorrentOptions) Category(cat string) *NewTorrentOptions {
	o.Data["category"] = cat
	return o
}

// Tags for the torrent
func (o *NewTorrentOptions) Tags(v []string) *NewTorrentOptions {
	o.Data["tags"] = strings.Join(v, ",")
	return o
}

// Skip hash checking
func (o *NewTorrentOptions) SkipChecking(v bool) *NewTorrentOptions {
	o.Data["skip_checking"] = strconv.FormatBool(v)
	return o
}

// Add torrents in the paused state.
func (o *NewTorrentOptions) Paused(v bool) *NewTorrentOptions {
	o.Data["paused"] = strconv.FormatBool(v)
	return o
}

// Create the root folder.
func (o *NewTorrentOptions) RootFolder(v bool) *NewTorrentOptions {
	o.Data["root_folder"] = strconv.FormatBool(v)
	return o
}

// Rename torrent
func (o *NewTorrentOptions) Rename(v string) *NewTorrentOptions {
	o.Data["rename"] = v
	return o
}

// Set torrent upload speed limit. Unit in bytes/second
func (o *NewTorrentOptions) UpLimit(v int) *NewTorrentOptions {
	o.Data["upLimit"] = strconv.Itoa(v)
	return o
}

// Set torrent download speed limit. Unit in bytes/second
func (o *NewTorrentOptions) DlLimit(v int) *NewTorrentOptions {
	o.Data["dlLimit"] = strconv.Itoa(v)
	return o
}

// Set torrent share ratio limit
func (o *NewTorrentOptions) RatioLimit(v float64) *NewTorrentOptions {
	o.Data["ratioLimit"] = strconv.FormatFloat(v, 'f', -1, 64)
	return o
}

// Set torrent seeding time limit. Unit in minutes
func (o *NewTorrentOptions) SeedingTimeLimit(v int) *NewTorrentOptions {
	o.Data["seedingTimeLimit"] = strconv.Itoa(v)
	return o
}

// Whether Automatic Torrent Management should be used
func (o *NewTorrentOptions) AutoTMM(v bool) *NewTorrentOptions {
	o.Data["autoTMM"] = strconv.FormatBool(v)
	return o
}

// Enable sequential download.
func (o *NewTorrentOptions) SequentialDownload(v string) *NewTorrentOptions {
	o.Data["sequentialDownload"] = v
	return o
}

// Prioritize download first last piece
func (o *NewTorrentOptions) FirstLastPiecePrio(v string) *NewTorrentOptions {
	o.Data["firstLastPiecePrio"] = v
	return o
}

func (o *NewTorrentOptions) AddFromFile(filePath string) *NewTorrentOptions {
	if path, ok := o.Data["torrents"]; ok {
		o.Data["torrents"] = path + "\n" + filePath
	} else {
		o.Data["torrents"] = filePath
	}

	return o
}
