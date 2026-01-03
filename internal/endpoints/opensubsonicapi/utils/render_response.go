// Package utils provides utility functions for the OpenSubsonic API handlers.
package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/util"
)

func RenderEmptyResponse(c echo.Context) error {
	base := util.GetBaseResponse()
	response := map[string]any{
		"subsonic-response": map[string]any{
			"status":        "ok",
			"version":       base.Version,
			"type":          base.Type,
			"serverVersion": base.ServerVersion,
			"openSubsonic":  base.OpenSubsonic,
		},
	}
	return c.JSON(http.StatusOK, response)
}

func RenderResponse(c echo.Context, key string, i any) error {
	base := util.GetBaseResponse()
	response := map[string]any{
		"subsonic-response": map[string]any{
			"status":        "ok",
			"version":       base.Version,
			"type":          base.Type,
			"serverVersion": base.ServerVersion,
			"openSubsonic":  base.OpenSubsonic,
			key:             i,
		},
	}
	return c.JSON(http.StatusOK, response)
}
