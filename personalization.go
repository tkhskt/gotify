package gotify

import (
	"encoding/json"
	"github.com/tkhskt/gotify/extensions"
	"github.com/tkhskt/gotify/models"
)

// GetRecentlyPlayedTracks : the method for GET https://api.spotify.com/v1/me/player/recently-played
func (g *Gotify) GetRecentlyPlayedTracks() (*models.RecentlyPlayedTracks, error) {
	/**
	https://developer.spotify.com/web-api/web-api-personalization-endpoints/get-recently-played/
	*/
	endpoint := "https://api.spotify.com/v1/me/player/recently-played"

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
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
