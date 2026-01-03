package system

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
)

// GetLicense Get details about the software license.
//
// https://opensubsonic.netlify.app/docs/endpoints/getlicense/
func (h *SystemHandler) GetLicense(c echo.Context) error {
	return utils.RenderResponse(c, "license", models.License{Valid: true})
}
