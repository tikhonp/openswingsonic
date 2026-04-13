package models

import "time"

type AlbumID3WithSongs struct {
	ID                  string        `json:"id" xml:"id,attr"`
	Name                string        `json:"name" xml:"name,attr"`
	Version             string        `json:"version" xml:"version,attr"`
	Artist              string        `json:"artist" xml:"artist,attr"`
	Year                int           `json:"year" xml:"year,attr"`
	CoverArt            string        `json:"coverArt" xml:"coverArt,attr"`
	Starred             *time.Time    `json:"starred,omitempty" xml:"starred,attr,omitempty"`
	Duration            int           `json:"duration" xml:"duration,attr"`
	PlayCount           int           `json:"playCount" xml:"playCount,attr"`
	Genre               string        `json:"genre" xml:"genre,attr"`
	Created             time.Time     `json:"created" xml:"created,attr"`
	ArtistID            string        `json:"artistId" xml:"artistId,attr"`
	SongCount           int           `json:"songCount" xml:"songCount,attr"`
	Played              time.Time     `json:"played" xml:"played,attr"`
	UserRating          int           `json:"userRating" xml:"userRating,attr"`
	RecordLabels        []RecordLabel `json:"recordLabels" xml:"recordLabels"`
	MusicBrainzID       string        `json:"musicBrainzId" xml:"musicBrainzId,attr"`
	Genres              []ItemGenre   `json:"genres" xml:"genres"`
	Artists             []ArtistID3   `json:"artists" xml:"artists"`
	DisplayArtist       string        `json:"displayArtist" xml:"displayArtist,attr"`
	ReleaseTypes        []string      `json:"releaseTypes" xml:"releaseTypes"`
	Moods               []string      `json:"moods" xml:"moods"`
	SortName            string        `json:"sortName" xml:"sortName,attr"`
	OriginalReleaseDate ItemDate      `json:"originalReleaseDate" xml:"originalReleaseDate"`
	ReleaseDate         ItemDate      `json:"releaseDate" xml:"releaseDate"`
	IsCompilation       bool          `json:"isCompilation" xml:"isCompilation,attr"`
	ExplicitStatus      string        `json:"explicitStatus" xml:"explicitStatus,attr"`
	DiscTitles          []DiscTitle   `json:"discTitles" xml:"discTitles"`
	Song                []Song        `json:"song" xml:"song"`

	Parent string `json:"parent" xml:"parent,attr"`
	Album  string `json:"album" xml:"album,attr"`
	Title  string `json:"title" xml:"title,attr"`
	IsDir  bool   `json:"isDir" xml:"isDir,attr"`
}

type Song struct {
	ID                 string           `json:"id" xml:"id,attr"`
	Parent             string           `json:"parent" xml:"parent,attr"`
	IsDir              bool             `json:"isDir" xml:"isDir,attr"`
	Title              string           `json:"title" xml:"title,attr"`
	Album              string           `json:"album" xml:"album,attr"`
	Artist             string           `json:"artist" xml:"artist,attr"`
	Track              int64            `json:"track" xml:"track,attr"`
	Year               int64            `json:"year" xml:"year,attr"`
	CoverArt           string           `json:"coverArt" xml:"coverArt,attr"`
	Size               int64            `json:"size" xml:"size,attr"`
	ContentType        string           `json:"contentType" xml:"contentType,attr"`
	Suffix             string           `json:"suffix" xml:"suffix,attr"`
	Starred            *time.Time       `json:"starred,omitempty" xml:"starred,attr,omitempty"`
	Duration           int64            `json:"duration" xml:"duration,attr"`
	BitRate            int64            `json:"bitRate" xml:"bitRate,attr"`
	BitDepth           int64            `json:"bitDepth" xml:"bitDepth,attr"`
	SamplingRate       int64            `json:"samplingRate" xml:"samplingRate,attr"`
	ChannelCount       int64            `json:"channelCount" xml:"channelCount,attr"`
	Path               string           `json:"path" xml:"path,attr"`
	PlayCount          int64            `json:"playCount" xml:"playCount,attr"`
	Played             time.Time        `json:"played" xml:"played,attr"`
	DiscNumber         int64            `json:"discNumber" xml:"discNumber,attr"`
	Created            time.Time        `json:"created" xml:"created,attr"`
	AlbumID            string           `json:"albumId" xml:"albumId,attr"`
	ArtistID           string           `json:"artistId" xml:"artistId,attr"`
	Type               string           `json:"type" xml:"type,attr"`
	MediaType          string           `json:"mediaType" xml:"mediaType,attr"`
	IsVideo            bool             `json:"isVideo" xml:"isVideo,attr"`
	BPM                int64            `json:"bpm" xml:"bpm,attr"`
	Comment            string           `json:"comment" xml:"comment,attr"`
	SortName           string           `json:"sortName" xml:"sortName,attr"`
	MusicBrainzID      string           `json:"musicBrainzId" xml:"musicBrainzId,attr"`
	Isrc               []string         `json:"isrc" xml:"isrc"`
	Genres             []ItemGenre      `json:"genres" xml:"genres"`
	Artists            []ArtistFromSong `json:"artists" xml:"artists"`
	DisplayArtist      string           `json:"displayArtist" xml:"displayArtist,attr"`
	AlbumArtists       []ArtistID3      `json:"albumArtists" xml:"albumArtists"`
	DisplayAlbumArtist string           `json:"displayAlbumArtist" xml:"displayAlbumArtist,attr"`
	Contributors       []Contributor    `json:"contributors" xml:"contributors"`
	DisplayComposer    string           `json:"displayComposer" xml:"displayComposer,attr"`
	Moods              []string         `json:"moods" xml:"moods"`
	ExplicitStatus     string           `json:"explicitStatus" xml:"explicitStatus,attr"`
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
