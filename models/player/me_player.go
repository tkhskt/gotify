package player

// MePlayer : the struct for GET https://api.spotify.com/v1/me/player
type MePlayer struct {
	Timestamp int64 `json:"timestamp"`
	Device    struct {
		ID            string `json:"id"`
		IsActive      bool   `json:"is_active"`
		IsRestricted  bool   `json:"is_restricted"`
		Name          string `json:"name"`
		Type          string `json:"type"`
		VolumePercent int    `json:"volume_percent"`
	} `json:"device"`
	ProgressMs string `json:"progress_ms"`
	IsPlaying  bool   `json:"is_playing"`
	Item       struct {
	} `json:"item"`
	ShuffleState bool   `json:"shuffle_state"`
	RepeatState  string `json:"repeat_state"`
	Context      struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"context"`
}
