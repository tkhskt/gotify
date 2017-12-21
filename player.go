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
