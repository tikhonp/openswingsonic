package browsing

import (
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/middleware"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

func mapTrackToChild(t *smmodels.Track) osmodels.Song {
	return osmodels.Song{
		ID:           t.Trackhash,
		Parent:       t.Albumhash,
		Title:        t.Title,
		IsDir:        false,
		IsVideo:      false,
		Type:         "music",
		AlbumID:      t.Albumhash,
		Album:        t.Album,
		ArtistID:     t.Artists[0].Artisthash,
		Artist:       t.Artists[0].Name,
		CoverArt:     t.Image,
		Duration:     t.Duration,
		BitRate:      t.Bitrate,
		BitDepth:     t.Extra.Bitdepth,
		SamplingRate: t.Extra.Samplerate,
		ChannelCount: t.Extra.Channels,
		Track:        t.Track,
		Year:         0, // not available per-track
		Genres:       mapSwingGenreToGenre(t.Extra.Genre),
		Size:         t.Extra.Filesize,
		DiscNumber:   t.Disc,
		Suffix:       filepath.Ext(t.Filepath)[1:],
		ContentType:  "", // optional
		Path:         t.Filepath,
	}
}

func mapSwingGenreToGenre(in []string) []osmodels.Genre {
	out := make([]osmodels.Genre, 0, len(in))
	for _, g := range in {
		out = append(out, osmodels.Genre{Name: g})
	}
	return out
}

func mapSwingAlbumToAlbumID3WithSongs(
	resp *smmodels.AlbumResponse,
) osmodels.AlbumID3WithSongs {

	info := resp.Info

	artistName := ""
	artistID := ""
	if len(info.AlbumArtists) > 0 {
		artistName = info.AlbumArtists[0].Name
		artistID = info.AlbumArtists[0].Artisthash
	}

	out := osmodels.AlbumID3WithSongs{

		ID:       info.AlbumHash,
		Parent:   info.PathHash,
		Album:    info.Title,
		Title:    info.Title,
		Name:     info.Title,
		IsDir:    true,
		CoverArt: info.Image,

		SongCount: info.TrackCount,
		Duration:  info.Duration,
		PlayCount: info.PlayCount,

		Artist:   artistName,
		ArtistID: artistID,

		Year:  info.Date,
		Genre: firstGenreName(info.Genres),

		Created: time.Unix(info.Date, 0),

		// OpenSubsonic extensions (supported but empty)
		// Version:             "",
		// Played:              "",
		// UserRating:          0,
		// MusicBrainzID:       "",
		DisplayArtist: artistName,
		ReleaseTypes:  []string{},
		Moods:         []string{},
		// SortName:            "",
		IsCompilation:  false,
		ExplicitStatus: albumExplicitStatus(resp.Tracks),
		DiscTitles:     []osmodels.DiscTitle{},
		// RecordLabels:        []osmodels.RecordLabel{},
		Artists: mapAlbumArtists(info.AlbumArtists),
		Genres:  mapGenres(info.Genres),
		// OriginalReleaseDate: nil,
		// ReleaseDate:         nil,
	}

	songs := make([]osmodels.Song, 0, len(resp.Tracks))
	for _, t := range resp.Tracks {
		songs = append(songs,
			mapTrackToChild(&t),
		)
	}
	out.Song = songs

	return out
}

func firstGenreName(genres []smmodels.Genre) string {
	if len(genres) > 0 {
		return genres[0].Name
	}
	return ""
}

func mapGenres(genres []smmodels.Genre) []osmodels.Genre {
	out := make([]osmodels.Genre, 0, len(genres))
	for _, g := range genres {
		out = append(out, osmodels.Genre{Name: g.Name})
	}
	return out
}

func mapAlbumArtists(in []smmodels.Artist) []osmodels.ArtistFromAlbum {
	out := make([]osmodels.ArtistFromAlbum, 0, len(in))
	for _, a := range in {
		out = append(out, osmodels.ArtistFromAlbum{
			ID:   a.Artisthash,
			Name: a.Name,
		})
	}
	return out
}

func albumExplicitStatus(tracks []smmodels.Track) string {
	hasExplicit := false
	hasClean := false

	for _, t := range tracks {
		if t.Explicit {
			hasExplicit = true
		} else {
			hasClean = true
		}
	}

	if hasExplicit {
		return "explicit"
	}
	if hasClean {
		return "clean"
	}
	return ""
}

// https://opensubsonic.netlify.app/docs/endpoints/getalbum/

func (h *BrowsingHandler) GetAlbum(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return middleware.RequiredParametrIsMissing
	}

	album, err := h.GetAuthedClient(c).Album(id, 0)
	if err != nil {
		return err
	}

	out := mapSwingAlbumToAlbumID3WithSongs(album)
	return utils.RenderResponse(c, "album", out)
}
