package constants

const (
	// PlaylistReadPrivate : Read access to user's private playlists.
	PlaylistReadPrivate = "playlist-read-private"

	// PlaylistReadCollaborative: Include collaborative playlists when requesting a user's playlists.
	PlaylistReadCollaborative = "playlist-read-collaborative"

	// PlaylistModifyPublic :  Write access to a user's public playlists.
	PlaylistModifyPublic = "playlist-modify-public"

	// PlaylistModifyPrivate : Write access to a user's private playlists.
	PlaylistModifyPrivate = "playlist-modify-private"

	// Streaming : Control playback of a Spotify track. This scope is currently only available to Spotify native SDKs (for example, the iOS SDK and the Android SDK). The user must have a Spotify Premium account.
	Streaming = "streaming"

	// UgcImageUpload : Upload playlist cover image.
	UgcImageUpload = "ugc-image-upload"

	// UserFollowModify : Write/delete access to the list of artists and other users that the user follows.
	UserFollowModify = "user-follow-modify"

	// UserFollowRead : Read access to the list of artists and other users that the user follows.
	UserFollowRead = "user-follow-read"

	// UserLibraryRead : Read access to a user's "Your Music" library.
	UserLibraryRead = "user-library-read"

	// UserLibraryModify : Write/delete access to a user's "Your Music" library.
	UserLibraryModify = "user-library-modify"

	// UserReadPrivate : Read access to user’s subscription details (type of user account).
	UserReadPrivate = "user-read-private"

	// UserReadBirthdate : Read access to the user's birthdate.
	UserReadBirthdate = "user-read-birthdate"

	// UserReadEmail : Read access to user’s email address.
	UserReadEmail = "user-read-email"

	// UserTopRead : Read access to a user's top artists and tracks.
	UserTopRead = "user-top-read"

	// UserReadPlaybackState : Read access to a user's playback state.
	UserReadPlaybackState = "user-read-playback-state"

	// UserModifyPlaybackState : Control playback on Spotify clients and Spotify Connect devices.
	UserModifyPlaybackState = "user-modify-playback-state"

	// UserReadCurrentlyPlaying : Read access to a user's currently playing track.
	UserReadCurrentlyPlaying = "user-read-currently-playing"

	// UserReadRecentlyPlayed : Read access to a user's recently played items.
	UserReadRecentlyPlayed = "user-read-recently-played"
)
