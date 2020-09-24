package gotify

import (
	"encoding/json"
	"github.com/tkhskt/gotify/extensions"
	"github.com/tkhskt/gotify/models"
)

// GetRecommendations : the method for GET https://api.spotify.com/v1/recommendations
func (g *Gotify) GetRecommendations() (*models.Recommendations, error) {
	/**
	https://developer.spotify.com/web-api/get-recommendations/
	*/

	endpoint := "https://api.spotify.com/v1/recommendations"

	res, err := extensions.GetRequest(endpoint, g.TokenInfo.GetAccessToken())
	if err != nil {
		return nil, err
	}

	recommendations := new(models.Recommendations)

	err = json.Unmarshal(res, recommendations)
	if err != nil {
		return nil, err
	}
	return recommendations, nil
}
