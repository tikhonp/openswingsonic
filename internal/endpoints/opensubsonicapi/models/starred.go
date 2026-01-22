package models

import "time"

type Starred struct {
	Artist []ArtistStared `json:"artist"`
	Album  []AlbumStarred `json:"album"`
	Song   []SongStarred  `json:"song"`
}

type AlbumStarred struct {
	ID        string    `json:"id"`
	Parent    string    `json:"parent"`
	Album     string    `json:"album"`
	Title     string    `json:"title"`
	Name      string    `json:"name"`
	IsDir     bool      `json:"isDir"`
	CoverArt  string    `json:"coverArt"`
	SongCount int64     `json:"songCount"`
	Created   time.Time `json:"created"`
	Duration  int64     `json:"duration"`
	PlayCount int64     `json:"playCount"`
	ArtistID  string    `json:"artistId"`
	Artist    string    `json:"artist"`
	Year      int       `json:"year"`
	Genre     string    `json:"genre"`
}

type ArtistStared struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	CoverArt string    `json:"coverArt"`
	Starred  time.Time `json:"starred"`
}

type SongStarred struct {
	ID           string    `json:"id"`
	Parent       string    `json:"parent"`
	IsDir        bool      `json:"isDir"`
	Title        string    `json:"title"`
	Album        string    `json:"album"`
	Artist       string    `json:"artist"`
	Track        int64     `json:"track"`
	Year         int64     `json:"year"`
	CoverArt     string    `json:"coverArt"`
	Size         int64     `json:"size"`
	ContentType  string    `json:"contentType"`
	Suffix       string    `json:"suffix"`
	Starred      time.Time `json:"starred"`
	Duration     int64     `json:"duration"`
	BitRate      int64     `json:"bitRate"`
	BitDepth     int64     `json:"bitDepth"`
	SamplingRate int64     `json:"samplingRate"`
	ChannelCount int64     `json:"channelCount"`
	Path         string    `json:"path"`
	PlayCount    int64     `json:"playCount"`
	DiscNumber   int64     `json:"discNumber"`
	Created      time.Time `json:"created"`
	AlbumID      string    `json:"albumId"`
	ArtistID     string    `json:"artistId"`
	Type         string    `json:"type"`
	IsVideo      bool      `json:"isVideo"`
}
