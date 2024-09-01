package qbittorrent

type Preferences struct {
	Data map[string]interface{}
}

func NewPreferences() *Preferences {
	return &Preferences{make(map[string]interface{})}
}

// Currently selected language (e.g. en_GB for English)
func (p *Preferences) Locale(v string) *Preferences {
	p.Data["locale"] = v
	return p
}

// True if a subfolder should be created when adding a torrent
func (p *Preferences) CreateSubfolderEnabled(v bool) *Preferences {
	p.Data["create_subfolder_enabled"] = v
	return p
}

// True if torrents should be added in a Paused state
func (p *Preferences) StartPausedEnabled(v bool) *Preferences {
	p.Data["start_paused_enabled"] = v
	return p
}

// TODO - documentation missing
func (p *Preferences) AutoDeleteMode(v int) *Preferences {
	p.Data["auto_delete_mode"] = v
	return p
}

// True if disk space should be pre-allocated for all files
func (p *Preferences) PreallocateAll(v bool) *Preferences {
	p.Data["preallocate_all"] = v
	return p
}

// True if ".!qB" should be appended to incomplete files
func (p *Preferences) IncompleteFilesExt(v bool) *Preferences {
	p.Data["incomplete_files_ext"] = v
	return p
}

// True if Automatic Torrent Management is enabled by default
func (p *Preferences) AutoTmmEnabled(v bool) *Preferences {
	p.Data["auto_tmm_enabled"] = v
	return p
}

// True if torrent should be relocated when its Category changes
func (p *Preferences) TorrentChangedTmmEnabled(v bool) *Preferences {
	p.Data["torrent_changed_tmm_enabled"] = v
	return p
}

// True if torrent should be relocated when the default save path changes
func (p *Preferences) SavePathChangedTmmEnabled(v bool) *Preferences {
	p.Data["save_path_changed_tmm_enabled"] = v
	return p
}

// True if torrent should be relocated when its Category's save path changes
func (p *Preferences) CategoryChangedTmmEnabled(v bool) *Preferences {
	p.Data["category_changed_tmm_enabled"] = v
	return p
}

// Default save path for torrents, separated by slashes
func (p *Preferences) SavePath(v string) *Preferences {
	p.Data["save_path"] = v
	return p
}

// True if folder for incomplete torrents is enabled
func (p *Preferences) TempPathEnabled(v bool) *Preferences {
	p.Data["temp_path_enabled"] = v
	return p
}

// Path for incomplete torrents, separated by slashes
func (p *Preferences) TempPath(v string) *Preferences {
	p.Data["temp_path"] = v
	return p
}

// Property: directory to watch for torrent files, value: where torrents loaded from this directory should be downloaded to. Slashes are used as path separators; multiple key/value pairs can be specified
//
// Possible values of scan_dirs
//   - `0`	Download to the monitored folder
//   - `1`	Download to the default save path
//   - "/path/to/download/to"	Download to this path
func (p *Preferences) ScanDirs(v map[string]interface{}) *Preferences {
	p.Data["scan_dirs"] = v
	return p
}

// Path to directory to copy .torrent files to. Slashes are used as path separators
func (p *Preferences) ExportDir(v string) *Preferences {
	p.Data["export_dir"] = v
	return p
}

// Path to directory to copy .torrent files of completed downloads to. Slashes are used as path separators
func (p *Preferences) ExportDirFin(v string) *Preferences {
	p.Data["export_dir_fin"] = v
	return p
}

// True if e-mail notification should be enabled
func (p *Preferences) MailNotificationEnabled(v bool) *Preferences {
	p.Data["mail_notification_enabled"] = v
	return p
}

// e-mail where notifications should originate from
func (p *Preferences) MailNotificationSender(v string) *Preferences {
	p.Data["mail_notification_sender"] = v
	return p
}

// e-mail to send notifications to
func (p *Preferences) MailNotificationEmail(v string) *Preferences {
	p.Data["mail_notification_email"] = v
	return p
}

// smtp server for e-mail notifications
func (p *Preferences) MailNotificationSmtp(v string) *Preferences {
	p.Data["mail_notification_smtp"] = v
	return p
}

// True if smtp server requires SSL connection
func (p *Preferences) MailNotificationSslEnabled(v bool) *Preferences {
	p.Data["mail_notification_ssl_enabled"] = v
	return p
}

// True if smtp server requires authentication
func (p *Preferences) MailNotificationAuthEnabled(v bool) *Preferences {
	p.Data["mail_notification_auth_enabled"] = v
	return p
}

// Username for smtp authentication
func (p *Preferences) MailNotificationUsername(v string) *Preferences {
	p.Data["mail_notification_username"] = v
	return p
}

// Password for smtp authentication
func (p *Preferences) MailNotificationPassword(v string) *Preferences {
	p.Data["mail_notification_password"] = v
	return p
}

// True if external program should be run after torrent has finished downloading
func (p *Preferences) AutorunEnabled(v bool) *Preferences {
	p.Data["autorun_enabled"] = v
	return p
}

// Program path/name/arguments to run if autorun_enabled is enabled; path is separated by slashes; you can use %f and %n arguments, which will be expanded by qBittorent as path_to_torrent_file and torrent_name (from the GUI; not the .torrent file name) respectively
func (p *Preferences) AutorunProgram(v string) *Preferences {
	p.Data["autorun_program"] = v
	return p
}

// True if torrent queuing is enabled
func (p *Preferences) QueueingEnabled(v bool) *Preferences {
	p.Data["queueing_enabled"] = v
	return p
}

// Maximum number of active simultaneous downloads
func (p *Preferences) MaxActiveDownloads(v int) *Preferences {
	p.Data["max_active_downloads"] = v
	return p
}

// Maximum number of active simultaneous downloads and uploads
func (p *Preferences) MaxActiveTorrents(v int) *Preferences {
	p.Data["max_active_torrents"] = v
	return p
}

// Maximum number of active simultaneous uploads
func (p *Preferences) MaxActiveUploads(v int) *Preferences {
	p.Data["max_active_uploads"] = v
	return p
}

// If true torrents w/o any activity (stalled ones) will not be counted towards max_active_* limits; see dont_count_slow_torrents for more information
func (p *Preferences) DontCountSlowTorrents(v bool) *Preferences {
	p.Data["dont_count_slow_torrents"] = v
	return p
}

// Download rate in KiB/s for a torrent to be considered "slow"
func (p *Preferences) SlowTorrentDlRateThreshold(v int) *Preferences {
	p.Data["slow_torrent_dl_rate_threshold"] = v
	return p
}

// Upload rate in KiB/s for a torrent to be considered "slow"
func (p *Preferences) SlowTorrentUlRateThreshold(v int) *Preferences {
	p.Data["slow_torrent_ul_rate_threshold"] = v
	return p
}

// Seconds a torrent should be inactive before considered "slow"
func (p *Preferences) SlowTorrentInactiveTimer(v int) *Preferences {
	p.Data["slow_torrent_inactive_timer"] = v
	return p
}

// True if share ratio limit is enabled
func (p *Preferences) MaxRatioEnabled(v bool) *Preferences {
	p.Data["max_ratio_enabled"] = v
	return p
}

// Get the global share ratio limit
func (p *Preferences) MaxRatio(v float64) *Preferences {
	p.Data["max_ratio"] = v
	return p
}

// Action performed when a torrent reaches the maximum share ratio. See list of possible values here below.
func (p *Preferences) MaxRatioAct(v DyndnsService) *Preferences {
	p.Data["max_ratio_act"] = v
	return p
}

// Port for incoming connections
func (p *Preferences) ListenPort(v int) *Preferences {
	p.Data["listen_port"] = v
	return p
}

// True if UPnP/NAT-PMP is enabled
func (p *Preferences) Upnp(v bool) *Preferences {
	p.Data["upnp"] = v
	return p
}

// True if the port is randomly selected
func (p *Preferences) RandomPort(v bool) *Preferences {
	p.Data["random_port"] = v
	return p
}

// Global download speed limit in KiB/s; -1 means no limit is applied
func (p *Preferences) DlLimit(v int) *Preferences {
	p.Data["dl_limit"] = v
	return p
}

// Global upload speed limit in KiB/s; -1 means no limit is applied
func (p *Preferences) UpLimit(v int) *Preferences {
	p.Data["up_limit"] = v
	return p
}

// Maximum global number of simultaneous connections
func (p *Preferences) MaxConnec(v int) *Preferences {
	p.Data["max_connec"] = v
	return p
}

// Maximum number of simultaneous connections per torrent
func (p *Preferences) MaxConnecPerTorrent(v int) *Preferences {
	p.Data["max_connec_per_torrent"] = v
	return p
}

// Maximum number of upload slots
func (p *Preferences) MaxUploads(v int) *Preferences {
	p.Data["max_uploads"] = v
	return p
}

// Maximum number of upload slots per torrent
func (p *Preferences) MaxUploadsPerTorrent(v int) *Preferences {
	p.Data["max_uploads_per_torrent"] = v
	return p
}

// Timeout in seconds for a stopped announce request to trackers
func (p *Preferences) StopTrackerTimeout(v int) *Preferences {
	p.Data["stop_tracker_timeout"] = v
	return p
}

// True if the advanced libtorrent option piece_extent_affinity is enabled
func (p *Preferences) EnablePieceExtentAffinity(v bool) *Preferences {
	p.Data["enable_piece_extent_affinity"] = v
	return p
}

// Bittorrent Protocol to use
func (p *Preferences) BittorrentProtocol(v BittorrentProtocol) *Preferences {
	p.Data["bittorrent_protocol"] = v
	return p
}

// True if [du]l_limit should be applied to uTP connections; this option is only available in qBittorent built against libtorrent version 0.16.X and higher
func (p *Preferences) LimitUtpRate(v bool) *Preferences {
	p.Data["limit_utp_rate"] = v
	return p
}

// True if [du]l_limit should be applied to estimated TCP overhead (service data: e.g. packet headers)
func (p *Preferences) LimitTcpOverhead(v bool) *Preferences {
	p.Data["limit_tcp_overhead"] = v
	return p
}

// True if [du]l_limit should be applied to peers on the LAN
func (p *Preferences) LimitLanPeers(v bool) *Preferences {
	p.Data["limit_lan_peers"] = v
	return p
}

// Alternative global download speed limit in KiB/s
func (p *Preferences) AltDlLimit(v int) *Preferences {
	p.Data["alt_dl_limit"] = v
	return p
}

// Alternative global upload speed limit in KiB/s
func (p *Preferences) AltUpLimit(v int) *Preferences {
	p.Data["alt_up_limit"] = v
	return p
}

// True if alternative limits should be applied according to schedule
func (p *Preferences) SchedulerEnabled(v bool) *Preferences {
	p.Data["scheduler_enabled"] = v
	return p
}

// Scheduler starting hour
func (p *Preferences) ScheduleFromHour(v int) *Preferences {
	p.Data["schedule_from_hour"] = v
	return p
}

// Scheduler starting minute
func (p *Preferences) ScheduleFromMin(v int) *Preferences {
	p.Data["schedule_from_min"] = v
	return p
}

// Scheduler ending hour
func (p *Preferences) ScheduleToHour(v int) *Preferences {
	p.Data["schedule_to_hour"] = v
	return p
}

// Scheduler ending minute
func (p *Preferences) ScheduleToMin(v int) *Preferences {
	p.Data["schedule_to_min"] = v
	return p
}

// Scheduler days. See possible values here below
func (p *Preferences) SchedulerDays(v SchedulerDays) *Preferences {
	p.Data["scheduler_days"] = v
	return p
}

// True if DHT is enabled
func (p *Preferences) Dht(v bool) *Preferences {
	p.Data["dht"] = v
	return p
}

// True if PeX is enabled
func (p *Preferences) Pex(v bool) *Preferences {
	p.Data["pex"] = v
	return p
}

// True if LSD is enabled
func (p *Preferences) Lsd(v bool) *Preferences {
	p.Data["lsd"] = v
	return p
}

// See list of possible values here below
func (p *Preferences) Encryption(v Encryption) *Preferences {
	p.Data["encryption"] = v
	return p
}

// If true anonymous mode will be enabled; read more here; this option is only available in qBittorent built against libtorrent version 0.16.X and higher
func (p *Preferences) AnonymousMode(v bool) *Preferences {
	p.Data["anonymous_mode"] = v
	return p
}

// See list of possible values here below
func (p *Preferences) ProxyType(v int) *Preferences {
	p.Data["proxy_type"] = v
	return p
}

// Proxy IP address or domain name
func (p *Preferences) ProxyIp(v string) *Preferences {
	p.Data["proxy_ip"] = v
	return p
}

// Proxy port
func (p *Preferences) ProxyPort(v int) *Preferences {
	p.Data["proxy_port"] = v
	return p
}

// True if peer and web seed connections should be proxified; this option will have any effect only in qBittorent built against libtorrent version 0.16.X and higher
func (p *Preferences) ProxyPeerConnections(v bool) *Preferences {
	p.Data["proxy_peer_connections"] = v
	return p
}

// True proxy requires authentication; doesn't apply to SOCKS4 proxies
func (p *Preferences) ProxyAuthEnabled(v bool) *Preferences {
	p.Data["proxy_auth_enabled"] = v
	return p
}

// Username for proxy authentication
func (p *Preferences) ProxyUsername(v string) *Preferences {
	p.Data["proxy_username"] = v
	return p
}

// Password for proxy authentication
func (p *Preferences) ProxyPassword(v string) *Preferences {
	p.Data["proxy_password"] = v
	return p
}

// True if proxy is only used for torrents
func (p *Preferences) ProxyTorrentsOnly(v bool) *Preferences {
	p.Data["proxy_torrents_only"] = v
	return p
}

// True if external IP filter should be enabled
func (p *Preferences) IpFilterEnabled(v bool) *Preferences {
	p.Data["ip_filter_enabled"] = v
	return p
}

// Path to IP filter file (.dat, .p2p, .p2b files are supported); path is separated by slashes
func (p *Preferences) IpFilterPath(v string) *Preferences {
	p.Data["ip_filter_path"] = v
	return p
}

// True if IP filters are applied to trackers
func (p *Preferences) IpFilterTrackers(v bool) *Preferences {
	p.Data["ip_filter_trackers"] = v
	return p
}

// Semicolon-separated list of domains to accept when performing Host header validation
func (p *Preferences) WebUiDomainList(v string) *Preferences {
	p.Data["web_ui_domain_list"] = v
	return p
}

// IP address to use for the WebUI
func (p *Preferences) WebUiAddress(v string) *Preferences {
	p.Data["web_ui_address"] = v
	return p
}

// WebUI port
func (p *Preferences) WebUiPort(v int) *Preferences {
	p.Data["web_ui_port"] = v
	return p
}

// True if UPnP is used for the WebUI port
func (p *Preferences) WebUiUpnp(v bool) *Preferences {
	p.Data["web_ui_upnp"] = v
	return p
}

// WebUI username
func (p *Preferences) WebUiUsername(v string) *Preferences {
	p.Data["web_ui_username"] = v
	return p
}

// For API ≥ v2.3.0: Plaintext WebUI password, not readable, write-only. For API < v2.3.0: MD5 hash of WebUI password, hash is generated from the following string: usernameUI Access
func (p *Preferences) WebUiPassword(v string) *Preferences {
	p.Data["web_ui_password"] = v
	return p
}

// True if WebUI CSRF protection is enabled
func (p *Preferences) WebUiCsrfProtectionEnabled(v bool) *Preferences {
	p.Data["web_ui_csrf_protection_enabled"] = v
	return p
}

// True if WebUI clickjacking protection is enabled
func (p *Preferences) WebUiClickjackingProtectionEnabled(v bool) *Preferences {
	p.Data["web_ui_clickjacking_protection_enabled"] = v
	return p
}

// True if WebUI cookie Secure flag is enabled
func (p *Preferences) WebUiSecureCookieEnabled(v bool) *Preferences {
	p.Data["web_ui_secure_cookie_enabled"] = v
	return p
}

// Maximum number of authentication failures before WebUI access ban
func (p *Preferences) WebUiMaxAuthFailCount(v int) *Preferences {
	p.Data["web_ui_max_auth_fail_count"] = v
	return p
}

// WebUI access ban duration in seconds
func (p *Preferences) WebUiBanDuration(v int) *Preferences {
	p.Data["web_ui_ban_duration"] = v
	return p
}

// Seconds until WebUI is automatically signed off
func (p *Preferences) WebUiSessionTimeout(v int) *Preferences {
	p.Data["web_ui_session_timeout"] = v
	return p
}

// True if WebUI host header validation is enabled
func (p *Preferences) WebUiHostHeaderValidationEnabled(v bool) *Preferences {
	p.Data["web_ui_host_header_validation_enabled"] = v
	return p
}

// True if authentication challenge for loopback address (127.0.0.1) should be disabled
func (p *Preferences) BypassLocalAuth(v bool) *Preferences {
	p.Data["bypass_local_auth"] = v
	return p
}

// True if webui authentication should be bypassed for clients whose ip resides within (at least) one of the subnets on the whitelist
func (p *Preferences) BypassAuthSubnetWhitelistEnabled(v bool) *Preferences {
	p.Data["bypass_auth_subnet_whitelist_enabled"] = v
	return p
}

// (White)list of ipv4/ipv6 subnets for which webui authentication should be bypassed; list entries are separated by commas
func (p *Preferences) BypassAuthSubnetWhitelist(v string) *Preferences {
	p.Data["bypass_auth_subnet_whitelist"] = v
	return p
}

// True if an alternative WebUI should be used
func (p *Preferences) AlternativeWebuiEnabled(v bool) *Preferences {
	p.Data["alternative_webui_enabled"] = v
	return p
}

// File path to the alternative WebUI
func (p *Preferences) AlternativeWebuiPath(v string) *Preferences {
	p.Data["alternative_webui_path"] = v
	return p
}

// True if WebUI HTTPS access is enabled
func (p *Preferences) UseHttps(v bool) *Preferences {
	p.Data["use_https"] = v
	return p
}

// For API < v2.0.1: SSL keyfile contents (this is a not a path)
func (p *Preferences) SslKey(v string) *Preferences {
	p.Data["ssl_key"] = v
	return p
}

// For API < v2.0.1: SSL certificate contents (this is a not a path)
func (p *Preferences) SslCert(v string) *Preferences {
	p.Data["ssl_cert"] = v
	return p
}

// For API ≥ v2.0.1: Path to SSL keyfile
func (p *Preferences) WebUiHttpsKeyPath(v string) *Preferences {
	p.Data["web_ui_https_key_path"] = v
	return p
}

// For API ≥ v2.0.1: Path to SSL certificate
func (p *Preferences) WebUiHttpsCertPath(v string) *Preferences {
	p.Data["web_ui_https_cert_path"] = v
	return p
}

// True if server DNS should be updated dynamically
func (p *Preferences) DyndnsEnabled(v bool) *Preferences {
	p.Data["dyndns_enabled"] = v
	return p
}

// See list of possible values here below
func (p *Preferences) DyndnsService(v DyndnsService) *Preferences {
	p.Data["dyndns_service"] = v
	return p
}

// Username for DDNS service
func (p *Preferences) DyndnsUsername(v string) *Preferences {
	p.Data["dyndns_username"] = v
	return p
}

// Password for DDNS service
func (p *Preferences) DyndnsPassword(v string) *Preferences {
	p.Data["dyndns_password"] = v
	return p
}

// Your DDNS domain name
func (p *Preferences) DyndnsDomain(v string) *Preferences {
	p.Data["dyndns_domain"] = v
	return p
}

// RSS refresh interval
func (p *Preferences) RssRefreshInterval(v int) *Preferences {
	p.Data["rss_refresh_interval"] = v
	return p
}

// Max stored articles per RSS feed
func (p *Preferences) RssMaxArticlesPerFeed(v int) *Preferences {
	p.Data["rss_max_articles_per_feed"] = v
	return p
}

// Enable processing of RSS feeds
func (p *Preferences) RssProcessingEnabled(v bool) *Preferences {
	p.Data["rss_processing_enabled"] = v
	return p
}

// Enable auto-downloading of torrents from the RSS feeds
func (p *Preferences) RssAutoDownloadingEnabled(v bool) *Preferences {
	p.Data["rss_auto_downloading_enabled"] = v
	return p
}

// For API ≥ v2.5.1: Enable downloading of repack/proper Episodes
func (p *Preferences) RssDownloadRepackProperEpisodes(v bool) *Preferences {
	p.Data["rss_download_repack_proper_episodes"] = v
	return p
}

// For API ≥ v2.5.1: List of RSS Smart Episode Filters
func (p *Preferences) RssSmartEpisodeFilters(v string) *Preferences {
	p.Data["rss_smart_episode_filters"] = v
	return p
}

// Enable automatic adding of trackers to new torrents
func (p *Preferences) AddTrackersEnabled(v bool) *Preferences {
	p.Data["add_trackers_enabled"] = v
	return p
}

// List of trackers to add to new torrent
func (p *Preferences) AddTrackers(v string) *Preferences {
	p.Data["add_trackers"] = v
	return p
}

// For API ≥ v2.5.1: Enable custom http headers
func (p *Preferences) WebUiUseCustomHttpHeadersEnabled(v bool) *Preferences {
	p.Data["web_ui_use_custom_http_headers_enabled"] = v
	return p
}

// For API ≥ v2.5.1: List of custom http headers
func (p *Preferences) WebUiCustomHttpHeaders(v string) *Preferences {
	p.Data["web_ui_custom_http_headers"] = v
	return p
}

// True enables max seeding time
func (p *Preferences) MaxSeedingTimeEnabled(v bool) *Preferences {
	p.Data["max_seeding_time_enabled"] = v
	return p
}

// Number of minutes to seed a torrent
func (p *Preferences) MaxSeedingTime(v int) *Preferences {
	p.Data["max_seeding_time"] = v
	return p
}

// TODO - add description
func (p *Preferences) AnnounceIp(v string) *Preferences {
	p.Data["announce_ip"] = v
	return p
}

// True always announce to all tiers
func (p *Preferences) AnnounceToAllTiers(v bool) *Preferences {
	p.Data["announce_to_all_tiers"] = v
	return p
}

// True always announce to all trackers in a tier
func (p *Preferences) AnnounceToAllTrackers(v bool) *Preferences {
	p.Data["announce_to_all_trackers"] = v
	return p
}

// Number of asynchronous I/O threads
func (p *Preferences) AsyncIoThreads(v int) *Preferences {
	p.Data["async_io_threads"] = v
	return p
}

// List of banned IPs
func (p *Preferences) BannedIPs(v string) *Preferences {
	p.Data["banned_IPs"] = v
	return p
}

// Outstanding memory when checking torrents in MiB
func (p *Preferences) CheckingMemoryUse(v int) *Preferences {
	p.Data["checking_memory_use"] = v
	return p
}

// IP Address to bind to. Empty String means All addresses
func (p *Preferences) CurrentInterfaceAddress(v string) *Preferences {
	p.Data["current_interface_address"] = v
	return p
}

// Network Interface used
func (p *Preferences) CurrentNetworkInterface(v string) *Preferences {
	p.Data["current_network_interface"] = v
	return p
}

// Disk cache used in MiB
func (p *Preferences) DiskCache(v int) *Preferences {
	p.Data["disk_cache"] = v
	return p
}

// Disk cache expiry interval in seconds
func (p *Preferences) DiskCacheTtl(v int) *Preferences {
	p.Data["disk_cache_ttl"] = v
	return p
}

// Port used for embedded tracker
func (p *Preferences) EmbeddedTrackerPort(v int) *Preferences {
	p.Data["embedded_tracker_port"] = v
	return p
}

// True enables coalesce reads & writes
func (p *Preferences) EnableCoalesceReadWrite(v bool) *Preferences {
	p.Data["enable_coalesce_read_write"] = v
	return p
}

// True enables embedded tracker
func (p *Preferences) EnableEmbeddedTracker(v bool) *Preferences {
	p.Data["enable_embedded_tracker"] = v
	return p
}

// True allows multiple connections from the same IP address
func (p *Preferences) EnableMultiConnectionsFromSameIP(v bool) *Preferences {
	p.Data["enable_multi_connections_from_same_ip"] = v
	return p
}

// True enables os cache
func (p *Preferences) EnableOsCache(v bool) *Preferences {
	p.Data["enable_os_cache"] = v
	return p
}

// True enables sending of upload piece suggestions
func (p *Preferences) EnableUploadSuggestions(v bool) *Preferences {
	p.Data["enable_upload_suggestions"] = v
	return p
}

// File pool size
func (p *Preferences) FilePoolSize(v int) *Preferences {
	p.Data["file_pool_size"] = v
	return p
}

// Maximal outgoing port (0: Disabled)
func (p *Preferences) OutgoingPortsMax(v int) *Preferences {
	p.Data["outgoing_ports_max"] = v
	return p
}

// Minimal outgoing port (0: Disabled)
func (p *Preferences) OutgoingPortsMin(v int) *Preferences {
	p.Data["outgoing_ports_min"] = v
	return p
}

// True rechecks torrents on completion
func (p *Preferences) RecheckCompletedTorrents(v bool) *Preferences {
	p.Data["recheck_completed_torrents"] = v
	return p
}

// True resolves peer countries
func (p *Preferences) ResolvePeerCountries(v bool) *Preferences {
	p.Data["resolve_peer_countries"] = v
	return p
}

// Save resume data interval in min
func (p *Preferences) SaveResumeDataInterval(v int) *Preferences {
	p.Data["save_resume_data_interval"] = v
	return p
}

// Send buffer low watermark in KiB
func (p *Preferences) SendBufferLowWatermark(v int) *Preferences {
	p.Data["send_buffer_low_watermark"] = v
	return p
}

// Send buffer watermark in KiB
func (p *Preferences) SendBufferWatermark(v int) *Preferences {
	p.Data["send_buffer_watermark"] = v
	return p
}

// Send buffer watermark factor in percent
func (p *Preferences) SocketBacklogSize(v int) *Preferences {
	p.Data["socket_backlog_size"] = v
	return p
}

// Upload choking algorithm used
func (p *Preferences) UploadChokingAlgorithm(v UploadChokingAlgorithm) *Preferences {
	p.Data["upload_choking_algorithm"] = v
	return p
}

// Upload slots behavior used
func (p *Preferences) UploadSlotsBehavior(v UploadSlotsBehavior) *Preferences {
	p.Data["upload_slots_behavior"] = v
	return p
}

// UPnP lease duration (0: Permanent lease)
func (p *Preferences) UpnpLeaseDuration(v int) *Preferences {
	p.Data["upnp_lease_duration"] = v
	return p
}

// μTP-TCP mixed mode algorithm
func (p *Preferences) UtpTcpMixedMode(v UtpTcpMixedMode) *Preferences {
	p.Data["utp_tcp_mixed_mode"] = v
	return p
}
