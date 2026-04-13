package models

import "time"

// https://opensubsonic.netlify.app/docs/responses/artists/

type ArtistsID3 struct {
	IgnoredArticles string     `json:"ignoredArticles" xml:"ignoredArticles"`
	Index           []IndexID3 `json:"index,omitempty" xml:"index,omitempty"`
}

type IndexID3 struct {
	Name   string      `json:"name" xml:"name"`
	Artist []ArtistID3 `json:"artist,omitempty" xml:"artist,omitempty"`
}

// ArtistID3 is an artist from ID3 tags.
type ArtistID3 struct {
	// The id of the artist.
	ID string `json:"id" xml:"id"`

	// The artist name.
	Name string `json:"name" xml:"name"`

	// A coverArt id.
	CoverArt string `json:"coverArt,omitempty" xml:"coverArt,omitempty"`

	// An url to an external image source.
	ArtistImageURL string `json:"artistImageUrl,omitempty" xml:"artistImageUrl,omitempty"`

	// Artist album count.
	AlbumCount int `json:"albumCount,omitempty" xml:"albumCount,omitempty"`

	// Date the artist was starred. [ISO 8601].
	Starred *time.Time `json:"starred,omitempty" xml:"starred,omitempty"`

	// The artist MusicBrainzID.
	MusicBrainzID *string `json:"musicBrainzId,omitempty" xml:"musicBrainzId,omitempty"`

	// The artist sort name.
	SortName string `json:"sortName,omitempty" xml:"sortName,omitempty"`

	// The list of all roles this artist has in the library.
	Roles []string `json:"roles,omitempty" xml:"roles,omitempty"`
}
