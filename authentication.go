package qbittorrent

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

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

	return
}

func (c *Client) Logout() (err error) {
	_, err = c.postReq("/api/v2/auth/logout", nil)
	if err != nil {
		return
	}

	return
}
