package gotify

import (
	"github.com/tkhskt/gotify/models"
)

type (
	Gotify struct {
		TokenInfo ITokenInfo
	}
	IGotify interface {
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
		SaveAlbums(albumIDs []string) error
		GetUsersSavedAlbums() (*models.UsersSavedAlbums, error)
		RemoveAlbumsForCurrentUser(albumIDs []string) error
		CheckUsersSavedAlbums(albumIDs []string) (*models.FollowAlbums, error)
		// personalization
		GetRecentlyPlayedTracks() (*models.RecentlyPlayedTracks, error)
		// player
		GetUsersAvailableDevices() (*models.UsersAvailableDevices, error)
		// FIXME
		// GetInformationAboutUsersCurrentPlayback() (*models.InformationAboutUsersCurrentPlayback, error)
		// FIXME
		//GetUsersCurrentlyPlayingTrack() (*models.UsersCurrentlyPlayingTrack, error)
		TransferUsersPlayback(deviceIDs []string) error
		StartResumeUsersPlayback() error
		PauseUsersPlayback() error
		SkipUsersPlaybackToNext() error
		SkipUsersPlaybackToPrevious() error
		SeekToPositionInCurrentlyPlayingTrack(position int) error
		SetRepeatModeUsersPlayback(state string) error
		SetVolumeUsersPlayback(volumePercent int) error
		ToggleShuffleUsersPlayback(state bool) error
		// search
		SearchArtists(keywords string) (*models.SearchArtists, error)
		SearchAlbums(keywords string) (*models.SearchAlbums, error)
		SearchPlaylists(keywords string) (*models.SearchPlaylists, error)
		SearchTracks(keywords string) (*models.SearchTracks, error)
		// tracks
		GetTracks(trackIDs []string) (*models.Tracks, error)
		GetAudioAnalysis(trackID string) (*models.AudioAnalysis, error)
		// profile
		GetCurrentUsersProfile() (*models.CurrentUsersProfile, error)
		GetUsersProfile(userID string) (*models.UsersProfile, error)
		// playlists
		GetUsersPlaylists(userID string) (*models.UsersPlaylists, error)
		GetCurrentUsersPlaylists() (*models.CurrentUsersProfile, error)
	}
)

func NewGotify(tokenInfo ITokenInfo) IGotify {
	return &Gotify{}
}
