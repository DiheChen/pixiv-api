package client

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var httpClient = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	},
}

type PixivClient struct {
	UserName     string
	UserID       int64
	Host         string
	RefreshToken string
	Headers      map[string]string
}

func Refresh(refreshToken string) (PixivClient, error) {
	localTime := time.Now().UTC().Format("2006-01-02T15:04:05+00:00")
	hashArr := md5.Sum([]byte(localTime + "28c1fdd170a5204386cb1313c7077b34f83e4aaf4aa829ce78c231e05b0bae2c"))
	hashStr := hex.EncodeToString(hashArr[:])
	urlValues := url.Values{
		"client_id":     {"MOBrBDS8blbauoSck0ZfDbtuzpyT"},
		"client_secret": {"lsACyCD94FhDUtGTXi3QzcFE2uU1hqtDaKeqrdwj"},
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
	}
	req, err := http.NewRequest("POST", "https://oauth.secure.pixiv.net/auth/token",
		strings.NewReader(urlValues.Encode()))
	if err != nil {
		return PixivClient{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("App-OS", "ios")
	req.Header.Set("App-OS-Version", "14.6")
	req.Header.Set("App-Version", "7.6.2")
	req.Header.Set("User-Agent", "PixivIOSApp/7.13.3 (iOS 14.6; iPhone13,2)")
	req.Header.Set("X-Client-Time", localTime)
	req.Header.Set("X-Client-Hash", hashStr)
	resp, err := httpClient.Do(req)
	if err != nil {
		return PixivClient{}, err
	}
	defer func() { _ = resp.Body.Close() }()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PixivClient{}, err
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(all, &res)
	if err != nil {
		return PixivClient{}, err
	}
	parseInt, err := strconv.ParseInt(res["user"].(map[string]interface{})["id"].(string), 10, 64)
	if err != nil {
		return PixivClient{}, err
	}
	return PixivClient{
		UserName:     res["user"].(map[string]interface{})["name"].(string),
		UserID:       parseInt,
		Host:         "https://app-api.pixiv.net",
		RefreshToken: refreshToken,
		Headers: map[string]string{
			"Authorization":   "Bearer " + res["access_token"].(string),
			"Accept-Language": "zh-CN,zh;q=0.9",
		},
	}, nil
}

func (c PixivClient) Get(endpoint string, params map[string]string) (map[string]interface{}, error) {
	urlValues := url.Values{}
	for k, v := range params {
		urlValues.Set(k, v)
	}
	req, err := http.NewRequest("GET", c.Host+endpoint+"?"+urlValues.Encode(), nil)
	if err != nil {
		return nil, err
	}
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(all, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c PixivClient) Post(endpoint string, params map[string]string) ([]byte, error) {
	urlValues := url.Values{}
	for k, v := range params {
		urlValues.Set(k, v)
	}
	req, err := http.NewRequest("POST", c.Host+endpoint, strings.NewReader(urlValues.Encode()))
	if err != nil {
		return nil, err
	}
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}
