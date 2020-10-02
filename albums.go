package gotify

import (
	"github.com/tkhskt/gotify/models/album"
)

// GetAlbums : method for GET https://api.spotify.com/v1/albums
func (g *Gotify) GetAlbums(albumIDs []string) (*album.Albums, error) {
	/**
	https://developer.spotify.com/web-api/get-several-albums/
	*/

	endpoint := "https://api.spotify.com/v1/albums/?ids="

	for k, v := range albumIDs {
		if k == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}
	var albums *album.Albums
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &albums)
	if err != nil {
		return nil, err
	}

	return albums, nil
}

// GetAlbumsTracks : the method for GET https://api.spotify.com/v1/albums/{id}/tracks
func (g *Gotify) GetAlbumsTracks(albumID string) (*album.Tracks, error) {
	/**
	https://developer.spotify.com/web-api/get-albums-tracks/
	*/

	endpoint := "https://api.spotify.com/v1/albums/" + albumID + "/tracks"

	var albumTracks *album.Tracks
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &albumTracks)
	if err != nil {
		return nil, err
	}
	return albumTracks, nil
}
