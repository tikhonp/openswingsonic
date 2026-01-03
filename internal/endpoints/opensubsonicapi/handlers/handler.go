// Package handlers provides HTTP request handlers for the opensubsonic API.
package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
)

type Handler struct {
	sm swingmusic.SwingMusicClient
}

func NewHandler(sm swingmusic.SwingMusicClient) *Handler {
	return &Handler{sm: sm}
}

func (h *Handler) GetAuthedClient(c echo.Context) swingmusic.SwingMusicClientAuthed {
	return h.sm.GetAuthed(c.Get("sessionKey").(string))
}
