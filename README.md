# qbittorrent-webapi-go

Golang wrapper for qBittorrent Web API (for versions above v4.1+).

This wrapper is based on the methods described in [qBittorrent WebUI APIwiki](<https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)>)

## Installation

```bash
go get github.com/alabsi91/qbittorrent-webapi-go
```

## Quick Guide

```go
package main

import (
    "fmt"

    "github.com/alabsi91/qbittorrent-webapi-go"
)

func main() {
    // First create the client
    client := qbittorrent.NewClient("http://localhost:8080")

    // Login
    err := client.Login("admin", "adminadmin")
    if err != nil {
        panic(err)
    }

    // Get torrents
    options := &qbittorrent.GetTorrentListOptions{
        Filter: qbittorrent.FilterAll,
    }
    torrents, err := client.GetTorrentList(options)
    if err != nil {
        panic(err)
    }

    for _, torrent := range torrents {
        fmt.Printf("%+v\n", torrent.Name)
    }
}
```

### Using builders

The following methods has parameters builders.

- `SetApplicationPreferences`
- `AddNewTorrent`
- `SetRSSAutoDownloadingRule`

```go
func main() {
    client := qbittorrent.NewClient("http://localhost:8080")

    // login
    err := client.Login("admin", "adminadmin")
    if err != nil {
        panic(err)
    }

    // set preferences
    params := qbittorrent.NewPreferences().
        Locale("en").
        SavePath("/home/user/Downloads")

    err = client.SetApplicationPreferences(&params.Data)
    if err != nil {
        panic(err)
    }

    // Add torrent
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
}
```

## Methods

### Authentication

- `Login(username, password string) (err error)`
- `Logout() (err error)`

### Application

- `GetApplicationVersion() (version string, err error)`
- `GetAPIVersion() (version string, err error)`
- `GetBuildInfo() (info BuildInfo, err error)`
- `ShutdownApplication() (err error)`
- `GetApplicationPreferences() (resultsApplicationPreferences, err error)`
- `SetApplicationPreferences(prefs *map[string]interfac{}) (err error)`
- `GetDefaultSavePath() (path string, err error)`

### Log

- `GetLog(params *GetLogParams) (results []GetLogResponse, err error)`
- `GetPeerLog(lastKnownId int) (results []GetPeerLogResponse, err error)`

### Sync

- `GetSyncMainData(rid int) (results SyncMainDataResponse, err error)`
- `GetSyncTorrentPeersData(hash string, rid int) (results map[string]interface{}, err error)`

### Transfer Info

- `GetGlobalTransferInfo() (results TransferInfoResponse, err error)`
- `GetAlternativeSpeedLimitsState() (results AlternativeSpeedLimitsStatus, err error)`
- `ToggleAlternativeSpeedLimits() (err error)`
- `GetGlobalDownloadLimit() (results int, err error)`
- `SetGlobalDownloadLimit(limit int) (err error)`
- `GetGlobalUploadLimit() (results int, err error)`
- `SetGlobalUploadLimit(limit int) (err error)`
- `BanPeers(peers []string) (err error)`

### Torrent Management

- `GetTorrentList(opts *GetTorrentListOptions) (results []TorrentListResponse, err error)`
- `GetTorrentGenericProperties(hash string) (results TorrentGenericProperties, err error)`
- `GetTorrentTrackers(hash string) (results []TorrentTracker, err error)`
- `GetTorrentWebSeeds(hash string) (results []TorrentSeed, err error)`
- `GetTorrentContents(hash string, indexes ...int) (results []TorrentFile, err error)`
- `GetTorrentPiecesStates(hash string) (results []TorrentPiecesState, err error)`
- `GetTorrentPiecesHashes(hash string) (results []string, err error)`
- `PauseTorrents(hashes []string) (err error)`
- `ResumeTorrents(hashes []string) (err error)`
- `DeleteTorrents(hashes []string, deleteFiles bool) (err error)`
- `RecheckTorrents(hashes []string) (err error)`
- `ReannounceTorrents(hashes []string) (err error)`
- `AddNewTorrent(formData map[string]string) (err error)`
- `AddTrackersToTorrent(hash string, trackers []string) (err error)`
- `EditTrackers(hash, origUrl, newUrl string) (err error)`
- `RemoveTrackers(hash string, urls []string) (err error)`
- `AddPeers(hashes, peers []string)`
- `IncreaseTorrentPriority(hashes []string) (err error)`
- `DecreaseTorrentPriority(hashes []string) (err error)`
- `MaximalTorrentPriority(hashes []string) (err error)`
- `MinimalTorrentPriority(hashes []string) (err error)`
- `SetFilePriority(hash string, ids []string, priority FilePriority) (err error)`
- `GetTorrentDownloadLimit(hashes []string) (results map[string]int, err error)`
- `SetTorrentDownloadLimit(hashes []string, limit int) (err error)`
- `SetTorrentShareLimit(hashes []string, ratioLimit float64, seedingTimeLimit, inactiveSeedingTimeLimit int) (err error)`
- `GetTorrentUploadLimit(hashes []string) (results map[string]int, err error)`
- `SetTorrentUploadLimit(hashes []string, limit int) (err error)`
- `SetTorrentLocation(hashes []string, location string) (err error)`
- `SetTorrentName(hash, name string) (err error)`
- `SetTorrentCategory(hashes []string, category string) (err error)`
- `GetAllCategories() (results map[string]Category, err error)`
- `AddNewCategory(name, savePath string) (err error)`
- `EditCategory(name, savePath string) (err error)`
- `RemoveCategories(categories []string) (err error)`
- `AddTorrentTags(hashes, tags []string) (err error)`
- `RemoveTorrentTags(hashes, tags []string) (err error)`
- `GetAllTags() (results []string, err error)`
- `CreateTags(tags []string) (err error)`
- `DeleteTags(tags []string) (err error)`
- `SetAutomaticTorrentManagement(hashes []string, enable bool) (err error)`
- `ToggleSequentialDownload(hashes []string) (err error)`
- `ToggleFirstLastPiecePriority(hashes []string) (err error)`
- `SetForceStart(hashes []string, enable bool) (err error)`
- `SetSuperSeeding(hashes []string, enable bool) (err error)`
- `RenameFile(hash, oldPath, newPath string) (err error)`
- `RenameFolder(hash, oldPath, newPath string) (err error)`

### RSS

- `AddRSSFolder(path string) (err error)`
- `AddRSSFeed(feedUrl, path string) (err error)`
- `RemoveRSSItem(path string) (err error)`
- `MoveRSSItem(itemPath, destPath string) (err error)`
- `GetAllRSSItems(withData bool) (results map[string]interface{}, err error)`
- `MarkRSSAsRead(itemPath, articleId string) (err error)`
- `RefreshRSSItem(itemPath string) (err error)`
- `SetRSSAutoDownloadingRule(ruleName string, ruleDef map[string]interface{}) (err error)`
- `RenameRSSAutoDownloadingRule(ruleName, newRuleName string) (err error)`
- `RemoveRSSAutoDownloadingRule(ruleName string) (err error)`
- `GetAllRSSDownloadingRules() (results map[string]RSSDownloadingRule, err error)`
- `GetAllRSSArticlesMatchingRule(ruleName string) (results map[string][]string, err error)`

### Search

- `StartSearch(pattern string, plugins []string, category []string) (results int, err error)`
- `StopSearch(id int) (err error)`
- `GetSearchStatus(id *int) (results []SearchStatusResponse, err error)`
- `GetSearchResults(id int, limit, offset *int) (results SearchResultsResponse, err error)`
- `DeleteSearch(id int) (err error)`
- `GetSearchPlugins() (results []SearchPluginsResponse, err error)`
- `InstallSearchPlugin(sources []string) (err error)`
- `UninstallSearchPlugin(names []string) (err error)`
- `EnableSearchPlugin(names []string, enable bool) (err error)`
- `UpdateSearchPlugins() (err error)`
