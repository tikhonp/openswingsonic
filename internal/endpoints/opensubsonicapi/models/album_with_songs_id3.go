package models

import "time"

type AlbumID3WithSongs struct {
	ID                  string        `json:"id" xml:"id"`
	Name                string        `json:"name" xml:"name"`
	Version             string        `json:"version" xml:"version"`
	Artist              string        `json:"artist" xml:"artist"`
	Year                int           `json:"year" xml:"year"`
	CoverArt            string        `json:"coverArt" xml:"coverArt"`
	Starred             *time.Time    `json:"starred,omitempty" xml:"starred,omitempty"`
	Duration            int           `json:"duration" xml:"duration"`
	PlayCount           int           `json:"playCount" xml:"playCount"`
	Genre               string        `json:"genre" xml:"genre"`
	Created             time.Time     `json:"created" xml:"created"`
	ArtistID            string        `json:"artistId" xml:"artistId"`
	SongCount           int           `json:"songCount" xml:"songCount"`
	Played              time.Time     `json:"played" xml:"played"`
	UserRating          int           `json:"userRating" xml:"userRating"`
	RecordLabels        []RecordLabel `json:"recordLabels" xml:"recordLabels"`
	MusicBrainzID       string        `json:"musicBrainzId" xml:"musicBrainzId"`
	Genres              []ItemGenre   `json:"genres" xml:"genres"`
	Artists             []ArtistID3   `json:"artists" xml:"artists"`
	DisplayArtist       string        `json:"displayArtist" xml:"displayArtist"`
	ReleaseTypes        []string      `json:"releaseTypes" xml:"releaseTypes"`
	Moods               []string      `json:"moods" xml:"moods"`
	SortName            string        `json:"sortName" xml:"sortName"`
	OriginalReleaseDate ItemDate      `json:"originalReleaseDate" xml:"originalReleaseDate"`
	ReleaseDate         ItemDate      `json:"releaseDate" xml:"releaseDate"`
	IsCompilation       bool          `json:"isCompilation" xml:"isCompilation"`
	ExplicitStatus      string        `json:"explicitStatus" xml:"explicitStatus"`
	DiscTitles          []DiscTitle   `json:"discTitles" xml:"discTitles"`
	Song                []Song        `json:"song" xml:"song"`

	Parent string `json:"parent" xml:"parent"`
	Album  string `json:"album" xml:"album"`
	Title  string `json:"title" xml:"title"`
	IsDir  bool   `json:"isDir" xml:"isDir"`
}

type Song struct {
	ID                 string           `json:"id" xml:"id"`
	Parent             string           `json:"parent" xml:"parent"`
	IsDir              bool             `json:"isDir" xml:"isDir"`
	Title              string           `json:"title" xml:"title"`
	Album              string           `json:"album" xml:"album"`
	Artist             string           `json:"artist" xml:"artist"`
	Track              int64            `json:"track" xml:"track"`
	Year               int64            `json:"year" xml:"year"`
	CoverArt           string           `json:"coverArt" xml:"coverArt"`
	Size               int64            `json:"size" xml:"size"`
	ContentType        string           `json:"contentType" xml:"contentType"`
	Suffix             string           `json:"suffix" xml:"suffix"`
	Starred            *time.Time       `json:"starred,omitempty" xml:"starred,omitempty"`
	Duration           int64            `json:"duration" xml:"duration"`
	BitRate            int64            `json:"bitRate" xml:"bitRate"`
	BitDepth           int64            `json:"bitDepth" xml:"bitDepth"`
	SamplingRate       int64            `json:"samplingRate" xml:"samplingRate"`
	ChannelCount       int64            `json:"channelCount" xml:"channelCount"`
	Path               string           `json:"path" xml:"path"`
	PlayCount          int64            `json:"playCount" xml:"playCount"`
	Played             time.Time        `json:"played" xml:"played"`
	DiscNumber         int64            `json:"discNumber" xml:"discNumber"`
	Created            time.Time        `json:"created" xml:"created"`
	AlbumID            string           `json:"albumId" xml:"albumId"`
	ArtistID           string           `json:"artistId" xml:"artistId"`
	Type               string           `json:"type" xml:"type"`
	MediaType          string           `json:"mediaType" xml:"mediaType"`
	IsVideo            bool             `json:"isVideo" xml:"isVideo"`
	BPM                int64            `json:"bpm" xml:"bpm"`
	Comment            string           `json:"comment" xml:"comment"`
	SortName           string           `json:"sortName" xml:"sortName"`
	MusicBrainzID      string           `json:"musicBrainzId" xml:"musicBrainzId"`
	Isrc               []string         `json:"isrc" xml:"isrc"`
	Genres             []ItemGenre      `json:"genres" xml:"genres"`
	Artists            []ArtistFromSong `json:"artists" xml:"artists"`
	DisplayArtist      string           `json:"displayArtist" xml:"displayArtist"`
	AlbumArtists       []ArtistID3      `json:"albumArtists" xml:"albumArtists"`
	DisplayAlbumArtist string           `json:"displayAlbumArtist" xml:"displayAlbumArtist"`
	Contributors       []Contributor    `json:"contributors" xml:"contributors"`
	DisplayComposer    string           `json:"displayComposer" xml:"displayComposer"`
	Moods              []string         `json:"moods" xml:"moods"`
	ExplicitStatus     string           `json:"explicitStatus" xml:"explicitStatus"`
	ReplayGain         ReplayGain       `json:"replayGain" xml:"replayGain"`
}

type ArtistFromSong struct {
	ID   string `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

type Contributor struct {
	Role    string    `json:"role" xml:"role"`
	Artist  ArtistID3 `json:"artist" xml:"artist"`
	SubRole *string   `json:"subRole,omitempty" xml:"subRole,omitempty"`
}

type ReplayGain struct {
	TrackGain float64 `json:"trackGain" xml:"trackGain"`
	AlbumGain float64 `json:"albumGain" xml:"albumGain"`
	TrackPeak float64 `json:"trackPeak" xml:"trackPeak"`
	AlbumPeak int64   `json:"albumPeak" xml:"albumPeak"`
	BaseGain  int64   `json:"baseGain" xml:"baseGain"`
}
