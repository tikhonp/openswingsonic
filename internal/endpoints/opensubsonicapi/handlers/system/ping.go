package system

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
)

// Ping Used to test connectivity with the server.
//
// https://opensubsonic.netlify.app/docs/endpoints/ping/
func (h *SystemHandler) Ping(c echo.Context) error {
	return utils.RenderEmptyResponse(c)
}
