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
	protected.GET("/ping.view", systemHandler.Ping)
	protected.GET("/getLicense.view", systemHandler.GetLicense)
	// Note: Unlike all other APIs getOpenSubsonicExtensions must be publicly accessible.
	g.GET("/getOpenSubsonicExtensions.view", systemHandler.GetOpenSubsonicExtensions)
	protected.GET("/tokenInfo.view", systemHandler.TokenInfo)

	// Browsing: getVideos getVideoInfo getSimilarSongs getSimilarSongs2
	browsingHandler := browsing.BrowsingHandler{Handler: handler}
	protected.GET("/getMusicFolders.view", browsingHandler.GetMusicFolders)
	protected.GET("/getIndexes.view", browsingHandler.GetIndexes)
	protected.GET("/getMusicDirectory.view", browsingHandler.GetMusicDirectory)
	protected.GET("/getGenres.view", browsingHandler.GetGenres)
	protected.GET("/getArtists.view", browsingHandler.GetArtists)
	protected.GET("/getArtist.view", browsingHandler.GetArtist)
	protected.GET("/getAlbum.view", browsingHandler.GetAlbum)
	protected.GET("/getSong.view", browsingHandler.GetSong)
	protected.GET("/getArtistInfo.view", browsingHandler.GetArtistInfo)
	protected.GET("/getArtistInfo2.view", browsingHandler.GetArtistInfo)
	protected.GET("/getAlbumInfo.view", browsingHandler.GetAlbumInfo)
	protected.GET("/getAlbumInfo2.view", browsingHandler.GetAlbumInfo)
	protected.GET("/getTopSongs.view", browsingHandler.GetTopSongs)

	// Album/song lists: getRandomSongs getSongsByGenre getNowPlaying
	albumSongListsHandler := albumsonglists.AlbumSongListsHandler{Handler: handler}
	protected.GET("/getAlbumList.view", albumSongListsHandler.GetAlbumList)
	protected.GET("/getAlbumList2.view", albumSongListsHandler.GetAlbumList2)
	protected.GET("/getStarred.view", albumSongListsHandler.GetStarred)
	protected.GET("/getStarred2.view", albumSongListsHandler.GetStarred2)

	// Searching:
	searchHandler := search.SearchHandler{Handler: handler}
	protected.GET("/search.view", searchHandler.Search)
	protected.GET("/search2.view", searchHandler.Search2)
	protected.GET("/search3.view", searchHandler.Search3)

	// Playlists: getPlaylist createPlaylist updatePlaylist deletePlaylist
	playlistsHandler := playlists.PlaylistsHandler{Handler: handler}
	protected.GET("/getPlaylists.view", playlistsHandler.GetPlaylists)

	// Media retrieval: hls getCaptions getLyrics getAvatar getLyricsBySongId
	mediaRetrivalHandler := mediaretrival.MediaRetrivalHandler{Handler: handler}
	protected.GET("/getCoverArt.view", mediaRetrivalHandler.GetCoverArt)
	protected.GET("/stream.view", mediaRetrivalHandler.Stream)
	protected.GET("/download.view", mediaRetrivalHandler.Stream)

	// Media annotation: setRating
	mediaAnnotationHandler := mediaannotation.MediaAnnotationHandler{Handler: handler}
	protected.GET("/star.view", mediaAnnotationHandler.Star)
	protected.GET("/unstar.view", mediaAnnotationHandler.Unstar)
	protected.GET("/scrobble.view", mediaAnnotationHandler.Scrobble)

	// Sharing	getShares createShare updateShare deleteShare

	// Podcast	getPodcasts getNewestPodcasts refreshPodcasts createPodcastChannel deletePodcastChannel deletePodcastEpisode downloadPodcastEpisode

	// Jukebox	jukeboxControl

	// Internet radio	getInternetRadioStations createInternetRadioStation updateInternetRadioStation deleteInternetRadioStation

	// Chat	getChatMessages addChatMessage

	// User management	getUsers createUser updateUser deleteUser changePassword
	userManagementHandler := usermanagement.SystemHandler{Handler: handler}
	protected.GET("/getUser.view", userManagementHandler.GetUser)

	// Bookmarks	getBookmarks createBookmark deleteBookmark getPlayQueue savePlayQueue

	// Media library scanning
	mediaLibraryScanningHandler := medialibraryscanning.MediaLibraryScanningHandler{Handler: handler}
	protected.GET("/getScanStatus.view", mediaLibraryScanningHandler.GetScanStatus)
	protected.GET("/startScan.view", mediaLibraryScanningHandler.StartScan)
}
