package gotify

import (
	"github.com/tkhskt/gotify/models/browse"
)

// GetRecommendations : the method for GET https://api.spotify.com/v1/recommendations
func (g *Gotify) GetRecommendations() (*browse.Recommendations, error) {
	/**
	https://developer.spotify.com/web-api/get-recommendations/
	*/

	endpoint := "https://api.spotify.com/v1/recommendations"

	var recommendations *browse.Recommendations
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &recommendations)
	if err != nil {
		return nil, err
	}

	return recommendations, nil
}
