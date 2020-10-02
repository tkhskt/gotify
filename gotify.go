package gotify

import (
	"github.com/tkhskt/gotify/models"
	"github.com/tkhskt/gotify/models/album"
	"github.com/tkhskt/gotify/models/artist"
	"github.com/tkhskt/gotify/models/browse"
	"github.com/tkhskt/gotify/models/follow"
	"github.com/tkhskt/gotify/models/library"
	"github.com/tkhskt/gotify/models/player"
	"github.com/tkhskt/gotify/models/playlist"
	"github.com/tkhskt/gotify/models/search"
	"github.com/tkhskt/gotify/models/track"
	"github.com/tkhskt/gotify/models/user"
)

type (
	Gotify struct {
		TokenInfo ITokenInfo
	}
	IGotify interface {
		// albums
		GetAlbums(albumIDs []string) (*album.Albums, error)
		GetAlbumsTracks(albumID string) (*album.Tracks, error)
		GetArtists(artistIDs []string) (*artist.Artists, error)
		GetArtistsAlbums(artistID string) (*artist.Albums, error)
		GetArtistsTopTracks(artistID string, country string) (*artist.TopTracks, error)
		GetArtistsRelatedArtists(artistID string) (*artist.RelatedArtists, error)
		// browse
		GetBrowseFeaturedPlaylists() (*browse.FeaturedPlaylists, error)
		GetBrowseNewReleases() (*browse.NewReleases, error)
		GetBrowseCategories() (*browse.Categories, error)
		GetBrowseCategory(categoryID string) (*browse.Category, error)
		GetBrowseCategorysPlaylists(categoryID string) (*browse.CategoriesPlaylists, error)
		// recommendations
		GetRecommendations() (*browse.Recommendations, error)
		// following
		GetFollowingArtists() (*follow.MeFollowingArtists, error)
		FollowArtistsOrUsers(followType string, IDs []string) error
		UnfollowArtistsOrUsers(unfollowType string, IDs []string) error
		CurrentFollowsArtistsOrUsers(followType string, IDs []string) (*follow.MeFollowingContains, error)
		FollowPlaylist(userID string, playlistID string) error
		UnfollowPlaylist(userID string, playlistID string) error
		CheckFollowPlaylist(ownerID string, playlistID string, userIDs []string) (*follow.PlaylistFollowersContainers, error)
		// library
		SaveTracks(trackIDs []string) error
		GetUsersSavedTracks() (*library.MeTracks, error)
		RemoveUsersSavedTracks(trackIDs []string) error
		CheckUsersSavedTracks(tracksIDs []string) (*models.FollowTracks, error)
		SaveAlbums(albumIDs []string) error
		GetUsersSavedAlbums() (*library.MeAlbums, error)
		RemoveAlbumsForCurrentUser(albumIDs []string) error
		CheckUsersSavedAlbums(albumIDs []string) (*library.FollowAlbums, error)
		// personalization
		GetRecentlyPlayedTracks() (*player.MePlayerRecentlyPlayed, error)
		// player
		GetUsersAvailableDevices() (*player.MePlayerDevices, error)
		// FIXME
		/// GetInformationAboutUsersCurrentPlayback() (*player.MePlayer, error)
		// FIXME
		//GetUsersCurrentlyPlayingTrack() (*player.MePlayerCurrentlyPlaying, error)
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
		SearchArtists(keywords string) (*search.Artists, error)
		SearchAlbums(keywords string) (*search.Albums, error)
		SearchPlaylists(keywords string) (*search.Playlists, error)
		SearchTracks(keywords string) (*search.Tracks, error)
		// tracks
		GetTracks(trackIDs []string) (*track.Tracks, error)
		GetAudioAnalysis(trackID string) (*track.AudioAnalysis, error)
		// profile
		GetCurrentUsersProfile() (*user.Me, error)
		GetUsersProfile(userID string) (*user.User, error)
		// playlists
		GetUsersPlaylists(userID string) (*playlist.UserPlaylists, error)
		GetCurrentUsersPlaylists() (*playlist.MePlaylists, error)
	}
)

func NewGotify(tokenInfo ITokenInfo) IGotify {
	return &Gotify{}
}
