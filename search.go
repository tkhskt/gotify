package gotify

import (
	"strings"

	"github.com/tkhskt/gotify/models/search"
)

func encodeQuery(q string) string {
	qstring := strings.Replace(q, " ", "%20", -1)
	return qstring
}

// SearchArtists : the method for GET https://api.spotify.com/v1/search
func (g *Gotify) SearchArtists(keywords string) (*search.Artists, error) {
	/**
	https://developer.spotify.com/web-api/search-item/
	*/

	query := encodeQuery(keywords)

	endpoint := "https://api.spotify.com/v1/search?q=" + query + "&type=artist"

	var artists *search.Artists
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

// SearchAlbums : the method for GET https://api.spotify.com/v1/search
func (g *Gotify) SearchAlbums(keywords string) (*search.Albums, error) {
	/**
	https://developer.spotify.com/web-api/search-item/
	*/

	query := encodeQuery(keywords)

	endpoint := "https://api.spotify.com/v1/search?q=" + query + "&type=album"

	var albums *search.Albums
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &albums)
	if err != nil {
		return nil, err
	}

	return albums, nil
}

// SearchPlaylists : the method for GET https://api.spotify.com/v1/search
func (g *Gotify) SearchPlaylists(keywords string) (*search.Playlists, error) {
	/**
	https://developer.spotify.com/web-api/search-item/
	*/

	query := encodeQuery(keywords)

	endpoint := "https://api.spotify.com/v1/search?q=" + query + "&type=playlist"

	var playlists *search.Playlists
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), playlists)
	if err != nil {
		return nil, err
	}

	return playlists, nil
}

// SearchTracks : the method for GET https://api.spotify.com/v1/search
func (g *Gotify) SearchTracks(keywords string) (*search.Tracks, error) {
	/**
	https://developer.spotify.com/web-api/search-item/
	*/

	query := encodeQuery(keywords)

	endpoint := "https://api.spotify.com/v1/search?q=" + query + "&type=track"

	var tracks *search.Tracks
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &tracks)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}
