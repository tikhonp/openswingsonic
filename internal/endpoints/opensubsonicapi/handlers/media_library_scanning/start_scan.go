package medialibraryscanning

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
)

func (h *MediaLibraryScanningHandler) StartScan(c echo.Context) error {
	err := h.GetAuthedClient(c).TriggerScan()
	if err != nil {
		return err
	}
	return utils.RenderResponse(c, "scanStatus", map[string]any{
		"scanning": true,
		"count":    1,
	})
}
