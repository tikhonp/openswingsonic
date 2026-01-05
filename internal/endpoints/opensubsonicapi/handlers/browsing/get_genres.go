package browsing

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
)

// GetGenres Returns all genres.
//
// https://opensubsonic.netlify.app/docs/endpoints/getgenres/
func (h *BrowsingHandler) GetGenres(c echo.Context) error {
	return utils.RenderResponse(c, "genres", map[string]any{
		"genre": []any{},
	})
}
