package gotify

import (
	"encoding/json"
	"fmt"

	"github.com/gericass/gotify/extensions"
	"github.com/gericass/gotify/models"
)

// GetFollowingArtists : the method for GET https://api.spotify.com/v1/me/following?type=artist
func (t *Tokens) GetFollowingArtists() (*models.FollowingArtists, error) {
	/**
	https://developer.spotify.com/web-api/get-followed-artists/
	*/

	endpoint := "https://api.spotify.com/v1/me/following?type=artist"

	res, err := extensions.GetRequest(endpoint, t.AccessToken)
	if err != nil {
		return nil, err
	}

	followingArtists := new(models.FollowingArtists)

	err = json.Unmarshal(res, followingArtists)
	if err != nil {
		return nil, err
	}
	return followingArtists, nil
}

// FollowArtistsOrUsers : the method for PUT https://api.spotify.com/v1/me/following
func (t *Tokens) FollowArtistsOrUsers(followType string, IDs []string) error {
	/**
	https://developer.spotify.com/web-api/follow-artists-users/
	*/

	endpoint := "https://api.spotify.com/v1/me/following?type=" + followType + "&ids="

	for i, v := range IDs {
		if i == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}

	res, err := extensions.PutRequest(endpoint, t.AccessToken)
	if err != nil {
		return err
	}
	if res != 204 {
		return fmt.Errorf("%d", 204)
	}
	return nil
}

// UnfollowArtistsOrUsers : the method for DELETE https://api.spotify.com/v1/me/following
func (t *Tokens) UnfollowArtistsOrUsers(unfollowType string, IDs []string) error {
	/**
	https://developer.spotify.com/web-api/follow-artists-users/
	*/

	endpoint := "https://api.spotify.com/v1/me/following?type=" + unfollowType + "&ids="

	for i, v := range IDs {
		if i == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}

	res, err := extensions.DeleteRequest(endpoint, t.AccessToken)
	if err != nil {
		return err
	}
	if res != 204 {
		return fmt.Errorf("%d", 204)
	}
	return nil
}
