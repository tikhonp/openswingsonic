// Package opensubsonicapi implements endpoints for the OpenSubsonic API.
package opensubsonicapi

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers"
	albumsonglists "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/album_song_lists"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/browsing"
	medialibraryscanning "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/media_library_scanning"
	mediaretrival "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/media_retrival"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/playlists"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/system"
	"github.com/tikhonp/openswingsonic/internal/middleware"
	opensubsonicauth "github.com/tikhonp/openswingsonic/internal/middleware/opensubsonic_auth"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
)

func ConfigureOpenSubsonicRoutes(
	g *echo.Group,
	osauth opensubsonicauth.OpenSubsonicAuth,
	sm swingmusic.SwingMusicClient,
) {
	// This function would typically set up the routes for the OpenSubsonic API.
	// Implementation details would depend on the specific web framework being used.

	g.Use(middleware.ErrorHandler)

	protected := g.Group("", opensubsonicauth.Middleware(osauth))

	handler := handlers.NewHandler(sm)

	// System
	systemHandler := system.SystemHandler{Handler: handler}
	protected.GET("/ping.view", systemHandler.Ping)
	protected.GET("/getLicense.view", systemHandler.GetLicense)
	// Note: Unlike all other APIs getOpenSubsonicExtensions must be publicly accessible.
	g.GET("/getOpenSubsonicExtensions.view", systemHandler.GetOpenSubsonicExtensions)
	protected.GET("/tokenInfo.view", systemHandler.TokenInfo)

	// Browsing	getVideos getVideoInfo getSimilarSongs getSimilarSongs2
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

	// Album/song lists	getRandomSongs getSongsByGenre getNowPlaying getStarred getStarred2
	albumSongListsHandler := albumsonglists.AlbumSongListsHandler{Handler: handler}
	protected.GET("/getAlbumList.view", albumSongListsHandler.GetAlbumList)
	protected.GET("/getAlbumList2.view", albumSongListsHandler.GetAlbumList2)

	// Searching	search search2 search3

	// Playlists	getPlaylist createPlaylist updatePlaylist deletePlaylist
	playlistsHandler := playlists.PlaylistsHandler{Handler: handler}
	protected.GET("/getPlaylists.view", playlistsHandler.GetPlaylists)

	// Media retrieval hls getCaptions getLyrics getAvatar getLyricsBySongId
	mediaRetrivalHandler := mediaretrival.MediaRetrivalHandler{Handler: handler}
	protected.GET("/getCoverArt.view", mediaRetrivalHandler.GetCoverArt)
	protected.GET("/stream.view", mediaRetrivalHandler.Stream)
	protected.GET("/download.view", mediaRetrivalHandler.Stream)

	// Media annotation	star unstar setRating scrobble

	// Sharing	getShares createShare updateShare deleteShare

	// Podcast	getPodcasts getNewestPodcasts refreshPodcasts createPodcastChannel deletePodcastChannel deletePodcastEpisode downloadPodcastEpisode

	// Jukebox	jukeboxControl

	// Internet radio	getInternetRadioStations createInternetRadioStation updateInternetRadioStation deleteInternetRadioStation

	// Chat	getChatMessages addChatMessage

	// User management	getUser getUsers createUser updateUser deleteUser changePassword

	// Bookmarks	getBookmarks createBookmark deleteBookmark getPlayQueue savePlayQueue

	// Media library scanning
	mediaLibraryScanningHandler := medialibraryscanning.MediaLibraryScanningHandler{Handler: handler}
	protected.GET("/getScanStatus.view", mediaLibraryScanningHandler.GetScanStatus)
	protected.GET("/startScan.view", mediaLibraryScanningHandler.StartScan)
}
