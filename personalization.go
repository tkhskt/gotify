package gotify

import (
	"github.com/tkhskt/gotify/models/player"
)

// GetRecentlyPlayedTracks : the method for GET https://api.spotify.com/v1/me/player/recently-played
func (g *Gotify) GetRecentlyPlayedTracks() (*player.MePlayerRecentlyPlayed, error) {
	/**
	https://developer.spotify.com/web-api/web-api-personalization-endpoints/get-recently-played/
	*/
	endpoint := "https://api.spotify.com/v1/me/player/recently-played"

	var recentlyPlayedTracks *player.MePlayerRecentlyPlayed
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), *recentlyPlayedTracks)
	if err != nil {
		return nil, err
	}

	return recentlyPlayedTracks, nil
}
