# gotify [![Go Report Card](https://goreportcard.com/badge/github.com/gericass/gotify)](https://goreportcard.com/report/github.com/gericass/gotify)

gotify is the wrapper library of [Spotify API](https://developer.spotify.com/web-api/)

## Requirements

- Go 1.9 or later

## Supported Authentication Flow

gotify supported [Authorization Code Flow](https://developer.spotify.com/web-api/authorization-guide/#authorization_code_flow)


## Supported Endpoint

### albums

*Endpoints that are not yet supported for optional parameters are also planned to be in order*

| Endpoint                              | Struct Name              | Optional param support |
|--------------------------------------|---------------------------|------------------------|
| /v1/albums?ids={ids}                 | Albums                    | ❌                      |
| /v1/albums/{id}/tracks               | AlbumsTracks              | ❌                      |


### artists

| Endpoint                              | Struct Name              | Optional param support |
|--------------------------------------|---------------------------|------------------------|
| /v1/artists?ids={ids}                | Artists                   | no option               |
| /v1/artists/{id}/albums              | ArtistsAlbums             | ✅                      |
| /v1/artists/{id}/top-tracks          | ArtistsTopTracks          | no option               |
| /v1/artists/{id}/related-artists     | ArtistsRelatedArtists     | no option               |

### browse

| Endpoint                              | Struct Name              | Optional param support |
|--------------------------------------|---------------------------|------------------------|
| /v1/browse/featured-playlists        | BrowseFeaturedPlaylists   | ❌                      |
| /v1/browse/new-releases              | BrowseNewReleases         | ❌                      |
| /v1/browse/categories                | BrowseCategories          | ❌                      |
| /v1/browse/categories/{id}           | BrowseCategory            | ❌                      |
| /v1/browse/categories/{id}/playlists | BrowseCategoriesPlaylists | ❌                      |
| /v1/recommendations                  | Recomendations            | ❌                      |

## Usage

1. Get the `client ID` and `client secret` of your application

2. Run `go get github.com/gericass/gotify`

## Sample

Please see [here](https://github.com/gericass/gotifySample) for samples
