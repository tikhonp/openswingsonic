package browsing

import (
	"strconv"

	"github.com/labstack/echo/v4"
	osmodels "github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/models"
	"github.com/tikhonp/openswingsonic/internal/endpoints/opensubsonicapi/utils"
	"github.com/tikhonp/openswingsonic/internal/middleware"
)

// GetTopSongs Returns top songs for the given artist.
//
// https://opensubsonic.netlify.app/docs/endpoints/gettopsongs/
//
// NOTE: This implementation has limitations due to SwingMusic API constraints:
// - Approach 1: Search for artist by name, then get their tracks (not sorted by popularity)
// - Approach 2: Get global top tracks and filter by artist name (may miss some artists)
func (h *BrowsingHandler) GetTopSongs(c echo.Context) error {
	artistName := c.QueryParam("artist")
	if artistName == "" {
		return middleware.RequiredParametrIsMissing
	}

	countStr := c.QueryParam("count")
	count := 50 // default
	if countStr != "" {
		if parsed, err := strconv.Atoi(countStr); err == nil && parsed > 0 {
			count = parsed
		}
	}

	client := h.GetAuthedClient(c)

	// Approach 1: Search for the artist, then get all their tracks
	// This won't be sorted by "top songs" but will return the artist's tracks
	searchResults, err := client.SearchArtists(artistName, 1)
	if err != nil {
		return err
	}

	var songs []osmodels.Song

	if searchResults.More || len(searchResults.Results) > 0 {
		// Found the artist, get their tracks
		if len(searchResults.Results) > 0 {
			artistHash := searchResults.Results[0].Artisthash

			tracks, err := client.ArtistTracks(artistHash)
			if err != nil {
				return err
			}

			// Convert tracks to songs and limit by count
			maxTracks := count
			if len(*tracks) < maxTracks {
				maxTracks = len(*tracks)
			}

			for i := 0; i < maxTracks; i++ {
				track := (*tracks)[i]
				song := mapTrackToChild(&track)
				songs = append(songs, song)
			}
		}
	}

	topSongs := osmodels.TopSongs{
		Song: songs,
	}

	return utils.RenderResponse(c, "topSongs", topSongs)
}

// Alternative implementation using global top tracks
// Uncomment if you prefer this approach
/*
func (h *BrowsingHandler) GetTopSongs(c echo.Context) error {
	artistName := c.QueryParam("artist")
	if artistName == "" {
		return middleware.RequiredParametrIsMissing
	}

	countStr := c.QueryParam("count")
	count := 50 // default
	if countStr != "" {
		if parsed, err := strconv.Atoi(countStr); err == nil && parsed > 0 {
			count = parsed
		}
	}

	client := h.GetAuthedClient(c)

	// Get global top tracks and filter by artist
	// Using a larger limit to ensure we find enough matching tracks
	topTracks, err := client.TopTracks("alltime", "playcount", count*3)
	if err != nil {
		return err
	}

	var songs []osmodels.Child

	// Filter tracks by artist name
	for _, track := range topTracks.Tracks {
		if len(songs) >= count {
			break
		}

		// Check if any of the track's artists match
		artistMatch := false
		for _, artist := range track.Artists {
			if artist.Name == artistName {
				artistMatch = true
				break
			}
		}

		// Also check album artists
		if !artistMatch {
			for _, artist := range track.Albumartists {
				if artist.Name == artistName {
					artistMatch = true
					break
				}
			}
		}

		if artistMatch {
			song := mapTrackToSong(&smmodels.Track{
				Album:        track.Album,
				Albumartists: track.Albumartists,
				Albumhash:    track.Albumhash,
				Artisthashes: track.Artisthashes,
				Artists:      track.Artists,
				Bitrate:      track.Bitrate,
				Duration:     track.Duration,
				Explicit:     track.Explicit,
				Extra:        track.Extra,
				Filepath:     track.Filepath,
				Folder:       track.Folder,
				Image:        track.Image,
				IsFavorite:   track.IsFavorite,
				Title:        track.Title,
				Trackhash:    track.Trackhash,
				Weakhash:     track.Weakhash,
			})
			songs = append(songs, song)
		}
	}

	topSongs := osmodels.TopSongs{
		Song: songs,
	}

	return utils.RenderResponse(c, "topSongs", topSongs)
}
*/
