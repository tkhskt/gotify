package gotify

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type (
	Client struct {
		ClientID     string
		ClientSecret string
		CallbackURI  string
	}

	Tokens struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		Scope        string `json:"scope"`
		ExpiresIn    int    `json:"expire_in"`
		RefreshToken string `json:"refresh_token"`
	}

	OAuth interface {
		Set(string, string, string) *Client
		Auth() string
		Token(*http.Request) (string, error)
	}
)

func (c *Client) getEncodedID() string {
	str := c.ClientID + ":" + c.ClientSecret
	enc := base64.StdEncoding.EncodeToString([]byte(str))
	return enc
}

func Set(clientID string, clientSecret string, callbackURI string) *Client {
	removeSlash := strings.Replace(callbackURI, "/", "%2F", -1)
	callback := strings.Replace(removeSlash, ":", "%3A", -1)
	client := &Client{ClientID: clientID, ClientSecret: clientSecret, CallbackURI: callback}
	return client
}

func (c *Client) AuthURL() string {
	responseType := "code"

	redirectURL := "https://accounts.spotify.com/authorize/?client_id=" + c.ClientID + "&response_type=" + responseType + "&redirect_uri=" + c.CallbackURI +
		"&scope=user-read-private%20user-library-read%20user-follow-read"

	return redirectURL
}

func (c *Client) GetToken(r *http.Request) (*Tokens, error) {

	client := &http.Client{}
	code := r.URL.Query().Get("code")

	removeSlash := strings.Replace(c.CallbackURI, "%2F", "/", -1)
	redirectUri := strings.Replace(removeSlash, "%3A", ":", -1)

	//Bodyに値を追加
	values := url.Values{}
	values.Set("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("redirect_uri", redirectUri)

	//リクエストオブジェクトを生成
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	auth := c.getEncodedID()
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	data := new(Tokens) //レスポンスのjsonをバインドする アクセストークン
	if err := json.Unmarshal(byteArray, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) Refresh(t *Tokens) error {
	client := &http.Client{}

	//Bodyに値を追加
	values := url.Values{}
	values.Set("grant_type", "refresh_token")
	values.Add("refresh_token", t.RefreshToken)

	//リクエストオブジェクトを生成
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	auth := c.getEncodedID()
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	data := new(Tokens) //レスポンスのjsonをバインドする アクセストークン
	if err := json.Unmarshal(byteArray, data); err != nil {
		return err
	}
	t.AccessToken = data.AccessToken
	return nil
}
