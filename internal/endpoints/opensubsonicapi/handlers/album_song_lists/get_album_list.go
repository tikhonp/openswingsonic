package albumsonglists

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

type AlbumListType string

const (
	AlbumListTypeRandom               AlbumListType = "random"
	AlbumListTypeNewest               AlbumListType = "newest"
	AlbumListTypeHighest              AlbumListType = "highest"
	AlbumListTypeFrequent             AlbumListType = "frequent"
	AlbumListTypeRecent               AlbumListType = "recent"
	AlbumListTypeAlphabeticalByName   AlbumListType = "alphabeticalByName"
	AlbumListTypeAlphabeticalByArtist AlbumListType = "alphabeticalByArtist"
	AlbumListTypeStarred              AlbumListType = "starred"
	AlbumListTypeByYear               AlbumListType = "byYear"
	AlbumListTypeByGenre              AlbumListType = "byGenre"
)

type GetAlbumListRquest struct {
	Type          AlbumListType `query:"type" form:"type" validate:"required,oneof=random newest highest frequent recent alphabeticalByName alphabeticalByArtist starred byYear byGenre"`
	Size          int           `query:"size" form:"size" validate:"omitempty,min=1,max=500"`
	Offset        int           `query:"offset" form:"offset" validate:"omitempty,min=0"`
	FromYear      string        `query:"fromYear" form:"fromYear" validate:"required_if=Type byYear"`
	ToYear        string        `query:"toYear" form:"toYear" validate:"required_if=Type byYear"`
	Genre         string        `query:"genre" form:"genre" validate:"required_if=Type byGenre"`
	MusicFolderID string        `query:"musicFolderId" form:"musicFolderId" validate:"omitempty"`
}

// GetAlbumList Returns a list of random, newest, highest rated etc. albums.
//
// https://opensubsonic.netlify.app/docs/endpoints/getalbumlist/
func (h *AlbumSongListsHandler) GetAlbumList(c echo.Context) error {
	albumList, err := h.FetchAlbumList(c)
	if err != nil {
		return err
	}
	return utils.RenderResponse(c, "albumList", albumList)
}

func (h *AlbumSongListsHandler) GetAlbumList2(c echo.Context) error {
	albumList, err := h.FetchAlbumList(c)
	if err != nil {
		return err
	}
	return utils.RenderResponse(c, "albumList2", albumList)
}

func (h *AlbumSongListsHandler) FetchAlbumList(c echo.Context) (*osmodels.AlbumList, error) {
	var req GetAlbumListRquest
	if err := c.Bind(&req); err != nil {
		return nil, err
	}
	if err := c.Validate(&req); err != nil {
		return nil, err
	}

	if req.Size == 0 {
		req.Size = 10
	}

	var albums []osmodels.AlbumID3
	var err error

	client := h.GetAuthedClient(c)

	switch req.Type {
	case AlbumListTypeRandom:
		albums, err = h.getRandomAlbums(&req, client)
	case AlbumListTypeNewest:
		albums, err = h.getAlbumsBySorttype(&req, client, "created_date", true)
	case AlbumListTypeHighest:
		albums, err = h.getAlbumsBySorttype(&req, client, "playcount", true)
	case AlbumListTypeFrequent:
		albums, err = h.getAlbumsBySorttype(&req, client, "playcount", true)
	case AlbumListTypeRecent:
		albums, err = h.getAlbumsBySorttype(&req, client, "lastplayed", true)
	case AlbumListTypeAlphabeticalByName:
		albums, err = h.getAlbumsBySorttype(&req, client, "title", false)
	case AlbumListTypeAlphabeticalByArtist:
		albums, err = h.getAlbumsBySorttype(&req, client, "albumartists", false)
	case "starred":
		return nil, echo.NewHTTPError(http.StatusNotImplemented, "Starred albums not implemented yet")
	case "byYear":
		albums, err = h.getAlbumsBySorttype(&req, client, "date", req.FromYear > req.ToYear)
	case "byGenre":
		return nil, echo.NewHTTPError(http.StatusNotImplemented, "Starred albums not implemented yet")
	default:
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid list type")
	}

	if err != nil {
		return nil, err
	}
	return &osmodels.AlbumList{Album: albums}, nil
}

func mapAlbumShortInfoToAlbum(album *smmodels.AlbumShortInfo, c swingmusic.SwingMusicClientAuthed) (*osmodels.AlbumID3, error) {
	albumData, err := c.Album(album.AlbumHash, 0)
	if err != nil {
		return nil, err
	}
	return mapAlbumInfoToAlbum(albumData)
}

// mapAlbumInfoToAlbum converts a SwingMusic AlbumInfo to OpenSubsonic AlbumID3
func mapAlbumInfoToAlbum(album *smmodels.AlbumResponse) (*osmodels.AlbumID3, error) {
	var artistName string
	for _, artist := range album.Info.AlbumArtists {
		if artistName != "" {
			artistName += ", "
		}
		artistName += artist.Name
	}
	var artistID string
	if len(album.Info.AlbumArtists) > 0 {
		artistID = album.Info.AlbumArtists[0].Artisthash
	}

	var starredTime *time.Time = nil
	if album.Info.IsFavorite {
		starredTime = &album.Info.CreatedDate.Time
	}

	var genre string
	for _, g := range album.Info.Genres {
		if genre != "" {
			genre += ", "
		}
		genre += g.Name
	}

	var recordLables = make([]osmodels.RecordLabel, 0, 1)
	for _, song := range album.Tracks {
		for _, label := range song.Extra.Label {
			exists := false
			for _, rl := range recordLables {
				if rl.Name == label {
					exists = true
					break
				}
			}
			if !exists {
				recordLables = append(recordLables, osmodels.RecordLabel{Name: label})
			}
		}
	}

	var genres = make([]osmodels.ItemGenre, 0, len(album.Info.Genres))
	for _, g := range album.Info.Genres {
		genres = append(genres, osmodels.ItemGenre{Name: g.Name})
	}

	var artists = make([]osmodels.ArtistID3, 0, len(album.Info.AlbumArtists))
	for _, a := range album.Info.AlbumArtists {
		artists = append(artists, osmodels.ArtistID3{
			ID:   a.Artisthash,
			Name: a.Name,
		})
	}

	var explicitStatus string
	hasExplicit := false
	hasClean := false
	for _, t := range album.Tracks {
		if t.Explicit {
			hasExplicit = true
		} else {
			hasClean = true
		}
	}
	if hasExplicit {
		explicitStatus = "explicit"
	} else if hasClean {
		explicitStatus = "clean"
	} else {
		explicitStatus = ""
	}

	var diskTitles = make([]osmodels.DiscTitle, 0)
	discsMap := make(map[int]string)
	for _, t := range album.Tracks {
		if _, exists := discsMap[int(t.Disc)]; !exists {
			discsMap[int(t.Disc)] = t.Folder
		}
	}

	osAlbum := osmodels.AlbumID3{
		ID:            album.Info.AlbumHash,
		Name:          album.Info.BaseTitle,
		Version:       strings.Join(album.Info.Versions, ", "),
		Artist:        artistName,
		ArtistID:      artistID,
		CoverArt:      album.Info.Image,
		SongCount:     album.Info.TrackCount,
		Duration:      album.Info.Duration,
		PlayCount:     album.Info.PlayCount,
		Created:       album.Info.CreatedDate.Time,
		Starred:       starredTime,
		Year:          album.Info.Date.Year(),
		Genre:         genre,
		Played:        album.Info.LastPlayed.Time,
		UserRating:    0, // SwingMusic doesn't have ratings
		RecordLabels:  recordLables,
		MusicBrainzID: "", // SwingMusic doesn't have MusicBrainzID
		Genres:        genres,
		Artists:       artists,
		DisplayArtist: artistName,
		RealeaseTypes: []string{}, // SwingMusic doesn't have release types
		Moods:         []string{}, // SwingMusic doesn't have moods
		SortName:      album.Info.BaseTitle,
		OriginalReleaseDate: osmodels.ItemDate{
			Year:  album.Info.Date.Year(),
			Month: int(album.Info.Date.Month()),
			Day:   album.Info.Date.Day(),
		},
		ReleaseDate: osmodels.ItemDate{
			Year:  album.Info.Date.Year(),
			Month: int(album.Info.Date.Month()),
			Day:   album.Info.Date.Day(),
		},
		IsCompilation:  false, // SwingMusic doesn't have compilation info
		ExplicitStatus: explicitStatus,
		DiskTitles:     diskTitles,

		// OLD FIELDS FOR COMPATIBILITY
		Title: album.Info.Title,
		Album: album.Info.Title,
		IsDir: true,
	}

	return &osAlbum, nil
}

func mapAlbums(allAlbums *smmodels.Albums, c swingmusic.SwingMusicClientAuthed) ([]osmodels.AlbumID3, error) {
	albums := make([]osmodels.AlbumID3, 0, len(allAlbums.Items))
	for _, album := range allAlbums.Items {
		fullAlbum, err := mapAlbumShortInfoToAlbum(&album, c)
		if err != nil {
			return nil, err
		}
		albums = append(albums, *fullAlbum)
	}
	return albums, nil
}

func (h *AlbumSongListsHandler) getAlbumsBySorttype(
	req *GetAlbumListRquest,
	c swingmusic.SwingMusicClientAuthed,
	sortBy string,
	reverse bool,
) ([]osmodels.AlbumID3, error) {
	// Sort by playcount descending
	allAlbums, err := c.AllAlbums(sortBy, reverse, req.Offset, req.Size)
	if err != nil {
		return nil, err
	}
	return mapAlbums(allAlbums, c)
}

func (h *AlbumSongListsHandler) getRandomAlbums(req *GetAlbumListRquest, c swingmusic.SwingMusicClientAuthed) ([]osmodels.AlbumID3, error) {
	sortTypes := []string{"duration", "created_date", "playcount", "playduration",
		"lastplayed", "trackcount", "title", "albumartists", "date"}
	randSortType := sortTypes[rand.Intn(len(sortTypes))]

	allAlbums, err := c.AllAlbums(randSortType, rand.Intn(2) == 1, 0, req.Size)
	if err != nil {
		return nil, err
	}

	rand.Shuffle(len(allAlbums.Items), func(i, j int) {
		allAlbums.Items[i], allAlbums.Items[j] = allAlbums.Items[j], allAlbums.Items[i]
	})

	return mapAlbums(allAlbums, c)
}
