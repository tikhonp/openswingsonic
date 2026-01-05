package models

type AlbumID3 struct {
	ID     string  `json:"id"`
	Parent *string `json:"parent,omitempty"`
	Name   string  `json:"name"`
	Title  string  `json:"title,omitempty"`
	Album  string  `json:"album,omitempty"`

	IsDir     bool    `json:"isDir"`
	CoverArt  *string `json:"coverArt,omitempty"`
	SongCount int     `json:"songCount"`
	Duration  int     `json:"duration"`

	Created *string `json:"created,omitempty"`
	Year    *int    `json:"year,omitempty"`
	Genre   *string `json:"genre,omitempty"`

	Artist   *string `json:"artist,omitempty"`
	ArtistId *string `json:"artistId,omitempty"`
}
