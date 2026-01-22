package albumsonglists

import (
	"mime"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

func mapArtistStarred(starred *smmodels.ArtistItem) osmodels.ArtistStared {
	return osmodels.ArtistStared{
		ID:       starred.Artisthash,
		Name:     starred.Name,
		CoverArt: starred.Image,
		Starred:  time.Now(),
	}
}

func mapAlbumStarred(starred *smmodels.AlbumShortInfo) osmodels.AlbumStarred {
	var artistID, artistName string
	if len(starred.AlbumArtists) > 0 {
		artistID = starred.AlbumArtists[0].Artisthash
		artistName = starred.AlbumArtists[0].Name
	}
	return osmodels.AlbumStarred{
		ID:       starred.AlbumHash,
		Parent:   starred.PathHash,
		Album:    starred.Title,
		Title:    starred.Title,
		Name:     starred.Title,
		IsDir:    true,
		CoverArt: starred.Image,
		Created:  starred.Date.Time,
		ArtistID: artistID,
		Artist:   artistName,
		Year:     starred.Date.Year(),
		Genre:    "", // Not available in AlbumShortInfo
	}
}

func mapSongStarred(starred *smmodels.Track) osmodels.SongStarred {
	// Implement the mapping logic here
	var artistID, artistName string
	if len(starred.Artists) > 0 {
		artistID = starred.Artists[0].Artisthash
		artistName = starred.Artists[0].Name
	}

	fileExtenstion := filepath.Ext(starred.Filepath)

	return osmodels.SongStarred{
		ID:           starred.Trackhash,
		Parent:       starred.Albumhash,
		IsDir:        false,
		Title:        starred.Title,
		Album:        starred.Album,
		Artist:       artistName,
		Track:        starred.Track,
		Year:         0, // Not available per-track
		CoverArt:     starred.Image,
		Size:         starred.Extra.Filesize,
		ContentType:  mime.TypeByExtension(fileExtenstion),
		Suffix:       fileExtenstion[1:],
		Starred:      time.Now(),
		Duration:     starred.Duration,
		BitRate:      starred.Bitrate,
		BitDepth:     starred.Extra.Bitdepth,
		SamplingRate: starred.Extra.Samplerate,
		ChannelCount: starred.Extra.Channels,
		Path:         starred.Filepath,
		DiscNumber:   starred.Disc,
		Created:      time.Now(), // Could be populated if needed
		AlbumID:      starred.Albumhash,
		ArtistID:     artistID,
		Type:         "music",
		IsVideo:      false,
	}
}

func mapStarred(starred *smmodels.Starred) osmodels.Starred {
	artists := make([]osmodels.ArtistStared, 0, len(starred.Artists))
	for _, a := range starred.Artists {
		artists = append(artists, mapArtistStarred(&a))
	}

	albums := make([]osmodels.AlbumStarred, 0, len(starred.Albums))
	for _, a := range starred.Albums {
		albums = append(albums, mapAlbumStarred(&a))
	}

	songs := make([]osmodels.SongStarred, 0, len(starred.Tracks))
	for _, s := range starred.Tracks {
		songs = append(songs, mapSongStarred(&s))
	}

	return osmodels.Starred{Artist: artists, Album: albums, Song: songs}
}

func (h *AlbumSongListsHandler) GetStarred(c echo.Context) error {
	starred, err := h.GetAuthedClient(c).Favorites()
	if err != nil {
		return err
	}
	return utils.RenderResponse(c, "starred", mapStarred(starred))
}
