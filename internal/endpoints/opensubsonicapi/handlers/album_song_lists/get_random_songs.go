package albumsonglists

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/browsing"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

type GetRandomSongsRequest struct {
	Size          int    `query:"size" form:"size" validate:"omitempty,min=1,max=500"`
	Genre         string `query:"genre" form:"genre" validate:"omitempty"`
	FromYear      int    `query:"fromYear" form:"fromYear" validate:"omitempty"`
	ToYear        int    `query:"toYear" form:"toYear" validate:"omitempty"`
	MusicFolderID string `query:"musicFolderId" form:"musicFolderId" validate:"omitempty"`
}

// GetRandomSongs returns random songs matching the given criteria.
//
// https://opensubsonic.netlify.app/docs/endpoints/getrandomsongs/
func (h *AlbumSongListsHandler) GetRandomSongs(c echo.Context) error {
	var req GetRandomSongsRequest
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(&req); err != nil {
		return err
	}

	if req.Size == 0 {
		req.Size = 10
	}

	if req.FromYear > 0 && req.ToYear > 0 && req.FromYear > req.ToYear {
		return echo.NewHTTPError(http.StatusBadRequest, "fromYear cannot be greater than toYear")
	}

	client := h.GetAuthedClient(c)

	musicFolderPrefix, err := h.resolveMusicFolderPrefix(&req, client)
	if err != nil {
		return err
	}

	songs, err := h.pickRandomSongs(&req, client, musicFolderPrefix)
	if err != nil {
		return err
	}

	rand.Shuffle(len(songs), func(i, j int) {
		songs[i], songs[j] = songs[j], songs[i]
	})

	return utils.RenderResponse(c, "randomSongs", osmodels.RandomSongs{Song: songs})
}

func (h *AlbumSongListsHandler) resolveMusicFolderPrefix(req *GetRandomSongsRequest, client swingmusic.SwingMusicClientAuthed) (string, error) {
	if req.MusicFolderID == "" {
		return "", nil
	}

	folderIdx, err := strconv.Atoi(req.MusicFolderID)
	if err != nil || folderIdx <= 0 {
		return "", echo.NewHTTPError(http.StatusBadRequest, "invalid musicFolderId")
	}

	folders, err := client.FolderContents("$home")
	if err != nil {
		return "", err
	}

	if folderIdx > len(folders.Folders) {
		return "", echo.NewHTTPError(http.StatusBadRequest, "musicFolderId not found")
	}

	return folders.Folders[folderIdx-1].Path, nil
}

func (h *AlbumSongListsHandler) pickRandomSongs(
	req *GetRandomSongsRequest,
	client swingmusic.SwingMusicClientAuthed,
	musicFolderPrefix string,
) ([]osmodels.Song, error) {
	// Fetch a random slice of albums as a source pool for tracks.
	sortTypes := []string{"duration", "created_date", "playcount", "playduration", "lastplayed", "trackcount", "title", "albumartists", "date"}
	randSortType := sortTypes[rand.Intn(len(sortTypes))]
	albumsToFetch := max(req.Size*2, 20)

	albums, err := client.AllAlbums(randSortType, rand.Intn(2) == 1, 0, albumsToFetch)
	if err != nil {
		return nil, err
	}

	rand.Shuffle(len(albums.Items), func(i, j int) {
		albums.Items[i], albums.Items[j] = albums.Items[j], albums.Items[i]
	})

	songs := make([]osmodels.Song, 0, req.Size)
	seen := make(map[string]struct{})

	for _, album := range albums.Items {
		if len(songs) >= req.Size {
			break
		}

		fullAlbum, err := client.Album(album.AlbumHash, 0)
		if err != nil {
			return nil, err
		}

		albumYear := fullAlbum.Info.Date.Year()
		if req.FromYear > 0 && albumYear < req.FromYear {
			continue
		}
		if req.ToYear > 0 && albumYear > req.ToYear {
			continue
		}
		if req.Genre != "" && !albumHasGenre(fullAlbum, req.Genre) {
			continue
		}

		tracks := append([]smmodels.Track(nil), fullAlbum.Tracks...)
		rand.Shuffle(len(tracks), func(i, j int) {
			tracks[i], tracks[j] = tracks[j], tracks[i]
		})

		for i := range tracks {
			if len(songs) >= req.Size {
				break
			}

			track := tracks[i]
			if len(track.Artists) == 0 {
				continue
			}
			if musicFolderPrefix != "" && !strings.HasPrefix(track.Filepath, musicFolderPrefix) {
				continue
			}

			mapped := browsing.MapTrackToChild(&track)
			mapped.Year = int64(albumYear)
			mapped.MediaType = "audio"

			if _, exists := seen[mapped.ID]; exists {
				continue
			}
			seen[mapped.ID] = struct{}{}
			songs = append(songs, mapped)
		}
	}

	return songs, nil
}

func albumHasGenre(album *smmodels.AlbumResponse, genre string) bool {
	for _, g := range album.Info.Genres {
		if strings.EqualFold(g.Name, genre) {
			return true
		}
	}
	return false
}
