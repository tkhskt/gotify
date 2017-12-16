package gotify

import (
	"encoding/json"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
)

// GetBrowseCategorysPlaylists : the method for GET https://api.spotify.com/v1/me/following?type=artist
func (t *Tokens) GetFollowingArtists() (*models.FollowingArtists, error) {
	/**
	https://developer.spotify.com/web-api/get-followed-artists/
	*/

	endpoint := "https://api.spotify.com/v1/me/following?type=artist"

	res, err := extensions.Request(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	followingArtists := new(models.FollowingArtists)

	err = json.Unmarshal(res, followingArtists)
	if err != nil {
		return nil, err
	}
	return followingArtists, nil
}
