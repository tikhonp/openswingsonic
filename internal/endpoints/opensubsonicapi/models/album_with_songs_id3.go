package models

import "time"

type AlbumID3WithSongs struct {
	ID                  string        `json:"id"`
	Name                string        `json:"name"`
	Version             string        `json:"version"`
	Artist              string        `json:"artist"`
	Year                int           `json:"year"`
	CoverArt            string        `json:"coverArt"`
	Starred             *time.Time    `json:"starred,omitempty"`
	Duration            int           `json:"duration"`
	PlayCount           int           `json:"playCount"`
	Genre               string        `json:"genre"`
	Created             time.Time     `json:"created"`
	ArtistID            string        `json:"artistId"`
	SongCount           int           `json:"songCount"`
	Played              time.Time     `json:"played"`
	UserRating          int           `json:"userRating"`
	RecordLabels        []RecordLabel `json:"recordLabels"`
	MusicBrainzID       string        `json:"musicBrainzId"`
	Genres              []ItemGenre   `json:"genres"`
	Artists             []ArtistID3   `json:"artists"`
	DisplayArtist       string        `json:"displayArtist"`
	ReleaseTypes        []string      `json:"releaseTypes"`
	Moods               []string      `json:"moods"`
	SortName            string        `json:"sortName"`
	OriginalReleaseDate ItemDate      `json:"originalReleaseDate"`
	ReleaseDate         ItemDate      `json:"releaseDate"`
	IsCompilation       bool          `json:"isCompilation"`
	ExplicitStatus      string        `json:"explicitStatus"`
	DiscTitles          []DiscTitle   `json:"discTitles"`
	Song                []Song        `json:"song"`

	Parent string `json:"parent"`
	Album  string `json:"album"`
	Title  string `json:"title"`
	IsDir  bool   `json:"isDir"`
}

type Song struct {
	ID                 string           `json:"id"`
	Parent             string           `json:"parent"`
	IsDir              bool             `json:"isDir"`
	Title              string           `json:"title"`
	Album              string           `json:"album"`
	Artist             string           `json:"artist"`
	Track              int64            `json:"track"`
	Year               int64            `json:"year"`
	CoverArt           string           `json:"coverArt"`
	Size               int64            `json:"size"`
	ContentType        string           `json:"contentType"`
	Suffix             string           `json:"suffix"`
	Starred            time.Time        `json:"starred"`
	Duration           int64            `json:"duration"`
	BitRate            int64            `json:"bitRate"`
	BitDepth           int64            `json:"bitDepth"`
	SamplingRate       int64            `json:"samplingRate"`
	ChannelCount       int64            `json:"channelCount"`
	Path               string           `json:"path"`
	PlayCount          int64            `json:"playCount"`
	Played             time.Time        `json:"played"`
	DiscNumber         int64            `json:"discNumber"`
	Created            time.Time        `json:"created"`
	AlbumID            string           `json:"albumId"`
	ArtistID           string           `json:"artistId"`
	Type               string           `json:"type"`
	MediaType          string           `json:"mediaType"`
	IsVideo            bool             `json:"isVideo"`
	BPM                int64            `json:"bpm"`
	Comment            string           `json:"comment"`
	SortName           string           `json:"sortName"`
	MusicBrainzID      string           `json:"musicBrainzId"`
	Isrc               []string         `json:"isrc"`
	Genres             []ItemGenre      `json:"genres"`
	Artists            []ArtistFromSong `json:"artists"`
	DisplayArtist      string           `json:"displayArtist"`
	AlbumArtists       []ArtistID3      `json:"albumArtists"`
	DisplayAlbumArtist string           `json:"displayAlbumArtist"`
	Contributors       []Contributor    `json:"contributors"`
	DisplayComposer    string           `json:"displayComposer"`
	Moods              []string         `json:"moods"`
	ExplicitStatus     string           `json:"explicitStatus"`
	ReplayGain         ReplayGain       `json:"replayGain"`
}

type ArtistFromSong struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Contributor struct {
	Role    string          `json:"role"`
	Artist  ArtistID3 `json:"artist"`
	SubRole *string         `json:"subRole,omitempty"`
}

type ReplayGain struct {
	TrackGain float64 `json:"trackGain"`
	AlbumGain float64 `json:"albumGain"`
	TrackPeak float64 `json:"trackPeak"`
	AlbumPeak int64   `json:"albumPeak"`
	BaseGain  int64   `json:"baseGain"`
}
