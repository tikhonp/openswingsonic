package models

// https://opensubsonic.netlify.app/docs/responses/directory/

type Directory struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Child []Child `json:"child,omitempty"`
}

type Child struct {
	ID       string `json:"id"`
	Parent   string `json:"parent,omitempty"`
	IsDir    bool   `json:"isDir"`
	Title    string `json:"title"`
	Artist   string `json:"artist,omitempty"`
	Album    string `json:"album,omitempty"`
	CoverArt string `json:"coverArt,omitempty"`
	Duration int64  `json:"duration,omitempty"`
	Track    int64  `json:"track,omitempty"`
	Path     string `json:"path,omitempty"`
}
