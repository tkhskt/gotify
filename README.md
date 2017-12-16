# gotify

gotify is the wrapper library of [Spotify API](https://developer.spotify.com/web-api/)

## Requirements

- Go 1.9~

## Supported Authentication Flow

gotify supported [Authorization Code Flow](https://developer.spotify.com/web-api/authorization-guide/#authorization_code_flow)


## Supported Endpoint

- albums
    - `/v1/albums?ids={ids}`
    - `/v1/albums/{id}/tracks`
- artists
    - `/v1/artists?ids={ids}`
    - `/v1/artists/{id}/albums`
    - `/v1/artists/{id}/top-tracks`
    - `/v1/artists/{id}/related-artists`
- browse
    - `/v1/browse/featured-playlists`

## Usage

1. Get the `client ID` and `client secret` of your application

2. Run `go get github.com/gericass/gotify`

## Sample

Please see [here](https://github.com/gericass/gotifySample) for samples
