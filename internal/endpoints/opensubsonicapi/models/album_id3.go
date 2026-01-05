package models

type AlbumList struct {
	Album []AlbumID3 `json:"album"`
}

type AlbumID3 struct {
	ID        string `json:"id"`
	Parent    string `json:"parent,omitempty"`
	Album     string `json:"album"`
	Title     string `json:"title"`
	Name      string `json:"name"`
	IsDir     bool   `json:"isDir"`
	CoverArt  string `json:"coverArt,omitempty"`
	SongCount int    `json:"songCount,omitempty"`
	Created   string `json:"created,omitempty"`
	Duration  int    `json:"duration,omitempty"`
	PlayCount int    `json:"playCount,omitempty"`
	ArtistID  string `json:"artistId,omitempty"`
	Artist    string `json:"artist,omitempty"`
	Year      int    `json:"year,omitempty"`
	Genre     string `json:"genre,omitempty"`
}
