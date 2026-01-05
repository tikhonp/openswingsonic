package browsing

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/middleware"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
)

// https://opensubsonic.netlify.app/docs/endpoints/getartistinfo/

func (h *BrowsingHandler) GetArtistInfo(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return middleware.RequiredParametrIsMissing
	}

	// Optional params (ignored but parsed for compatibility)
	_ = c.QueryParam("count")
	_ = c.QueryParam("includeNotPresent")

	info := models.ArtistInfo{
		Biography:      "",
		MusicBrainzID:  "",
		SmallImageURL:  h.GetClient().GetArtistImageURL(id, swingmusic.ImageSizeSmall),
		MediumImageURL: h.GetClient().GetArtistImageURL(id, swingmusic.ImageSizeMedium),
		LargeImageURL:  h.GetClient().GetArtistImageURL(id, swingmusic.ImageSizeLarge),
	}
	return utils.RenderResponse(c, "artistInfo", info)
}
