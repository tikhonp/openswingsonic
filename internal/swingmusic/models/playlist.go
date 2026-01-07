package models

import "time"

type Playlists struct {
	Playlist []Playlist `json:"data"`
}

type Playlist struct {
	LastUpdated_ string   `json:"_last_updated"`
	Score        int64    `json:"_score"`
	Count        int      `json:"count"`
	Duration     int      `json:"duration"`
	Extra        any      `json:"extra"`
	HasImage     bool     `json:"has_image"`
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Images       []Image  `json:"images"`
	LastUpdated  string   `json:"last_updated"`
	Name         string   `json:"name"`
	Pinned       bool     `json:"pinned"`
	Settings     Settings `json:"settings"`
	Thumb        string   `json:"thumb"`
	Trackhashes  []any    `json:"trackhashes"`
	Userid       int64    `json:"userid"`
}

type Image struct {
	Color any    `json:"color"`
	Image string `json:"image"`
}

type Settings struct {
	BannerPos int64 `json:"banner_pos"`
	HasGIF    bool  `json:"has_gif"`
	Pinned    bool  `json:"pinned"`
	SquareImg bool  `json:"square_img"`
}

func (p *Playlist) GetLastUpdatedTime() (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", p.LastUpdated)
}
