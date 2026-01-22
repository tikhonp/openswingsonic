// Package models contains data structures for swing music client API.
package models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AlbumRequest represents POST /album request body
type AlbumRequest struct {
	AlbumHash  string `json:"albumhash"`
	AlbumLimit int    `json:"albumlimit,omitempty"`
}

type AlbumOtherVersionsRequest struct {
	AlbumHash    string `json:"albumhash"`
	OgAlbumTitle string `json:"og_album_title"`
}

type StreamedFileHeaders struct {
	ContentType        string
	ContentLength      int
	ContentDisposition string
	ContentRange       string
}

type Starred struct {
	Albums  []AlbumShortInfo `json:"albums"`
	Artists []ArtistItem     `json:"artists"`
	Count   StarredCount     `json:"count"`
	Recents []StarredRecent  `json:"recents"`
	Tracks  []Track          `json:"tracks"`
}

type StarredRecent struct {
	// Item Item   `json:"item"`
	Type string `json:"type"`
}

type StarredCount struct {
	Albums  int64 `json:"albums"`
	Artists int64 `json:"artists"`
	Tracks  int64 `json:"tracks"`
}
