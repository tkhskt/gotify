package gotify

import (
	"encoding/json"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
)

// GetCurrentUsersProfile : the method for GET https://api.spotify.com/v1/me
func (t *Tokens) GetCurrentUsersProfile() (*models.CurrentUsersProfile, error) {
	/**
	https://developer.spotify.com/web-api/get-current-users-profile/
	*/

	endpoint := "https://api.spotify.com/v1/me"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	profile := new(models.CurrentUsersProfile)

	err = json.Unmarshal(res, profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// GetUsersProfile : the method for GET https://api.spotify.com/v1/me
func (t *Tokens) GetUsersProfile(userID string) (*models.UsersProfile, error) {
	/**
	https://developer.spotify.com/web-api/get-users-profile/
	*/

	endpoint := "https://api.spotify.com/v1/users/" + userID

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	profile := new(models.UsersProfile)

	err = json.Unmarshal(res, profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}
