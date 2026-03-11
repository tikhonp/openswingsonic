package mediaannotation

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/middleware"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

type ScrobbleRequest struct {
	ID         string `query:"id" form:"id" validate:"required"`
	Time       int64  `query:"time" form:"time" validate:"required"`
	Submission bool   `query:"submission" form:"submission"`
}

func (h *MediaAnnotationHandler) Scrobble(c echo.Context) error {
	var req ScrobbleRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	if req.Submission {
		client := h.GetAuthedClient(c)

		searchResult, err := client.SearchTracks(req.ID, 1)
		if err != nil {
			return err
		}
		if len(searchResult.Results) == 0 {
			return middleware.TheRequestedDataWasNotFound
		}

		track := searchResult.Results[0]
		duration := track.Duration
		if duration == 0 {
			duration = 240
		}

		timestamp := req.Time
		if timestamp > 10_000_000_000 {
			timestamp = timestamp / 1000
		}

		logReq := &smmodels.LogTrackRequest{
			Duration:  duration,
			Source:    "openswingsonic",
			Timestamp: timestamp,
			Trackhash: track.Trackhash,
		}

		if err := client.LogTrack(logReq); err != nil {
			return err
		}
	}

	return utils.RenderEmptyResponse(c)
}
