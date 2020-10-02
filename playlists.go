package gotify

import (
	"github.com/tkhskt/gotify/models/playlist"
)

// GetUsersPlaylists : the method for GET https://api.spotify.com/v1/users/{user_id}/playlists
func (g *Gotify) GetUsersPlaylists(userID string) (*playlist.UserPlaylists, error) {
	/**
	https://developer.spotify.com/web-api/get-list-users-playlists/
	*/

	endpoint := "https://api.spotify.com/v1/users/" + userID + "/playlists"

	var playlists *playlist.UserPlaylists
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &playlists)
	if err != nil {
		return nil, err
	}

	return playlists, nil
}

// GetUsersPlaylists : the method for GET https://api.spotify.com/v1/me/playlists
func (g *Gotify) GetCurrentUsersPlaylists() (*playlist.MePlaylists, error) {
	/**
	https://developer.spotify.com/web-api/get-a-list-of-current-users-playlists/
	*/

	endpoint := "https://api.spotify.com/v1/me/playlists"

	var playlists *playlist.MePlaylists
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &playlists)
	if err != nil {
		return nil, err
	}

	return playlists, nil
}
