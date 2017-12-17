package extensions

import (
	"io/ioutil"
	"net/http"
)

// GetRequest : receives the endpoint and returns the result of the request
func GetRequest(url string, token string) ([]byte, error) {
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

// PutRequest : request to endpoint with PUT method
func PutReqest(url string, token string) (int, error) {
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	//byteArray, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, nil
}

// DeleteRequest : request to endpoint with DELETE method
func DeleteReqest(url string, token string) (int, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	//byteArray, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, nil
}
