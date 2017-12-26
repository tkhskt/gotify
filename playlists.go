package gotify

import (
	"encoding/json"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
)

// GetUsersPlaylists : the method for GET https://api.spotify.com/v1/users/{user_id}/playlists
func (t *Tokens) GetUsersPlaylists(userID string) (*models.UsersPlaylists, error) {
	/**
	https://developer.spotify.com/web-api/get-list-users-playlists/
	*/

	endpoint := "https://api.spotify.com/v1/users/" + userID + "/playlists"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	playlists := new(models.UsersPlaylists)

	err = json.Unmarshal(res, playlists)
	if err != nil {
		return nil, err
	}
	return playlists, nil
}
