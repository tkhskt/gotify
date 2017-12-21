package gotify

import (
	"encoding/json"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
)

// GetRecentlyPlayedTracks : the method for GET https://api.spotify.com/v1/me/player/recently-played
func (t *Tokens) GetRecentlyPlayedTracks() (*models.RecentlyPlayedTracks, error) {
	/**
	https://developer.spotify.com/web-api/web-api-personalization-endpoints/get-recently-played/
	*/
	endpoint := "https://api.spotify.com/v1/me/player/recently-played"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}
	recentlyPlayedTracks := new(models.RecentlyPlayedTracks)

	err = json.Unmarshal(res, recentlyPlayedTracks)
	if err != nil {
		return nil, err
	}
	return recentlyPlayedTracks, nil
}
