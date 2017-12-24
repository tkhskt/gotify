package gotify

import (
	"encoding/json"
	"strings"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
)

func encodeQuery(q string) string {
	qstring := strings.Replace(q, " ", "%20", -1)
	return qstring
}

// SearchArtist : the method for GET https://api.spotify.com/v1/search
func (t *Tokens) SearchArtist(keywords string) (*models.SearchArtists, error) {
	/**
	https://developer.spotify.com/web-api/search-item/
	*/

	query := encodeQuery(keywords)

	endpoint := "https://api.spotify.com/v1/search?q=" + query + "&type=artist"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	artists := new(models.SearchArtists)

	err = json.Unmarshal(res, artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}
