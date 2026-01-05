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
	contentType, imgReader, err := h.GetClient().GetThumbnailByID(id)
	if err != nil {
		return err
	}
	return c.Stream(http.StatusOK, contentType, imgReader)
}
