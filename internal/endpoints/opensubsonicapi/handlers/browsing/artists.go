package browsing

import (
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/labstack/echo/v4"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/swingmusic"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

const ignoredArticles = "The An A Die Das Ein Eine Les Le La"

func fetchArtistAlbumsCount(artistID string, c swingmusic.SwingMusicClientAuthed, cb swingmusic.SwingMusicClient) (*osmodels.ArtistID3, error) {
	artistDetail, err := c.Artist(artistID)
	if err != nil {
		return nil, err
	}
	artist := artistDetail.Artist
	var starred string
	if artist.IsFavorite {
		starred = time.Unix(0, 0).Format(time.RFC3339)
	}
	return &osmodels.ArtistID3{
		ID:             artist.ArtistHash,
		Name:           artist.Name,
		CoverArt:       artist.Image,
		AlbumCount:     artist.AlbumCount,
		ArtistImageURL: cb.GetThumbnailURL(artist.Image),
		Starred:        starred,
		MusicBrainzID:  "", // SwingMusic does not have MusicBrainz integration
		SortName:       artist.Name,
		Roles:          []string{},
	}, nil
}

func mapArtistsToArtistsID3(in *smmodels.Artists, c swingmusic.SwingMusicClientAuthed, cb swingmusic.SwingMusicClient) (*osmodels.ArtistsID3, error) {
	grouped := make(map[string][]osmodels.ArtistID3)

	for _, a := range in.Items {
		if a.Name == "" {
			continue
		}

		first := strings.ToUpper(string([]rune(a.Name)[0]))
		if !unicode.IsLetter([]rune(first)[0]) {
			first = "#"
		}

		artist, err := fetchArtistAlbumsCount(a.Artisthash, c, cb)
		if err != nil {
			return nil, err
		}

		grouped[first] = append(grouped[first], *artist)
	}

	indexes := make([]osmodels.IndexID3, 0, len(grouped))

	for letter, artists := range grouped {
		sort.Slice(artists, func(i, j int) bool {
			return artists[i].Name < artists[j].Name
		})

		indexes = append(indexes, osmodels.IndexID3{
			Name:   letter,
			Artist: artists,
		})
	}

	sort.Slice(indexes, func(i, j int) bool {
		return indexes[i].Name < indexes[j].Name
	})

	return &osmodels.ArtistsID3{
		IgnoredArticles: ignoredArticles,
		Index:           indexes,
	}, nil
}

// GetArtists Returns all artists organized according to ID3 tags.
//
// https://opensubsonic.netlify.app/docs/endpoints/getartists/
func (h *BrowsingHandler) GetArtists(c echo.Context) error {
	// musicFolderId is optional and ignored for now
	artists, err := h.GetAuthedClient(c).AllArtists("name", false)
	if err != nil {
		return err
	}
	data, err := mapArtistsToArtistsID3(artists, h.GetAuthedClient(c), h.GetClient())
	if err != nil {
		return err
	}
	return utils.RenderResponse(c, "artists", data)
}
