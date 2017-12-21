package gotify

import (
	"encoding/json"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
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
