package browsing

import (
	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

func mapModels(in *smmodels.Folders) (out models.MusicFolders) {
	out.MusicFolder = make([]models.MusicFolder, 0, len(in.Folders))
	for i, folder := range in.Folders {
		out.MusicFolder = append(out.MusicFolder, models.MusicFolder{
			ID:   int64(i + 1),
			Name: &folder.Path,
		})
	}
	return out
}

// GetMusicFolders Returns all configured top-level music folders.
//
// https://opensubsonic.netlify.app/docs/endpoints/getmusicfolders/
func (h *BrowsingHandler) GetMusicFolders(c echo.Context) error {
	data, err := h.GetAuthedClient(c).FolderContents("$home")
	if err != nil {
		return err
	}
	return utils.RenderResponse(c, "musicFolders", mapModels(data))
}
