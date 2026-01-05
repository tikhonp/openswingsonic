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
