package qbittorrent

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

/*
# Params
  - "username" Username used to access the WebUI
  - "password" Password used to access the WebUI

# Http Error Codes
  - 403 User's IP is banned for too many failed login attempts

https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#login
*/
func (c *Client) Login(username, password string) (err error) {
	query, err := url.JoinPath(c.ServerURL, "/api/v2/auth/login")
	if err != nil {
		return
	}

	data := fmt.Sprintf("username=%s&password=%s", username, password)
	req, err := http.NewRequest("POST", query, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return
	}

	req.Header.Set("Referer", c.ServerURL)
	req.Header.Set("User-Agent", "qbittorrent-webapi-go v1.0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	// 403 User's IP is banned for too many failed login attempts
	if res.StatusCode == 403 {
		return fmt.Errorf("login failed: %s", res.Status)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("login failed: %s", res.Status)
	}

	if cookies := res.Cookies(); len(cookies) > 0 {
		cookieURL, _ := url.Parse(c.ServerURL)
		c.Jar.SetCookies(cookieURL, cookies)
	}

	c.http = &http.Client{Jar: c.Jar}

	c.username = username
	c.password = password

	return
}

/*
https://github.com/qbittorrent/qBittorrent/wiki/WebUI-API-(qBittorrent-4.1)#logout
*/
func (c *Client) Logout() (err error) {
	_, err = c.postReq("/api/v2/auth/logout", nil)
	if err != nil {
		return
	}

	c.username = ""
	c.password = ""

	return
}
