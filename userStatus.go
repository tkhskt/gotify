package gotify

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetUserStatus : return string
func (g *Gotify) GetUserStatus() (string, error) {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/following?type=artist", nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", g.TokenInfo.GetAccessToken()))
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	return string(byteArray), nil
}
