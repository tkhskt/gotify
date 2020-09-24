package gotify

import (
	"encoding/json"
	"github.com/tkhskt/gotify/extensions"
	"github.com/tkhskt/gotify/models"
)

// GetArtists : the method for GET https://api.spotify.com/v1/artists
func (g *Gotify) GetArtists(artistIDs []string) (*models.Artists, error) {
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

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
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

// GetArtistsAlbums : the method for GET https://api.spotify.com/v1/artists/{id}/albums
func (g *Gotify) GetArtistsAlbums(artistID string) (*models.ArtistsAlbums, error) {
	/**
	https://developer.spotify.com/web-api/get-artists-albums/
	*/

	endpoint := "https://api.spotify.com/v1/artists/" + artistID + "/albums"

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
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

// GetArtistsTopTracks : the method for GET https://api.spotify.com/v1/artists/{id}/top-tracks
func (g *Gotify) GetArtistsTopTracks(artistID string, country string) (*models.ArtistsTopTracks, error) {
	/**
	https://developer.spotify.com/web-api/get-artists-top-tracks/
	*/

	endpoint := "https://api.spotify.com/v1/artists/" + artistID + "/top-tracks?country=" + country

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
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

// GetArtistsRelatedArtists : the method for GET https://api.spotify.com/v1/artists/{id}/related-artists
func (g *Gotify) GetArtistsRelatedArtists(artistID string) (*models.ArtistsRelatedArtists, error) {
	/**
	https://developer.spotify.com/web-api/get-related-artists/
	*/

	endpoint := "https://api.spotify.com/v1/artists/" + artistID + "/related-artists"

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
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
