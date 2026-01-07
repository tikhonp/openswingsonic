package medialibraryscanning

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
)

// GetScanStatus returns the current status for media library scanning.
//
// https://opensubsonic.netlify.app/docs/endpoints/getscanstatus/
func (h *MediaLibraryScanningHandler) GetScanStatus(c echo.Context) error {
	return utils.RenderResponse(c, "scanStatus", map[string]bool{
		"scanning": false,
	})
}
