package models

import "time"

type AlbumList struct {
	Album []AlbumID3 `json:"album"`
}

// AlbumID3 is an album from ID3 tags.
type AlbumID3 struct {
	// The id of the album
	ID string `json:"id"`
	// The album name.
	Name string `json:"name"`
	// The album version name (Remastered, Anniversary Box Set, …).
	Version string `json:"version"`
	// Artist name.
	Artist string `json:"artist,omitempty"`
	// The id of the artist.
	ArtistID string `json:"artistId,omitempty"`
	// A covertArt id.
	CoverArt string `json:"coverArt,omitempty"`
	// Number of songs.
	SongCount int `json:"songCount"`
	// Total duration of the album in seconds.
	Duration int `json:"duration"`
	// Number of play of the album.
	PlayCount int `json:"playCount,omitempty"`
	// Date the album was added. [ISO 8601].
	Created time.Time `json:"created"`
	// Date the album was starred. [ISO 8601].
	Starred *time.Time `json:"starred"`
	// The album year.
	Year int `json:"year,omitempty"`
	// The album genre.
	Genre string `json:"genre,omitempty"`
	// Date the album was last played. [ISO 8601].
	Played time.Time `json:"played"`
	// The user rating of the album. [1-5].
	UserRating int `json:"userRating"`
	// The labels producing the album.
	RecordLabels []RecordLabel `json:"recordLabels"`
	// The album MusicBrainzID.
	MusicBrainzID string `json:"musicBrainzId"`
	// The list of all genres of the album.
	Genres []ItemGenre `json:"genres"`
	// The list of all album artists of the album.
	// (Note: Only the required ArtistID3 fields should be returned by default).
	Artists []ArtistID3 `json:"artists"`
	// The single value display artist.
	DisplayArtist string `json:"displayArtist"`
	// The types of this album release. (Album, Compilation, EP, Remix, …).
	RealeaseTypes []string `json:"releaseTypes"`
	// The list of all moods of the album.
	Moods []string `json:"moods"`
	// The album sort name.
	SortName string `json:"sortName"`
	// Date the album was originally released.
	OriginalReleaseDate ItemDate `json:"originalReleaseDate"`
	// Date the specific edition of the album was released.
	// Note: for files using ID3 tags, releaseDate should generally be read from the TDRL tag.
	// Servers that use a different source for this field should document the behavior.
	ReleaseDate ItemDate `json:"releaseDate"`
	// True if the album is a compilation.
	IsCompilation bool `json:"isCompilation"`
	// Returns “explicit” if at least one song is explicit,
	// “clean” if no song is explicit and at least one is “clean” else “”.
	ExplicitStatus string `json:"explicitStatus"`
	// The list of all disc titles of the album.
	DiskTitles []DiscTitle `json:"discTitles"`

	Parent string `json:"parent,omitempty"`
	Album  string `json:"album"`
	Title  string `json:"title"`
	IsDir  bool   `json:"isDir"`
}

type RecordLabel struct {
	// The record label name.
	Name string `json:"name"`
}

type ItemGenre struct {
	// The genre name.
	Name string `json:"name"`
}

type ItemDate struct {
	// The year
	Year int `json:"year"`
	// The month (1-12)
	Month int `json:"month"`
	// The day (1-31)
	Day int `json:"day"`
}

type DiscTitle struct {
	// The disc number.
	DiscNumber int `json:"discNumber"`
	// The disc title.
	Title string `json:"title"`
}
