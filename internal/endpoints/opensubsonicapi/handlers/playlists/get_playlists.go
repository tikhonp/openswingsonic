package playlists

import (
	"strconv"

	"github.com/labstack/echo/v4"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

func mapPlaylistsResponse(playlists []smmodels.Playlist) []osmodels.Playlist {
	mapped := make([]osmodels.Playlist, len(playlists))
	for i, pl := range playlists {
		date, _ := pl.GetLastUpdatedTime()
		mapped[i] = osmodels.Playlist{
			ID:        strconv.Itoa(pl.ID),
			Name:      pl.Name,
			SongCount: pl.Count,
			Duration:  pl.Duration,
			Created:   date,
			Changed:   date,
		}
	}
	return mapped
}

func (h *PlaylistsHandler) GetPlaylists(c echo.Context) error {
	playlists, err := h.GetAuthedClient(c).Playlists()
	if err != nil {
		return err
	}
	return utils.RenderResponse(c, "playlists", osmodels.Playlists{
		Playlist: mapPlaylistsResponse(playlists.Playlist),
	})
}
