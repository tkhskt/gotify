package gotify

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gericass/gotify/constants"
	"github.com/gericass/gotify/models"
)

type (
	// Client : basic client data
	Client struct {
		ClientID     string
		ClientSecret string
		CallbackURI  string
	}

	// Tokens : tokens for using API
	Tokens struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		Scope        string `json:"scope"`
		ExpiresIn    int    `json:"expire_in"`
		RefreshToken string `json:"refresh_token"`
		client       *Client
	}

	// OAuth : describe method for OAuth
	OAuth interface {
		AuthURL() string
		GetToken(*http.Request) (Gotify, error)
	}

	// Gotify : interface for each endpoint of Spotify API
	Gotify interface {
		// method for token
		Refresh() error
		// albums
		GetAlbums(albumIDs []string) (*models.Albums, error)
		GetAlbumsTracks(albumID string) (*models.AlbumsTracks, error)
		GetArtists(artistIDs []string) (*models.Artists, error)
		GetArtistsAlbums(artistID string) (*models.ArtistsAlbums, error)
		GetArtistsTopTracks(artistID string, country string) (*models.ArtistsTopTracks, error)
		GetArtistsRelatedArtists(artistID string) (*models.ArtistsRelatedArtists, error)
		// browse
		GetBrowseFeaturedPlaylists() (*models.BrowseFeaturedPlaylists, error)
		GetBrowseNewReleases() (*models.BrowseNewReleases, error)
		GetBrowseCategories() (*models.BrowseCategories, error)
		GetBrowseCategory(categoryID string) (*models.BrowseCategory, error)
		GetBrowseCategorysPlaylists(categoryID string) (*models.BrowseCategorysPlaylists, error)
		// recommendations
		GetRecommendations() (*models.Recommendations, error)
		// following
		GetFollowingArtists() (*models.FollowingArtists, error)
		FollowArtistsOrUsers(followType string, IDs []string) error
		UnfollowArtistsOrUsers(unfollowType string, IDs []string) error
		CurrentFollowsArtistsOrUsers(followType string, IDs []string) (*models.CurrentFollowsArtistsOrUsers, error)
		FollowPlaylist(userID string, playlistID string) error
		UnfollowPlaylist(userID string, playlistID string) error
		CheckFollowPlaylist(ownerID string, playlistID string, userIDs []string) (*models.FollowPlaylist, error)

		// library
		SaveTracks(trackIDs []string) error
		GetUsersSavedTracks() (*models.UsersSavedTracks, error)
		RemoveUsersSavedTracks(trackIDs []string) error
		CheckUsersSavedTracks(tracksIDs []string) (*models.FollowTracks, error)
	}
)

func (c *Client) getEncodedID() string {
	str := c.ClientID + ":" + c.ClientSecret
	enc := base64.StdEncoding.EncodeToString([]byte(str))
	return enc
}

// Set : generates and returns Client object
func Set(clientID string, clientSecret string, callbackURI string) OAuth {
	removeSlash := strings.Replace(callbackURI, "/", "%2F", -1)
	callback := strings.Replace(removeSlash, ":", "%3A", -1)
	client := &Client{ClientID: clientID, ClientSecret: clientSecret, CallbackURI: callback}
	var oauth OAuth = client
	return oauth
}

// AuthURL : returns URL for authorizing app
func (c *Client) AuthURL() string {
	responseType := "code"
	and := "%20"
	redirectURL := "https://accounts.spotify.com/authorize/?client_id=" + c.ClientID + "&response_type=" + responseType + "&redirect_uri=" + c.CallbackURI + "&scope=" //"user-read-private%20user-library-read%20user-follow-read"

	scopes := []string{constants.PlaylistReadPrivate,
		constants.PlaylistReadCollaborative,
		constants.PlaylistModifyPublic,
		constants.PlaylistModifyPrivate,
		constants.Streaming,
		constants.UgcImageUpload,
		constants.UserFollowModify,
		constants.UserFollowRead,
		constants.UserLibraryRead,
		constants.UserLibraryModify,
		constants.UserReadPrivate,
		constants.UserReadBirthdate,
		constants.UserReadEmail,
		constants.UserTopRead,
		constants.UserReadPlaybackState,
		constants.UserModifyPlaybackState,
		constants.UserReadCurrentlyPlaying,
		constants.UserReadRecentlyPlayed}
	for i, v := range scopes {
		if i == 0 {
			redirectURL += v
		} else {
			redirectURL += and + v
		}
	}

	return redirectURL
}

// GetToken : returns Tokens object
func (c *Client) GetToken(r *http.Request) (Gotify, error) {

	client := &http.Client{}
	code := r.URL.Query().Get("code")

	removeSlash := strings.Replace(c.CallbackURI, "%2F", "/", -1)
	redirectUri := strings.Replace(removeSlash, "%3A", ":", -1)

	//Bodyに値を追加
	values := url.Values{}
	values.Set("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("redirect_uri", redirectUri)

	//リクエストオブジェクトを生成
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	auth := c.getEncodedID()
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	data := new(Tokens) //レスポンスのjsonをバインドする アクセストークン
	if err := json.Unmarshal(byteArray, data); err != nil {
		return nil, err
	}
	data.client = c

	var gotify Gotify = data
	return gotify, nil
}

// Refresh : refreshes AccessToken
func (t *Tokens) Refresh() error {
	client := &http.Client{}

	//Bodyに値を追加
	values := url.Values{}
	values.Set("grant_type", "refresh_token")
	values.Add("refresh_token", t.RefreshToken)

	//リクエストオブジェクトを生成
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	auth := t.client.getEncodedID()
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	data := new(Tokens) //レスポンスのjsonをバインドする アクセストークン
	if err := json.Unmarshal(byteArray, data); err != nil {
		return err
	}

	t.AccessToken = data.AccessToken
	return nil
}
