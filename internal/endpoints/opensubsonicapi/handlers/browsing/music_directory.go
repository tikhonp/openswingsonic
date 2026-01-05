package browsing

import (
	"path/filepath"

	"github.com/labstack/echo/v4"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/middleware"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

func mapFolderToDirectory(folderID string, in *smmodels.Folders) osmodels.Directory {
	dir := osmodels.Directory{
		ID:   folderID,
		Name: filepath.Base(in.Path),
	}

	children := make([]osmodels.Child, 0)

	// Subfolders
	for _, f := range in.Folders {
		children = append(children, osmodels.Child{
			ID:     f.Path,
			Parent: folderID,
			IsDir:  true,
			Title:  f.Name,
		})
	}

	// Tracks
	for _, t := range in.Tracks {
		artist := ""
		if len(t.Artists) > 0 {
			artist = t.Artists[0].Name
		}

		children = append(children, osmodels.Child{
			ID:       t.Trackhash,
			Parent:   folderID,
			IsDir:    false,
			Title:    t.Title,
			Artist:   artist,
			Album:    t.Album,
			Duration: t.Duration,
			Path:     t.Filepath,
		})
	}

	dir.Child = children
	return dir
}

// GetMusicDirectory Returns a listing of all files in a music directory.
//
// https://opensubsonic.netlify.app/docs/endpoints/getmusicdirectory/
func (h *BrowsingHandler) GetMusicDirectory(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return middleware.RequiredParametrIsMissing
	}

	data, err := h.GetAuthedClient(c).FolderContents(id)
	if err != nil {
		return err
	}

	directory := mapFolderToDirectory(id, data)
	return utils.RenderResponse(c, "directory", directory)
}
