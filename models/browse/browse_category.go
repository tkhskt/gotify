package browse

// Category : the struct for GET https://api.spotify.com/v1/browse/categories/{category_id}
type Category struct {
	Href  string `json:"href"`
	Icons []struct {
		Height int    `json:"height"`
		URL    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"icons"`
	ID   string `json:"id"`
	Name string `json:"name"`
}
