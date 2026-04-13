package models

import "time"

type Starred struct {
	Artist []ArtistStared `json:"artist" xml:"artist"`
	Album  []AlbumStarred `json:"album" xml:"album"`
	Song   []SongStarred  `json:"song" xml:"song"`
}

type AlbumStarred struct {
	ID        string    `json:"id" xml:"id"`
	Parent    string    `json:"parent" xml:"parent"`
	Album     string    `json:"album" xml:"album"`
	Title     string    `json:"title" xml:"title"`
	Name      string    `json:"name" xml:"name"`
	IsDir     bool      `json:"isDir" xml:"isDir"`
	CoverArt  string    `json:"coverArt" xml:"coverArt"`
	SongCount int64     `json:"songCount" xml:"songCount"`
	Created   time.Time `json:"created" xml:"created"`
	Duration  int64     `json:"duration" xml:"duration"`
	PlayCount int64     `json:"playCount" xml:"playCount"`
	ArtistID  string    `json:"artistId" xml:"artistId"`
	Artist    string    `json:"artist" xml:"artist"`
	Year      int       `json:"year" xml:"year"`
	Genre     string    `json:"genre" xml:"genre"`
}

type ArtistStared struct {
	ID       string     `json:"id" xml:"id"`
	Name     string     `json:"name" xml:"name"`
	CoverArt string     `json:"coverArt" xml:"coverArt"`
	Starred  *time.Time `json:"starred,omitempty" xml:"starred,omitempty"`
}

type SongStarred struct {
	ID           string     `json:"id" xml:"id"`
	Parent       string     `json:"parent" xml:"parent"`
	IsDir        bool       `json:"isDir" xml:"isDir"`
	Title        string     `json:"title" xml:"title"`
	Album        string     `json:"album" xml:"album"`
	Artist       string     `json:"artist" xml:"artist"`
	Track        int64      `json:"track" xml:"track"`
	Year         int64      `json:"year" xml:"year"`
	CoverArt     string     `json:"coverArt" xml:"coverArt"`
	Size         int64      `json:"size" xml:"size"`
	ContentType  string     `json:"contentType" xml:"contentType"`
	Suffix       string     `json:"suffix" xml:"suffix"`
	Starred      *time.Time `json:"starred,omitempty" xml:"starred,omitempty"`
	Duration     int64      `json:"duration" xml:"duration"`
	BitRate      int64      `json:"bitRate" xml:"bitRate"`
	BitDepth     int64      `json:"bitDepth" xml:"bitDepth"`
	SamplingRate int64      `json:"samplingRate" xml:"samplingRate"`
	ChannelCount int64      `json:"channelCount" xml:"channelCount"`
	Path         string     `json:"path" xml:"path"`
	PlayCount    int64      `json:"playCount" xml:"playCount"`
	DiscNumber   int64      `json:"discNumber" xml:"discNumber"`
	Created      time.Time  `json:"created" xml:"created"`
	AlbumID      string     `json:"albumId" xml:"albumId"`
	ArtistID     string     `json:"artistId" xml:"artistId"`
	Type         string     `json:"type" xml:"type"`
	IsVideo      bool       `json:"isVideo" xml:"isVideo"`
}
