package models

import "time"

// https://opensubsonic.netlify.app/docs/responses/artists/

type ArtistsID3 struct {
	IgnoredArticles string     `json:"ignoredArticles"`
	Index           []IndexID3 `json:"index,omitempty"`
}

type IndexID3 struct {
	Name   string      `json:"name"`
	Artist []ArtistID3 `json:"artist,omitempty"`
}

// ArtistID3 is an artist from ID3 tags.
type ArtistID3 struct {
	// The id of the artist.
	ID string `json:"id"`
	// The artist name.
	Name string `json:"name"`
	// A covertArt id.
	CoverArt string `json:"coverArt,omitempty"`
	// An url to an external image source.
	ArtistImageURL string `json:"artistImageUrl,omitempty"`
	// Artist album count.
	AlbumCount int `json:"albumCount,omitempty"`
	// Date the artist was starred. [ISO 8601].
	Starred time.Time `json:"starred"`
	// The artist MusicBrainzID.
	MusicBrainzID string `json:"musicBrainzId"`
	// The artist sort name.
	SortName string `json:"sortName,omitempty"`
	// The list of all roles this artist has in the library.
	Roles []string `json:"roles"`
}
