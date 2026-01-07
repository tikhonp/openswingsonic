package browsing

import (
	"time"

	"github.com/labstack/echo/v4"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/middleware"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

func mapAlbumShortToAlbumID3(
	album smmodels.AlbumShortInfo,
	artistID string,
) osmodels.AlbumID3 {
	artist := firstArtistName(album.AlbumArtists)
	return osmodels.AlbumID3{
		ID:       album.AlbumHash,
		Parent:   album.PathHash,
		Album:    album.Title,
		Title:    album.Title,
		Name:     album.Title,
		IsDir:    true,
		CoverArt: album.Image,
		ArtistID: artistID,
		Artist:   artist,
		Year:     album.Date,
	}
}

func firstArtistName(artists []smmodels.Artist) string {
	if len(artists) > 0 {
		return artists[0].Name
	}
	return ""
}

func mapArtistToArtistID3(
	artist *smmodels.ArtistDetail,
	c swingmusic.SwingMusicClient,
) *osmodels.ArtistWithAlbumsID3 {
	var starred string
	if artist.IsFavorite {
		starred = time.Unix(0, 0).Format(time.RFC3339)
	}
	return &osmodels.ArtistWithAlbumsID3{
		ID:             artist.ArtistHash,
		Name:           artist.Name,
		CoverArt:       artist.Image,
		AlbumCount:     artist.AlbumCount,
		ArtistImageURL: c.GetThumbnailURL(artist.Image),
		Starred:        starred,
		MusicBrainzID:  "", // SwingMusic does not have MusicBrainz integration
		SortName:       artist.Name,
		Roles:          []string{},
	}
}

func mapSwingArtistToArtistWithAlbumsID3(
	artist *smmodels.ArtistDetail,
	albums *smmodels.ArtistAlbumsResponse,
	c swingmusic.SwingMusicClient,
) *osmodels.ArtistWithAlbumsID3 {
	out := mapArtistToArtistID3(artist, c)

	albumList := make([]osmodels.AlbumID3, 0)

	for _, a := range albums.Albums {
		albumList = append(albumList,
			mapAlbumShortToAlbumID3(a, artist.ArtistHash),
		)
	}

	// Also include compilations / appearances if you want
	for _, a := range albums.Compilations {
		albumList = append(albumList,
			mapAlbumShortToAlbumID3(a, artist.ArtistHash),
		)
	}

	for _, a := range albums.SinglesAndEPs {
		albumList = append(albumList,
			mapAlbumShortToAlbumID3(a, artist.ArtistHash),
		)
	}

	out.Album = albumList
	return out
}

// https://opensubsonic.netlify.app/docs/endpoints/getartist/

func (h *BrowsingHandler) GetArtist(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return middleware.RequiredParametrIsMissing
	}

	client := h.GetAuthedClient(c)
	artistResp, err := client.Artist(id)
	if err != nil {
		return err
	}
	albumsResp, err := client.ArtistAlbums(id)
	if err != nil {
		return err
	}

	artist := mapSwingArtistToArtistWithAlbumsID3(&artistResp.Artist, albumsResp, h.GetClient())

	return utils.RenderResponse(c, "artist", artist)
}
