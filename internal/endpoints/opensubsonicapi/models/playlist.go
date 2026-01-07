package models

import "time"

type Playlists struct {
	// A list of playlists
	Playlist []Playlist `json:"playlist"`
}

type Playlist struct {
	// Id of the playlist
	ID string `json:"id"`
	// Name of the playlist
	Name string `json:"name"`
	// A commnet
	Comment string `json:"comment,omitempty"`
	// Owner of the playlist
	Owner string `json:"owner,omitempty"`
	// Is the playlist public
	Public bool `json:"public,omitempty"`
	// number of songs
	SongCount int `json:"songCount"`
	// Playlist duration in seconds
	Duration int `json:"duration"`
	// Creation date [ISO 8601]
	Created time.Time `json:"created"`
	// Last changed date [ISO 8601]
	Changed time.Time `json:"changed"`
	// 	A cover Art Id
	CoverArt string `json:"coverArt,omitempty"`
	// 	A list of allowed usernames
	AllowedUser []string `json:"allowedUser,omitempty"`
}
