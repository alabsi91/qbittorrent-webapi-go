package qbittorrent

type Filters = string

const (
	FilterAll                Filters = "all"
	FilterDownloading        Filters = "downloading"
	FilterSeeding            Filters = "seeding"
	FilterCompleted          Filters = "completed"
	FilterPaused             Filters = "paused"
	FilterActive             Filters = "active"
	FilterInactive           Filters = "inactive"
	FilterResumed            Filters = "resumed"
	FilterStalled            Filters = "stalled"
	FilterStalledUploading   Filters = "stalled_uploading"
	FilterStalledDownloading Filters = "stalled_downloading"
	FilterErrored            Filters = "errored"
)

type TorrentState = string

const (
	TorrentStateError              TorrentState = "error"              // Some error occurred, applies to paused torrents
	TorrentStateMissingFiles       TorrentState = "missingFiles"       // Torrent data files is missing
	TorrentStateUploading          TorrentState = "uploading"          // Torrent is being seeded and data is being transferred
	TorrentStatePausedUP           TorrentState = "pausedUP"           // Torrent is paused and has finished downloading
	TorrentStateQueuedUP           TorrentState = "queuedUP"           // Queuing is enabled and torrent is queued for upload
	TorrentStateStalledUP          TorrentState = "stalledUP"          // Torrent is being seeded, but no connection were made
	TorrentStateCheckingUP         TorrentState = "checkingUP"         // Torrent has finished downloading and is being checked
	TorrentStateForcedUP           TorrentState = "forcedUP"           // Torrent is forced to uploading and ignore queue limit
	TorrentStateAllocating         TorrentState = "allocating"         // Torrent is allocating disk space for download
	TorrentStateDownloading        TorrentState = "downloading"        // Torrent is being downloaded and data is being transferred
	TorrentStateMetaDL             TorrentState = "metaDL"             // Torrent has just started downloading and is fetching metadata
	TorrentStatePausedDL           TorrentState = "pausedDL"           // Torrent is paused and has NOT finished downloading
	TorrentStateQueuedDL           TorrentState = "queuedDL"           // Queuing is enabled and torrent is queued for download
	TorrentStateStalledDL          TorrentState = "stalledDL"          // Torrent is being downloaded, but no connection were made
	TorrentStateCheckingDL         TorrentState = "checkingDL"         // Same as checkingUP, but torrent has NOT finished downloading
	TorrentStateForcedDL           TorrentState = "forcedDL"           // Torrent is forced to downloading to ignore queue limit
	TorrentStateCheckingResumeData TorrentState = "checkingResumeData" // Checking resume data on qBt startup
	TorrentStateMoving             TorrentState = "moving"             // Torrent is moving to another location
	TorrentStateUnknown            TorrentState = "unknown"            // Unknown status
)

type SchedulerDays int

const (
	SchedulerEveryDay       SchedulerDays = 0
	SchedulerEveryWeekday   SchedulerDays = 1
	SchedulerEveryWeekend   SchedulerDays = 2
	SchedulerEveryMonday    SchedulerDays = 3
	SchedulerEveryTuesday   SchedulerDays = 4
	SchedulerEveryWednesday SchedulerDays = 5
	SchedulerEveryThursday  SchedulerDays = 6
	SchedulerEveryFriday    SchedulerDays = 7
	SchedulerEverySaturday  SchedulerDays = 8
	SchedulerEverySunday    SchedulerDays = 9
)

type Encryption int

const (
	EncryptionPrefer   Encryption = 0
	EncryptionForceOn  Encryption = 1
	EncryptionForceOff Encryption = 2
)

type ProxyType int

const (
	ProxyIsDisabled                  ProxyType = -1
	ProxyHTTPWithoutAuthentication   ProxyType = 1
	ProxySOCKS5WithoutAuthentication ProxyType = 2
	ProxyHTTPWithAuthentication      ProxyType = 3
	ProxySOCKS5WithAuthentication    ProxyType = 4
	ProxySOCKS4WithoutAuthentication ProxyType = 5
)

type DyndnsService int

const (
	DyndnsServiceUseDyDNS DyndnsService = 0
	DyndnsServiceUseNOIP  DyndnsService = 1
)

type MaxRatioAct int

const (
	MaxRatioActPause  MaxRatioAct = 0
	MaxRatioActRemove MaxRatioAct = 1
)

type BittorrentProtocol int

const (
	BittorrentProtocolBoth BittorrentProtocol = 0
	BittorrentProtocolTCP  BittorrentProtocol = 1
	BittorrentProtocolUTP  BittorrentProtocol = 2
)

type UploadChokingAlgorithm int

const (
	UploadChokingAlgorithmRoundRobin    UploadChokingAlgorithm = 0
	UploadChokingAlgorithmFastestUpload UploadChokingAlgorithm = 1
	UploadChokingAlgorithmAntiLeech     UploadChokingAlgorithm = 2
)

type UploadSlotsBehavior int

const (
	UploadSlotsBehaviorFixedSlots      UploadSlotsBehavior = 0
	UploadSlotsBehaviorUploadRateBased UploadSlotsBehavior = 1
)

type UtpTcpMixedMode int

const (
	UtpTcpMixedModePreferTCP        UtpTcpMixedMode = 0
	UtpTcpMixedModePeerProportional UtpTcpMixedMode = 1
)

type TorrentPiecesState int

const (
	TorrentPiecesNotDownloaded    TorrentPiecesState = 0 // Not downloaded yet
	TorrentPiecesStateDownloading TorrentPiecesState = 1 // Downloading
	TorrentPiecesStateDone        TorrentPiecesState = 2 // Already downloaded
)

type LogMessageType int

const (
	LogNormal   LogMessageType = 1
	LogInfo     LogMessageType = 2
	LogWarn     LogMessageType = 4
	LogCritical LogMessageType = 8
)

type ConnectionStatus string

const (
	Connected    ConnectionStatus = "connected"
	FireWalled   ConnectionStatus = "firewalled"
	Disconnected ConnectionStatus = "disconnected"
)

type AlternativeSpeedLimitsStatus int

const (
	AlternativeSpeedLimitsDisabled AlternativeSpeedLimitsStatus = 0
	AlternativeSpeedLimitsEnabled  AlternativeSpeedLimitsStatus = 1
)

type TrackerStatus int

const (
	TrackerDisabled     TrackerStatus = 0 // Tracker is disabled (used for DHT, PeX, and LSD)
	TrackerNotContacted TrackerStatus = 1 // Tracker has not been contacted yet
	TrackerWorking      TrackerStatus = 2 // Tracker has been contacted and is working
	TrackerUpdating     TrackerStatus = 3 // Tracker is updating
	TrackerNotWorking   TrackerStatus = 4 // Tracker has been contacted, but it is not working (or doesn't send proper replies)
)

type FilePriority int

const (
	FilePriorityDoNotDownload FilePriority = 0 // Do not download
	FilePriorityNormal        FilePriority = 1 // Normal priority
	FilePriorityHigh          FilePriority = 6 // High priority
	FilePriorityMaximal       FilePriority = 7 // Maximal priority
)

type SearchStatus = string

const (
	SearchStatusRunning SearchStatus = "Running"
	SearchStatusStopped SearchStatus = "Stopped"
)

// -------------------------------------------------------------------------

type BuildInfo struct {
	QT         string `json:"qt"`         // QT version
	LibTorrent string `json:"libtorrent"` // libtorrent version
	Boost      string `json:"boost"`      // Boost version
	OpenSSL    string `json:"openssl"`    // OpenSSL version
	Bitness    int    `json:"bitness"`    // Application bitness (e.g. 64-bit)
}

type ApplicationPreferences struct {
	Locale                             string                 `json:"locale"`                                 // Currently selected language (e.g. en_GB for English)
	CreateSubfolderEnabled             bool                   `json:"create_subfolder_enabled"`               // True if a subfolder should be created when adding a torrent
	StartPausedEnabled                 bool                   `json:"start_paused_enabled"`                   // True if torrents should be added in a Paused state
	AutoDeleteMode                     int                    `json:"auto_delete_mode"`                       // TODO
	PreallocateAll                     bool                   `json:"preallocate_all"`                        // True if disk space should be pre-allocated for all files
	IncompleteFilesExt                 bool                   `json:"incomplete_files_ext"`                   // True if ".!qB" should be appended to incomplete files
	AutoTmmEnabled                     bool                   `json:"auto_tmm_enabled"`                       // True if Automatic Torrent Management is enabled by default
	TorrentChangedTmmEnabled           bool                   `json:"torrent_changed_tmm_enabled"`            // True if torrent should be relocated when its Category changes
	SavePathChangedTmmEnabled          bool                   `json:"save_path_changed_tmm_enabled"`          // True if torrent should be relocated when the default save path changes
	CategoryChangedTmmEnabled          bool                   `json:"category_changed_tmm_enabled"`           // True if torrent should be relocated when its Category's save path changes
	SavePath                           string                 `json:"save_path"`                              // Default save path for torrents, separated by slashes
	TempPathEnabled                    bool                   `json:"temp_path_enabled"`                      // True if folder for incomplete torrents is enabled
	TempPath                           string                 `json:"temp_path"`                              // Path for incomplete torrents, separated by slashes
	ScanDirs                           map[string]interface{} `json:"scan_dirs"`                              // Property: directory to watch for torrent files, value: where torrents loaded from this directory should be downloaded to. Slashes are used as path separators; multiple key/value pairs can be specified Possible values of scan_dirs - `0` Download to the monitored folder - `1` Download to the default save path - "/path/to/download/to"	Download to this path
	ExportDir                          string                 `json:"export_dir"`                             // Path to directory to copy .torrent files to. Slashes are used as path separators
	ExportDirFin                       string                 `json:"export_dir_fin"`                         // Path to directory to copy .torrent files of completed downloads to. Slashes are used as path separators
	MailNotificationEnabled            bool                   `json:"mail_notification_enabled"`              // True if e-mail notification should be enabled
	MailNotificationSender             string                 `json:"mail_notification_sender"`               // e-mail where notifications should originate from
	MailNotificationEmail              string                 `json:"mail_notification_email"`                // e-mail to send notifications to
	MailNotificationSmtp               string                 `json:"mail_notification_smtp"`                 // smtp server for e-mail notifications
	MailNotificationSslEnabled         bool                   `json:"mail_notification_ssl_enabled"`          // True if smtp server requires SSL connection
	MailNotificationAuthEnabled        bool                   `json:"mail_notification_auth_enabled"`         // True if smtp server requires authentication
	MailNotificationUsername           string                 `json:"mail_notification_username"`             // Username for smtp authentication
	MailNotificationPassword           string                 `json:"mail_notification_password"`             // Password for smtp authentication
	AutorunEnabled                     bool                   `json:"autorun_enabled"`                        // True if external program should be run after torrent has finished downloading
	AutorunProgram                     string                 `json:"autorun_program"`                        // Program path/name/arguments to run if autorun_enabled is enabled; path is separated by slashes; you can use %f and %n arguments, which will be expanded by qBittorent as path_to_torrent_file and torrent_name (from the GUI; not the .torrent file name) respectively
	QueueingEnabled                    bool                   `json:"queueing_enabled"`                       // True if torrent queuing is enabled
	MaxActiveDownloads                 int                    `json:"max_active_downloads"`                   // Maximum number of active simultaneous downloads
	MaxActiveTorrents                  int                    `json:"max_active_torrents"`                    // Maximum number of active simultaneous downloads and uploads
	MaxActiveUploads                   int                    `json:"max_active_uploads"`                     // Maximum number of active simultaneous uploads
	DontCountSlowTorrents              bool                   `json:"dont_count_slow_torrents"`               // If true torrents w/o any activity (stalled ones) will not be counted towards max_active_* limits; see dont_count_slow_torrents for more information
	SlowTorrentDlRateThreshold         int                    `json:"slow_torrent_dl_rate_threshold"`         // Download rate in KiB/s for a torrent to be considered "slow"
	SlowTorrentUlRateThreshold         int                    `json:"slow_torrent_ul_rate_threshold"`         // Upload rate in KiB/s for a torrent to be considered "slow"
	SlowTorrentInactiveTimer           int                    `json:"slow_torrent_inactive_timer"`            // Seconds a torrent should be inactive before considered "slow"
	MaxRatioEnabled                    bool                   `json:"max_ratio_enabled"`                      // True if share ratio limit is enabled
	MaxRatio                           float64                `json:"max_ratio"`                              // Get the global share ratio limit
	MaxRatioAct                        DyndnsService          `json:"max_ratio_act"`                          // Action performed when a torrent reaches the maximum share ratio. See list of possible values here below.
	ListenPort                         int                    `json:"listen_port"`                            // Port for incoming connections
	Upnp                               bool                   `json:"upnp"`                                   // True if UPnP/NAT-PMP is enabled
	RandomPort                         bool                   `json:"random_port"`                            // True if the port is randomly selected
	DlLimit                            int                    `json:"dl_limit"`                               // Global download speed limit in KiB/s; -1 means no limit is applied
	UpLimit                            int                    `json:"up_limit"`                               // Global upload speed limit in KiB/s; -1 means no limit is applied
	MaxConnec                          int                    `json:"max_connec"`                             // Maximum global number of simultaneous connections
	MaxConnecPerTorrent                int                    `json:"max_connec_per_torrent"`                 // Maximum number of simultaneous connections per torrent
	MaxUploads                         int                    `json:"max_uploads"`                            // Maximum number of upload slots
	MaxUploadsPerTorrent               int                    `json:"max_uploads_per_torrent"`                // Maximum number of upload slots per torrent
	StopTrackerTimeout                 int                    `json:"stop_tracker_timeout"`                   // Timeout in seconds for a stopped announce request to trackers
	EnablePieceExtentAffinity          bool                   `json:"enable_piece_extent_affinity"`           // True if the advanced libtorrent option piece_extent_affinity is enabled
	BittorrentProtocol                 BittorrentProtocol     `json:"bittorrent_protocol"`                    // Bittorrent Protocol to use
	LimitUtpRate                       bool                   `json:"limit_utp_rate"`                         // True if [du]l_limit should be applied to uTP connections; this option is only available in qBittorent built against libtorrent version 0.16.X and higher
	LimitTcpOverhead                   bool                   `json:"limit_tcp_overhead"`                     // True if [du]l_limit should be applied to estimated TCP overhead (service data: e.g. packet headers)
	LimitLanPeers                      bool                   `json:"limit_lan_peers"`                        // True if [du]l_limit should be applied to peers on the LAN
	AltDlLimit                         int                    `json:"alt_dl_limit"`                           // Alternative global download speed limit in KiB/s
	AltUpLimit                         int                    `json:"alt_up_limit"`                           // Alternative global upload speed limit in KiB/s
	SchedulerEnabled                   bool                   `json:"scheduler_enabled"`                      // True if alternative limits should be applied according to schedule
	ScheduleFromHour                   int                    `json:"schedule_from_hour"`                     // Scheduler starting hour
	ScheduleFromMin                    int                    `json:"schedule_from_min"`                      // Scheduler starting minute
	ScheduleToHour                     int                    `json:"schedule_to_hour"`                       // Scheduler ending hour
	ScheduleToMin                      int                    `json:"schedule_to_min"`                        // Scheduler ending minute
	SchedulerDays                      SchedulerDays          `json:"scheduler_days"`                         // Scheduler days. See possible values here below
	Dht                                bool                   `json:"dht"`                                    // True if DHT is enabled
	Pex                                bool                   `json:"pex"`                                    // True if PeX is enabled
	Lsd                                bool                   `json:"lsd"`                                    // True if LSD is enabled
	Encryption                         Encryption             `json:"encryption"`                             // See list of possible values here below
	AnonymousMode                      bool                   `json:"anonymous_mode"`                         // If true anonymous mode will be enabled; read more here; this option is only available in qBittorent built against libtorrent version 0.16.X and higher
	ProxyType                          int                    `json:"proxy_type"`                             // See list of possible values here below
	ProxyIp                            string                 `json:"proxy_ip"`                               // Proxy IP address or domain name
	ProxyPort                          int                    `json:"proxy_port"`                             // Proxy port
	ProxyPeerConnections               bool                   `json:"proxy_peer_connections"`                 // True if peer and web seed connections should be proxified; this option will have any effect only in qBittorent built against libtorrent version 0.16.X and higher
	ProxyAuthEnabled                   bool                   `json:"proxy_auth_enabled"`                     // True proxy requires authentication; doesn't apply to SOCKS4 proxies
	ProxyUsername                      string                 `json:"proxy_username"`                         // Username for proxy authentication
	ProxyPassword                      string                 `json:"proxy_password"`                         // Password for proxy authentication
	ProxyTorrentsOnly                  bool                   `json:"proxy_torrents_only"`                    // True if proxy is only used for torrents
	IpFilterEnabled                    bool                   `json:"ip_filter_enabled"`                      // True if external IP filter should be enabled
	IpFilterPath                       string                 `json:"ip_filter_path"`                         // Path to IP filter file (.dat, .p2p, .p2b files are supported); path is separated by slashes
	IpFilterTrackers                   bool                   `json:"ip_filter_trackers"`                     // True if IP filters are applied to trackers
	WebUiDomainList                    string                 `json:"web_ui_domain_list"`                     // Semicolon-separated list of domains to accept when performing Host header validation
	WebUiAddress                       string                 `json:"web_ui_address"`                         // IP address to use for the WebUI
	WebUiPort                          int                    `json:"web_ui_port"`                            // WebUI port
	WebUiUpnp                          bool                   `json:"web_ui_upnp"`                            // True if UPnP is used for the WebUI port
	WebUiUsername                      string                 `json:"web_ui_username"`                        // WebUI username
	WebUiPassword                      string                 `json:"web_ui_password"`                        // For API ≥ v2.3.0: Plaintext WebUI password, not readable, write-only. For API < v2.3.0: MD5 hash of WebUI password, hash is generated from the following string: username:Web UI Access:plain_text_web_ui_password
	WebUiCsrfProtectionEnabled         bool                   `json:"web_ui_csrf_protection_enabled"`         // True if WebUI CSRF protection is enabled
	WebUiClickjackingProtectionEnabled bool                   `json:"web_ui_clickjacking_protection_enabled"` // True if WebUI clickjacking protection is enabled
	WebUiSecureCookieEnabled           bool                   `json:"web_ui_secure_cookie_enabled"`           // True if WebUI cookie Secure flag is enabled
	WebUiMaxAuthFailCount              int                    `json:"web_ui_max_auth_fail_count"`             // Maximum number of authentication failures before WebUI access ban
	WebUiBanDuration                   int                    `json:"web_ui_ban_duration"`                    // WebUI access ban duration in seconds
	WebUiSessionTimeout                int                    `json:"web_ui_session_timeout"`                 // Seconds until WebUI is automatically signed off
	WebUiHostHeaderValidationEnabled   bool                   `json:"web_ui_host_header_validation_enabled"`  // True if WebUI host header validation is enabled
	BypassLocalAuth                    bool                   `json:"bypass_local_auth"`                      // True if authentication challenge for loopback address (127.0.0.1) should be disabled
	BypassAuthSubnetWhitelistEnabled   bool                   `json:"bypass_auth_subnet_whitelist_enabled"`   // True if webui authentication should be bypassed for clients whose ip resides within (at least) one of the subnets on the whitelist
	BypassAuthSubnetWhitelist          string                 `json:"bypass_auth_subnet_whitelist"`           // (White)list of ipv4/ipv6 subnets for which webui authentication should be bypassed; list entries are separated by commas
	AlternativeWebuiEnabled            bool                   `json:"alternative_webui_enabled"`              // True if an alternative WebUI should be used
	AlternativeWebuiPath               string                 `json:"alternative_webui_path"`                 // File path to the alternative WebUI
	UseHttps                           bool                   `json:"use_https"`                              // True if WebUI HTTPS access is enabled
	SslKey                             string                 `json:"ssl_key"`                                // For API < v2.0.1: SSL keyfile contents (this is a not a path)
	SslCert                            string                 `json:"ssl_cert"`                               // For API < v2.0.1: SSL certificate contents (this is a not a path)
	WebUiHttpsKeyPath                  string                 `json:"web_ui_https_key_path"`                  // For API ≥ v2.0.1: Path to SSL keyfile
	WebUiHttpsCertPath                 string                 `json:"web_ui_https_cert_path"`                 // For API ≥ v2.0.1: Path to SSL certificate
	DyndnsEnabled                      bool                   `json:"dyndns_enabled"`                         // True if server DNS should be updated dynamically
	DyndnsService                      DyndnsService          `json:"dyndns_service"`                         // See list of possible values here below
	DyndnsUsername                     string                 `json:"dyndns_username"`                        // Username for DDNS service
	DyndnsPassword                     string                 `json:"dyndns_password"`                        // Password for DDNS service
	DyndnsDomain                       string                 `json:"dyndns_domain"`                          // Your DDNS domain name
	RssRefreshInterval                 int                    `json:"rss_refresh_interval"`                   // RSS refresh interval
	RssMaxArticlesPerFeed              int                    `json:"rss_max_articles_per_feed"`              // Max stored articles per RSS feed
	RssProcessingEnabled               bool                   `json:"rss_processing_enabled"`                 // Enable processing of RSS feeds
	RssAutoDownloadingEnabled          bool                   `json:"rss_auto_downloading_enabled"`           // Enable auto-downloading of torrents from the RSS feeds
	RssDownloadRepackProperEpisodes    bool                   `json:"rss_download_repack_proper_episodes"`    // For API ≥ v2.5.1: Enable downloading of repack/proper Episodes
	RssSmartEpisodeFilters             string                 `json:"rss_smart_episode_filters"`              // For API ≥ v2.5.1: List of RSS Smart Episode Filters
	AddTrackersEnabled                 bool                   `json:"add_trackers_enabled"`                   // Enable automatic adding of trackers to new torrents
	AddTrackers                        string                 `json:"add_trackers"`                           // List of trackers to add to new torrent
	WebUiUseCustomHttpHeadersEnabled   bool                   `json:"web_ui_use_custom_http_headers_enabled"` // For API ≥ v2.5.1: Enable custom http headers
	WebUiCustomHttpHeaders             string                 `json:"web_ui_custom_http_headers"`             // For API ≥ v2.5.1: List of custom http headers
	MaxSeedingTimeEnabled              bool                   `json:"max_seeding_time_enabled"`               // True enables max seeding time
	MaxSeedingTime                     int                    `json:"max_seeding_time"`                       // Number of minutes to seed a torrent
	AnnounceIp                         string                 `json:"announce_ip"`                            // TODO
	AnnounceToAllTiers                 bool                   `json:"announce_to_all_tiers"`                  // True always announce to all tiers
	AnnounceToAllTrackers              bool                   `json:"announce_to_all_trackers"`               // True always announce to all trackers in a tier
	AsyncIoThreads                     int                    `json:"async_io_threads"`                       // Number of asynchronous I/O threads
	BannedIPs                          string                 `json:"banned_IPs"`                             // List of banned IPs
	CheckingMemoryUse                  int                    `json:"checking_memory_use"`                    // Outstanding memory when checking torrents in MiB
	CurrentInterfaceAddress            string                 `json:"current_interface_address"`              // IP Address to bind to. Empty String means All addresses
	CurrentNetworkInterface            string                 `json:"current_network_interface"`              // Network Interface used
	DiskCache                          int                    `json:"disk_cache"`                             // Disk cache used in MiB
	DiskCacheTtl                       int                    `json:"disk_cache_ttl"`                         // Disk cache expiry interval in seconds
	EmbeddedTrackerPort                int                    `json:"embedded_tracker_port"`                  // Port used for embedded tracker
	EnableCoalesceReadWrite            bool                   `json:"enable_coalesce_read_write"`             // True enables coalesce reads & writes
	EnableEmbeddedTracker              bool                   `json:"enable_embedded_tracker"`                // True enables embedded tracker
	EnableMultiConnectionsFromSameIP   bool                   `json:"enable_multi_connections_from_same_ip"`  // True allows multiple connections from the same IP address
	EnableOsCache                      bool                   `json:"enable_os_cache"`                        // True enables os cache
	EnableUploadSuggestions            bool                   `json:"enable_upload_suggestions"`              // True enables sending of upload piece suggestions
	FilePoolSize                       int                    `json:"file_pool_size"`                         // File pool size
	OutgoingPortsMax                   int                    `json:"outgoing_ports_max"`                     // Maximal outgoing port (0: Disabled)
	OutgoingPortsMin                   int                    `json:"outgoing_ports_min"`                     // Minimal outgoing port (0: Disabled)
	RecheckCompletedTorrents           bool                   `json:"recheck_completed_torrents"`             // True rechecks torrents on completion
	ResolvePeerCountries               bool                   `json:"resolve_peer_countries"`                 // True resolves peer countries
	SaveResumeDataInterval             int                    `json:"save_resume_data_interval"`              // Save resume data interval in min
	SendBufferLowWatermark             int                    `json:"send_buffer_low_watermark"`              // Send buffer low watermark in KiB
	SendBufferWatermark                int                    `json:"send_buffer_watermark"`                  // Send buffer watermark in KiB
	SendBufferWatermarkFactor          int                    `json:"send_buffer_watermark_factor"`           // Send buffer watermark factor in percent
	SocketBacklogSize                  int                    `json:"socket_backlog_size"`                    // Socket backlog size
	UploadChokingAlgorithm             UploadChokingAlgorithm `json:"upload_choking_algorithm"`               // Upload choking algorithm used
	UploadSlotsBehavior                UploadSlotsBehavior    `json:"upload_slots_behavior"`                  // Upload slots behavior used
	UpnpLeaseDuration                  int                    `json:"upnp_lease_duration"`                    // UPnP lease duration (0: Permanent lease)
	UtpTcpMixedMode                    UtpTcpMixedMode        `json:"utp_tcp_mixed_mode"`                     // μTP-TCP mixed mode algorithm
}

type GetTorrentListOptions struct {
	Filter   Filters  // Filter torrent list by state. Allowed state filters see [qbittorrent.Filters]
	Category *string  // Get torrents with the given category. - Pass an empty string for "without category" torrents. - Leave it or pass nil for "any category" torrents.
	Tag      *string  // Get torrents with the given tag. - Pass an empty string for "without tag" torrents. - Leave it or pass nil for "any tag" torrents.
	Sort     string   // Sort torrents by given key. They can be sorted using any field of the response's array as the sort key.
	Reverse  bool     // Enable reverse sorting. Defaults to false
	Limit    int      // Limit the number of torrents returned
	Offset   int      // Set offset (if less than 0, offset from end)
	Hashes   []string // Filter by hashes.
}

type DeleteOptions struct {
	Hashes      []string // The hashes of the torrents you want to delete. hashes can contain multiple hashes, to delete multiple torrents, or set to []{"all"}, to delete all torrents.
	DeleteFiles bool     // If set to true, the downloaded data will also be deleted, otherwise has no effect.
}

type TorrentListResponse struct {
	AddedOn           int          `json:"added_on"`           // Time (Unix Epoch) when the torrent was added to the client
	AmountLeft        int          `json:"amount_left"`        // Amount of data left to download (bytes)
	AutoTmm           bool         `json:"auto_tmm"`           // Whether this torrent is managed by Automatic Torrent Management
	Availability      float64      `json:"availability"`       // Percentage of file pieces currently available
	Category          string       `json:"category"`           // Category of the torrent
	Completed         int          `json:"completed"`          // Amount of transfer data completed (bytes)
	CompletionOn      int          `json:"completion_on"`      // Time (Unix Epoch) when the torrent completed
	ContentPath       string       `json:"content_path"`       // Absolute path of torrent content (root path for multifile torrents, absolute file path for singlefile torrents)
	DlLimit           float64      `json:"dl_limit"`           // Torrent download speed limit (bytes/s). -1 if unlimited.
	DlSpeed           float64      `json:"dlspeed"`            // Torrent download speed (bytes/s)
	Downloaded        float64      `json:"downloaded"`         // Amount of data downloaded
	DownloadedSession float64      `json:"downloaded_session"` // Amount of data downloaded this session
	ETA               int          `json:"eta"`                // Torrent ETA (seconds)
	FLPiecePrio       bool         `json:"f_l_piece_prio"`     // True if first last piece are prioritized
	ForceStart        bool         `json:"force_start"`        // True if force start is enabled for this torrent
	Hash              string       `json:"hash"`               // Torrent hash
	IsPrivate         bool         `json:"is_private"`         // True if torrent is from a private tracker (added in 5.0.0)
	LastActivity      int          `json:"last_activity"`      // Last time (Unix Epoch) when a chunk was downloaded/uploaded
	MagnetURI         string       `json:"magnet_uri"`         // Magnet URI corresponding to this torrent
	MaxRatio          float64      `json:"max_ratio"`          // Maximum share ratio until torrent is stopped from seeding/uploading
	MaxSeedingTime    int          `json:"max_seeding_time"`   // Maximum seeding time (seconds) until torrent is stopped from seeding
	Name              string       `json:"name"`               // Torrent name
	NumComplete       int          `json:"num_complete"`       // Number of seeds in the swarm
	NumIncomplete     int          `json:"num_incomplete"`     // Number of leechers in the swarm
	NumLeechs         int          `json:"num_leechs"`         // Number of leechers connected to
	NumSeeds          int          `json:"num_seeds"`          // Number of seeds connected to
	Priority          int          `json:"priority"`           // Torrent priority. Returns -1 if queuing is disabled or torrent is in seed mode
	Progress          float64      `json:"progress"`           // Torrent progress (percentage/100)
	Ratio             float64      `json:"ratio"`              // Torrent share ratio. Max ratio value: 9999.
	RatioLimit        float64      `json:"ratio_limit"`        // TODO (what is different from max_ratio?)
	SavePath          string       `json:"save_path"`          // Path where this torrent's data is stored
	SeedingTime       int          `json:"seeding_time"`       // Torrent elapsed time while complete (seconds)
	SeedingTimeLimit  int          `json:"seeding_time_limit"` // TODO (what is different from max_seeding_time?)
	SeenComplete      int          `json:"seen_complete"`      // Time (Unix Epoch) when this torrent was last seen complete
	SeqDL             bool         `json:"seq_dl"`             // True if sequential download is enabled
	Size              int          `json:"size"`               // Total size (bytes) of files selected for download
	State             TorrentState `json:"state"`              // Torrent state. See [TorrentState] possible values
	SuperSeeding      bool         `json:"super_seeding"`      // True if super seeding is enabled
	Tags              string       `json:"tags"`               // Comma-concatenated tag list of the torrent
	TimeActive        int          `json:"time_active"`        // Total active time (seconds)
	TotalSize         int          `json:"total_size"`         // Total size (bytes) of all file in this torrent (including unselected ones)
	Tracker           string       `json:"tracker"`            // The first tracker with working status. Returns empty string if no tracker is working.
	UpLimit           int          `json:"up_limit"`           // Torrent upload speed limit (bytes/s). -1 if unlimited.
	Uploaded          int          `json:"uploaded"`           // Amount of data uploaded
	UploadedSession   int          `json:"uploaded_session"`   // Amount of data uploaded this session
	UpSpeed           int          `json:"upspeed"`            // Torrent upload speed (bytes/s)
}

type GetLogParams struct {
	Normal      *bool // Include normal messages (default: true)
	Info        *bool // Include info messages (default: true)
	Warning     *bool // Include warning messages (default: true)
	Critical    *bool // Include critical messages (default: true)
	LastKnownId *int  // integer	Exclude messages with "message id" <= last_known_id (default: -1)
}

type GetLogResponse struct {
	Id        int            `json:"id"`        // ID of the message
	Message   string         `json:"message"`   // Text of the message
	Timestamp int            `json:"timestamp"` // Seconds since epoch (Note: switched from milliseconds to seconds in v4.5.0)
	Type      LogMessageType `json:"type"`      // Type of the message: Log::NORMAL: 1, Log::INFO: 2, Log::WARNING: 4, Log::CRITICAL: 8
}

type GetPeerLogResponse struct {
	Id        int    `json:"id"`        // ID of the peer
	Ip        string `json:"ip"`        // IP of the peer
	Timestamp int    `json:"timestamp"` // Seconds since epoch
	Blocked   bool   `json:"blocked"`   // Whether or not the peer was blocked
	Reason    string `json:"reason"`    // Reason of the block
}

type SyncMainDataResponse struct {
	Rid               int                   `json:"rid"`                // Response ID
	FullUpdate        bool                  `json:"full_update"`        // Whether the response contains all the data or partial data
	Torrents          TorrentListResponse   `json:"torrents"`           // Property: torrent hash, value: same as torrent list
	TorrentsRemoved   []string              `json:"torrents_removed"`   // List of hashes of torrents removed since last request
	Categories        map[string]Category   `json:"categories"`         // Info for categories added since last request
	CategoriesRemoved []map[string]Category `json:"categories_removed"` // List of categories removed since last request
	Tags              []string              `json:"tags"`               // List of tags added since last request
	TagsRemoved       []string              `json:"tags_removed"`       // array	List of tags removed since last request
	ServerState       ServerState           `json:"server_state"`       // object	Global transfer info
}

type Category struct {
	Name     string `json:"name"`
	SavePath string `json:"savePath"`
}

type ServerState struct {
	AllTimeDL            int64            `json:"alltime_dl"`
	AllTimeUL            int64            `json:"alltime_ul"`
	AverageTimeQueue     int              `json:"average_time_queue"`
	ConnectionStatus     ConnectionStatus `json:"connection_status"`
	DHTNodes             int              `json:"dht_nodes"`
	DLInfoData           int64            `json:"dl_info_data"`
	DLInfoSpeed          int64            `json:"dl_info_speed"`
	DLRateLimit          int64            `json:"dl_rate_limit"`
	FreeSpaceOnDisk      int64            `json:"free_space_on_disk"`
	GlobalRatio          string           `json:"global_ratio"`
	QueuedIOJobs         int              `json:"queued_io_jobs"`
	Queueing             bool             `json:"queueing"`
	ReadCacheHits        string           `json:"read_cache_hits"`
	ReadCacheOverload    string           `json:"read_cache_overload"`
	RefreshInterval      int              `json:"refresh_interval"`
	TotalBuffersSize     int64            `json:"total_buffers_size"`
	TotalPeerConnections int              `json:"total_peer_connections"`
	TotalQueuedSize      int64            `json:"total_queued_size"`
	TotalWastedSession   int64            `json:"total_wasted_session"`
	UPInfoData           int64            `json:"up_info_data"`
	UPInfoSpeed          int64            `json:"up_info_speed"`
	UPRateLimit          int64            `json:"up_rate_limit"`
	UseAltSpeedLimits    bool             `json:"use_alt_speed_limits"`
	WriteCacheOverload   string           `json:"write_cache_overload"`
}

type TransferInfoResponse struct {
	DLInfoSpeed       int64            `json:"dl_info_speed"`        // Global download rate in bytes per second (bytes/s)
	DLInfoData        int64            `json:"dl_info_data"`         // Data downloaded this session in bytes
	UPInfoSpeed       int64            `json:"up_info_speed"`        // Global upload rate in bytes per second (bytes/s)
	UPInfoData        int64            `json:"up_info_data"`         // Data uploaded this session in bytes
	DLRateLimit       int64            `json:"dl_rate_limit"`        // Download rate limit in bytes per second (bytes/s)
	UPRateLimit       int64            `json:"up_rate_limit"`        // Upload rate limit in bytes per second (bytes/s)
	DHTNodes          int              `json:"dht_nodes"`            // Number of DHT nodes connected to
	ConnectionStatus  ConnectionStatus `json:"connection_status"`    // Connection status (e.g., "connected", "disconnected")
	Queueing          bool             `json:"queueing"`             // True if torrent queueing is enabled
	UseAltSpeedLimits bool             `json:"use_alt_speed_limits"` // True if alternative speed limits are enabled
	RefreshInterval   int              `json:"refresh_interval"`     // Transfer list refresh interval (milliseconds)
}

type TorrentGenericProperties struct {
	SavePath               string  `json:"save_path"`                // Torrent save path
	CreationDate           int64   `json:"creation_date"`            // Torrent creation date (Unix timestamp)
	PieceSize              int64   `json:"piece_size"`               // Torrent piece size (bytes)
	Comment                string  `json:"comment"`                  // Torrent comment
	TotalWasted            int64   `json:"total_wasted"`             // Total data wasted for torrent (bytes)
	TotalUploaded          int64   `json:"total_uploaded"`           // Total data uploaded for torrent (bytes)
	TotalUploadedSession   int64   `json:"total_uploaded_session"`   // Total data uploaded this session (bytes)
	TotalDownloaded        int64   `json:"total_downloaded"`         // Total data downloaded for torrent (bytes)
	TotalDownloadedSession int64   `json:"total_downloaded_session"` // Total data downloaded this session (bytes)
	UpLimit                int64   `json:"up_limit"`                 // Torrent upload limit (bytes/s)
	DlLimit                int64   `json:"dl_limit"`                 // Torrent download limit (bytes/s)
	TimeElapsed            int64   `json:"time_elapsed"`             // Torrent elapsed time (seconds)
	SeedingTime            int64   `json:"seeding_time"`             // Torrent elapsed time while complete (seconds)
	NbConnections          int     `json:"nb_connections"`           // Torrent connection count
	NbConnectionsLimit     int     `json:"nb_connections_limit"`     // Torrent connection count limit
	ShareRatio             float64 `json:"share_ratio"`              // Torrent share ratio
	AdditionDate           int64   `json:"addition_date"`            // When this torrent was added (unix timestamp)
	CompletionDate         int64   `json:"completion_date"`          // Torrent completion date (unix timestamp)
	CreatedBy              string  `json:"created_by"`               // Torrent creator
	DlSpeedAvg             int64   `json:"dl_speed_avg"`             // Torrent average download speed (bytes/second)
	DlSpeed                int64   `json:"dl_speed"`                 // Torrent download speed (bytes/second)
	ETA                    int64   `json:"eta"`                      // Torrent ETA (seconds)
	LastSeen               int64   `json:"last_seen"`                // Last seen complete date (unix timestamp)
	Peers                  int     `json:"peers"`                    // Number of peers connected to
	PeersTotal             int     `json:"peers_total"`              // Number of peers in the swarm
	PiecesHave             int     `json:"pieces_have"`              // Number of pieces owned
	PiecesNum              int     `json:"pieces_num"`               // Number of pieces of the torrent
	Reannounce             int64   `json:"reannounce"`               // Number of seconds until the next announce
	Seeds                  int     `json:"seeds"`                    // Number of seeds connected to
	SeedsTotal             int     `json:"seeds_total"`              // Number of seeds in the swarm
	TotalSize              int64   `json:"total_size"`               // Torrent total size (bytes)
	UpSpeedAvg             int64   `json:"up_speed_avg"`             // Torrent average upload speed (bytes/second)
	UpSpeed                int64   `json:"up_speed"`                 // Torrent upload speed (bytes/second)
	IsPrivate              bool    `json:"isPrivate"`                // True if torrent is from a private tracker
}

type TorrentTracker struct {
	URL           string        `json:"url"`            // Tracker URL
	Status        TrackerStatus `json:"status"`         // Tracker status (see TrackerStatus enum for possible values)
	Tier          int           `json:"tier"`           // Tracker priority tier; lower tiers are tried before higher tiers
	NumPeers      int           `json:"num_peers"`      // Number of peers for the current torrent, as reported by the tracker
	NumSeeds      int           `json:"num_seeds"`      // Number of seeds for the current torrent, as reported by the tracker
	NumLeeches    int           `json:"num_leeches"`    // Number of leeches for the current torrent, as reported by the tracker
	NumDownloaded int           `json:"num_downloaded"` // Number of completed downloads for the current torrent, as reported by the tracker
	Msg           string        `json:"msg"`            // Tracker message; content depends on the tracker admins
}

type TorrentSeed struct {
	Url string `json:"url"` // URL of the web seed
}

type TorrentFile struct {
	Index        int          `json:"index"`        // File index
	Name         string       `json:"name"`         // File name (including relative path)
	Size         int64        `json:"size"`         // File size in bytes
	Progress     float64      `json:"progress"`     // File progress as a fraction (percentage/100)
	Priority     FilePriority `json:"priority"`     // File priority (see FilePriority enum for possible values)
	IsSeed       bool         `json:"is_seed"`      // True if file is seeding/complete
	PieceRange   []int        `json:"piece_range"`  // Array with the first number as the starting piece index and the second as the ending piece index (inclusive)
	Availability float64      `json:"availability"` // Percentage of file pieces currently available (percentage/100)
}

type RSSDownloadingRule struct {
	Enabled                   bool     `json:"enabled"`                   // Whether the rule is enabled
	MustContain               string   `json:"mustContain"`               // The substring that the torrent name must contain
	MustNotContain            string   `json:"mustNotContain"`            // The substring that the torrent name must not contain
	UseRegex                  bool     `json:"useRegex"`                  // Enable regex mode in "mustContain" and "mustNotContain"
	EpisodeFilter             string   `json:"episodeFilter"`             // Episode filter definition
	SmartFilter               bool     `json:"smartFilter"`               // Enable smart episode filter
	PreviouslyMatchedEpisodes []int    `json:"previouslyMatchedEpisodes"` // The list of episode IDs already matched by smart filter
	AffectedFeeds             []string `json:"affectedFeeds"`             // The feed URLs the rule applied to
	IgnoreDays                int      `json:"ignoreDays"`                // Ignore subsequent rule matches for these days
	LastMatch                 string   `json:"lastMatch"`                 // The rule's last match time
	AddPaused                 bool     `json:"addPaused"`                 // Add matched torrent in paused mode
	AssignedCategory          string   `json:"assignedCategory"`          // Assign category to the torrent
	SavePath                  string   `json:"savePath"`                  // Save torrent to the given directory
}

type SearchId struct {
	Id int `json:"id"` // the id of the search job
}

type SearchStatusResponse struct {
	Id     int          `json:"id"`     // ID of the search job
	Status SearchStatus `json:"status"` // Current status of the search job (either Running or Stopped)
	Total  int          `json:"total"`  // Total number of results. If the status is Running this number may contineu to increase
}

type SearchResult struct {
	DescrLink  string `json:"descrLink"`  // URL of the torrent's description page
	FileName   string `json:"fileName"`   // Name of the file
	FileSize   int    `json:"fileSize"`   // Size of the file in Bytes
	FileUrl    string `json:"fileUrl"`    // Torrent download link (usually either .torrent file or magnet link)
	NbLeechers int    `json:"nbLeechers"` // Number of leechers
	NbSeeders  int    `json:"nbSeeders"`  // Number of seeders
	SiteUrl    string `json:"siteUrl"`    // URL of the torrent site
}

type SearchResultsResponse struct {
	Results []SearchResult `json:"results"`
	Status  SearchStatus   // Current status of the search job (either Running or Stopped)
	Total   int            // Total number of results. If the status is Running this number may continue to increase
}

type SearchPluginsResponse struct {
	Enabled             bool             `json:"enabled"`             // Whether the plugin is enabled
	FullName            string           `json:"fullName"`            // Full name of the plugin
	Name                string           `json:"name"`                // Short name of the plugin
	SupportedCategories []SearchCategory `json:"supportedCategories"` // List of category objects
	Url                 string           `json:"url"`                 // URL of the torrent site
	Version             string           `json:"version"`             // Installed version of the plugin
}

type SearchCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
