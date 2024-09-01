package qbittorrent

type RssRule struct {
	Data map[string]interface{}
}

func NewRSSRule() *RssRule {
	return &RssRule{
		Data: make(map[string]interface{}),
	}
}

// Whether the rule is enabled
func (r *RssRule) Enabled(v bool) *RssRule {
	r.Data["enabled"] = v
	return r
}

// The substring that the torrent name must contain
func (r *RssRule) MustContain(v string) *RssRule {
	r.Data["mustContain"] = v
	return r
}

// The substring that the torrent name must not contain
func (r *RssRule) MustNotContain(v string) *RssRule {
	r.Data["mustNotContain"] = v
	return r
}

// Enable regex mode in "mustContain" and "mustNotContain"
func (r *RssRule) UseRegex(v bool) *RssRule {
	r.Data["useRegex"] = v
	return r
}

// Episode filter definition
func (r *RssRule) EpisodeFilter(v string) *RssRule {
	r.Data["episodeFilter"] = v
	return r
}

// Enable smart episode filter
func (r *RssRule) SmartFilter(v bool) *RssRule {
	r.Data["smartFilter"] = v
	return r
}

// The list of episode IDs already matched by smart filter
func (r *RssRule) PreviouslyMatchedEpisodes(v []int) *RssRule {
	r.Data["previouslyMatchedEpisodes"] = v
	return r
}

// The feed URLs the rule applied to
func (r *RssRule) AffectedFeeds(v []string) *RssRule {
	r.Data["affectedFeeds"] = v
	return r
}

// Ignore sunsequent rule matches
func (r *RssRule) IgnoreDays(v int) *RssRule {
	r.Data["ignoreDays"] = v
	return r
}

// The rule last match time
func (r *RssRule) LastMatch(v string) *RssRule {
	r.Data["lastMatch"] = v
	return r
}

// Add matched torrent in paused mode
func (r *RssRule) AddPaused(v bool) *RssRule {
	r.Data["addPaused"] = v
	return r
}

// Assign category to the torrent
func (r *RssRule) AssignedCategory(v string) *RssRule {
	r.Data["assignedCategory"] = v
	return r
}

// Save torrent to the given directory
func (r *RssRule) SavePath(v string) *RssRule {
	r.Data["savePath"] = v
	return r
}
