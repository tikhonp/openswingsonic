package models

import "time"

type AlbumList struct {
	Album []AlbumID3 `json:"album" xml:"album"`
}

// AlbumID3 is an album from ID3 tags.
type AlbumID3 struct {
	// The id of the album
	ID string `json:"id" xml:"id"`
	// The album name.
	Name string `json:"name" xml:"name"`
	// The album version name (Remastered, Anniversary Box Set, …).
	Version string `json:"version" xml:"version"`
	// Artist name.
	Artist string `json:"artist,omitempty" xml:"artist,omitempty"`
	// The id of the artist.
	ArtistID string `json:"artistId,omitempty" xml:"artistId,omitempty"`
	// A covertArt id.
	CoverArt string `json:"coverArt,omitempty" xml:"coverArt,omitempty"`
	// Number of songs.
	SongCount int `json:"songCount" xml:"songCount"`
	// Total duration of the album in seconds.
	Duration int `json:"duration" xml:"duration"`
	// Number of play of the album.
	PlayCount int `json:"playCount,omitempty" xml:"playCount,omitempty"`
	// Date the album was added. [ISO 8601].
	Created time.Time `json:"created" xml:"created"`
	// Date the album was starred. [ISO 8601].
	Starred *time.Time `json:"starred" xml:"starred"`
	// The album year.
	Year int `json:"year,omitempty" xml:"year,omitempty"`
	// The album genre.
	Genre string `json:"genre,omitempty" xml:"genre,omitempty"`
	// Date the album was last played. [ISO 8601].
	Played time.Time `json:"played" xml:"played"`
	// The user rating of the album. [1-5].
	UserRating int `json:"userRating" xml:"userRating"`
	// The labels producing the album.
	RecordLabels []RecordLabel `json:"recordLabels" xml:"recordLabels"`
	// The album MusicBrainzID.
	MusicBrainzID string `json:"musicBrainzId" xml:"musicBrainzId"`
	// The list of all genres of the album.
	Genres []ItemGenre `json:"genres" xml:"genres"`
	// The list of all album artists of the album.
	// (Note: Only the required ArtistID3 fields should be returned by default).
	Artists []ArtistID3 `json:"artists" xml:"artists"`
	// The single value display artist.
	DisplayArtist string `json:"displayArtist" xml:"displayArtist"`
	// The types of this album release. (Album, Compilation, EP, Remix, …).
	RealeaseTypes []string `json:"releaseTypes" xml:"releaseTypes"`
	// The list of all moods of the album.
	Moods []string `json:"moods" xml:"moods"`
	// The album sort name.
	SortName string `json:"sortName" xml:"sortName"`
	// Date the album was originally released.
	OriginalReleaseDate ItemDate `json:"originalReleaseDate" xml:"originalReleaseDate"`
	// Date the specific edition of the album was released.
	// Note: for files using ID3 tags, releaseDate should generally be read from the TDRL tag.
	// Servers that use a different source for this field should document the behavior.
	ReleaseDate ItemDate `json:"releaseDate" xml:"releaseDate"`
	// True if the album is a compilation.
	IsCompilation bool `json:"isCompilation" xml:"isCompilation"`
	// Returns “explicit” if at least one song is explicit,
	// “clean” if no song is explicit and at least one is “clean” else “”.
	ExplicitStatus string `json:"explicitStatus" xml:"explicitStatus"`
	// The list of all disc titles of the album.
	DiskTitles []DiscTitle `json:"discTitles" xml:"discTitles"`

	Parent string `json:"parent,omitempty" xml:"parent,omitempty"`
	Album  string `json:"album" xml:"album"`
	Title  string `json:"title" xml:"title"`
	IsDir  bool   `json:"isDir" xml:"isDir"`
}

type RecordLabel struct {
	// The record label name.
	Name string `json:"name" xml:"name"`
}

type ItemGenre struct {
	// The genre name.
	Name string `json:"name" xml:"name"`
}

type ItemDate struct {
	// The year
	Year int `json:"year" xml:"year"`
	// The month (1-12)
	Month int `json:"month" xml:"month"`
	// The day (1-31)
	Day int `json:"day" xml:"day"`
}

type DiscTitle struct {
	// The disc number.
	DiscNumber int `json:"discNumber" xml:"discNumber"`
	// The disc title.
	Title string `json:"title" xml:"title"`
}
