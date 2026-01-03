// Package opensubsonicapi implements endpoints for the OpenSubsonic API.
package opensubsonicapi

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/browsing"
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

	// Browsing	getIndexes getMusicDirectory getGenres getArtists getArtist getAlbum getSong getVideos getVideoInfo getArtistInfo getArtistInfo2 getAlbumInfo getAlbumInfo2 getSimilarSongs getSimilarSongs2 getTopSongs
	browsingHandler := browsing.BrowsingHandler{Handler: handler}
	protected.GET("/getMusicFolders.view", browsingHandler.GetMusicFolders)

	// Album/song lists	getAlbumList getAlbumList2 getRandomSongs getSongsByGenre getNowPlaying getStarred getStarred2

	// Searching	search search2 search3

	// Playlists	getPlaylists getPlaylist createPlaylist updatePlaylist deletePlaylist

	// Media retrieval	stream download hls getCaptions getCoverArt getLyrics getAvatar getLyricsBySongId

	// Media annotation	star unstar setRating scrobble

	// Sharing	getShares createShare updateShare deleteShare

	// Podcast	getPodcasts getNewestPodcasts refreshPodcasts createPodcastChannel deletePodcastChannel deletePodcastEpisode downloadPodcastEpisode

	// Jukebox	jukeboxControl

	// Internet radio	getInternetRadioStations createInternetRadioStation updateInternetRadioStation deleteInternetRadioStation

	// Chat	getChatMessages addChatMessage

	// User management	getUser getUsers createUser updateUser deleteUser changePassword

	// Bookmarks	getBookmarks createBookmark deleteBookmark getPlayQueue savePlayQueue

	// Media library scanning	getScanStatus startScan
}
