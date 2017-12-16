package gotify

import (
	"encoding/json"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
)

// GetBrowseFeaturedPlaylists is the method for GET https://api.spotify.com/v1/browse/featured-playlists
func (t *Tokens) GetBrowseFeaturedPlaylists() (*models.BrowseFetaturedPlaylists, error) {
	/**
	https://developer.spotify.com/web-api/get-list-featured-playlists/
	*/

	endpoint := "https://api.spotify.com/v1/browse/featured-playlists"

	res, err := extensions.Request(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	browseFeaturedPlaylists := new(models.BrowseFetaturedPlaylists)

	err = json.Unmarshal(res, browseFeaturedPlaylists)
	if err != nil {
		return nil, err
	}
	return browseFeaturedPlaylists, nil
}
