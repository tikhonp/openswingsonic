package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
)

// CutViewSuffix is a middleware that removes the "view" suffix from the request path if it exists.
func CutViewSuffix(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		if before, ok := strings.CutSuffix(path, ".view"); ok {
			c.Request().URL.Path = before
		}
		return next(c)
	}
}
