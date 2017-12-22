package gotify

import (
	"encoding/json"

	"fmt"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
	"github.com/gericass/gotify/values"
)

// GetUsersAvailableDevices : the method for GET https://api.spotify.com/v1/me/player/devices
func (t *Tokens) GetUsersAvailableDevices() (*models.UsersAvailableDevices, error) {
	/**
	https://developer.spotify.com/web-api/get-a-users-available-devices/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/devices"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}
	usersAvailableDevices := new(models.UsersAvailableDevices)

	err = json.Unmarshal(res, usersAvailableDevices)
	if err != nil {
		return nil, err
	}

	return usersAvailableDevices, nil
}

// FIXME parse json failing
// GetInformationAboutUsersCurrentPlayback : the method for GET https://api.spotify.com/v1/me/player
func (t *Tokens) GetInformationAboutUsersCurrentPlayback() (*models.InformationAboutUsersCurrentPlayback, error) {
	/**
	https://developer.spotify.com/web-api/get-information-about-the-users-current-playback/
	*/

	endpoint := "https://api.spotify.com/v1/me/player"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	informationAboutUsersCurrentPlayback := new(models.InformationAboutUsersCurrentPlayback)

	err = json.Unmarshal(res, informationAboutUsersCurrentPlayback)
	if err != nil {
		return nil, err
	}

	return informationAboutUsersCurrentPlayback, nil
}

// FIXME parse json failing
// GetUsersCurrentyPlayingTrack : the method for GET https://api.spotify.com/v1/me/player/currently-playing
func (t *Tokens) GetUsersCurrentlyPlayingTrack() (*models.UsersCurrentlyPlayingTrack, error) {
	/**
	https://developer.spotify.com/web-api/get-the-users-currently-playing-track/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/currently-playing"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	usersCurrentlyPlayingTrack := new(models.UsersCurrentlyPlayingTrack)

	err = json.Unmarshal(res, usersCurrentlyPlayingTrack)
	if err != nil {
		return nil, err
	}

	return usersCurrentlyPlayingTrack, nil
}

// TransferUsersPlayback : the method for PUT https://api.spotify.com/v1/me/player
func (t *Tokens) TransferUsersPlayback(deviceIDs []string) error {
	/**
	https://developer.spotify.com/web-api/transfer-a-users-playback/
	*/
	type body struct {
		DeviceIDs []string `json:"device_ids"`
	}

	endpoint := "https://api.spotify.com/v1/me/player"

	reqBody := &body{DeviceIDs: deviceIDs}
	b, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	res, err := extensions.PutRequestWithBody(endpoint, t.AccessToken, string(b))
	if err != nil {
		return err
	}
	if res != values.NoContent {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// StartResumeUsersPlayback : the method for PUT https://api.spotify.com/v1/me/player/play
func (t *Tokens) StartResumeUsersPlayback() error {
	/**
	https://developer.spotify.com/web-api/start-a-users-playback/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/play"

	res, err := extensions.PutRequest(endpoint, t.AccessToken)
	if err != nil {
		return err
	}
	if res != values.NoContent {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// PauseUsersPlayback : the method for PUT https://api.spotify.com/v1/me/player/pause
func (t *Tokens) PauseUsersPlayback() error {
	/**
	https://developer.spotify.com/web-api/pause-a-users-playback/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/pause"

	res, err := extensions.PutRequest(endpoint, t.AccessToken)
	if err != nil {
		return err
	}
	if res != values.NoContent {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// SkipUsersPlayback : the method for POST https://api.spotify.com/v1/me/player/next
func (t *Tokens) SkipUsersPlayback() error {
	/**
	https://developer.spotify.com/web-api/pause-a-users-playback/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/next"

	res, err := extensions.PostRequest(endpoint, t.AccessToken)
	if err != nil {
		return err
	}
	if res != values.NoContent {
		return fmt.Errorf("%d", res)
	}
	return nil
}
