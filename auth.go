package gotify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tkhskt/gotify/values"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type (
	OAuthClientInfo struct {
		ClientID     string
		ClientSecret string
		CallbackURI  string
	}
	IOAuthClientInfo interface {
		GetRequestAuthorizationURL() (*url.URL, error)
		GetCallBackURL() string
		GetEncodedID() string
		parseURL()
	}
)

func NewOAuthClientInfo(clientID, clientSecret, callbackURI string) IOAuthClientInfo {
	oAuthClientInfo := new(OAuthClientInfo)
	oAuthClientInfo.ClientID = clientID
	oAuthClientInfo.ClientSecret = clientSecret
	oAuthClientInfo.CallbackURI = callbackURI
	oAuthClientInfo.parseURL()
	return oAuthClientInfo
}

func (O OAuthClientInfo) GetCallBackURL() string {
	return O.CallbackURI
}

func (O OAuthClientInfo) GetEncodedID() string {
	str := O.ClientID + ":" + O.ClientSecret
	enc := base64.StdEncoding.EncodeToString([]byte(str))
	return enc
}

func (O OAuthClientInfo) parseURL() {
	firstReplacedSlash := strings.Replace(O.CallbackURI, "/", "%2F", -1)
	secondReplacedColon := strings.Replace(firstReplacedSlash, ":", "%3A", -1)
	O.CallbackURI = secondReplacedColon
}

func (O OAuthClientInfo) GetRequestAuthorizationURL() (*url.URL, error) {
	const (
		baseAuthorizeURL = "https://accounts.spotify.com/authorize"
		option           = "?client_id=%s&response_type=%s&redirect_uri=%s&scope=" //"user-read-private%20user-library-read%20user-follow-read"
		responseType     = "code"
		and              = "%20"
	)

	redirectURL := fmt.Sprintf(baseAuthorizeURL+option, O.ClientID, responseType, O.CallbackURI)
	// Setup Scope
	scopes := []string{
		values.PlaylistReadPrivate,
		values.PlaylistReadCollaborative,
		values.PlaylistModifyPublic,
		values.PlaylistModifyPrivate,
		values.Streaming,
		values.UgcImageUpload,
		values.UserFollowModify,
		values.UserFollowRead,
		values.UserLibraryRead,
		values.UserLibraryModify,
		values.UserReadPrivate,
		values.UserReadEmail,
		values.UserTopRead,
		values.UserReadPlaybackState,
		values.UserModifyPlaybackState,
		values.UserReadCurrentlyPlaying,
		values.UserReadRecentlyPlayed}
	for i, v := range scopes {
		if i == 0 {
			redirectURL += v
		} else {
			redirectURL += and + v
		}
	}
	parsedURL, err := url.Parse(redirectURL)
	if err != nil {
		return nil, fmt.Errorf("url parse error on /OAuthClientInfo/GetRequestAuthorizationURL: %w", err)
	}
	return parsedURL, nil
}

type (
	TokenInfo struct {
		AccessToken     string `json:"access_token"`
		TokenType       string `json:"token_type"`
		Scope           string `json:"scope"`
		ExpiresIn       int    `json:"expire_in"`
		RefreshToken    string `json:"refresh_token"`
		ExpirationTime  time.Time
		OAuthClientInfo IOAuthClientInfo
	}
	ITokenInfo interface {
		GetAccessToken() string
		ValidateExpireIn() error
		refresh() error
		getTokenOnHTTP(redirectedURL url.URL) (io.ReadCloser, error)
	}
)

func NewTokenInfo(redirectedURL url.URL) (ITokenInfo, error) {
	tokenInfo := new(TokenInfo)
	tokenBody, err := tokenInfo.getTokenOnHTTP(redirectedURL)
	if err != nil {
		return nil, err
	}
	byteBody, err := ioutil.ReadAll(tokenBody)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(byteBody, tokenInfo); err != nil {
		return nil, err
	}
	tokenInfo.ExpirationTime = time.Now().Add(time.Duration(tokenInfo.ExpiresIn) * time.Second)
	return tokenInfo, nil
}

func (t *TokenInfo) GetAccessToken() string {
	return t.AccessToken
}

func (t *TokenInfo) ValidateExpireIn() error {
	now := time.Now()
	if t.ExpirationTime.After(now) {
		if err := t.refresh(); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (t *TokenInfo) refresh() error {
	const (
		requestAccessTokenURL = "https://accounts.spotify.com/api/token"
	)
	var (
		client    = new(http.Client)
		formValue = new(url.Values)
	)

	formValue.Set("grant_type", "refresh_token")
	formValue.Add("refresh_token", t.RefreshToken)

	req, err := http.NewRequest("POST", requestAccessTokenURL, strings.NewReader(formValue.Encode()))
	if err != nil {
		return err
	}

	clientEncodedID := t.OAuthClientInfo.GetEncodedID()
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", clientEncodedID))
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	data := new(TokenInfo)
	if err := json.Unmarshal(byteArray, data); err != nil {
		return err
	}
	t.AccessToken = data.AccessToken
	t.ExpiresIn = data.ExpiresIn
	t.ExpirationTime = time.Now().Add(time.Duration(t.ExpiresIn) * time.Second)
	return nil
}

func (t TokenInfo) getTokenOnHTTP(redirectedURL url.URL) (io.ReadCloser, error) {
	// 2. Have your application request refresh and access tokens; Spotify returns access and refresh tokens
	const (
		requestAccessTokenURL = "https://accounts.spotify.com/api/token"
	)
	var (
		client    = new(http.Client)
		formValue = new(url.Values)
		code      = redirectedURL.Query().Get("code")
	)

	formValue.Set("grant_type", "authorization_code")
	formValue.Add("code", code)
	formValue.Add("redirect_uri", t.OAuthClientInfo.GetCallBackURL())

	req, err := http.NewRequest("POST", requestAccessTokenURL, strings.NewReader(formValue.Encode()))
	if err != nil {
		return nil, fmt.Errorf("error: create request error on /TokenInfo/getTokenOnHTTP: %w", err)
	}

	clientEncodedID := t.OAuthClientInfo.GetEncodedID()
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", clientEncodedID))
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error: request access token error on /TokenInfo/getTokenOnHTTP: %w", err)
	}
	return resp.Body, nil
}
