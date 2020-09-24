package gotify

import (
	"encoding/json"
	"github.com/tkhskt/gotify/extensions"
	"github.com/tkhskt/gotify/models"
)

// GetTracks : the method for GET https://api.spotify.com/v1/tracks
func (g *Gotify) GetTracks(trackIDs []string) (*models.Tracks, error) {
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

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return nil, err
	}

	tracks := new(models.Tracks)

	err = json.Unmarshal(res, tracks)
	if err != nil {
		return nil, err
	}
	return tracks, nil
}

// GetAudioAnalysis : the method for GET https://api.spotify.com/v1/audio-analysis/{id}
func (g *Gotify) GetAudioAnalysis(trackID string) (*models.AudioAnalysis, error) {
	/**
	https://developer.spotify.com/web-api/get-audio-analysis/
	*/

	endpoint := "https://api.spotify.com/v1/audio-analysis/" + trackID

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return nil, err
	}

	analysis := new(models.AudioAnalysis)

	err = json.Unmarshal(res, analysis)
	if err != nil {
		return nil, err
	}
	return analysis, nil
}
