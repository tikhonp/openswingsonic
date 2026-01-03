package system

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/middleware"
)

func (h *SystemHandler) TokenInfo(c echo.Context) error {
	return middleware.InvalidAPIKey
}
