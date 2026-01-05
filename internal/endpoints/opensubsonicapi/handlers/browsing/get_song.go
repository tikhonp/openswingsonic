package browsing

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/middleware"
)

// https://opensubsonic.netlify.app/docs/endpoints/getsong/

func (h *BrowsingHandler) GetSong(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return middleware.RequiredParametrIsMissing
	}

	// Fast path: search by trackhash
	result, err := h.GetAuthedClient(c).SearchTracks(id, 1)
	if err != nil {
		return err
	}

	if len(result.Results) == 0 {
		return middleware.TheRequestedDataWasNotFound
	}

	songResult := result.Results[0]

	song := mapTrackToChild(&songResult)
	return utils.RenderResponse(c, "song", song)
}
