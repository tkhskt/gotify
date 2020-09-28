package gotify

import (
	"encoding/json"
	"github.com/tkhskt/gotify/extensions"
	"github.com/tkhskt/gotify/models"
)

// GetCurrentUsersProfile : the method for GET https://api.spotify.com/v1/me
func (g *Gotify) GetCurrentUsersProfile() (*models.CurrentUsersProfile, error) {
	/**
	https://developer.spotify.com/web-api/get-current-users-profile/
	*/

	endpoint := "https://api.spotify.com/v1/me"

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
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
func (g *Gotify) GetUsersProfile(userID string) (*models.UsersProfile, error) {
	/**
	https://developer.spotify.com/web-api/get-users-profile/
	*/

	endpoint := "https://api.spotify.com/v1/users/" + userID

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
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
