package models

// User represents an OpenSubsonic user response.
type User struct {
	// Username
	Username string `json:"username"`

	// Whether scrobbling is enabled
	ScrobblingEnabled bool `json:"scrobblingEnabled"`

	// Maximum bitrate allowed for streaming
	MaxBitRate *int `json:"maxBitRate,omitempty"`

	// Roles / permissions

	// Whether the user is an admin
	AdminRole bool `json:"adminRole"`
	// Whether the user is can edit settings
	SettingsRole bool `json:"settingsRole"`
	// Whether the user can download
	DownloadRole bool `json:"downloadRole"`
	// Whether the user can upload
	UploadRole bool `json:"uploadRole"`
	// Whether the user can create playlists
	PlaylistRole bool `json:"playlistRole"`
	// Whether the user can get cover art
	CoverArtRole bool `json:"coverArtRole"`
	// Whether the user can create comments
	CommentRole bool `json:"commentRole"`
	// Whether the user can create/refresh podcasts
	PodcastRole bool `json:"podcastRole"`
	// Whether the user can stream
	StreamRole bool `json:"streamRole"`
	// Whether the user can control the jukebox
	JukeboxRole bool `json:"jukeboxRole"`
	// Whether the user can create a stream
	ShareRole bool `json:"shareRole"`
	// Whether the user can convert videos
	VideoConversionRole bool `json:"videoConversionRole"`

	// Last time avatar changed (ISO 8601)
	AvatarLastChanged *string `json:"avatarLastChanged,omitempty"`

	// Accessible folder IDs
	Folder *[]int `json:"folder,omitempty"`

	// Optional email (present in example but not required in spec)
	Email *string `json:"email,omitempty"`
}
