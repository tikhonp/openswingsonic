// Package opensubsonicauth provides authentication parameters for middleware.
package opensubsonicauth

import (
	"github.com/labstack/echo/v4"
)

func Middleware(auth OpenSubsonicAuth) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authParams := new(authParams)
			if err := c.Bind(authParams); err != nil {
				return err
			}
			sessionKey, err := auth.Authentificate(*authParams)
			if err != nil {
				return err
			}
			c.Set("sessionKey", sessionKey)
			return next(c)
		}
	}
}
