package gotify

import (
	"encoding/json"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
)

func (t *Tokens) GetArtists(artistIDs []string) (*models.Artists, error) {

	endpoint := "https://api.spotify.com/v1/artists/?ids="

	for k, v := range artistIDs {
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

	artists := new(models.Artists)

	err = json.Unmarshal(res, artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func (t *Tokens) GetArtistsAlbums(artistID string) (*models.ArtistsAlbums, error) {

	endpoint := "https://api.spotify.com/v1/artists/" + artistID + "/albums"

	res, err := extensions.Request(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	artistsAlbums := new(models.ArtistsAlbums)

	err = json.Unmarshal(res, artistsAlbums)
	if err != nil {
		return nil, err
	}
	return artistsAlbums, nil
}
