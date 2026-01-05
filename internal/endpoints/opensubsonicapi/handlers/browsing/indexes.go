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

func mapArtistsToIndexes(in *smmodels.Artists) osmodels.Indexes {
	grouped := make(map[string][]osmodels.Artist)

	for _, a := range in.Items {
		if a.Name == "" {
			continue
		}

		first := strings.ToUpper(string([]rune(a.Name)[0]))
		if !unicode.IsLetter([]rune(first)[0]) {
			first = "#"
		}

		grouped[first] = append(grouped[first], osmodels.Artist{
			ID:   a.Artisthash,
			Name: a.Name,
		})
	}

	indexes := make([]osmodels.Index, 0, len(grouped))

	for letter, artists := range grouped {
		sort.Slice(artists, func(i, j int) bool {
			return artists[i].Name < artists[j].Name
		})

		indexes = append(indexes, osmodels.Index{
			Name:   letter,
			Artist: artists,
		})
	}

	sort.Slice(indexes, func(i, j int) bool {
		return indexes[i].Name < indexes[j].Name
	})

	return osmodels.Indexes{
		Index: indexes,
	}
}

// GetIndexes Returns an indexed structure of all artists.
//
// https://opensubsonic.netlify.app/docs/endpoints/getindexes/
//
// TODO: implement full api
//
//	You only need additional endpoints if you later want:
//
// child entries (songs inside indexes) → not required
// shortcut folders → needs folder metadata
// starred artists → needs favorites API
// For now: this implementation is complete and correct.
func (h *BrowsingHandler) GetIndexes(c echo.Context) error {
	// Parameters are optional and ignored for now:
	// musicFolderId
	// ifModifiedSince

	artists, err := h.GetAuthedClient(c).AllArtists("name", false)
	if err != nil {
		return err
	}

	println("Total artists fetched:", len(artists.Items))

	indexes := mapArtistsToIndexes(artists)

	return utils.RenderResponse(c, "indexes", indexes)
}
