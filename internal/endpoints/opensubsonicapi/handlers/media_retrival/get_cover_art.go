package mediaretrival

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/middleware"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
)

func (h *MediaRetrivalHandler) GetCoverArt(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return middleware.RequiredParametrIsMissing
	}
	client := h.GetAuthedClient(c)

	// if already valid id like de62d74ea26abc18.webp?pathhash=6a454f11d079fbdd
	if strings.Contains(id, "webp?pathhash") {
		contentType, imgReader, err := client.GetThumbnailByID(id)
		if err != nil {
			return err
		}
		return c.Stream(http.StatusOK, contentType, imgReader)
	}

	album, err := client.Album(id, 0)
	if err == nil {
		contentType, imgReader, err := client.GetThumbnailByID(album.Info.Image)
		if err != nil {
			return err
		}
		return c.Stream(http.StatusOK, contentType, imgReader)
	}

	if !errors.Is(err, swingmusic.ErrNotFound) {
		return err
	}

	artist, err := client.Artist(id)
	if err != nil {
		return err
	}

	contentType, imgReader, err := client.GetArtistImageByID(artist.Artist.Image)
	if err != nil {
		return err
	}

	return c.Stream(http.StatusOK, contentType, imgReader)
}
