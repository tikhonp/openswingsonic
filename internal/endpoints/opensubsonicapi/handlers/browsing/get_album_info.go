package browsing

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/middleware"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
)

// https://opensubsonic.netlify.app/docs/endpoints/getalbuminfo/

func (h *BrowsingHandler) GetAlbumInfo(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return middleware.RequiredParametrIsMissing
	}

	// id may be albumId or songId — ignored but accepted
	info := models.AlbumInfo{
		SmallImageURL:  h.GetClient().GetAlbumImageURL(id, swingmusic.ImageSizeSmall),
		MediumImageURL: h.GetClient().GetAlbumImageURL(id, swingmusic.ImageSizeMedium),
		LargeImageURL:  h.GetClient().GetAlbumImageURL(id, swingmusic.ImageSizeLarge),
	}
	return utils.RenderResponse(c, "albumInfo", info)
}
