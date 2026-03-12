package mediaretrival

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/middleware"
)

func (h *MediaRetrivalHandler) Stream(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return middleware.RequiredParametrIsMissing
	}

	result, err := h.GetAuthedClient(c).SearchTracks(id, 1)
	if err != nil {
		return err
	}
	if len(result.Results) == 0 {
		return middleware.TheRequestedDataWasNotFound
	}
	track := result.Results[0]

	headers, reader, err := h.GetAuthedClient(c).Stream(
		track.Trackhash,
		track.Folder,
		c.Request().Header.Get("Range"),
	)
	if err != nil {
		return err
	}

	if headers.ContentLength != 0 {
		c.Response().Header().Set("Content-Length", strconv.Itoa(headers.ContentLength))
	}
	if headers.ContentRange != "" {
		c.Response().Header().Set("Content-Range", headers.ContentRange)
	}
	if headers.ContentDisposition != "" {
		c.Response().Header().Set("Content-Disposition", headers.ContentDisposition)
	}
	return c.Stream(http.StatusOK, headers.ContentType, reader)
}
