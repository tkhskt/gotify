package models

// FollowingArtists : the struct for GET https://api.spotify.com/v1/me/following?type=artist
type FollowingArtists struct {
	Artists struct {
		Items []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls,omitempty"`
			Followers struct {
				Href  []string `json:"href"`
				Total int      `json:"total"`
			} `json:"followers,omitempty"`
			Genres []string `json:"genres,omitempty"`
			Href   string   `json:"href,omitempty"`
			ID     string   `json:"id,omitempty"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images,omitempty"`
			Name       string `json:"name,omitempty"`
			Popularity int    `json:"popularity,omitempty"`
			Type       string `json:"type,omitempty"`
			URI        string `json:"uri,omitempty"`
		} `json:"items"`
		Next    string `json:"next"`
		Total   int    `json:"total"`
		Cursors struct {
			After string `json:"after"`
		} `json:"cursors"`
		Limit int    `json:"limit"`
		Href  string `json:"href"`
	} `json:"artists"`
}
