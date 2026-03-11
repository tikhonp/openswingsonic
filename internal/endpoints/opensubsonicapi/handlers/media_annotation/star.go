package mediaannotation

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/middleware"
)

func parseIDList(raw string) []string {
	if raw == "" {
		return nil
	}

	parts := strings.Split(raw, ",")
	ids := make([]string, 0, len(parts))
	for _, part := range parts {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			ids = append(ids, trimmed)
		}
	}
	return ids
}

type favoriteUpdater func(itemType, hash string) error

func applyFavorites(update favoriteUpdater, itemType string, ids []string) error {
	for _, id := range ids {
		if err := update(itemType, id); err != nil {
			return err
		}
	}
	return nil
}

func (h *MediaAnnotationHandler) Star(c echo.Context) error {
	trackIDs := parseIDList(c.QueryParam("id"))
	albumIDs := parseIDList(c.QueryParam("albumId"))
	artistIDs := parseIDList(c.QueryParam("artistId"))

	if len(trackIDs)+len(albumIDs)+len(artistIDs) == 0 {
		return middleware.RequiredParametrIsMissing
	}

	client := h.GetAuthedClient(c)

	if err := applyFavorites(client.AddFavorite, "track", trackIDs); err != nil {
		return err
	}
	if err := applyFavorites(client.AddFavorite, "album", albumIDs); err != nil {
		return err
	}
	if err := applyFavorites(client.AddFavorite, "artist", artistIDs); err != nil {
		return err
	}

	return utils.RenderEmptyResponse(c)
}

func (h *MediaAnnotationHandler) Unstar(c echo.Context) error {
	trackIDs := parseIDList(c.QueryParam("id"))
	albumIDs := parseIDList(c.QueryParam("albumId"))
	artistIDs := parseIDList(c.QueryParam("artistId"))

	if len(trackIDs)+len(albumIDs)+len(artistIDs) == 0 {
		return middleware.RequiredParametrIsMissing
	}

	client := h.GetAuthedClient(c)

	if err := applyFavorites(client.RemoveFavorite, "track", trackIDs); err != nil {
		return err
	}
	if err := applyFavorites(client.RemoveFavorite, "album", albumIDs); err != nil {
		return err
	}
	if err := applyFavorites(client.RemoveFavorite, "artist", artistIDs); err != nil {
		return err
	}

	return utils.RenderEmptyResponse(c)
}
