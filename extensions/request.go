package extensions

import (
	"io/ioutil"
	"net/http"
)

// Request : receives the endpoint and returns the result of the request
func Request(url string, token string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}
