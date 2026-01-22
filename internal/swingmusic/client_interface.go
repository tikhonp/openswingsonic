package swingmusic

import (
	"io"
	"net/http"

	"github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

type ImageSize string

const (
	ImageSizeSmall  ImageSize = "small"
	ImageSizeMedium ImageSize = "medium"
	ImageSizeLarge  ImageSize = "large"
)

// SwingMusicClient defines the interface for interacting with the Swing Music server.
type SwingMusicClient interface {
	// Login performs user authentication and returns the session cookie.
	Login(username, password string) (authCookie *http.Cookie, err error)

	// GetAuthed returns an authenticated client using the provided authentication cookie.
	GetAuthed(authCookie string) SwingMusicClientAuthed

	// GetAlbumCoverURL returns the URL for the album cover image of the specified size.
	GetAlbumImageURL(albumHash string, size ImageSize) string

	// GetArtistImageURL returns the URL for the artist image of the specified size.
	GetArtistImageURL(artistHash string, size ImageSize) string

	// GetThumbnailURL returns the URL for the thumbnail image with the given ID.
	GetThumbnailURL(thumbnailID string) string
}

// SwingMusicClientAuthed defines the interface for an authenticated Swing Music client.
type SwingMusicClientAuthed interface {

	// Album returns album info and tracks for the given albumhash.
	// albumLimit specifies the maximum number of albums to return in the "more from" section.
	Album(albumHash string, albumLimit int) (*models.AlbumResponse, error)

	// AlbumOtherVersions returns other versions of the given album.
	// ogAlbumTitle is the original album title (album.og_title)
	AlbumOtherVersions(albumHash, ogAlbumTitle string) (*[]models.AlbumShortInfo, error)

	// AlbumTracks returns all the tracks in the given album, sorted by disc and track number.
	// NOTE: No album info is returned.
	AlbumTracks(albumHash string) (*[]models.Track, error)

	// AllAlbums returns all albums in the library.
	// sortBy arg sort keys:
	// 		duration, created_date, playcount, playduration, lastplayed, trackcount, title, albumartists, date
	// reverse arg indicates whether to sort in descending order.
	// offset and limit args are used for pagination.
	AllAlbums(sortBy string, reverse bool, offset, limit int) (*models.Albums, error)

	// Artist returns artist data, tracks and genres for the given artisthash.
	Artist(artistHash string) (*models.ArtistResponse, error)

	// ArtistAlbums returns all albums for the given artist.
	ArtistAlbums(artistHash string) (*models.ArtistAlbumsResponse, error)

	// ArtistTracks returns all tracks for the given artist.
	ArtistTracks(artistHash string) (*[]models.Track, error)

	// AllArtists returns all artists in the library.
	// sortBy arg sort keys:
	// 		duration, created_date, playcount, playduration, lastplayed, trackcount, name, albumcount
	// reverse arg indicates whether to sort in descending order.
	AllArtists(sortBy string, reverse bool) (*models.Artists, error)

	// FolderContents returns a list of all the folders and tracks in the given folder.
	// folder arg is the path to the folder. Use "$home" to get root folder.
	FolderContents(folder string) (*models.Folders, error)

	// FolderDirBrowser returns a list of all the folders in the given folder.
	// Used when selecting root dirs.
	// folder arg is the path to the folder. Use "$home" to get root folder.
	FolderDirBrowser(folder string) (*models.DirBrowserResponse, error)

	// FolderTracksAll Gets all (or a max of 300) tracks from the given path and its subdirectories.
	// path arg is the path to the folder. Use "$home" to get root folder.
	FolderTracksAll(path string) (*models.FolderTrackResponse, error)

	// SearchTracks performs a search for tracks matching the given query.
	// limit arg specifies the maximum number of results to return.
	SearchTracks(query string, limit int) (*models.SearchedTracks, error)

	// SearchAlbums performs a search for albums matching the given query.
	// limit arg specifies the maximum number of results to return.
	SearchAlbums(query string, limit int) (*models.SearchedAlbums, error)

	// SearchArtists performs a search for artists matching the given query.
	// limit arg specifies the maximum number of results to return.
	SearchArtists(query string, limit int) (*models.SearchedArtists, error)

	// SearchAll performs a search for tracks, albums, and artists matching the given query.
	// limit arg specifies the maximum number of results to return for each type.
	SearchAll(query string, limit int) (*models.SearchedAll, error)

	// Stats returns overall user statistics.
	Stats() (*models.Stats, error)

	// TopTracks returns the top N albums played within a given duration.
	// duration arg can be "week", "month", "year", "alltime"
	// orderBy arg can be "playcount" or "playduration"
	// limit arg specifies the maximum number of results to return.
	TopTracks(duration, orderBy string, limit int) (*models.TopTracks, error)

	// TopAlbums returns the top N albums played within a given duration.
	// duration arg can be "week", "month", "year", "alltime"
	// orderBy arg can be "playcount" or "playduration"
	// limit arg specifies the maximum number of results to return.
	TopAlbums(duration, orderBy string, limit int) (*models.TopAlbums, error)

	// TopArtists returns the top N artists played within a given duration.
	// duration arg can be "week", "month", "year", "alltime"
	// orderBy arg can be "playcount" or "playduration"
	// limit arg specifies the maximum number of results to return.
	TopArtists(duration, orderBy string, limit int) (*models.TopArtists, error)

	// LogTrack logs a track play to the server.
	// use it to implement scrobbles.
	LogTrack(req *models.LogTrackRequest) error

	// Stream streams the media file located at the given filepath and with assosiated trackhash.
	// rangeHeader arg is the value of the "Range" HTTP header for partial content requests.
	Stream(trackhash, filepath, rangeHeader string) (*models.StreamedFileHeaders, io.ReadCloser, error)

	// Favorites returns all starred/favorite albums, artists, tracks and recents for the authenticated user.
	Favorites() (*models.Starred, error)

	// Playlists returns all playlists for the authenticated user.
	Playlists() (*models.Playlists, error)

	// TriggerScan triggers a media library scan on the server.
	TriggerScan() error

	// GetThumbnailByID returns the URL for the thumbnail image with the given ID.
	// returns content type, image reader and error if any.
	GetThumbnailByID(thumbnailID string) (string, io.ReadCloser, error)
}
