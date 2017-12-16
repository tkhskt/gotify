package gotify

import (
	"encoding/json"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
)

// GetAlbums : method for GET https://api.spotify.com/v1/albums
func (t *Tokens) GetAlbums(albumIDs []string) (*models.Albums, error) {
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

	res, err := extensions.Request(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}
	albums := new(models.Albums)

	err = json.Unmarshal(res, albums)
	if err != nil {
		return nil, err
	}

	return albums, nil
}

// GetAlbumsTracks : the method for GET https://api.spotify.com/v1/albums/{id}/tracks
func (t *Tokens) GetAlbumsTracks(albumID string) (*models.AlbumsTracks, error) {
	/**
	https://developer.spotify.com/web-api/get-albums-tracks/
	*/

	endpoint := "https://api.spotify.com/v1/albums/" + albumID + "/tracks"

	res, err := extensions.Request(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	albumTracks := new(models.AlbumsTracks)

	err = json.Unmarshal(res, albumTracks)
	if err != nil {
		return nil, err
	}
	return albumTracks, nil
}
