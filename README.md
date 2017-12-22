# gotify [![Go Report Card](https://goreportcard.com/badge/github.com/gericass/gotify)](https://goreportcard.com/report/github.com/gericass/gotify) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/0951a711ac0a4f5fa9309cfdf41d8e9d)](https://www.codacy.com/app/gericass/gotify?utm_source=github.com&utm_medium=referral&utm_content=gericass/gotify&utm_campaign=badger)

gotify is the wrapper library for [Spotify API](https://developer.spotify.com/web-api/)

## Requirements

- Go 1.9 or later

## Supported Authentication Flow

gotify supported [Authorization Code Flow](https://developer.spotify.com/web-api/authorization-guide/#authorization_code_flow)


## Supported Endpoint

*Endpoints that are not yet supported for optional parameters are also planned to be in order*

### albums

| Endpoint                                 | Struct                    | Method                       | Optional param support |
|------------------------------------------|---------------------------|------------------------------|------------------------|
| GET /v1/albums?ids={ids}                 | Albums                    | GetAlbums                    | ❌                     |
| GET /v1/albums/{id}/tracks               | AlbumsTracks              | GetAlbumsTracks              | ❌                     |


### artists

| Endpoint                                 | Struct                    | Method                       | Optional param support |
|------------------------------------------|---------------------------|------------------------------|------------------------|
| GET /v1/artists?ids={ids}                | Artists                   | GetArtists                   | no option              |
| GET /v1/artists/{id}/albums              | ArtistsAlbums             | GetArtistsAlbums             | ✅                     |
| GET /v1/artists/{id}/top-tracks          | ArtistsTopTracks          | GetArtistsTopTracks          | no option              |
| GET /v1/artists/{id}/related-artists     | ArtistsRelatedArtists     | GetArtistsRelatedArtists     | no option              |

### browse

| Endpoint                                 | Struct                    | Method                       | Optional param support |
|------------------------------------------|---------------------------|------------------------------|------------------------|
| GET /v1/browse/featured-playlists        | BrowseFeaturedPlaylists   | GetBrowseFeaturedPlaylists   | ❌                     |
| GET /v1/browse/new-releases              | BrowseNewReleases         | GetBrowseNewReleases         | ❌                     |
| GET /v1/browse/categories                | BrowseCategories          | GetBrowseCategories          | ❌                     |
| GET /v1/browse/categories/{id}           | BrowseCategory            | GetBrowseCategory            | ❌                     |
| GET /v1/browse/categories/{id}/playlists | BrowseCategorysPlaylists  | GetBrowseCategorysPlaylists  | ❌                     |
| GET /v1/recommendations                  | Recommendations           | GetRecomendations            | ❌                     |

### following

| Endpoint                                 | Struct                        | Method                             | Optional param support |
|------------------------------------------|-------------------------------|------------------------------------|------------------------|
| GET /v1/me/following?type=artist         | FollowingArtists              | GetFollowingArtists                | ❌                     |
| PUT /v1/me/following                     | -                             | FollowArtistsOrUsers               | ✅                     |
| DELETE /v1/me/following                  | -                             | UnfollowArtistsOrUsers             | ✅                     |
| GET /v1/me/following/contains            | CurrentFollowsArtistsOrUsers  | GetCurrentFollowsArtistsOrUsers    | ✅                     |
| PUT /v1/users/{owner_id}/playlists/{playlist_id}/followers                  | -                            | FollowPlaylists            | ❌                     |
| DELETE /v1/users/{owner_id}/playlists/{playlist_id}/followers               | -                            | UnFollowPlaylists          | ❌                     |
| GET /v1/users/{owner_id}/playlists/{playlist_id}/followers/contains         | FollowPlaylist               | CheckFollowPlaylist        | ✅                     |


### library

| Endpoint                                 | Struct                        | Method                             | Optional param support |
|------------------------------------------|-------------------------------|------------------------------------|------------------------|
| PUT /v1/me/tracks                        | -                             | SaveTracks                         | ✅                     |
| GET /v1/me/tracks                        | UsersSavedTracks              | GetUsersSavedTracks                | ❌                     |
| DELETE /v1/me/tracks                     | -                             | RemoveUsersSavedTracks             | ✅                     |
| GET /v1/me/tracks/contains               | FollowTracks                  | CheckUsersSavedTracks              | no option              |
| PUT /v1/me/albums?ids={ids}              | -                             | SaveAlbums                         | ✅                     |
| GET /v1/me/albums                        | UsersSavedAlbums              | GetUsersSavedAlbums                | ❌                     |
| DELETE /v1/me/albums?ids={ids}           | -                             | RemoveAlbumsForCurrentUser         | ✅                     |
| GET /v1/me/albums/contains               | FollowAlbums                  | CheckUsersSavedAlbums              | no option              |

### personalization

| Endpoint                                 | Struct                        | Method                             | Optional param support |
|------------------------------------------|-------------------------------|------------------------------------|------------------------|
| GET /v1/me/player/recently-played        | RecentlyPlayedTracks          | GetRecentlyPlayedTracks            | ❌                     |

### player

| Endpoint                                 | Struct                        | Method                             | Optional param support |
|------------------------------------------|-------------------------------|------------------------------------|------------------------|
| GET /v1/me/player/devices                | UsersAvailableDevices         | GetUsersAvailableDevices           | ❌                     |
| PUT /v1/me/player                        | -                             | TransferUsersPlayback              | ❌                     |
| PUT /v1/me/player/play                   | -                             | StartResumeUsersPlayback           | ❌                     |
| PUT /v1/me/player/pause                  | -                             | PauseUsersPlayback                 | ❌                     |
| POST /v1/me/player/next                  | -                             | SkipUsersPlaybackToNext            | ❌                     |
| POST /v1/me/player/previous              | -                             | SkipUsersPlaybackToPrevious        | ❌                     |

## Installation

1. Get the `client ID` and `client secret` of your application

2. Run `go get github.com/gericass/gotify`

## Usage

```go
package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/gericass/gotify"
)

var Auth gotify.OAuth
var Token gotify.Gotify

// Handler : Controller for https://localhost:3000/
func Handler(c echo.Context) error {
	Auth = gotify.Set(clientID, clientSecret, callbackURI) // Set and get the basic data for using Spotify API
	url := Auth.AuthURL() // Get the Redirect URL for authenticate
	return c.Redirect(301, url)
}

// CallbackHandler : Controller for https://localhost:3000/callback/
func CallbackHandler(c echo.Context) error {

	t, err := Auth.GetToken(c.Request()) // Get the token for using Spotify API
	if err != nil {
		return err
	}
	Token = t

	return c.String(http.StatusOK, "Authentication success")
}

// RefreshHandler : Controller for https://localhost:3000/refresh/
func RefreshHandler(c echo.Context) error {

	err := Token.Refresh() // Refreshing token for using Spotify API
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "AccessToken Refreshed")
}
```
## Sample

Please see [here](https://github.com/gericass/gotifySample) for samples