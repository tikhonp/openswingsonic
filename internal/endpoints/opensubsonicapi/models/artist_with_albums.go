package models

// https://opensubsonic.netlify.app/docs/responses/artistwithalbumsid3/

type ArtistWithAlbumsID3 struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	CoverArt   *string     `json:"coverArt,omitempty"`
	AlbumCount int         `json:"albumCount"`

	ArtistImageUrl *string `json:"artistImageUrl,omitempty"`
	Starred        *string `json:"starred,omitempty"`

	// OpenSubsonic optional extensions
	MusicBrainzId *string  `json:"musicBrainzId,omitempty"`
	SortName      *string  `json:"sortName,omitempty"`
	Roles         []string `json:"roles,omitempty"`

	Album []AlbumID3 `json:"album,omitempty"`
}
