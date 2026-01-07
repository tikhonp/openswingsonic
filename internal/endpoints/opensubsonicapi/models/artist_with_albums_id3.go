package models

// https://opensubsonic.netlify.app/docs/responses/artistwithalbumsid3/

type ArtistWithAlbumsID3 struct {
	// The id of the artist
	ID string `json:"id"`
	// The artist name.
	Name string `json:"name"`
	// A covertArt id.
	CoverArt string `json:"coverArt,omitempty"`
	// An url to an external image source.
	ArtistImageURL string `json:"artistImageUrl,omitempty"`
	// 	Artist album count.
	AlbumCount int `json:"albumCount,omitempty"`
	// Date the artist was starred. [ISO 8601].
	Starred string `json:"starred"`
	// The artist MusicBrainzID.
	MusicBrainzID string `json:"musicBrainzId"`
	// The artist sort name.
	SortName string `json:"sortName"`
	// The list of all roles this artist has in the library.
	Roles []string `json:"roles"`

	Album []AlbumID3 `json:"album,omitempty"`
}
