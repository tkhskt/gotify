package gotify

import (
	"github.com/tkhskt/gotify/models/track"
)

// GetTracks : the method for GET https://api.spotify.com/v1/tracks
func (g *Gotify) GetTracks(trackIDs []string) (*track.Tracks, error) {
	/**
	https://developer.spotify.com/web-api/get-several-tracks/
	*/

	endpoint := "https://api.spotify.com/v1/tracks/?ids="

	for k, v := range trackIDs {
		if k == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}

	var tracks *track.Tracks
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &tracks)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

// GetAudioAnalysis : the method for GET https://api.spotify.com/v1/audio-analysis/{id}
func (g *Gotify) GetAudioAnalysis(trackID string) (*track.AudioAnalysis, error) {
	/**
	https://developer.spotify.com/web-api/get-audio-analysis/
	*/

	endpoint := "https://api.spotify.com/v1/audio-analysis/" + trackID

	var analysis *track.AudioAnalysis
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &analysis)
	if err != nil {
		return nil, err
	}

	return analysis, nil
}
