package models

// https://opensubsonic.netlify.app/docs/responses/artists/

type ArtistsID3 struct {
	IgnoredArticles string     `json:"ignoredArticles"`
	Index           []IndexID3 `json:"index,omitempty"`
}

type IndexID3 struct {
	Name   string      `json:"name"`
	Artist []ArtistID3 `json:"artist,omitempty"`
}

type ArtistID3 struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	CoverArt   *string `json:"coverArt,omitempty"`
	AlbumCount int     `json:"albumCount,omitempty"`

	// OpenSubsonic optional extensions (not supported yet)
	// MusicBrainzId *string  `json:"musicBrainzId,omitempty"`
	// SortName      *string  `json:"sortName,omitempty"`
	// Roles         []string `json:"roles,omitempty"`
}
