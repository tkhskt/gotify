package gotify

import (
	"encoding/json"

	"github.com/tkhskt/gotify/extensions"
	"github.com/tkhskt/gotify/models/player"
	"github.com/tkhskt/gotify/values"

	"fmt"
)

// GetUsersAvailableDevices : the method for GET https://api.spotify.com/v1/me/player/devices
func (g *Gotify) GetUsersAvailableDevices() (*player.MePlayerDevices, error) {
	/**
	https://developer.spotify.com/web-api/get-a-users-available-devices/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/devices"

	var usersAvailableDevices *player.MePlayerDevices
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &usersAvailableDevices)
	if err != nil {
		return nil, err
	}

	return usersAvailableDevices, nil
}

// FIXME parse json failing
// GetInformationAboutUsersCurrentPlayback : the method for GET https://api.spotify.com/v1/me/player
func (g *Gotify) GetInformationAboutUsersCurrentPlayback() (*player.MePlayer, error) {
	/**
	https://developer.spotify.com/web-api/get-information-about-the-users-current-playback/
	*/

	endpoint := "https://api.spotify.com/v1/me/player"

	var informationAboutUsersCurrentPlayback *player.MePlayer
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &informationAboutUsersCurrentPlayback)
	if err != nil {
		return nil, err
	}

	return informationAboutUsersCurrentPlayback, nil
}

// FIXME parse json failing
// GetUsersCurrentyPlayingTrack : the method for GET https://api.spotify.com/v1/me/player/currently-playing
func (g *Gotify) GetUsersCurrentlyPlayingTrack() (*player.MePlayerCurrentlyPlaying, error) {
	/**
	https://developer.spotify.com/web-api/get-the-users-currently-playing-track/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/currently-playing"

	var usersCurrentlyPlayingTrack *player.MePlayerCurrentlyPlaying
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &usersCurrentlyPlayingTrack)
	if err != nil {
		return nil, err
	}

	return usersCurrentlyPlayingTrack, nil
}

// TransferUsersPlayback : the method for PUT https://api.spotify.com/v1/me/player
func (g *Gotify) TransferUsersPlayback(deviceIDs []string) error {
	/**
	https://developer.spotify.com/web-api/transfer-a-users-playback/
	*/
	type body struct {
		DeviceIDs []string `json:"device_ids"`
	}

	endpoint := "https://api.spotify.com/v1/me/player"

	reqBody := &body{DeviceIDs: deviceIDs}
	b, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	res, err := extensions.PutRequestWithBody(endpoint, g.TokenInfo.GetAccessToken(), string(b))
	if err != nil {
		return err
	}
	if res != values.NoContent {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// StartResumeUsersPlayback : the method for PUT https://api.spotify.com/v1/me/player/play
func (g *Gotify) StartResumeUsersPlayback() error {
	/**
	https://developer.spotify.com/web-api/start-a-users-playback/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/play"

	res, err := extensions.PutRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return err
	}
	if res != values.NoContent {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// PauseUsersPlayback : the method for PUT https://api.spotify.com/v1/me/player/pause
func (g *Gotify) PauseUsersPlayback() error {
	/**
	https://developer.spotify.com/web-api/pause-a-users-playback/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/pause"

	res, err := extensions.PutRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return err
	}
	if res != values.NoContent {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// SkipUsersPlaybackToNext : the method for POST https://api.spotify.com/v1/me/player/next
func (g *Gotify) SkipUsersPlaybackToNext() error {
	/**
	https://developer.spotify.com/web-api/pause-a-users-playback/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/next"

	res, err := extensions.PostRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return err
	}
	if res != values.NoContent {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// SkipUsersPlaybackToPrevious : the method for POST https://api.spotify.com/v1/me/player/previous
func (g *Gotify) SkipUsersPlaybackToPrevious() error {
	/**
	https://developer.spotify.com/web-api/skip-users-playback-to-previous-track/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/previous"

	res, err := extensions.PostRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return err
	}
	if res != values.NoContent {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// SeekToPositionInCurrentlyPlayingTrack : the method for PUT https://api.spotify.com/v1/me/player/seek
func (g *Gotify) SeekToPositionInCurrentlyPlayingTrack(position int) error {
	/**
	https://developer.spotify.com/web-api/seek-to-position-in-currently-playing-track/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/seek?position_ms=" + string(position)

	res, err := extensions.PutRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return err
	}
	if res != values.NoContent {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// SetRepeatModeUsersPlayback : the method for PUT https://api.spotify.com/v1/me/player/repeat
func (g *Gotify) SetRepeatModeUsersPlayback(state string) error {
	/**
	https://developer.spotify.com/web-api/set-repeat-mode-on-users-playback/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/repeat?state=" + state

	res, err := extensions.PutRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return err
	}
	if res != values.NoContent {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// SetVolumeUsersPlayback : the method for PUT https://api.spotify.com/v1/me/player/volume
func (g *Gotify) SetVolumeUsersPlayback(volumePercent int) error {
	/**
	https://developer.spotify.com/web-api/set-volume-for-users-playback/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/volume?volume_percent" + string(volumePercent)

	res, err := extensions.PutRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return err
	}
	if res != values.NoContent {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// ToggleShuffleUsersPlayback : the method for PUT https://api.spotify.com/v1/me/player/shuffle
func (g *Gotify) ToggleShuffleUsersPlayback(state bool) error {
	/**
	https://developer.spotify.com/web-api/toggle-shuffle-for-users-playback/
	*/

	endpoint := "https://api.spotify.com/v1/me/player/shuffle?state="
	if state {
		endpoint += "true"
	} else {
		endpoint += "false"
	}

	res, err := extensions.PutRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return err
	}
	if res != values.NoContent {
		return fmt.Errorf("%d", res)
	}
	return nil
}
