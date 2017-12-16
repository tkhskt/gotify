package models

// BrowseNewReleases : the struct for GET https://api.spotify.com/v1/browse/new-releases
type BrowseNewReleases struct {
	Albums struct {
		Href  string `json:"href"`
		Items []struct {
			AlbumType string `json:"album_type,omitempty"`
			Artists   []struct {
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
			ExternalUrls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls,omitempty"`
			Href   string `json:"href,omitempty"`
			ID     string `json:"id,omitempty"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images,omitempty"`
			Name string `json:"name,omitempty"`
			Type string `json:"type,omitempty"`
			URI  string `json:"uri,omitempty"`
		} `json:"items"`
		Limit    int         `json:"limit"`
		Next     string      `json:"next"`
		Offset   int         `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int         `json:"total"`
	} `json:"albums"`
}
