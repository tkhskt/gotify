package models

// Albums : the struct for GET https://api.spotify.com/v1/albums
type Albums struct {
	Albums []struct {
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
		Copyrights       []struct {
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"copyrights"`
		ExternalIds struct {
			Upc string `json:"upc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Genres []string `json:"genres"`
		Href   string   `json:"href"`
		ID     string   `json:"id"`
		Images []struct {
			Height int    `json:"height"`
			URL    string `json:"url"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name                 string `json:"name"`
		Popularity           int    `json:"popularity"`
		ReleaseDate          string `json:"release_date"`
		ReleaseDatePrecision string `json:"release_date_precision"`
		Tracks               struct {
			Href  string `json:"href"`
			Items []struct {
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
				ExternalUrls     struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls,omitempty"`
				Href        string `json:"href,omitempty"`
				ID          string `json:"id,omitempty"`
				Name        string `json:"name,omitempty"`
				PreviewURL  string `json:"preview_url,omitempty"`
				TrackNumber int    `json:"track_number,omitempty"`
				Type        string `json:"type,omitempty"`
				URI         string `json:"uri,omitempty"`
			} `json:"items"`
			Limit    int         `json:"limit"`
			Next     interface{} `json:"next"`
			Offset   int         `json:"offset"`
			Previous interface{} `json:"previous"`
			Total    int         `json:"total"`
		} `json:"tracks"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"albums"`
}
