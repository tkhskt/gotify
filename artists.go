package gotify

import (
	"encoding/json"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
)

func (t *Tokens) GetArtists(artistIDs []string) (*models.Artists, error) {
	/**
	https://developer.spotify.com/web-api/get-several-artists/
	*/

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
	/**
	https://developer.spotify.com/web-api/get-artists-albums/
	*/

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

func (t *Tokens) GetArtistsTopTracks(artistID string, country string) (*models.ArtistsTopTracks, error) {
	/**
	https://developer.spotify.com/web-api/get-artists-top-tracks/
	*/

	endpoint := "https://api.spotify.com/v1/artists/" + artistID + "/top-tracks?country=" + country

	res, err := extensions.Request(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	artistsTopTracks := new(models.ArtistsTopTracks)

	err = json.Unmarshal(res, artistsTopTracks)
	if err != nil {
		return nil, err
	}
	return artistsTopTracks, nil
}

func (t *Tokens) GetArtistsRelatedArtists(artistID string) (*models.ArtistsRelatedArtists, error) {
	/**
	https://developer.spotify.com/web-api/get-related-artists/
	*/

	endpoint := "https://api.spotify.com/v1/artists/" + artistID + "/related-artists"

	res, err := extensions.Request(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	artistsRelatedArtists := new(models.ArtistsRelatedArtists)

	err = json.Unmarshal(res, artistsRelatedArtists)
	if err != nil {
		return nil, err
	}
	return artistsRelatedArtists, nil
}
