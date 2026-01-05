package browsing

import (
	"sort"
	"strings"
	"unicode"

	"github.com/labstack/echo/v4"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	smmodels "github.com/tikhonp/openswingsonic/internal/swingmusic/models"
)

const ignoredArticles = "The An A Die Das Ein Eine Les Le La"

func mapArtistsToArtistsID3(in *smmodels.Artists) osmodels.ArtistsID3 {
	grouped := make(map[string][]osmodels.ArtistID3)

	for _, a := range in.Items {
		if a.Name == "" {
			continue
		}

		first := strings.ToUpper(string([]rune(a.Name)[0]))
		if !unicode.IsLetter([]rune(first)[0]) {
			first = "#"
		}

		grouped[first] = append(grouped[first], osmodels.ArtistID3{
			ID:         a.Artisthash,
			Name:       a.Name,
			AlbumCount: 0,
		})
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

	return osmodels.ArtistsID3{
		IgnoredArticles: ignoredArticles,
		Index:           indexes,
	}
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
	data := mapArtistsToArtistsID3(artists)
	return utils.RenderResponse(c, "artists", data)
}
