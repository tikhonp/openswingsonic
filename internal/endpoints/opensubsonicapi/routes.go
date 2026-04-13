// Package opensubsonicapi implements endpoints for the OpenSubsonic API.
package opensubsonicapi

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers"
	albumsonglists "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/album_song_lists"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/browsing"
	mediaannotation "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/media_annotation"
	medialibraryscanning "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/media_library_scanning"
	mediaretrival "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/media_retrival"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/playlists"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/search"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/system"
	usermanagement "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/user_management"
	"github.com/tikhonp/openswingsonic/internal/middleware"
	opensubsonicauth "github.com/tikhonp/openswingsonic/internal/middleware/opensubsonic_auth"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
)

func ConfigureOpenSubsonicRoutes(
	g *echo.Group,
	osauth opensubsonicauth.OpenSubsonicAuth,
	sm swingmusic.SwingMusicClient,
) {
	g.Use(middleware.ErrorHandler)

	protected := g.Group("", opensubsonicauth.Middleware(osauth))

	handler := handlers.NewHandler(sm)

	// System:
	systemHandler := system.SystemHandler{Handler: handler}
	protected.GET("/ping", systemHandler.Ping)
	protected.GET("/getLicense", systemHandler.GetLicense)
	// Note: Unlike all other APIs getOpenSubsonicExtensions must be publicly accessible.
	g.GET("/getOpenSubsonicExtensions", systemHandler.GetOpenSubsonicExtensions)
	protected.GET("/tokenInfo", systemHandler.TokenInfo)

	// Browsing: getVideos getVideoInfo getSimilarSongs getSimilarSongs2
	browsingHandler := browsing.BrowsingHandler{Handler: handler}
	protected.GET("/getMusicFolders", browsingHandler.GetMusicFolders)
	protected.GET("/getIndexes", browsingHandler.GetIndexes)
	protected.GET("/getMusicDirectory", browsingHandler.GetMusicDirectory)
	protected.GET("/getGenres", browsingHandler.GetGenres)
	protected.GET("/getArtists", browsingHandler.GetArtists)
	protected.GET("/getArtist", browsingHandler.GetArtist)
	protected.GET("/getAlbum", browsingHandler.GetAlbum)
	protected.GET("/getSong", browsingHandler.GetSong)
	protected.GET("/getArtistInfo", browsingHandler.GetArtistInfo)
	protected.GET("/getArtistInfo2", browsingHandler.GetArtistInfo)
	protected.GET("/getAlbumInfo", browsingHandler.GetAlbumInfo)
	protected.GET("/getAlbumInfo2", browsingHandler.GetAlbumInfo)
	protected.GET("/getTopSongs", browsingHandler.GetTopSongs)

	// Album/song lists: getSongsByGenre getNowPlaying
	albumSongListsHandler := albumsonglists.AlbumSongListsHandler{Handler: handler}
	protected.GET("/getAlbumList", albumSongListsHandler.GetAlbumList)
	protected.GET("/getAlbumList2", albumSongListsHandler.GetAlbumList2)
	protected.GET("/getRandomSongs", albumSongListsHandler.GetRandomSongs)
	protected.GET("/getStarred", albumSongListsHandler.GetStarred)
	protected.GET("/getStarred2", albumSongListsHandler.GetStarred2)

	// Searching:
	searchHandler := search.SearchHandler{Handler: handler}
	protected.GET("/search", searchHandler.Search)
	protected.GET("/search2", searchHandler.Search2)
	protected.GET("/search3", searchHandler.Search3)

	// Playlists: createPlaylist updatePlaylist deletePlaylist
	playlistsHandler := playlists.PlaylistsHandler{Handler: handler}
	protected.GET("/getPlaylists", playlistsHandler.GetPlaylists)
	protected.GET("/getPlaylist", playlistsHandler.GetPlaylist)

	// Media retrieval: hls getCaptions getLyrics getAvatar getLyricsBySongId
	mediaRetrivalHandler := mediaretrival.MediaRetrivalHandler{Handler: handler}
	protected.GET("/getCoverArt", mediaRetrivalHandler.GetCoverArt)
	protected.GET("/stream", mediaRetrivalHandler.Stream)
	protected.GET("/download", mediaRetrivalHandler.Stream)

	// Media annotation: setRating
	mediaAnnotationHandler := mediaannotation.MediaAnnotationHandler{Handler: handler}
	protected.GET("/star", mediaAnnotationHandler.Star)
	protected.GET("/unstar", mediaAnnotationHandler.Unstar)
	protected.GET("/scrobble", mediaAnnotationHandler.Scrobble)

	// Sharing	getShares createShare updateShare deleteShare

	// Podcast	getPodcasts getNewestPodcasts refreshPodcasts createPodcastChannel deletePodcastChannel deletePodcastEpisode downloadPodcastEpisode

	// Jukebox	jukeboxControl

	// Internet radio	getInternetRadioStations createInternetRadioStation updateInternetRadioStation deleteInternetRadioStation

	// Chat	getChatMessages addChatMessage

	// User management	getUsers createUser updateUser deleteUser changePassword
	userManagementHandler := usermanagement.SystemHandler{Handler: handler}
	protected.GET("/getUser", userManagementHandler.GetUser)

	// Bookmarks	getBookmarks createBookmark deleteBookmark getPlayQueue savePlayQueue

	// Media library scanning
	mediaLibraryScanningHandler := medialibraryscanning.MediaLibraryScanningHandler{Handler: handler}
	protected.GET("/getScanStatus", mediaLibraryScanningHandler.GetScanStatus)
	protected.GET("/startScan", mediaLibraryScanningHandler.StartScan)
}
