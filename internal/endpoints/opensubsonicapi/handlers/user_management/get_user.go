package usermanagement

import (
	"github.com/labstack/echo/v4"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
)

func (h *SystemHandler) GetUser(c echo.Context) error {
	smUser, err := h.GetAuthedClient(c).User()
	if err != nil {
		return err
	}

	settings, err := h.GetAuthedClient(c).NotSettings()
	if err != nil {
		return err
	}
	scrobblingEnabled := settings.LastfmSessionKey != ""

	return utils.RenderResponse(c, "user", osmodels.User{
		Username:            smUser.Username,
		ScrobblingEnabled:   scrobblingEnabled,
		AdminRole:           false,
		SettingsRole:        false,
		DownloadRole:        true,
		UploadRole:          false,
		PlaylistRole:        false,
		CoverArtRole:        true,
		CommentRole:         false,
		PodcastRole:         false,
		StreamRole:          true,
		JukeboxRole:         false,
		ShareRole:           false,
		VideoConversionRole: false,
	})
}
