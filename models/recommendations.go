package models

// Recommendations : the struct for GET https://api.spotify.com/v1/recommendations
type Recommendations struct {
	Tracks []struct {
		Artists []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists"`
		DiscNumber   int  `json:"disc_number"`
		DurationMs   int  `json:"duration_ms"`
		Explicit     bool `json:"explicit"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		IsPlayable  bool   `json:"is_playable"`
		Name        string `json:"name"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
	} `json:"tracks"`
	Seeds []struct {
		InitialPoolSize    int    `json:"initialPoolSize"`
		AfterFilteringSize int    `json:"afterFilteringSize"`
		AfterRelinkingSize int    `json:"afterRelinkingSize"`
		Href               string `json:"href"`
		ID                 string `json:"id"`
		Type               string `json:"type"`
	} `json:"seeds"`
}
