package albumsonglists

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/middleware"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

// mapAlbumShortInfoToAlbum converts a SwingMusic AlbumShortInfo to OpenSubsonic AlbumID3
func mapAlbumShortInfoToAlbum(album *smmodels.AlbumShortInfo) osmodels.AlbumID3 {
	osAlbum := osmodels.AlbumID3{
		ID:       album.AlbumHash,
		Name:     album.Title,
		Title:    album.Title,
		Album:    album.Title,
		IsDir:    true,
		CoverArt: album.Image,
		Year:     album.Date,
	}

	// Set artist information
	if len(album.AlbumArtists) > 0 {
		osAlbum.Artist = album.AlbumArtists[0].Name
		osAlbum.ArtistID = album.AlbumArtists[0].Artisthash
		osAlbum.Parent = album.AlbumArtists[0].Artisthash
	}

	return osAlbum
}

// mapAlbumInfoToAlbum converts a SwingMusic AlbumInfo to OpenSubsonic AlbumID3
func mapAlbumInfoToAlbum(album *smmodels.AlbumInfo) osmodels.AlbumID3 {
	osAlbum := osmodels.AlbumID3{
		ID:        album.AlbumHash,
		Name:      album.Title,
		Title:     album.Title,
		Album:     album.Title,
		IsDir:     true,
		CoverArt:  album.Image,
		SongCount: int(album.TrackCount),
		Duration:  int(album.Duration),
		PlayCount: int(album.PlayCount),
	}

	// Convert Unix timestamp to ISO8601 format
	if album.CreatedDate > 0 {
		osAlbum.Created = time.Unix(album.CreatedDate, 0).Format(time.RFC3339)
	}

	// Extract year from Date field (assuming it's a Unix timestamp or year)
	if album.Date > 0 {
		if album.Date > 3000 {
			// Likely a Unix timestamp
			osAlbum.Year = time.Unix(album.Date, 0).Year()
		} else {
			// Likely already a year
			osAlbum.Year = int(album.Date)
		}
	}

	// Set artist information
	if len(album.AlbumArtists) > 0 {
		osAlbum.Artist = album.AlbumArtists[0].Name
		osAlbum.ArtistID = album.AlbumArtists[0].Artisthash
		osAlbum.Parent = album.AlbumArtists[0].Artisthash
	}

	// Set genre
	if len(album.Genres) > 0 {
		osAlbum.Genre = album.Genres[0].Name
	}

	return osAlbum
}

// GetAlbumList Returns a list of random, newest, highest rated etc. albums.
//
// https://opensubsonic.netlify.app/docs/endpoints/getalbumlist/
func (h *AlbumSongListsHandler) GetAlbumList(c echo.Context) error {
	listType := c.QueryParam("type")
	if listType == "" {
		return middleware.RequiredParametrIsMissing
	}

	sizeStr := c.QueryParam("size")
	size := 10 // default
	if sizeStr != "" {
		if parsed, err := strconv.Atoi(sizeStr); err == nil && parsed > 0 {
			if parsed > 500 {
				size = 500 // max limit
			} else {
				size = parsed
			}
		}
	}

	offsetStr := c.QueryParam("offset")
	offset := 0
	if offsetStr != "" {
		if parsed, err := strconv.Atoi(offsetStr); err == nil && parsed >= 0 {
			offset = parsed
		}
	}

	client := h.GetAuthedClient(c)

	var albums []osmodels.AlbumID3
	var err error

	switch listType {
	case "random":
		albums, err = h.getRandomAlbums(client, size, offset)
	case "newest":
		albums, err = h.getNewestAlbums(client, size, offset)
	case "highest":
		albums, err = h.getHighestRatedAlbums(client, size, offset)
	case "frequent":
		albums, err = h.getFrequentAlbums(client, size, offset)
	case "recent":
		albums, err = h.getRecentAlbums(client, size, offset)
	case "alphabeticalByName":
		albums, err = h.getAlphabeticalByName(client, size, offset)
	case "alphabeticalByArtist":
		albums, err = h.getAlphabeticalByArtist(client, size, offset)
	case "starred":
		albums, err = h.getStarredAlbums(client, size, offset)
	case "byYear":
		fromYear := c.QueryParam("fromYear")
		toYear := c.QueryParam("toYear")
		if fromYear == "" || toYear == "" {
			return middleware.RequiredParametrIsMissing
		}
		albums, err = h.getAlbumsByYear(client, size, offset, fromYear, toYear)
	case "byGenre":
		genre := c.QueryParam("genre")
		if genre == "" {
			return middleware.RequiredParametrIsMissing
		}
		albums, err = h.getAlbumsByGenre(client, size, offset, genre)
	default:
		return echo.NewHTTPError(400, "Invalid list type")
	}

	if err != nil {
		return err
	}

	albumList := osmodels.AlbumList{
		Album: albums,
	}

	return utils.RenderResponse(c, "albumList", albumList)
}

func (h *AlbumSongListsHandler) getRandomAlbums(client interface {
	AllAlbums(sortBy string, reverse bool) (*smmodels.Albums, error)
}, size, offset int) ([]osmodels.AlbumID3, error) {
	// Get all albums and shuffle
	allAlbums, err := client.AllAlbums("title", false)
	if err != nil {
		return nil, err
	}

	// Shuffle the albums
	rand.Shuffle(len(allAlbums.Items), func(i, j int) {
		allAlbums.Items[i], allAlbums.Items[j] = allAlbums.Items[j], allAlbums.Items[i]
	})

	// Apply offset and size
	start := offset
	end := offset + size
	if start >= len(allAlbums.Items) {
		return []osmodels.AlbumID3{}, nil
	}
	if end > len(allAlbums.Items) {
		end = len(allAlbums.Items)
	}

	albums := make([]osmodels.AlbumID3, 0, end-start)
	for i := start; i < end; i++ {
		album := mapAlbumShortInfoToAlbum(&allAlbums.Items[i])
		albums = append(albums, album)
	}

	return albums, nil
}

func (h *AlbumSongListsHandler) getNewestAlbums(client interface {
	AllAlbums(sortBy string, reverse bool) (*smmodels.Albums, error)
}, size, offset int) ([]osmodels.AlbumID3, error) {
	// Sort by created_date descending
	allAlbums, err := client.AllAlbums("created_date", true)
	if err != nil {
		return nil, err
	}

	return h.paginateAlbums(allAlbums.Items, size, offset), nil
}

func (h *AlbumSongListsHandler) getHighestRatedAlbums(client interface {
	AllAlbums(sortBy string, reverse bool) (*smmodels.Albums, error)
}, size, offset int) ([]osmodels.AlbumID3, error) {
	// SwingMusic doesn't have ratings, so we'll use playcount as a proxy
	allAlbums, err := client.AllAlbums("playcount", true)
	if err != nil {
		return nil, err
	}

	return h.paginateAlbums(allAlbums.Items, size, offset), nil
}

func (h *AlbumSongListsHandler) getFrequentAlbums(client interface {
	AllAlbums(sortBy string, reverse bool) (*smmodels.Albums, error)
}, size, offset int) ([]osmodels.AlbumID3, error) {
	// Sort by playcount descending
	allAlbums, err := client.AllAlbums("playcount", true)
	if err != nil {
		return nil, err
	}

	return h.paginateAlbums(allAlbums.Items, size, offset), nil
}

func (h *AlbumSongListsHandler) getRecentAlbums(client interface {
	AllAlbums(sortBy string, reverse bool) (*smmodels.Albums, error)
}, size, offset int) ([]osmodels.AlbumID3, error) {
	// Sort by lastplayed descending
	allAlbums, err := client.AllAlbums("lastplayed", true)
	if err != nil {
		return nil, err
	}

	return h.paginateAlbums(allAlbums.Items, size, offset), nil
}

func (h *AlbumSongListsHandler) getAlphabeticalByName(client interface {
	AllAlbums(sortBy string, reverse bool) (*smmodels.Albums, error)
}, size, offset int) ([]osmodels.AlbumID3, error) {
	// Sort by title ascending
	allAlbums, err := client.AllAlbums("title", false)
	if err != nil {
		return nil, err
	}

	return h.paginateAlbums(allAlbums.Items, size, offset), nil
}

func (h *AlbumSongListsHandler) getAlphabeticalByArtist(client interface {
	AllAlbums(sortBy string, reverse bool) (*smmodels.Albums, error)
}, size, offset int) ([]osmodels.AlbumID3, error) {
	// Sort by albumartists ascending
	allAlbums, err := client.AllAlbums("albumartists", false)
	if err != nil {
		return nil, err
	}

	return h.paginateAlbums(allAlbums.Items, size, offset), nil
}

func (h *AlbumSongListsHandler) getStarredAlbums(client interface {
	AllAlbums(sortBy string, reverse bool) (*smmodels.Albums, error)
}, size, offset int) ([]osmodels.AlbumID3, error) {
	// Get all albums and filter by favorites
	// Note: SwingMusic uses is_favorite field
	// allAlbums, err := client.AllAlbums("title", false)
	// if err != nil {
	// 	return nil, err
	// }

	// Filter starred albums - we need to get full album info
	// This is a limitation as AlbumShortInfo doesn't have is_favorite
	// We'll return empty for now, or you could fetch each album individually
	return []osmodels.AlbumID3{}, nil
}

func (h *AlbumSongListsHandler) getAlbumsByYear(client interface {
	AllAlbums(sortBy string, reverse bool) (*smmodels.Albums, error)
}, size, offset int, fromYearStr, toYearStr string) ([]osmodels.AlbumID3, error) {
	fromYear, err := strconv.Atoi(fromYearStr)
	if err != nil {
		return nil, err
	}
	toYear, err := strconv.Atoi(toYearStr)
	if err != nil {
		return nil, err
	}

	// Get all albums
	allAlbums, err := client.AllAlbums("date", fromYear > toYear)
	if err != nil {
		return nil, err
	}

	// Filter by year range
	var filtered []smmodels.AlbumShortInfo
	for _, album := range allAlbums.Items {
		year := album.Date
		if year > 3000 {
			// Unix timestamp, convert to year
			year = time.Unix(int64(year), 0).Year()
		}

		if fromYear <= toYear {
			// Normal range
			if year >= fromYear && year <= toYear {
				filtered = append(filtered, album)
			}
		} else {
			// Reverse range
			if year <= fromYear && year >= toYear {
				filtered = append(filtered, album)
			}
		}
	}

	return h.paginateAlbums(filtered, size, offset), nil
}

func (h *AlbumSongListsHandler) getAlbumsByGenre(client interface {
	AllAlbums(sortBy string, reverse bool) (*smmodels.Albums, error)
}, size, offset int, genre string) ([]osmodels.AlbumID3, error) {
	// Get all albums
	// allAlbums, err := client.AllAlbums("title", false)
	// if err != nil {
	// 	return nil, err
	// }

	// Filter by genre - note that AlbumShortInfo doesn't have genre info
	// This is a limitation of the SwingMusic API
	// You might need to fetch full album info for each album
	// For now, returning empty
	return []osmodels.AlbumID3{}, nil
}

func (h *AlbumSongListsHandler) paginateAlbums(albums []smmodels.AlbumShortInfo, size, offset int) []osmodels.AlbumID3 {
	start := offset
	end := offset + size
	if start >= len(albums) {
		return []osmodels.AlbumID3{}
	}
	if end > len(albums) {
		end = len(albums)
	}

	result := make([]osmodels.AlbumID3, 0, end-start)
	for i := start; i < end; i++ {
		album := mapAlbumShortInfoToAlbum(&albums[i])
		result = append(result, album)
	}

	return result
}
