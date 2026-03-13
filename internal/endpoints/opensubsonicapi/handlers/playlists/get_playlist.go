package playlists

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/browsing"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/middleware"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

func mapPlaylistWithEntries(pl smmodels.Playlist, tracks []smmodels.Track) osmodels.PlaylistWithEntries {
	date, _ := pl.GetLastUpdatedTime()
	entries := make([]osmodels.Song, 0, len(tracks))
	for i := range tracks {
		entries = append(entries, browsing.MapTrackToChild(&tracks[i]))
	}

	songCount := pl.Count
	if songCount == 0 {
		songCount = len(entries)
	}

	coverArt := pl.Image
	if coverArt == "None" {
		coverArt = ""
	}

	return osmodels.PlaylistWithEntries{
		Playlist: osmodels.Playlist{
			ID:        strconv.Itoa(pl.ID),
			Name:      pl.Name,
			Owner:     "",
			Public:    false,
			SongCount: songCount,
			Duration:  pl.Duration,
			Created:   date,
			Changed:   date,
			CoverArt:  coverArt,
		},
		Entry: entries,
	}
}

// https://opensubsonic.netlify.app/docs/endpoints/getplaylist/
func (h *PlaylistsHandler) GetPlaylist(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return middleware.RequiredParametrIsMissing
	}

	playlist, err := h.GetAuthedClient(c).Playlist(id, true, 0, 0)
	if err != nil {
		return err
	}

	return utils.RenderResponse(c, "playlist", mapPlaylistWithEntries(playlist.Info, playlist.Tracks))
}
