package gotify

import (
	"fmt"

	"encoding/json"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
)

// SaveTracks : the method for PUT https://api.spotify.com/v1/me/tracks
func (t *Tokens) SaveTracks(trackIDs []string) error {
	/**
	https://developer.spotify.com/web-api/unfollow-playlist/
	*/

	endpoint := "https://api.spotify.com/v1/me/tracks?ids="

	for i, v := range trackIDs {
		if i == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}

	res, err := extensions.PutRequest(endpoint, t.AccessToken)
	if err != nil {
		return err
	}
	if res != 200 {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// GetUsersSavedTracks : the method for GET https://api.spotify.com/v1/me/tracks
func (t *Tokens) GetUsersSavedTracks() (*models.UsersSavedTracks, error) {
	/**
	https://developer.spotify.com/web-api/get-users-saved-tracks/
	*/

	endpoint := "https://api.spotify.com/v1/me/tracks"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}
	usersSavedTracks := new(models.UsersSavedTracks)

	err = json.Unmarshal(res, usersSavedTracks)
	if err != nil {
		return nil, err
	}
	return usersSavedTracks, nil
}

// RemoveUsersSavedTracks : the method for DELETE https://api.spotify.com/v1/me/tracks
func (t *Tokens) RemoveUsersSavedTracks(trackIDs []string) error {
	/**
	https://developer.spotify.com/web-api/remove-tracks-user/
	*/

	endpoint := "https://api.spotify.com/v1/me/tracks?ids="

	for i, v := range trackIDs {
		if i == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}

	res, err := extensions.DeleteRequest(endpoint, t.AccessToken)
	if err != nil {
		return err
	}
	if res != 200 {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// CheckUsersSavedTracks : the method for GET https://api.spotify.com/v1/me/tracks/contains
func (t *Tokens) CheckUsersSavedTracks(trackIDs []string) (*models.FollowTracks, error) {
	/**
	https://developer.spotify.com/web-api/check-users-saved-tracks/
	*/

	endpoint := "https://api.spotify.com/v1/me/tracks/contains?ids="

	for i, v := range trackIDs {
		if i == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}
	followTracks := new(models.FollowTracks)

	err = json.Unmarshal(res, followTracks)
	if err != nil {
		return nil, err
	}

	return followTracks, nil
}

// SaveAlbums : the method for PUT https://api.spotify.com/v1/me/albums?ids={ids}
func (t *Tokens) SaveAlbums(albumIDs []string) error {
	/**
	https://developer.spotify.com/web-api/save-albums-user/
	*/

	endpoint := "https://api.spotify.com/v1/me/albums?ids="

	for i, v := range albumIDs {
		if i == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}

	res, err := extensions.PutRequest(endpoint, t.AccessToken)
	if err != nil {
		return err
	}
	fmt.Println("uuuuu")
	fmt.Println(res)
	if res != 200 {
		return fmt.Errorf("%d", res)
	}
	return nil
}
