package gotify

import (
	"github.com/tkhskt/gotify/models/user"
)

// GetCurrentUsersProfile : the method for GET https://api.spotify.com/v1/me
func (g *Gotify) GetCurrentUsersProfile() (*user.Me, error) {
	/**
	https://developer.spotify.com/web-api/get-current-users-profile/
	*/

	endpoint := "https://api.spotify.com/v1/me"

	var profile *user.Me
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

// GetUsersProfile : the method for GET https://api.spotify.com/v1/me
func (g *Gotify) GetUsersProfile(userID string) (*user.User, error) {
	/**
	https://developer.spotify.com/web-api/get-users-profile/
	*/

	endpoint := "https://api.spotify.com/v1/users/" + userID

	var profile *user.User
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
