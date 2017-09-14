package handler

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo"
)

const clientID string = ""
const clientSecret string = ""

type Tokens struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expire_in"`
	RefreshToken string `json:"refresh_token"`
}

func getUserStatus(token string) (string, error) {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/following?type=artist", nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)

	return string(byteArray), nil
}

func getEncodedID() string {
	str := clientID + ":" + clientSecret
	enc := base64.StdEncoding.EncodeToString([]byte(str))
	return enc
}

func OAuthHandler(c echo.Context) error {
	responseType := "code"
	redirectURI := "https%3A%2F%2Flocalhost%3A3000%2Fcallback%2F"
	url := "https://accounts.spotify.com/authorize/?client_id=" + clientID + "&response_type=" + responseType + "&redirect_uri=" + redirectURI + "&scope=user-read-private%20user-library-read%20user-follow-read"
	return c.Redirect(http.StatusSeeOther, url)
}

func CallbackHandler(c echo.Context) error {

	client := &http.Client{}
	qr := c.QueryParams()
	str := qr["code"]

	//Bpdyに値を追加
	values := url.Values{}
	values.Set("grant_type", "authorization_code")
	values.Add("code", str[0])
	values.Add("redirect_uri", "https://localhost:3000/callback/")

	//リクエストオブジェクトを生成
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	auth := getEncodedID()
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	data := new(Tokens) //レスポンスのjsonをバインドする
	if err := json.Unmarshal(byteArray, data); err != nil {
		return err
	}
	result, err := getUserStatus(data.AccessToken)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, result)
}
