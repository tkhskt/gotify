package models

// UsersAvailableDevices : the struct for GET https://api.spotify.com/v1/me/player/devices
type UsersAvailableDevices struct {
	Devices []struct {
		ID            string `json:"id"`
		IsActive      bool   `json:"is_active"`
		IsRestricted  bool   `json:"is_restricted"`
		Name          string `json:"name"`
		Type          string `json:"type"`
		VolumePercent int    `json:"volume_percent"`
	} `json:"devices"`
}
