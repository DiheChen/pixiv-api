package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

// Generate a random token
func generateURLSafeToken(length int) string {

	str := "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
	sb := strings.Builder{}
	sb.Grow(length)
	for i := 0; i < length; i++ {
		sb.WriteByte(str[rand.Intn(len(str))])
	}
	return sb.String()
}

// S256 transformation method.
func s256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

// Proof Key for Code Exchange by OAuth Public Clients (RFC7636).
func oauthPkce() (string, string) {
	codeVerifier := generateURLSafeToken(32)
	codeChallenge := s256(codeVerifier)
	return codeVerifier, codeChallenge
}

func getLoginURL() (string, string) {
	codeVerifier, codeChallenge := oauthPkce()
	urlValues := url.Values{
		"code_challenge":        {codeChallenge},
		"code_challenge_method": {"S256"},
		"client":                {"pixiv-android"},
	}
	return codeVerifier, "https://app-api.pixiv.net/web/v1/login" + "?" + urlValues.Encode()
}

func loginPixiv(codeVerifier, code string) (string, error) {
	urlValues := url.Values{
		"client_id":      {"MOBrBDS8blbauoSck0ZfDbtuzpyT"},
		"client_secret":  {"lsACyCD94FhDUtGTXi3QzcFE2uU1hqtDaKeqrdwj"},
		"code":           {code},
		"code_verifier":  {codeVerifier},
		"grant_type":     {"authorization_code"},
		"include_policy": {"true"},
		"redirect_uri":   {"https://app-api.pixiv.net/web/v1/users/auth/pixiv/callback"},
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}
	req, err := http.NewRequest("POST", "https://oauth.secure.pixiv.net/auth/token", strings.NewReader(urlValues.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "PixivAndroidApp/5.0.234 (Android 11; Pixel 5)")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	res := make(map[string]interface{})
	err = json.Unmarshal(all, &res)
	if err != nil {
		return "", err
	}
	return res["refresh_token"].(string), nil
}

func Login() (string, error) {
	codeVerifier, loginURL := getLoginURL()
	str := []string{"请使用现代化浏览器打开以下链接:", loginURL,
		"请先按下 f12 打开开发者控制台, 并切换到 network 选项卡。",
		"现在, 请发送剩下的请求的 request url 中请求参数里 code 的值。",
		"登录成功后, 请输入 code 值:",
		"注意, code 的生命周期非常短, 请确保上一步被迅速完成。"}
	fmt.Println(strings.Join(str, "\n"))
	var code string
	_, err := fmt.Scanln(&code)
	if err == nil {
		refreshToken, err := loginPixiv(codeVerifier, code)
		if err == nil {
			return refreshToken, nil
		}
	}
	return "", err
}
