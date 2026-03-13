package albumsonglists

import (
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/handlers/browsing"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

func mapArtistStarred2(artist *smmodels.ArtistItem, client swingmusic.SwingMusicClient) osmodels.ArtistID3 {
	starred := time.Now().UTC().Format(time.RFC3339)

	return osmodels.ArtistID3{
		ID:             artist.Artisthash,
		Name:           artist.Name,
		CoverArt:       artist.Image,
		ArtistImageURL: client.GetArtistImageURL(artist.Artisthash, swingmusic.ImageSizeLarge),
		Starred:        starred,
		SortName:       artist.Name,
		Roles:          []string{},
	}
}

func mapAlbumStarred2(album *smmodels.AlbumShortInfo) osmodels.AlbumID3 {
	starred := time.Now().UTC()

	artistName := ""
	artistID := ""
	artistNames := make([]string, 0, len(album.AlbumArtists))
	for _, a := range album.AlbumArtists {
		artistNames = append(artistNames, a.Name)
		if artistID == "" {
			artistID = a.Artisthash
			artistName = a.Name
		}
	}

	return osmodels.AlbumID3{
		ID:            album.AlbumHash,
		Parent:        album.PathHash,
		Album:         album.Title,
		Title:         album.Title,
		Name:          album.Title,
		IsDir:         true,
		CoverArt:      album.Image,
		SongCount:     0,
		Duration:      0,
		PlayCount:     0,
		Created:       album.Date.Time,
		Starred:       &starred,
		ArtistID:      artistID,
		Artist:        artistName,
		Year:          album.Date.Year(),
		Genre:         "",
		DisplayArtist: strings.Join(artistNames, ", "),
	}
}

func mapSongStarred2(track *smmodels.Track) osmodels.Song {
	starred := time.Now().UTC()

	song := browsing.MapTrackToChild(track)
	song.Starred = starred
	song.Created = starred
	song.PlayCount = 0
	song.MediaType = "audio"

	return song
}

func mapStarred2(starred *smmodels.Starred, client swingmusic.SwingMusicClient) osmodels.Starred2 {
	artists := make([]osmodels.ArtistID3, 0, len(starred.Artists))
	for _, a := range starred.Artists {
		artists = append(artists, mapArtistStarred2(&a, client))
	}

	albums := make([]osmodels.AlbumID3, 0, len(starred.Albums))
	for _, a := range starred.Albums {
		albums = append(albums, mapAlbumStarred2(&a))
	}

	songs := make([]osmodels.Song, 0, len(starred.Tracks))
	for _, s := range starred.Tracks {
		songs = append(songs, mapSongStarred2(&s))
	}

	return osmodels.Starred2{Artist: artists, Album: albums, Song: songs}
}

// GetStarred2 returns starred songs, albums and artists organized by ID3 tags.
//
// https://opensubsonic.netlify.app/docs/endpoints/getstarred2/
func (h *AlbumSongListsHandler) GetStarred2(c echo.Context) error {
	starred, err := h.GetAuthedClient(c).Favorites()
	if err != nil {
		return err
	}

	data := mapStarred2(starred, h.GetClient())
	return utils.RenderResponse(c, "starred2", data)
}
