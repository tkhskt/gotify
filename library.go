package gotify

import (
	"fmt"
	"github.com/tkhskt/gotify/extensions"
	"github.com/tkhskt/gotify/models"
	"github.com/tkhskt/gotify/values"

	"encoding/json"
)

// SaveTracks : the method for PUT https://api.spotify.com/v1/me/tracks
func (g *Gotify) SaveTracks(trackIDs []string) error {
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

	res, err := extensions.PutRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return err
	}
	if res != values.OK {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// GetUsersSavedTracks : the method for GET https://api.spotify.com/v1/me/tracks
func (g *Gotify) GetUsersSavedTracks() (*models.UsersSavedTracks, error) {
	/**
	https://developer.spotify.com/web-api/get-users-saved-tracks/
	*/

	endpoint := "https://api.spotify.com/v1/me/tracks"

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
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
func (g *Gotify) RemoveUsersSavedTracks(trackIDs []string) error {
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

	res, err := extensions.DeleteRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return err
	}
	if res != values.OK {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// CheckUsersSavedTracks : the method for GET https://api.spotify.com/v1/me/tracks/contains
func (g *Gotify) CheckUsersSavedTracks(trackIDs []string) (*models.FollowTracks, error) {
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

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
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

// SaveAlbums : the method for GET https://api.spotify.com/v1/me/albums
func (g *Gotify) SaveAlbums(albumIDs []string) error {
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

	res, err := extensions.PutRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return err
	}
	if res != values.OK {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// GetUsersSavedAlbums : the method for GET https://api.spotify.com/v1/me/albums
func (g *Gotify) GetUsersSavedAlbums() (*models.UsersSavedAlbums, error) {
	/**
	https://developer.spotify.com/web-api/get-users-saved-albums/
	*/

	endpoint := "https://api.spotify.com/v1/me/albums"

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return nil, err
	}
	usersSavedAlbums := new(models.UsersSavedAlbums)

	err = json.Unmarshal(res, usersSavedAlbums)
	if err != nil {
		return nil, err
	}

	return usersSavedAlbums, nil
}

// RemoveAlbumsForCurrentUser : the method for DELETE https://api.spotify.com/v1/me/albums?ids={ids}
func (g *Gotify) RemoveAlbumsForCurrentUser(albumIDs []string) error {
	/**
	https://developer.spotify.com/web-api/remove-albums-user/
	*/
	endpoint := "https://api.spotify.com/v1/me/albums?ids="

	for i, v := range albumIDs {
		if i == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}

	res, err := extensions.DeleteRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return err
	}
	if res != values.OK {
		return fmt.Errorf("%d", res)
	}
	return nil
}

// CheckUsersSavedAlbums : the method for GET https://api.spotify.com/v1/me/albums/contains
func (g *Gotify) CheckUsersSavedAlbums(albumIDs []string) (*models.FollowAlbums, error) {
	/**
	https://developer.spotify.com/web-api/check-users-saved-albums/
	*/

	endpoint := "https://api.spotify.com/v1/me/albums/contains?ids="

	for i, v := range albumIDs {
		if i == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return nil, err
	}
	followAlbums := new(models.FollowAlbums)

	err = json.Unmarshal(res, followAlbums)
	if err != nil {
		return nil, err
	}

	return followAlbums, nil
}
