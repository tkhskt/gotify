package gotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)


func (g *Gotify) get(url string, token string, result interface{}) error {
	const methodTypeIsGet = "GET"
	req, err := http.NewRequest(methodTypeIsGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := g.HttpClient.Do(req)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}
	return nil
}

func (g *Gotify) put(url string, token string, obj interface{}) error {
	const methodTypeIsPut = "PUT"
	return baseMethod(g.HttpClient, methodTypeIsPut, url, token, obj)
}

func (g *Gotify) delete(url string, token string, obj interface{}) error {
	const methodTypeIsDelete = "DELETE"
	return baseMethod(g.HttpClient, methodTypeIsDelete, url, token, obj)
}

func (g *Gotify) post(url string, token string, obj interface{}) error {
	const methodTypeIsPost = "POST"
	return baseMethod(g.HttpClient, methodTypeIsPost, url, token, obj)
}

func baseMethod(httpClient *http.Client, methodType string, url string, token string, obj interface{}) error {
	byteSliceObj, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(methodType, url, bytes.NewBuffer(byteSliceObj))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	_, err = httpClient.Do(req)
	if err != nil {
		return err
	}
	return nil
}