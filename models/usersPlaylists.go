package models

// UsersPlaylists : the struct for GET https://api.spotify.com/v1/users/{user_id}/playlists
type UsersPlaylists struct {
	Href  string `json:"href"`
	Items []struct {
		Collaborative bool `json:"collaborative"`
		ExternalUrls  struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href   string        `json:"href"`
		ID     string        `json:"id"`
		Images []interface{} `json:"images"`
		Name   string        `json:"name"`
		Owner  struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"owner"`
		Public     bool   `json:"public"`
		SnapshotID string `json:"snapshot_id"`
		Tracks     struct {
			Href  string `json:"href"`
			Total int    `json:"total"`
		} `json:"tracks"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"items"`
	Limit    int         `json:"limit"`
	Next     interface{} `json:"next"`
	Offset   int         `json:"offset"`
	Previous interface{} `json:"previous"`
	Total    int         `json:"total"`
}
