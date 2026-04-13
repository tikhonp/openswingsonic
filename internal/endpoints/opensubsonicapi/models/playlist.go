package models

import "time"

type Playlists struct {
	// A list of playlists
	Playlist []Playlist `json:"playlist" xml:"playlist"`
}

type Playlist struct {
	// Id of the playlist
	ID string `json:"id" xml:"id"`
	// Name of the playlist
	Name string `json:"name" xml:"name"`
	// A commnet
	Comment string `json:"comment,omitempty" xml:"comment,omitempty"`
	// Owner of the playlist
	Owner string `json:"owner,omitempty" xml:"owner,omitempty"`
	// Is the playlist public
	Public bool `json:"public,omitempty" xml:"public,omitempty"`
	// number of songs
	SongCount int `json:"songCount" xml:"songCount"`
	// Playlist duration in seconds
	Duration int `json:"duration" xml:"duration"`
	// Creation date [ISO 8601]
	Created time.Time `json:"created" xml:"created"`
	// Last changed date [ISO 8601]
	Changed time.Time `json:"changed" xml:"changed"`
	// 	A cover Art Id
	CoverArt string `json:"coverArt,omitempty" xml:"coverArt,omitempty"`
	// 	A list of allowed usernames
	AllowedUser []string `json:"allowedUser,omitempty" xml:"allowedUser,omitempty"`
}

type PlaylistWithEntries struct {
	Playlist
	Entry []Song `json:"entry" xml:"entry"`
}
