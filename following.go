package gotify

import (
	"github.com/tkhskt/gotify/models/follow"
)

// GetFollowingArtists : the method for GET https://api.spotify.com/v1/me/following?type=artist
func (g *Gotify) GetFollowingArtists() (*follow.MeFollowingArtists, error) {
	/**
	https://developer.spotify.com/web-api/get-followed-artists/
	*/

	endpoint := "https://api.spotify.com/v1/me/following?type=artist"

	var followingArtists *follow.MeFollowingArtists
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), *followingArtists)
	if err != nil {
		return nil, err
	}

	return followingArtists, nil
}

// FollowArtistsOrUsers : the method for PUT https://api.spotify.com/v1/me/following
func (g *Gotify) FollowArtistsOrUsers(followType string, IDs []string) error {
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

	// FIXME
	var body *interface{}
	err := g.put(endpoint, g.TokenInfo.GetAccessToken(), &body)
	if err != nil {
		return err
	}
	// if res != values.NoContent {
	// 	return fmt.Errorf("%d", 204)
	// }
	return nil
}

// UnfollowArtistsOrUsers : the method for DELETE https://api.spotify.com/v1/me/following
func (g *Gotify) UnfollowArtistsOrUsers(unfollowType string, IDs []string) error {
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

	// FIXME
	var body *interface{}
	err := g.delete(endpoint, g.TokenInfo.GetAccessToken(), &body)
	if err != nil {
		return err
	}
	// if res != values.NoContent {
	// 	return fmt.Errorf("%d", 204)
	// }
	return nil
}

// CurrentFollowsArtistsOrUsers : the method for GET https://api.spotify.com/v1/me/following/contains
func (g *Gotify) CurrentFollowsArtistsOrUsers(followType string, IDs []string) (*follow.MeFollowingContains, error) {
	/**
	https://developer.spotify.com/web-api/check-current-user-follows/
	*/

	endpoint := "https://api.spotify.com/v1/me/following/contains?type=" + followType + "&ids="

	for i, v := range IDs {
		if i == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}

	var currentFollowArtistsOrUsers *follow.MeFollowingContains
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &currentFollowArtistsOrUsers)
	if err != nil {
		return nil, err
	}

	return currentFollowArtistsOrUsers, nil
}

// FollowPlaylist : the method for PUT https://api.spotify.com/v1/users/{owner_id}/playlists/{playlist_id}/followers
func (g *Gotify) FollowPlaylist(ownerID string, playlistID string) error {
	/**
	https://developer.spotify.com/web-api/follow-playlist/
	*/

	endpoint := "https://api.spotify.com/v1/users/" + ownerID + "/playlists/" + playlistID + "/followers"

	// FIXME
	var body *interface{}
	err := g.put(endpoint, g.TokenInfo.GetAccessToken(), &body)
	if err != nil {
		return err
	}
	// if res != values.OK {
	// 	return fmt.Errorf("%d", res)
	// }
	return nil
}

// UnfollowPlaylist : the method for DELETE https://api.spotify.com/v1/users/{owner_id}/playlists/{playlist_id}/followers
func (g *Gotify) UnfollowPlaylist(ownerID string, playlistID string) error {
	/**
	https://developer.spotify.com/web-api/unfollow-playlist/
	*/

	endpoint := "https://api.spotify.com/v1/users/" + ownerID + "/playlists/" + playlistID + "/followers"
	// FIXME
	var body *interface{}
	err := g.delete(endpoint, g.TokenInfo.GetAccessToken(), &body)
	if err != nil {
		return err
	}
	// if res != values.OK {
	// 	return fmt.Errorf("%d", res)
	// }
	return nil
}

// CheckFollowPlaylist : the method for GET https://api.spotify.com/v1/users/{owner_id}/playlists/{playlist_id}/followers/contains
func (g *Gotify) CheckFollowPlaylist(ownerID string, playlistID string, userIDs []string) (*follow.PlaylistFollowersContainers, error) {
	/**
	https://developer.spotify.com/web-api/check-user-following-playlist/
	*/

	endpoint := "https://api.spotify.com/v1/users/" + ownerID + "/playlists/" + playlistID + "/followers/contains?ids="
	for i, v := range userIDs {
		if i == 0 {
			endpoint += v
		} else {
			endpoint += "," + v
		}
	}

	var followsPlaylist *follow.PlaylistFollowersContainers
	err := g.get(endpoint, g.TokenInfo.GetAccessToken(), &followsPlaylist)
	if err != nil {
		return nil, err
	}

	return followsPlaylist, nil
}
