package system

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
)

// GetOpenSubsonicExtensions List the OpenSubsonic extensions supported by this server.
//
// https://opensubsonic.netlify.app/docs/endpoints/getopensubsonicextensions/
func (h *SystemHandler) GetOpenSubsonicExtensions(c echo.Context) error {
	return utils.RenderResponse(c, "openSubsonicExtensions", []models.OpenSubsonicExtension{})
}
