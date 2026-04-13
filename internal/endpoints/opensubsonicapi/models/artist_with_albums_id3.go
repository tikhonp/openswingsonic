package models

import "time"

// https://opensubsonic.netlify.app/docs/responses/artistwithalbumsid3/

type ArtistWithAlbumsID3 struct {
	// The id of the artist
	ID string `json:"id" xml:"id"`
	// The artist name.
	Name string `json:"name" xml:"name"`
	// A covertArt id.
	CoverArt string `json:"coverArt,omitempty" xml:"coverArt,omitempty"`
	// An url to an external image source.
	ArtistImageURL string `json:"artistImageUrl,omitempty" xml:"artistImageUrl,omitempty"`
	// 	Artist album count.
	AlbumCount int `json:"albumCount,omitempty" xml:"albumCount,omitempty"`
	// Date the artist was starred. [ISO 8601].
	Starred *time.Time `json:"starred,omitempty" xml:"starred,omitempty"`
	// The artist MusicBrainzID.
	MusicBrainzID *string `json:"musicBrainzId,omitempty" xml:"musicBrainzId,omitempty"`
	// The artist sort name.
	SortName string `json:"sortName" xml:"sortName"`
	// The list of all roles this artist has in the library.
	Roles []string `json:"roles" xml:"roles"`

	Album []AlbumID3 `json:"album,omitempty" xml:"album,omitempty"`
}
