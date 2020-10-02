package gotify

import (
	"fmt"

	"github.com/tkhskt/gotify/extensions"
	"github.com/tkhskt/gotify/models/library"
	"github.com/tkhskt/gotify/values"
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
func (g *Gotify) GetUsersSavedTracks() (*library.MeTracks, error) {
	/**
	https://developer.spotify.com/web-api/get-users-saved-tracks/
	*/

	endpoint := "https://api.spotify.com/v1/me/tracks"

	var usersSavedTracks *library.MeTracks
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), *usersSavedTracks)
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
func (g *Gotify) CheckUsersSavedTracks(trackIDs []string) (*library.FollowTracks, error) {
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

	var followTracks *library.FollowTracks
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &followTracks)
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
func (g *Gotify) GetUsersSavedAlbums() (*library.MeAlbums, error) {
	/**
	https://developer.spotify.com/web-api/get-users-saved-albums/
	*/

	endpoint := "https://api.spotify.com/v1/me/albums"

	var usersSavedAlbums *library.MeAlbums
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &usersSavedAlbums)
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
func (g *Gotify) CheckUsersSavedAlbums(albumIDs []string) (*library.FollowAlbums, error) {
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

	var followAlbums *library.FollowAlbums
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &followAlbums)
	if err != nil {
		return nil, err
	}

	return followAlbums, nil
}
