package search

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/browsing"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/middleware"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

const defaultSearchCount = 20

func parseCountAndOffset(c echo.Context, countKey, offsetKey string, defaultCount int) (int, int) {
	count := defaultCount
	offset := 0

	if val := c.QueryParam(countKey); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil && parsed >= 0 {
			count = parsed
		}
	}

	if val := c.QueryParam(offsetKey); val != "" {
		if parsed, err := strconv.Atoi(val); err == nil && parsed >= 0 {
			offset = parsed
		}
	}

	return count, offset
}

func sliceWithOffset[T any](items []T, offset, count int) []T {
	if count <= 0 || offset >= len(items) {
		return []T{}
	}
	end := min(offset + count, len(items))
	return items[offset:end]
}

func mapSearchArtistsBasic(items []smmodels.ArtistItem, offset, count int) []osmodels.Artist {
	selected := sliceWithOffset(items, offset, count)
	out := make([]osmodels.Artist, 0, len(selected))
	for _, artist := range selected {
		out = append(out, osmodels.Artist{
			ID:   artist.Artisthash,
			Name: artist.Name,
		})
	}
	return out
}

func mapSearchArtistsID3(items []smmodels.ArtistItem, offset, count int) []osmodels.ArtistID3 {
	selected := sliceWithOffset(items, offset, count)
	out := make([]osmodels.ArtistID3, 0, len(selected))
	for _, artist := range selected {
		out = append(out, osmodels.ArtistID3{
			ID:       artist.Artisthash,
			Name:     artist.Name,
			CoverArt: artist.Image,
			SortName: artist.Name,
		})
	}
	return out
}

func mapAlbumShortInfos(items []smmodels.AlbumShortInfo, offset, count int) []osmodels.AlbumID3 {
	selected := sliceWithOffset(items, offset, count)
	out := make([]osmodels.AlbumID3, 0, len(selected))
	for _, album := range selected {
		var artistName, artistID string
		if len(album.AlbumArtists) > 0 {
			artistName = album.AlbumArtists[0].Name
			artistID = album.AlbumArtists[0].Artisthash
		}

		releaseDate := osmodels.ItemDate{
			Year:  album.Date.Year(),
			Month: int(album.Date.Month()),
			Day:   album.Date.Day(),
		}

		out = append(out, osmodels.AlbumID3{
			ID:                  album.AlbumHash,
			Name:                album.Title,
			Album:               album.Title,
			Title:               album.Title,
			Artist:              artistName,
			ArtistID:            artistID,
			CoverArt:            album.Image,
			SongCount:           0,
			Duration:            0,
			Created:             album.Date.Time,
			Year:                album.Date.Year(),
			Genres:              []osmodels.ItemGenre{},
			Artists:             []osmodels.ArtistID3{},
			DisplayArtist:       artistName,
			RealeaseTypes:       []string{},
			Moods:               []string{},
			SortName:            album.Title,
			OriginalReleaseDate: releaseDate,
			ReleaseDate:         releaseDate,
			IsCompilation:       false,
			ExplicitStatus:      "",
			DiskTitles:          []osmodels.DiscTitle{},
			IsDir:               true,
		})
	}
	return out
}

func mapTracks(tracks []smmodels.Track, offset, count int) []osmodels.Song {
	selected := sliceWithOffset(tracks, offset, count)
	out := make([]osmodels.Song, 0, len(selected))
	for i := range selected {
		out = append(out, browsing.MapTrackToChild(&selected[i]))
	}
	return out
}

func (h *SearchHandler) fetchLimits(count, offset int) int {
	limit := count + offset
	if limit == 0 {
		limit = defaultSearchCount
	}
	return limit
}

func (h *SearchHandler) Search(c echo.Context) error {
	query := c.QueryParam("query")

	artistCount, artistOffset := parseCountAndOffset(c, "artistCount", "artistOffset", defaultSearchCount)
	albumCount, albumOffset := parseCountAndOffset(c, "albumCount", "albumOffset", defaultSearchCount)
	songCount, songOffset := parseCountAndOffset(c, "songCount", "songOffset", defaultSearchCount)

	client := h.GetAuthedClient(c)

	artists, err := client.SearchArtists(query, h.fetchLimits(artistCount, artistOffset))
	if err != nil {
		return err
	}
	albums, err := client.SearchAlbums(query, h.fetchLimits(albumCount, albumOffset))
	if err != nil {
		return err
	}
	songs, err := client.SearchTracks(query, h.fetchLimits(songCount, songOffset))
	if err != nil {
		return err
	}

	result := osmodels.SearchResult{
		Artist: mapSearchArtistsBasic(artists.Results, artistOffset, artistCount),
		Album:  mapAlbumShortInfos(albums.Results, albumOffset, albumCount),
		Song:   mapTracks(songs.Results, songOffset, songCount),
	}

	return utils.RenderResponse(c, "searchResult", result)
}

func (h *SearchHandler) Search2(c echo.Context) error {
	query := c.QueryParam("query")
	if query == "" {
		return middleware.RequiredParametrIsMissing
	}

	artistCount, artistOffset := parseCountAndOffset(c, "artistCount", "artistOffset", defaultSearchCount)
	albumCount, albumOffset := parseCountAndOffset(c, "albumCount", "albumOffset", defaultSearchCount)
	songCount, songOffset := parseCountAndOffset(c, "songCount", "songOffset", defaultSearchCount)

	client := h.GetAuthedClient(c)

	artists, err := client.SearchArtists(query, h.fetchLimits(artistCount, artistOffset))
	if err != nil {
		return err
	}
	albums, err := client.SearchAlbums(query, h.fetchLimits(albumCount, albumOffset))
	if err != nil {
		return err
	}
	songs, err := client.SearchTracks(query, h.fetchLimits(songCount, songOffset))
	if err != nil {
		return err
	}

	result := osmodels.SearchResult2{
		Artist: mapSearchArtistsBasic(artists.Results, artistOffset, artistCount),
		Album:  mapAlbumShortInfos(albums.Results, albumOffset, albumCount),
		Song:   mapTracks(songs.Results, songOffset, songCount),
	}

	return utils.RenderResponse(c, "searchResult2", result)
}

func (h *SearchHandler) Search3(c echo.Context) error {
	query := c.QueryParam("query")

	artistCount, artistOffset := parseCountAndOffset(c, "artistCount", "artistOffset", defaultSearchCount)
	albumCount, albumOffset := parseCountAndOffset(c, "albumCount", "albumOffset", defaultSearchCount)
	songCount, songOffset := parseCountAndOffset(c, "songCount", "songOffset", defaultSearchCount)

	client := h.GetAuthedClient(c)

	artists, err := client.SearchArtists(query, h.fetchLimits(artistCount, artistOffset))
	if err != nil {
		return err
	}
	albums, err := client.SearchAlbums(query, h.fetchLimits(albumCount, albumOffset))
	if err != nil {
		return err
	}
	songs, err := client.SearchTracks(query, h.fetchLimits(songCount, songOffset))
	if err != nil {
		return err
	}

	result := osmodels.SearchResult3{
		Artist: mapSearchArtistsID3(artists.Results, artistOffset, artistCount),
		Album:  mapAlbumShortInfos(albums.Results, albumOffset, albumCount),
		Song:   mapTracks(songs.Results, songOffset, songCount),
	}

	return utils.RenderResponse(c, "searchResult3", result)
}
