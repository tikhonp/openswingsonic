package mediaretrival

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/middleware"
)

func (h *MediaRetrivalHandler) GetCoverArt(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return middleware.RequiredParametrIsMissing
	}
	client := h.GetAuthedClient(c)
	album, err := client.Album(id, 0)
	if err != nil {
		return err
	}
	contentType, data, err := client.GetThumbnailByID(album.Info.Image)
	if err != nil {
		return err
	}
	return c.Blob(http.StatusOK, contentType, data)
}
