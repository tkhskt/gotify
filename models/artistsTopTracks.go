package models

// GET https://api.spotify.com/v1/artists/{id}/top-tracks

type ArtistsTopTracks struct {
	Tracks []struct {
		Album struct {
			AlbumType string `json:"album_type"`
			Artists   []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"album,omitempty"`
		Artists []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists,omitempty"`
		AvailableMarkets []string `json:"available_markets,omitempty"`
		DiscNumber       int      `json:"disc_number,omitempty"`
		DurationMs       int      `json:"duration_ms,omitempty"`
		Explicit         bool     `json:"explicit,omitempty"`
		ExternalIds      struct {
			Isrc string `json:"isrc"`
		} `json:"external_ids,omitempty"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls,omitempty"`
		Href        string `json:"href,omitempty"`
		ID          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Popularity  int    `json:"popularity,omitempty"`
		PreviewURL  string `json:"preview_url,omitempty"`
		TrackNumber int    `json:"track_number,omitempty"`
		Type        string `json:"type,omitempty"`
		URI         string `json:"uri,omitempty"`
	} `json:"tracks"`
}
