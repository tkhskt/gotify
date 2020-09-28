package browse

// Categories : the struct for GET https://api.spotify.com/v1/browse/categories
type Categories struct {
	Categories struct {
		Href  string `json:"href"`
		Items []struct {
			Href  string `json:"href"`
			Icons []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"icons"`
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"items"`
		Limit    int         `json:"limit"`
		Next     string      `json:"next"`
		Offset   int         `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int         `json:"total"`
	} `json:"categories"`
}
