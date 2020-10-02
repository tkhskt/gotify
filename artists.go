package gotify

import (
	"github.com/tkhskt/gotify/models/artist"
)

// GetArtists : the method for GET https://api.spotify.com/v1/artists
func (g *Gotify) GetArtists(artistIDs []string) (*artist.Artists, error) {
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

	var artists *artist.Artists
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

// GetArtistsAlbums : the method for GET https://api.spotify.com/v1/artists/{id}/albums
func (g *Gotify) GetArtistsAlbums(artistID string) (*artist.Albums, error) {
	/**
	https://developer.spotify.com/web-api/get-artists-albums/
	*/

	endpoint := "https://api.spotify.com/v1/artists/" + artistID + "/albums"

	var artistsAlbums *artist.Albums
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &artistsAlbums)
	if err != nil {
		return nil, err
	}
	return artistsAlbums, nil
}

// GetArtistsTopTracks : the method for GET https://api.spotify.com/v1/artists/{id}/top-tracks
func (g *Gotify) GetArtistsTopTracks(artistID string, country string) (*artist.TopTracks, error) {
	/**
	https://developer.spotify.com/web-api/get-artists-top-tracks/
	*/

	endpoint := "https://api.spotify.com/v1/artists/" + artistID + "/top-tracks?country=" + country

	var artistsTopTracks *artist.TopTracks
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &artistsTopTracks)
	if err != nil {
		return nil, err
	}
	return artistsTopTracks, nil
}

// GetArtistsRelatedArtists : the method for GET https://api.spotify.com/v1/artists/{id}/related-artists
func (g *Gotify) GetArtistsRelatedArtists(artistID string) (*artist.RelatedArtists, error) {
	/**
	https://developer.spotify.com/web-api/get-related-artists/
	*/

	endpoint := "https://api.spotify.com/v1/artists/" + artistID + "/related-artists"

	var artistsRelatedArtists *artist.RelatedArtists
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &artistsRelatedArtists)
	if err != nil {
		return nil, err
	}
	return artistsRelatedArtists, nil
}
