package browsing

import (
	"mime"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/middleware"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

func mapTrackToChild(t *smmodels.Track) osmodels.Song {
	fileExtenstion := filepath.Ext(t.Filepath)

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
		Suffix:       fileExtenstion[1:],
		ContentType:  mime.TypeByExtension(fileExtenstion),
		Path:         t.Filepath,
	}
}

func mapSwingGenreToGenre(in []string) []osmodels.ItemGenre {
	out := make([]osmodels.ItemGenre, 0, len(in))
	for _, g := range in {
		out = append(out, osmodels.ItemGenre{Name: g})
	}
	return out
}

func mapSwingAlbumToAlbumID3WithSongs(
	resp *smmodels.AlbumResponse,
) osmodels.AlbumID3WithSongs {

	info := resp.Info

	var artistName string
	for _, artist := range info.AlbumArtists {
		if artistName != "" {
			artistName += ", "
		}
		artistName += artist.Name
	}
	var artistID string
	if len(info.AlbumArtists) > 0 {
		artistID = info.AlbumArtists[0].Artisthash
	}

	var starredTime *time.Time = nil
	if info.IsFavorite {
		starredTime = &info.CreatedDate.Time
	}

	var genre string
	for _, g := range info.Genres {
		if genre != "" {
			genre += ", "
		}
		genre += g.Name
	}

	var recordLables = make([]osmodels.RecordLabel, 0, 1)
	for _, song := range resp.Tracks {
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

	var genres = make([]osmodels.ItemGenre, 0, len(info.Genres))
	for _, g := range info.Genres {
		genres = append(genres, osmodels.ItemGenre{Name: g.Name})
	}

	var artists = make([]osmodels.ArtistID3, 0, len(info.AlbumArtists))
	for _, a := range info.AlbumArtists {
		artists = append(artists, osmodels.ArtistID3{
			ID:   a.Artisthash,
			Name: a.Name,
		})
	}

	var explicitStatus string
	hasExplicit := false
	hasClean := false
	for _, t := range resp.Tracks {
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
	for _, t := range resp.Tracks {
		if _, exists := discsMap[int(t.Disc)]; !exists {
			discsMap[int(t.Disc)] = t.Folder
		}
	}

	out := osmodels.AlbumID3WithSongs{
		ID:            info.AlbumHash,
		Name:          info.BaseTitle,
		Version:       strings.Join(info.Versions, ", "),
		Artist:        artistName,
		ArtistID:      artistID,
		CoverArt:      info.Image,
		SongCount:     info.TrackCount,
		Duration:      info.Duration,
		PlayCount:     info.PlayCount,
		Created:       info.CreatedDate.Time,
		Starred:       starredTime,
		Year:          info.Date.Year(),
		Genre:         genre,
		Played:        info.LastPlayed.Time,
		UserRating:    0, // SwingMusic doesn't have ratings
		RecordLabels:  recordLables,
		MusicBrainzID: "", // SwingMusic doesn't have MusicBrainzID
		Genres:        genres,
		Artists:       artists,
		DisplayArtist: artistName,
		Moods:         []string{}, // SwingMusic doesn't have moods
		SortName:      info.BaseTitle,
		OriginalReleaseDate: osmodels.ItemDate{
			Year:  info.Date.Year(),
			Month: int(info.Date.Month()),
			Day:   info.Date.Day(),
		},
		ReleaseDate: osmodels.ItemDate{
			Year:  info.Date.Year(),
			Month: int(info.Date.Month()),
			Day:   info.Date.Day(),
		},
		IsCompilation:  false, // SwingMusic doesn't have compilation info
		ExplicitStatus: explicitStatus,
		DiscTitles:     diskTitles,

		// OLD FIELDS FOR COMPATIBILITY
		Title: info.Title,
		Album: info.Title,
		IsDir: true,
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
