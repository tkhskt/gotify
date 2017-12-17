[![Codacy Badge](https://api.codacy.com/project/badge/Grade/0951a711ac0a4f5fa9309cfdf41d8e9d)](https://www.codacy.com/app/gericass/gotify?utm_source=github.com&utm_medium=referral&utm_content=gericass/gotify&utm_campaign=badger)
# gotify [![Go Report Card](https://goreportcard.com/badge/github.com/gericass/gotify)](https://goreportcard.com/report/github.com/gericass/gotify)

gotify is the wrapper library for [Spotify API](https://developer.spotify.com/web-api/)

## Requirements

- Go 1.9 or later

## Supported Authentication Flow

gotify supported [Authorization Code Flow](https://developer.spotify.com/web-api/authorization-guide/#authorization_code_flow)


## Supported Endpoint

*Endpoints that are not yet supported for optional parameters are also planned to be in order*

### albums

| Endpoint                             | Struct Name               | Optional param support |
|--------------------------------------|---------------------------|------------------------|
| /v1/albums?ids={ids}                 | Albums                    | ❌                     |
| /v1/albums/{id}/tracks               | AlbumsTracks              | ❌                     |


### artists

| Endpoint                             | Struct Name               | Optional param support |
|--------------------------------------|---------------------------|------------------------|
| /v1/artists?ids={ids}                | Artists                   | no option              |
| /v1/artists/{id}/albums              | ArtistsAlbums             | ✅                     |
| /v1/artists/{id}/top-tracks          | ArtistsTopTracks          | no option              |
| /v1/artists/{id}/related-artists     | ArtistsRelatedArtists     | no option              |

### browse

| Endpoint                             | Struct Name               | Optional param support |
|--------------------------------------|---------------------------|------------------------|
| /v1/browse/featured-playlists        | BrowseFeaturedPlaylists   | ❌                     |
| /v1/browse/new-releases              | BrowseNewReleases         | ❌                     |
| /v1/browse/categories                | BrowseCategories          | ❌                     |
| /v1/browse/categories/{id}           | BrowseCategory            | ❌                     |
| /v1/browse/categories/{id}/playlists | BrowseCategorysPlaylists  | ❌                     |
| /v1/recommendations                  | Recomendations            | ❌                     |

### following

| Endpoint                             | Struct Name               | Optional param support |
|--------------------------------------|---------------------------|------------------------|
| /v1/me/following?type=artist         | FollowingArtists          | ❌                     |


## Usage

1. Get the `client ID` and `client secret` of your application

2. Run `go get github.com/gericass/gotify`

## Sample

Please see [here](https://github.com/gericass/gotifySample) for samples

```go
package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/gericass/gotify"
)

var Auth *gotify.Client
var Token *gotify.Tokens

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
```