package qbittorrent

import (
	"encoding/json"
	"net/url"
	"strconv"
)

/*
# Params
  - "path" Full path of added folder (e.g. "The Pirate Bay\Top100")

# Http Error Codes
  - 409 Failure to add folder
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#add-folder
*/
func (c *Client) AddRSSFolder(path string) (err error) {
	form := url.Values{}
	form.Add("path", path)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/rss/addFolder", &formDataBytes)
	return
}

/*
# Params
  - "url" URL of RSS feed (e.g. "http://thepiratebay.org/rss//top100/200")
  - "path" Full path of added folder (e.g. "The Pirate Bay\Top100\Video")

# Http Error Codes
  - 409 Failure to add feed
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#add-feed
*/
func (c *Client) AddRSSFeed(feedUrl, path string) (err error) {
	form := url.Values{}
	form.Add("url", feedUrl)
	form.Add("path", path)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/rss/addFeed", &formDataBytes)
	return
}

/*
# Params
  - "path" Full path of removed item (e.g. "The Pirate Bay\Top100")

# Http Error Codes
  - 409 Failure to remove item
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#remove-item
*/
func (c *Client) RemoveRSSItem(path string) (err error) {
	form := url.Values{}
	form.Add("path", path)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/rss/removeItem", &formDataBytes)
	return
}

/*
# Params
  - "itemPath" Current full path of item (e.g. "The Pirate Bay\Top100")
  - "destPath" New full path of item (e.g. "The Pirate Bay")

# Http Error Codes
  - 409 Failure to move item
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#move-item
*/
func (c *Client) MoveRSSItem(itemPath, destPath string) (err error) {
	form := url.Values{}
	form.Add("itemPath", itemPath)
	form.Add("destPath", destPath)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/rss/moveItem", &formDataBytes)
	return
}

/*
# Params
  - "withData" True if you need current feed articles

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-all-items
*/
func (c *Client) GetAllRSSItems(withData bool) (results map[string]interface{}, err error) {
	form := url.Values{}
	form.Add("withData", strconv.FormatBool(withData))

	body, err := c.getReq("/api/v2/rss/items", &form)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &results)

	return
}

/*
# Params
  - "itemPath" Current full path of item (e.g. "The Pirate Bay\Top100")
  - "articleId" ID of article

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#mark-as-read
*/
func (c *Client) MarkRSSAsRead(itemPath, articleId string) (err error) {
	form := url.Values{}
	form.Add("itemPath", itemPath)
	form.Add("articleId", articleId)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/rss/markAsRead", &formDataBytes)
	return
}

/*
# Params
  - "itemPath" Current full path of item (e.g. "The Pirate Bay\Top100")

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#refresh-item
*/
func (c *Client) RefreshRSSItem(itemPath string) (err error) {
	form := url.Values{}
	form.Add("itemPath", itemPath)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/rss/refreshItem", &formDataBytes)
	return
}

/*
# Params
  - "ruleName" Name of rule
  - "ruleDef" Definition of rule. Use [NewRSSRule] to build it then pass its `data` field

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#set-auto-downloading-rule
*/
func (c *Client) SetRSSAutoDownloadingRule(ruleName string, ruleDef map[string]interface{}) (err error) {
	ruleDefStr, err := json.Marshal(ruleDef)
	if err != nil {
		return
	}

	form := url.Values{}
	form.Add("ruleName", ruleName)
	form.Add("ruleDef", string(ruleDefStr))

	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/rss/setRule", &formDataBytes)

	return
}

/*
# Params
  - "ruleName" Name of rule
  - "newRuleName" New name of rule

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#rename-auto-downloading-rule
*/
func (c *Client) RenameRSSAutoDownloadingRule(ruleName, newRuleName string) (err error) {
	form := url.Values{}
	form.Add("ruleName", ruleName)
	form.Add("newRuleName", newRuleName)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/rss/renameRule", &formDataBytes)
	return
}

/*
# Params
  - "ruleName" Name of rule

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#remove-auto-downloading-rule
*/
func (c *Client) RemoveRSSAutoDownloadingRule(ruleName string) (err error) {
	form := url.Values{}
	form.Add("ruleName", ruleName)
	formDataBytes := []byte(form.Encode())
	_, err = c.postReq("/api/v2/rss/removeRule", &formDataBytes)
	return
}

/*
Returns all auto-downloading rules

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-all-auto-downloading-rules
*/
func (c *Client) GetAllRSSDownloadingRules() (results map[string]RSSDownloadingRule, err error) {
	body, err := c.getReq("/api/v2/rss/rules", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &results)
	return
}

/*
Returns all articles that match a rule by feed name

# Params
  - "ruleName" Name of rule

# Http Error Codes
  - 403 Forbidden, if the client is not authorized

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#get-all-articles-matching-a-rule
*/
func (c *Client) GetAllRSSArticlesMatchingRule(ruleName string) (results map[string][]string, err error) {
	form := url.Values{}
	form.Add("ruleName", ruleName)

	body, err := c.getReq("/api/v2/rss/matchingArticles", &form)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &results)

	return
}
