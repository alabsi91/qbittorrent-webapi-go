package qbittorrent

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path"

	"golang.org/x/net/publicsuffix"
)

type Client struct {
	ServerURL string
	http      *http.Client
	Jar       http.CookieJar
}

func NewClient(serverURL string) *Client {
	client := &Client{ServerURL: serverURL}
	client.Jar, _ = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	client.http = &http.Client{Jar: client.Jar}
	return client
}

func (c *Client) getReq(endpoint string, params *url.Values) (body []byte, err error) {
	u, err := url.Parse(c.ServerURL)
	if err != nil {
		fmt.Println("Error parsing base URL:", err)
		return
	}

	u.Path, err = url.JoinPath(u.Path, endpoint)
	if err != nil {
		return
	}

	if params != nil {
		u.RawQuery = params.Encode()
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return
	}
	req.Header.Set("User-Agent", "qbittorrent-webapi-go v1.0")

	resp, err := c.http.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%d: %s", resp.StatusCode, string(body))
		return
	}

	return body, nil
}

func (c *Client) postReq(endpoint string, data *[]byte) (body []byte, err error) {
	fullUrl, err := url.JoinPath(c.ServerURL, endpoint)
	if err != nil {
		return
	}

	var payload io.Reader = nil
	if data != nil {
		payload = bytes.NewReader(*data)
	}

	req, err := http.NewRequest("POST", fullUrl, payload)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", "qbittorrent-webapi-go v1.0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.http.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%d: %s", resp.StatusCode, string(body))
		return
	}

	return
}

func (c *Client) postMultipart(endpoint string, buffer bytes.Buffer, contentType string) (body []byte, err error) {
	fullUrl, err := url.JoinPath(c.ServerURL, endpoint)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", fullUrl, &buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", "qbittorrent-webapi-go v1.0")

	resp, err := c.http.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%d: %s", resp.StatusCode, string(body))
		return
	}

	return
}

func (c *Client) addFileToWriter(writer *multipart.Writer, filePath string) (err error) {
	formWriter, err := writer.CreateFormFile("torrents", path.Base(filePath))
	if err != nil {
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(formWriter, file)
	if err != nil {
		return
	}

	return
}
