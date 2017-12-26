package gotify

import (
	"encoding/json"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
)

// GetTracks : the method for GET https://api.spotify.com/v1/tracks
func (t *Tokens) GetTracks(trackIDs []string) (*models.Tracks, error) {
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

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
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
func (t *Tokens) GetAudioAnalysis(trackID string) (*models.AudioAnalysis, error) {
	/**
	https://developer.spotify.com/web-api/get-audio-analysis/
	*/

	endpoint := "https://api.spotify.com/v1/audio-analysis/" + trackID

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
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
