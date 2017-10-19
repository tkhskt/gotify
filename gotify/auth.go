package gotify

import "strings"

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

	Endpoints interface {
		GetUserStatus() string
	}
)

func Set(clientID string, clientSecret string, callbackURI string) (*Client) {
	client := &Client{ClientID: clientID, ClientSecret: clientSecret, CallbackURI: callbackURI}
	return client
}

func (c *Client) AuthCode() string {
	responseType := "code"
	removeSlash := strings.Replace(c.CallbackURI, "/", "%2F", -1)
	redirectURI := strings.Replace(removeSlash, ":", "%3A", -1)

	url := "https://accounts.spotify.com/authorize/?client_id=" + c.ClientID + "&response_type=" + responseType + "&redirect_uri=" + redirectURI +
		"&scope=user-read-private%20user-library-read%20user-follow-read"

	return url
}
